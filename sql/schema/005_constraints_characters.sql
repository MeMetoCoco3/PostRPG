-- +goose Up
ALTER TABLE characters ADD CONSTRAINT fk_characters_weapon 
FOREIGN KEY(weapon_id) REFERENCES weapons(id);

ALTER TABLE characters ADD CONSTRAINT fk_characters_skill
FOREIGN KEY(skill_id) REFERENCES skills(id);


-- +goose Down
ALTER TABLE characters
DROP CONSTRAINT fk_characters_weapon;

ALTER TABLE characters 
DROP CONSTRAINT fk_characters_skill;
