-- +goose Up
CREATE TYPE terrain_type AS ENUM ('WALL', 'GRASS', 'WATER');
CREATE TABLE map (
    x int NOT NULL,
    y int NOT NULL,
    terrain terrain_type,
    character TEXT,
    CONSTRAINT map_pk PRIMARY KEY(x,y)
);

-- +goose Down
DROP TABLE map;
DROP TYPE terrain_type;
