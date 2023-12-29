-- name: CreateEquipment :one
INSERT INTO equipment (
    model_number,
    description,
    location,
    parent,
    powered_from,
    controlled_from
    ) VALUES (
    $1, $2, $3, $4, $5, $6
    ) RETURNING *;
    
-- name: UpdateEquipment :exec
UPDATE equipment 
SET
    model_number = $1,
    description = $2,
    location = $3,
    parent = $4,
    powered_from = $5,
    controlled_from = $6
WHERE id = $7
RETURNING *;

-- name: GetEquipment :one
SELECT * FROM equipment WHERE id = $1 LIMIT 1;

-- name: ListEquipment :many
SELECT * FROM equipment
ORDER BY id
LIMIT $1
OFFSET $2;



-- name: DeleteEquipment :exec
DELETE FROM equipment 
WHERE id = $1;