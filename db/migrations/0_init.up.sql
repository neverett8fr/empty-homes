CREATE TABLE IF NOT EXISTS properties(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    postcode VARCHAR NOT NULL,
    street VARCHAR,
    type VARCHAR,
    bedrooms int,
    inhabited BOOLEAN,
    safe BOOLEAN,
    owner_name VARCHAR,
    owner_contact VARCHAR,
    date_added timestamp,
    date_last_checked timestamp
);
