-- name: CreateNewCharacter :one
INSERT INTO characters (id, health, mana, stamina, strength, job, name, skill_id, weapon_id, icon) 
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9
)
RETURNING *;

-- name: GetCharacter :one
SELECT * FROM CHARACTERS WHERE id = $1;

-- name: DeleteAllCharacters :exec
DELETE FROM characters;

-- name: AssignWeapon :exec
UPDATE characters SET weapon_id = $1 WHERE characters.id = $2;

-- name: AssignSkill :exec
UPDATE characters SET skill_id = $1 WHERE characters.id = $2;

-- name: DeleteOneCharacter :exec
DELETE FROM characters WHERE id = $1;
