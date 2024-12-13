-- +goose Up
CREATE TABLE characters (
    id UUID PRIMARY KEY,
    health INT NOT NULL,
    mana INT  NOT NULL,
    stamina INT NOT NULL,
    strength INT NOT NULL,
    job INT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    skill_id UUID,
    weapon_id UUID,
    icon TEXT NOT NULL
);

-- +goose Down
DROP TABLE characters;


