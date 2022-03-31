CREATE TABLE IF NOT EXISTS "event_ticket_options_lots" (
    id UUID UNIQUE NOT NULL,
    event_ticket_option_id UUID NOT NULL,
    number INTEGER NOT NULL,
    price DECIMAL NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (event_ticket_option_id) REFERENCES event_ticket_options(id) ON DELETE CASCADE
);