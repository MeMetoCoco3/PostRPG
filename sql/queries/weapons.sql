-- name: CreateNewWeapon :one
INSERT INTO weapons (id, damage, reach, name, description, role) 
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	$5
)
RETURNING *;

-- name: GetWeapon :one
SELECT * FROM weapons WHERE id = $1;


-- name: DeleteAllWeapons :exec
DELETE FROM weapons;

-- name: DeleteOneWeapon :exec
DELETE FROM weapons WHERE id = $1;
