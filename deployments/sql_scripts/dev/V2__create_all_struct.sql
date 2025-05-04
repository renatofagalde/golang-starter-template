SET search_path TO public;
-- Create extension for UUID support
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Article source table
CREATE TABLE IF NOT EXISTS article_source (
                                              id SERIAL PRIMARY KEY,
                                              ide UUID DEFAULT gen_random_uuid(),
                                              created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                              updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                              deleted BOOLEAN DEFAULT FALSE,
                                              source_id VARCHAR(255),
                                              name VARCHAR(255) NOT NULL,
                                              CONSTRAINT uni_article_source_ide UNIQUE (ide)
);

COMMENT ON TABLE article_source IS 'Stores article sources with optional source ID';

-- Articles table
CREATE TABLE IF NOT EXISTS article (
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
                                       CONSTRAINT uni_article_ide UNIQUE (ide),
                                       CONSTRAINT uni_article_url UNIQUE (url),
                                       CONSTRAINT fk_article_source FOREIGN KEY (source_id) REFERENCES article_source(id) ON DELETE CASCADE
);

COMMENT ON TABLE article IS 'Stores articles with their metadata';

-- Note table
CREATE TABLE IF NOT EXISTS note (
                                    id SERIAL PRIMARY KEY,
                                    ide UUID DEFAULT gen_random_uuid(),
                                    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    deleted BOOLEAN DEFAULT FALSE,
                                    status VARCHAR(100) NOT NULL,
                                    total_results INT NOT NULL DEFAULT 0,
                                    query_text VARCHAR(255),
                                    CONSTRAINT uni_note_ide UNIQUE (ide)
);

COMMENT ON TABLE note IS 'Stores search queries with result status and count';

-- Note-article association table
CREATE TABLE IF NOT EXISTS note_article (
                                            id SERIAL PRIMARY KEY,
                                            ide UUID DEFAULT gen_random_uuid(),
                                            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            deleted BOOLEAN DEFAULT FALSE,
                                            note_id INT NOT NULL,
                                            article_id INT NOT NULL,
                                            CONSTRAINT fk_note_id FOREIGN KEY (note_id) REFERENCES note(id) ON DELETE CASCADE,
                                            CONSTRAINT fk_article_id FOREIGN KEY (article_id) REFERENCES article(id) ON DELETE CASCADE,
                                            CONSTRAINT uni_note_article UNIQUE (note_id, article_id)
);

COMMENT ON TABLE note_article IS 'Relates notes to their articles';

-- History tables
CREATE TABLE IF NOT EXISTS article_source_history AS TABLE article_source WITH NO DATA;
CREATE TABLE IF NOT EXISTS article_history AS TABLE article WITH NO DATA;
CREATE TABLE IF NOT EXISTS note_history AS TABLE note WITH NO DATA;
CREATE TABLE IF NOT EXISTS note_article_history AS TABLE note_article WITH NO DATA;

-- History update function
CREATE OR REPLACE FUNCTION insert_into_history() RETURNS TRIGGER AS $$
DECLARE
    target_table TEXT;
BEGIN
    -- Define target history table
    target_table := format('%I_history', TG_TABLE_NAME);

    -- Insert into history using the structure of the OLD row
    EXECUTE format('INSERT INTO %s SELECT ($1).*', target_table)
        USING OLD;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Configure history triggers
DO $$
    DECLARE
        tbl RECORD;
        trigger_name TEXT;
    BEGIN
        FOR tbl IN
            SELECT tablename
            FROM pg_tables
            WHERE schemaname = 'public'
              AND tablename NOT LIKE '%_history'
              AND tablename IN ('article_source', 'article', 'note', 'note_article')
            LOOP
                trigger_name := tbl.tablename || '_history';
                EXECUTE format('DROP TRIGGER IF EXISTS %I ON %I', trigger_name, tbl.tablename);
                EXECUTE format(
                        'CREATE TRIGGER %I
                         BEFORE UPDATE ON %I
                         FOR EACH ROW
                         EXECUTE FUNCTION insert_into_history()',
                        trigger_name, tbl.tablename
                        );
            END LOOP;
    END;
$$;