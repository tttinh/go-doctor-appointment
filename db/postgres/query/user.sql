-- name: InsertDoctor :one
INSERT INTO
  doctor (username, email, hashed_password)
VALUES
  ($1, $2, $3)
RETURNING
  *;


-- name: FetchDoctorByUsername :one
SELECT
  *
FROM
  doctor d
WHERE
  d.username = $1;


-- name: InsertPatient :one
INSERT INTO
  patient (username, email, hashed_password)
VALUES
  ($1, $2, $3)
RETURNING
  *;


-- name: FetchPatientByUsername :one
SELECT
  *
FROM
  patient p
WHERE
  p.username = $1;