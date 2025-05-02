DROP database IF EXISTS note;
create database note;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- uuid

CREATE TABLE IF NOT EXISTS "note"."article_source" (
                                                       id SERIAL PRIMARY KEY,
                                                       ide UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    source_id VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    CONSTRAINT "uni_note_article_source_ide" UNIQUE ("ide")
    );

COMMENT ON TABLE "note"."article_source" IS 'Stores article sources with optional source ID';

-- Tabela de artigos
CREATE TABLE IF NOT EXISTS "note"."article" (
                                                id SERIAL PRIMARY KEY,
                                                ide UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    source_id INT NOT NULL,
    author VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    url VARCHAR(512) NOT NULL,
    url_to_image VARCHAR(512),
    published_at TIMESTAMP,
    content TEXT,
    CONSTRAINT "uni_note_article_ide" UNIQUE ("ide"),
    CONSTRAINT "uni_note_article_url" UNIQUE ("url"),
    CONSTRAINT fk_article_source FOREIGN KEY (source_id) REFERENCES note.article_source(id) ON DELETE CASCADE
    );

COMMENT ON TABLE "note"."article" IS 'Stores note articles with their metadata';

-- Tabela de notas (consultas para agrupar artigos)
CREATE TABLE IF NOT EXISTS "note"."note" (
                                             id SERIAL PRIMARY KEY,
                                             ide UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    status VARCHAR(100) NOT NULL,
    total_results INT NOT NULL DEFAULT 0,
    query_text VARCHAR(255),
    CONSTRAINT "uni_note_note_ide" UNIQUE ("ide")
    );

COMMENT ON TABLE "note"."note" IS 'Stores search queries with result status and count';

-- Tabela de associação entre notas e artigos
CREATE TABLE IF NOT EXISTS "note"."note_article" (
                                                     id SERIAL PRIMARY KEY,
                                                     ide UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN DEFAULT FALSE,
    note_id INT NOT NULL,
    article_id INT NOT NULL,
    CONSTRAINT fk_note_id FOREIGN KEY (note_id) REFERENCES note.note(id) ON DELETE CASCADE,
    CONSTRAINT fk_article_id FOREIGN KEY (article_id) REFERENCES note.article(id) ON DELETE CASCADE,
    CONSTRAINT uni_note_article UNIQUE (note_id, article_id)
    );

COMMENT ON TABLE "note"."note_article" IS 'Relates notes to their articles';

-- Criação das tabelas de histórico
CREATE TABLE IF NOT EXISTS note.article_source_history AS TABLE note.article_source WITH NO DATA;
CREATE TABLE IF NOT EXISTS note.article_history AS TABLE note.article WITH NO DATA;
CREATE TABLE IF NOT EXISTS note.note_history AS TABLE note.note WITH NO DATA;
CREATE TABLE IF NOT EXISTS note.note_article_history AS TABLE note.note_article WITH NO DATA;

-- Função para inserir em histórico antes de update
CREATE OR REPLACE FUNCTION note.insert_into_history() RETURNS TRIGGER AS $$
DECLARE
target_table TEXT;
BEGIN
    -- Definindo a tabela de destino
    target_table := format('%I.%I_history', TG_TABLE_SCHEMA, TG_TABLE_NAME);

    -- Inserindo em histórico usando a estrutura da linha `OLD`
EXECUTE format('INSERT INTO %s SELECT ($1).*', target_table)
    USING OLD;

RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Configuração dos triggers de histórico
DO $$
    DECLARE
tbl RECORD;
        trigger_name TEXT;
BEGIN
FOR tbl IN
SELECT tablename
FROM pg_tables
WHERE schemaname = 'note'
  AND tablename NOT LIKE '%_history'
    LOOP
                trigger_name := tbl.tablename || '_history';
EXECUTE format('DROP TRIGGER IF EXISTS %I ON note.%I', trigger_name, tbl.tablename);
EXECUTE format(
        'CREATE TRIGGER %I
         BEFORE UPDATE ON note.%I
         FOR EACH ROW
         EXECUTE FUNCTION note.insert_into_history()',
        trigger_name, tbl.tablename
        );
END LOOP;
END;
$$;