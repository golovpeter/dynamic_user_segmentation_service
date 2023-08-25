CREATE TABLE segments
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    slug       VARCHAR(256) UNIQUE      NOT NULL,
    deleted    BOOLEAN                  NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE users_to_segments
(
    user_id    BIGINT NOT NULL,
    segment_id BIGINT NOT NULL REFERENCES segments (id),
    CONSTRAINT unique_user_segment UNIQUE (user_id, segment_id)
);

CREATE TYPE operation AS ENUM ('create', 'delete');

CREATE TABLE users_to_segments_history
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id    BIGINT                   NOT NULL,
    segment_id BIGINT                   NOT NULL REFERENCES segments (id),
    operation  operation                NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- TODO: проверить - два разных индекса лучше или составной
CREATE INDEX users_to_segments_history_user_id_created_at_idx
    ON users_to_segments_history (user_id, created_at);


