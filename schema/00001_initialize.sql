-------------------------------------------------------------------
------------------------------ USERS ------------------------------
-------------------------------------------------------------------
DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
    "id" BIGINT PRIMARY KEY,
    "first_name" TEXT NOT NULL,
    "last_name" TEXT,
    "birth_date" DATE,
    "gender" TEXT,
    "gender_interest" TEXT,
    "phone_number" VARCHAR(15) UNIQUE,
    "relationship_status" TEXT,
    "relationship_goal" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX ON "users" ("phone_number");
CREATE INDEX ON "users" ("deleted_at");
CREATE INDEX ON "users" ("created_at");



-------------------------------------------------------------------
--------------------------- CHANNELS ------------------------------
-------------------------------------------------------------------
DROP TABLE IF EXISTS channels CASCADE;

CREATE TABLE channels (
    "id" BIGINT PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX ON "channels" ("user_id", "created_at" ASC );



-------------------------------------------------------------------
--------------------------- MESSAGES ------------------------------
-------------------------------------------------------------------
DROP TABLE IF EXISTS messages CASCADE;

CREATE TABLE messages (
    "id" BIGINT PRIMARY KEY,
    "channel_id" BIGINT NOT NULL,
    "sender_type" TEXT NOT NULL, -- user, assistant
    "sender_id" BIGINT NOT NULL,
    "content_type" TEXT NOT NULL, -- text, image
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX ON "messages" ("channel_id", "created_at" DESC);
CREATE INDEX ON "messages" ("channel_id", "content_type", created_at DESC);
CREATE INDEX ON "messages" ("channel_id", "sender_id", created_at ASC);



-------------------------------------------------------------------
----------------------- MESSAGE INPUTS ----------------------------
-------------------------------------------------------------------
DROP TABLE IF EXISTS message_inputs CASCADE;

CREATE TABLE message_inputs (
    "id" BIGINT PRIMARY KEY,
    "channel_id" BIGINT NOT NULL,
    "source" TEXT NOT NULL, -- imessage, whatsapp, image
    "sender" TEXT NOT NULL,
    "receiver" TEXT NOT NULL,
    "receiver_pronoun" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX ON "message_inputs" ("channel_id");
CREATE INDEX ON "message_inputs" ("channel_id", created_at DESC );



-------------------------------------------------------------------
---------------------- MESSAGE SOURCES ----------------------------
-------------------------------------------------------------------
DROP TABLE IF EXISTS message_sources CASCADE;

CREATE TABLE message_sources (
    "id" BIGINT PRIMARY KEY,
    "message_input_id" BIGINT NOT NULL,
    "sender" TEXT NOT NULL,
    "content_type" TEXT NOT NULL,
    "content" TEXT NOT NULL,
    "sent_at" TIMESTAMPTZ NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX ON "message_sources" ("message_input_id", "content_type", "content", "sent_at");
CREATE INDEX ON "message_sources" ("message_input_id", "sent_at" ASC);
CREATE INDEX ON "message_sources" ("message_input_id", sent_at DESC );