--
-- PostgreSQL database dump
--

\restrict XJHFgvgEcENKPjIZT0uCT4kl1GgOfRyLVkQ23ya84Olpagxw7qULvgeyRfjL6wL

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
-- Name: notes; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA notes;


--
-- Name: set_updated_at(); Type: FUNCTION; Schema: notes; Owner: -
--

CREATE FUNCTION notes.set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$;


SET default_table_access_method = heap;

--
-- Name: casbin_rule; Type: TABLE; Schema: notes; Owner: -
--

CREATE TABLE notes.casbin_rule (
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
-- Name: schema_migrations; Type: TABLE; Schema: notes; Owner: -
--

CREATE TABLE notes.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: sessions; Type: TABLE; Schema: notes; Owner: -
--

CREATE TABLE notes.sessions (
    token text NOT NULL,
    data bytea NOT NULL,
    expiry timestamp with time zone NOT NULL
);


--
-- Name: user_sessions; Type: TABLE; Schema: notes; Owner: -
--

CREATE TABLE notes.user_sessions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    last_token text,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: notes; Owner: -
--

CREATE TABLE notes.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    auth0_sub text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: casbin_rule casbin_rule_pkey; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.casbin_rule
    ADD CONSTRAINT casbin_rule_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);


--
-- Name: user_sessions user_sessions_pkey; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.user_sessions
    ADD CONSTRAINT user_sessions_pkey PRIMARY KEY (id);


--
-- Name: user_sessions user_sessions_user_id_key; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.user_sessions
    ADD CONSTRAINT user_sessions_user_id_key UNIQUE (user_id);


--
-- Name: users users_auth0_sub_key; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.users
    ADD CONSTRAINT users_auth0_sub_key UNIQUE (auth0_sub);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: sessions_last_seen_idx; Type: INDEX; Schema: notes; Owner: -
--

CREATE INDEX sessions_last_seen_idx ON notes.user_sessions USING btree (updated_at);


--
-- Name: sessions_user_idx; Type: INDEX; Schema: notes; Owner: -
--

CREATE INDEX sessions_user_idx ON notes.user_sessions USING btree (user_id);


--
-- Name: user_sessions user_sessions_updated_at; Type: TRIGGER; Schema: notes; Owner: -
--

CREATE TRIGGER user_sessions_updated_at BEFORE UPDATE ON notes.user_sessions FOR EACH ROW EXECUTE FUNCTION notes.set_updated_at();


--
-- Name: users users_updated_at; Type: TRIGGER; Schema: notes; Owner: -
--

CREATE TRIGGER users_updated_at BEFORE UPDATE ON notes.users FOR EACH ROW EXECUTE FUNCTION notes.set_updated_at();


--
-- Name: user_sessions user_sessions_last_token_fkey; Type: FK CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.user_sessions
    ADD CONSTRAINT user_sessions_last_token_fkey FOREIGN KEY (last_token) REFERENCES notes.sessions(token) ON DELETE CASCADE;


--
-- Name: user_sessions user_sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: notes; Owner: -
--

ALTER TABLE ONLY notes.user_sessions
    ADD CONSTRAINT user_sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES notes.users(id);


--
-- PostgreSQL database dump complete
--

\unrestrict XJHFgvgEcENKPjIZT0uCT4kl1GgOfRyLVkQ23ya84Olpagxw7qULvgeyRfjL6wL

