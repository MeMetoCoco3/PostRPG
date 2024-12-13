-- name: CreateNewSkill :one
INSERT INTO skills (id, damage, reach, coin, amount_to_pay, name, description, role) 
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
)
RETURNING *;

-- name: GetSkill :one
SELECT * FROM skills WHERE id = $1;

-- name: DeleteAllSkills :exec
DELETE FROM skills;

-- name: DeleteOneSkill :exec
DELETE FROM skills WHERE id = $1;
