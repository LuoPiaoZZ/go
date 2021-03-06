--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: test; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test (
    test json
);


ALTER TABLE public.test OWNER TO postgres;

--
-- Data for Name: test; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test (test) FROM stdin;
[{"a":1,"b":2},{"a":1,"b":2}]
[{"a":1,"b":2},{"a":1,"b":3}]
[{"a":1,"b":4},{"a":1,"b":3}]
\.


--
-- PostgreSQL database dump complete
--

