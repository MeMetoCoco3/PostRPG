-- +goose Up
CREATE TABLE skills (
    id UUID PRIMARY KEY,
    coin TEXT  NOT NULL,
    amount_to_pay INT  NOT NULL,
    damage INT NOT NULL,
    role INT NOT NULL,
    reach INT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    description  TEXT  NOT NULL
);


-- +goose Down
DROP TABLE skills;



