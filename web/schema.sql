--
-- PostgreSQL database dump
--

\restrict JkYRcKWKvyYMxhc4TQ0fJsNdQE6x69IMqcaP9N5dvUfETfBhrT1rSoPFpxJ2yDH

-- Dumped from database version 17.6 (Debian 17.6-2.pgdg13+1)
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: blog; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA blog;


--
-- Name: current_user_id(); Type: FUNCTION; Schema: blog; Owner: -
--

CREATE FUNCTION blog.current_user_id() RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
    SELECT
        nullif (current_setting('app.current_user_id', TRUE), '')::uuid
$$;


--
-- Name: is_owner(uuid); Type: FUNCTION; Schema: blog; Owner: -
--

CREATE FUNCTION blog.is_owner(user_id uuid) RETURNS boolean
    LANGUAGE sql STABLE
    AS $$
    SELECT
        user_id = blog.current_user_id ()
$$;


--
-- Name: set_updated_at(); Type: FUNCTION; Schema: blog; Owner: -
--

CREATE FUNCTION blog.set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$;


SET default_table_access_method = heap;

--
-- Name: casbin_rule; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.casbin_rule (
    id character varying NOT NULL,
    p_type character varying(255),
    v0 character varying(255),
    v1 character varying(255),
    v2 character varying(255),
    v3 character varying(255),
    v4 character varying(255),
    v5 character varying(255)
);


--
-- Name: fleeting; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.fleeting (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    date text NOT NULL,
    tag_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    user_id uuid
);


--
-- Name: tags; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.tags (
    id integer NOT NULL,
    name text NOT NULL
);


--
-- Name: fleeting_note_type_id_seq; Type: SEQUENCE; Schema: blog; Owner: -
--

CREATE SEQUENCE blog.fleeting_note_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: fleeting_note_type_id_seq; Type: SEQUENCE OWNED BY; Schema: blog; Owner: -
--

ALTER SEQUENCE blog.fleeting_note_type_id_seq OWNED BY blog.tags.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: sessions; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.sessions (
    token text NOT NULL,
    data bytea NOT NULL,
    expiry timestamp with time zone NOT NULL
);


--
-- Name: user_sessions; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.user_sessions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    last_token text,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: blog; Owner: -
--

CREATE TABLE blog.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    auth0_sub text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: tags id; Type: DEFAULT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.tags ALTER COLUMN id SET DEFAULT nextval('blog.fleeting_note_type_id_seq'::regclass);


--
-- Name: casbin_rule casbin_rule_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.casbin_rule
    ADD CONSTRAINT casbin_rule_pkey PRIMARY KEY (id);


--
-- Name: tags fleeting_note_type_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.tags
    ADD CONSTRAINT fleeting_note_type_pkey PRIMARY KEY (id);


--
-- Name: fleeting fleeting_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.fleeting
    ADD CONSTRAINT fleeting_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);


--
-- Name: user_sessions user_sessions_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.user_sessions
    ADD CONSTRAINT user_sessions_pkey PRIMARY KEY (id);


--
-- Name: user_sessions user_sessions_user_id_key; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.user_sessions
    ADD CONSTRAINT user_sessions_user_id_key UNIQUE (user_id);


--
-- Name: users users_auth0_sub_key; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.users
    ADD CONSTRAINT users_auth0_sub_key UNIQUE (auth0_sub);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_fleeting_date; Type: INDEX; Schema: blog; Owner: -
--

CREATE INDEX idx_fleeting_date ON blog.fleeting USING btree (date);


--
-- Name: sessions_last_seen_idx; Type: INDEX; Schema: blog; Owner: -
--

CREATE INDEX sessions_last_seen_idx ON blog.user_sessions USING btree (updated_at);


--
-- Name: sessions_user_idx; Type: INDEX; Schema: blog; Owner: -
--

CREATE INDEX sessions_user_idx ON blog.user_sessions USING btree (user_id);


--
-- Name: user_sessions user_sessions_updated_at; Type: TRIGGER; Schema: blog; Owner: -
--

CREATE TRIGGER user_sessions_updated_at BEFORE UPDATE ON blog.user_sessions FOR EACH ROW EXECUTE FUNCTION blog.set_updated_at();


--
-- Name: users users_updated_at; Type: TRIGGER; Schema: blog; Owner: -
--

CREATE TRIGGER users_updated_at BEFORE UPDATE ON blog.users FOR EACH ROW EXECUTE FUNCTION blog.set_updated_at();


--
-- Name: fleeting fleeting_note_type_id_fkey; Type: FK CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.fleeting
    ADD CONSTRAINT fleeting_note_type_id_fkey FOREIGN KEY (tag_id) REFERENCES blog.tags(id);


--
-- Name: fleeting fleeting_user_id_fkey; Type: FK CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.fleeting
    ADD CONSTRAINT fleeting_user_id_fkey FOREIGN KEY (user_id) REFERENCES blog.users(id);


--
-- Name: user_sessions user_sessions_last_token_fkey; Type: FK CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.user_sessions
    ADD CONSTRAINT user_sessions_last_token_fkey FOREIGN KEY (last_token) REFERENCES blog.sessions(token) ON DELETE CASCADE;


--
-- Name: user_sessions user_sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: blog; Owner: -
--

ALTER TABLE ONLY blog.user_sessions
    ADD CONSTRAINT user_sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES blog.users(id);


--
-- Name: fleeting; Type: ROW SECURITY; Schema: blog; Owner: -
--

ALTER TABLE blog.fleeting ENABLE ROW LEVEL SECURITY;

--
-- Name: fleeting fleeting_owner; Type: POLICY; Schema: blog; Owner: -
--

CREATE POLICY fleeting_owner ON blog.fleeting USING (blog.is_owner(user_id)) WITH CHECK (blog.is_owner(user_id));


--
-- PostgreSQL database dump complete
--

\unrestrict JkYRcKWKvyYMxhc4TQ0fJsNdQE6x69IMqcaP9N5dvUfETfBhrT1rSoPFpxJ2yDH

