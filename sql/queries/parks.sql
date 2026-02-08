-- name: CreateUser :one 
INSERT INTO parks (id, nps_id, full_name, park_code, states, description, designation, activities, topics, details, created_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11
)
RETURNING *;

-- name: CheckExists :one
SELECT EXISTS (
  SELECT  full_name
  FROM parks
  WHERE nps_id = $1
  LIMIT 1
);
