CREATE TABLE IF NOT EXISTS users (
    id         text,
    balance    NUMERIC NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);


CREATE TYPE transaction_type AS ENUM (
    'credit',
    'debit'
);

CREATE TABLE IF NOT EXISTS transactions (
    id          text NOT NULL,
    user_id     text NOT NULL,
    amount      NUMERIC NOT NULL,
    type        transaction_type NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT transactions_pkey PRIMARY KEY (id)
);
