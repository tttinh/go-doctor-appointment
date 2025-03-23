BEGIN;


CREATE TABLE IF NOT EXISTS doctor (
  id serial,
  uname TEXT NOT NULL,
  upassword TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS patient (
  id serial,
  uname TEXT NOT NULL,
  upassword TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS slot (
  id BIGSERIAL,
  doctor_id INTEGER,
  from_time TIMESTAMP,
  to_time TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  available BOOLEAN DEFAULT TRUE,
  PRIMARY KEY (id),
  FOREIGN key (doctor_id) REFERENCES doctor (id)
);


CREATE TABLE IF NOT EXISTS appointment (
  id BIGINT,
  patient_id INTEGER,
  note TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id),
  FOREIGN key (id) REFERENCES slot (id),
  FOREIGN key (patient_id) REFERENCES patient (id)
);


COMMIT;