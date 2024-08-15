-- name: GetAUser :one
SELECT * FROM "User"
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM "User"
LIMIT 5;

-- name: InsertInstructor :one
INSERT INTO "User" (
                    name, user_id, email_id, phone, past_experiences, user_type
) VALUES (
          $1, $2, $3, $4, $5, 'Instructor'
         ) RETURNING *;

-- name: InsertLearner :one
INSERT INTO "User" (
    name, user_id, email_id, phone, user_type
) VALUES (
             $1, $2, $3, $4, 'Learner'
         ) RETURNING *;
