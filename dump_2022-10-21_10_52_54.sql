--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE chat_app;




pg_dumpall: error: query failed: ERROR:  permission denied for table pg_authid
pg_dumpall: error: query was: SELECT rolname FROM pg_authid WHERE rolname !~ '^pg_' ORDER BY 1
