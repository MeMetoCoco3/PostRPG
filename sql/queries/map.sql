-- name: GetCharacterPosition :one
SELECT x, y, terrain FROM map WHERE map.character = $1;

-- name: CheckPosition :one
SELECT terrain, character FROM map WHERE map.x = $1 AND map.y = $2;

-- name: MoveCharacter :exec
UPDATE map SET character = $1 WHERE map.x = $2 AND map.y = $3;

-- name: SetNull :exec
UPDATE map SET character = NULL WHERE map.x = $1 AND map.y = $2;

-- name: SetAllNull :exec
UPDATE map SET character = NULL;
