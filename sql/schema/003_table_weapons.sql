-- +goose Up

CREATE TABLE weapons(
    id UUID PRIMARY KEY,
    name TEXT,
    description TEXT,
    damage INT,
    reach INT,
    role INT
);

-- +goose Down
DROP TABLE weapons;


