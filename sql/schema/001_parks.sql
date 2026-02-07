-- +goose Up
CREATE TABLE parks(
  id UUID PRIMARY KEY,
  nps_id TEXT UNIQUE NOT NULL,
  full_name TEXT NOT NULL,
  park_code TEXT UNIQUE NOT NULL,
  states TEXT NOT NULL,
  description TEXT,
  designation TEXT,
  activities JSONB,
  topics JSONB,
  details JSONB,
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE parks;
