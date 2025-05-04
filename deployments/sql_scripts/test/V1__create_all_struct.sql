-- SELECT pg_terminate_backend(pg_stat_activity.pid)
-- FROM pg_stat_activity
-- WHERE pg_stat_activity.datname = 'note'
--   AND pid <> pg_backend_pid();

DROP schema IF EXISTS public;
CREATE schema public;


DO $$
    BEGIN
        PERFORM 1;
    END $$;