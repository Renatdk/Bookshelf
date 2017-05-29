--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: library; Type: TABLE; Schema: public; Owner: tbzobuabpgyvwj (heroku db); Tablespace:
--

CREATE TABLE library (
    id integer NOT NULL,
    user_id integer,
    title character varying,
    address text,
    phone text,
    updated_at integer,
    created_at integer
);


ALTER TABLE library OWNER TO tbzobuabpgyvwj;

--
-- Name: library_id_seq; Type: SEQUENCE; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

CREATE SEQUENCE library_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE library_id_seq OWNER TO tbzobuabpgyvwj;

--
-- Name: library_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

ALTER SEQUENCE library_id_seq OWNED BY library.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: tbzobuabpgyvwj (heroku db); Tablespace:
--

CREATE TABLE "user" (
    id integer NOT NULL,
    email character varying,
    password character varying,
    name character varying,
    updated_at integer,
    created_at integer
);


ALTER TABLE "user" OWNER TO tbzobuabpgyvwj;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO tbzobuabpgyvwj;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

ALTER TABLE ONLY library ALTER COLUMN id SET DEFAULT nextval('library_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- Data for Name: library; Type: TABLE DATA; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

COPY library (id, user_id, title, content, updated_at, created_at) FROM stdin;
\.


--
-- Name: library_id_seq; Type: SEQUENCE SET; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

SELECT pg_catalog.setval('library_id_seq', 1, false);


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

COPY "user" (id, email, password, name, updated_at, created_at) FROM stdin;
\.


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

SELECT pg_catalog.setval('user_id_seq', 1, false);


--
-- Name: library_id; Type: CONSTRAINT; Schema: public; Owner: tbzobuabpgyvwj (heroku db); Tablespace:
--

ALTER TABLE ONLY library
    ADD CONSTRAINT library_id PRIMARY KEY (id);


--
-- Name: user_id; Type: CONSTRAINT; Schema: public; Owner: tbzobuabpgyvwj (heroku db); Tablespace:
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_id PRIMARY KEY (id);


--
-- Name: article_user_id; Type: FK CONSTRAINT; Schema: public; Owner: tbzobuabpgyvwj (heroku db)
--

ALTER TABLE ONLY library
    ADD CONSTRAINT article_user_id FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: public; Type: ACL; Schema: -; Owner: tbzobuabpgyvwj (heroku db)
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM tbzobuabpgyvwj;
GRANT ALL ON SCHEMA public TO tbzobuabpgyvwj;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- tbzobuabpgyvwj (heroku db)QL database dump complete
--