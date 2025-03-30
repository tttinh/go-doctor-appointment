-- name: InsertSlots :copyfrom
INSERT INTO
  slot (doctor_id, start_time)
VALUES
  ($1, $2);


-- name: UpdateSlotAvailability :exec
UPDATE slot
SET
  available = $2
WHERE
  id = $1;


-- name: FetchSlots :many
SELECT
  id,
  doctor_id,
  start_time,
  available
FROM
  slot;


-- name: FetchSlotByID :one
SELECT
  id,
  doctor_id,
  start_time,
  available
FROM
  slot
WHERE
  id = $1;


-- name: UpdateSlotAvailable :exec
UPDATE slot
SET
  available = $2
WHERE
  id = $1;