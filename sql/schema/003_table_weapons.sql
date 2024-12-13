-- +goose Up

CREATE TABLE weapons(
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL,
    damage INT NOT NULL,
    reach INT NOT NULL,
    role INT NOT NULL
);

-- +goose Down
DROP TABLE weapons;


