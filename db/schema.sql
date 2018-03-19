-- TABLES
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255)
);

CREATE TABLE tasks (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id uuid,
  title varchar(255),
  note text,
  complete boolean,
  created_at timestamp
);

-- SEEDS
INSERT INTO users (name) VALUES ('Adam');
INSERT INTO users (name) VALUES ('Bran');
