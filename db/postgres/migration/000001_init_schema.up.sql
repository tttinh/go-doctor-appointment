BEGIN;


CREATE TABLE IF NOT EXISTS doctor (
  id SERIAL,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  hashed_password TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id),
  UNIQUE (username),
  UNIQUE (email)
);


CREATE TABLE IF NOT EXISTS patient (
  id SERIAL,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  hashed_password TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id),
  UNIQUE (username),
  UNIQUE (email)
);


CREATE TABLE IF NOT EXISTS slot (
  id SERIAL,
  doctor_id INTEGER,
  start_time TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  available BOOLEAN DEFAULT TRUE,
  PRIMARY KEY (id),
  FOREIGN KEY (doctor_id) REFERENCES doctor (id),
  UNIQUE (doctor_id, start_time)
);


CREATE TABLE IF NOT EXISTS appointment (
  id INTEGER,
  patient_id INTEGER,
  note TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id),
  FOREIGN key (id) REFERENCES slot (id),
  FOREIGN key (patient_id) REFERENCES patient (id)
);


COMMIT;