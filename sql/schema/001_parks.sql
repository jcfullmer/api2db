-- +goose Up
CREATE TABLE parks(
  id UUID PRIMARY KEY, 
  nps_id TEXT UNIQUE NOT NULL,
  full_name TEXT NOT NULL,
  park_code TEXT UNIQUE NOT NULL,
  states TEXT NOT NULL,
  description TEXT NOT NULL,
  designation TEXT NOT NULL,
  activities JSONB NOT NULL,
  topics JSONB NOT NULL,
  details JSONB NOT NULL,
  created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE parks;
