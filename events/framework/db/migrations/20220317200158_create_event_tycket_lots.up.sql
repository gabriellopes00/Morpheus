CREATE TABLE IF NOT EXISTS "event_tycket_lots" (
    id UUID UNIQUE NOT NULL,
    event_tycket_option_id UUID NOT NULL,
    number INTEGER NOT NULL,
    tycket_price DECIMAL NOT NULL,
    tycket_amount INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (event_tycket_option_id) REFERENCES event_tycket_options(id) ON DELETE CASCADE
);