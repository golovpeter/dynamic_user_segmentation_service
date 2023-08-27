CREATE TABLE IF NOT EXISTS segments
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    slug       VARCHAR(256) UNIQUE      NOT NULL,
    deleted    BOOLEAN                  NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS segments_search_by_slug_and_deleted_idx
    ON segments (slug, deleted)
    WHERE deleted = false;

CREATE TABLE IF NOT EXISTS users_to_segments
(
    user_id    BIGINT NOT NULL,
    segment_id BIGINT NOT NULL REFERENCES segments (id),
    expired_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT unique_user_segment UNIQUE (user_id, segment_id)
);

CREATE INDEX IF NOT EXISTS users_to_segments_expired_at_idx
    ON users_to_segments (expired_at) WHERE expired_at is not null;

CREATE INDEX IF NOT EXISTS users_to_segments_segment_id_idx
    ON users_to_segments (segment_id);

DROP TYPE IF EXISTS operation;
CREATE TYPE operation AS ENUM ('create', 'delete');

CREATE TABLE IF NOT EXISTS users_to_segments_history
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id    BIGINT                   NOT NULL,
    segment_id BIGINT                   NOT NULL REFERENCES segments (id),
    operation  operation                NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS users_to_segments_history_created_at_idx
    ON users_to_segments_history (created_at);

CREATE OR REPLACE FUNCTION users_to_segments_history_trg()
    RETURNS trigger
    LANGUAGE 'plpgsql'
AS $BODY$
BEGIN
    IF TG_OP = 'INSERT'
    THEN
        INSERT INTO users_to_segments_history (user_id, segment_id, operation) VALUES (NEW.user_id, NEW.segment_id, 'create');
    ELSIF TG_OP = 'DELETE'
    THEN
        INSERT INTO users_to_segments_history (user_id, segment_id, operation) VALUES (OLD.user_id, OLD.segment_id, 'delete');
    END IF;
    RETURN NEW;
END;
$BODY$;

CREATE OR REPLACE TRIGGER users_to_segments_history_created_at_idx
    AFTER INSERT OR DELETE
    ON users_to_segments
    FOR EACH ROW
    EXECUTE PROCEDURE users_to_segments_history_trg();