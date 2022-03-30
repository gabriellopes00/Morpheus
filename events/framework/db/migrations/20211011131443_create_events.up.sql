DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'EVENT_STATUS') THEN
        CREATE TYPE EVENT_STATUS AS ENUM ('available', 'finished', 'canceled', 'sold_out');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'EVENT_AGE_GROUP') THEN
        CREATE TYPE EVENT_AGE_GROUP AS ENUM ('0', '10', '12', '14', '16', '18');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'EVENT_VISIBILITY') THEN
        CREATE TYPE EVENT_VISIBILITY AS ENUM ('private', 'public', 'invited_only');
    END IF;
END$$;


CREATE TABLE IF NOT EXISTS "events" (
    id UUID UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    description TEXT DEFAULT NULL,
    cover_url VARCHAR DEFAULT NULL,
    organizer_account_id UUID NOT NULL,
    age_group EVENT_AGE_GROUP NOT NULL,
    status EVENT_STATUS NOT NULL,
    start_datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    end_datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    -- category_id UUID DEFAULT NULL,
    -- subject_id UUID DEFAULT NULL,
    visibility EVENT_VISIBILITY NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (organizer_account_id) REFERENCES accounts(id) ON DELETE SET NULL
    -- FOREIGN KEY (subject_id) REFERENCES subjects(id) ON DELETE SET NULL,
    -- FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
);

CREATE INDEX search_events_stts ON events (status);
CREATE INDEX search_events_vsbt ON events (visibility);