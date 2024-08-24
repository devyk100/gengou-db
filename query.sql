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

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE user_id = $1;

-- name: CreateFlashcardDeck :one
WITH new_deck AS (
INSERT INTO "FlashcardDeck" (title)
VALUES ($1)
    RETURNING id
    )
INSERT INTO "FlashcardDeckToEditors" (deck_id, user_id)
SELECT id, $2
FROM new_deck RETURNING *;



-- name: CreateFlashcard :one
INSERT INTO "Flashcard" (
                         front_side, rear_side, deck_id
) VALUES (
          $1,$2, $3
         ) RETURNING *;

-- name: UpdateFlashcardFrontSide :one
UPDATE "Flashcard"
SET front_side = $2
WHERE id = $1 RETURNING *;

-- name: UpdateFlashcardRearSide :one
UPDATE "Flashcard"
SET rear_side = $2
WHERE id = $1 RETURNING *;

-- name: UpdateFlashcardFrontAudio :one
UPDATE "Flashcard"
SET front_audio = $2
WHERE id = $1 RETURNING *;

-- name: UpdateFlashcardRearAudio :one
UPDATE "Flashcard"
SET rear_audio = $2
WHERE id = $1 RETURNING *;

-- name: UpdateFlashcardFrontImage :one
UPDATE "Flashcard"
SET front_image = $2
WHERE id = $1 RETURNING *;

-- name: UpdateFlashcardRearImage :one
UPDATE "Flashcard"
SET rear_image = $2
WHERE id = $1
RETURNING *;

-- name: GetAllFlashcards :many
SELECT *
FROM "Flashcard"
WHERE deck_id = $1
LIMIT $2
OFFSET $3;

-- name: GetAFlashcard :many
SELECT *
FROM "Flashcard"
WHERE deck_id = $1
LIMIT $2
OFFSET $3;

-- name: GetFlashcardDecks :many
SELECT *
FROM "FlashcardDeck"
WHERE id IN (
    SELECT deck_id
    FROM "FlashcardDeckToEditors"
    WHERE user_id = $1
)
ORDER BY id
    LIMIT $2
OFFSET $3;

-- name: CopyFlashcardDeck: one
INSERT INTO "FlashcardDeck" (
    title,
    max_review_limit_per_day,
    graduating_interval,
    learning_steps,
    new_cards_limit_per_day,
    easy_interval
)
SELECT
    title,
    max_review_limit_per_day,
    graduating_interval,
    learning_steps,
    new_cards_limit_per_day,
    easy_interval
FROM "FlashcardDeck"
WHERE id = $1
    RETURNING id AS new_deck_id;

-- name: CreateCopyFlashcardDecKMapping :one
INSERT INTO "FlashcardDeckToCopiers" (deck_id, user_id, copied_deck_id)
VALUES ($1, $3, $2);

-- name: CopyFlashcardsForDeck :many
INSERT INTO "Flashcard"
(front_side, rear_side, front_audio, rear_audio, front_image, rear_image, review_factor, review_interval,
 priority_num, unreviewed_priority_num, deck_id)
SELECT
    front_side, rear_side, front_audio, rear_audio, front_image, rear_image, review_factor, review_interval,
    priority_num, unreviewed_priority_num, $2 AS deck_id  --> new deck ID
FROM "Flashcard"
WHERE deck_id = $1;  --> old one


-- name: FlashcardReview :one
UPDATE "Flashcard"
SET review_factor = $2, review_interval = $3, priority_num  $4, unreviewed_priority_num = $5
WHERE id = $1 RETURNING *;