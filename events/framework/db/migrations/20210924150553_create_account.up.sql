CREATE TABLE IF NOT EXISTS "events" (
  id VARCHAR(36) UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  is_available BOOLEAN NOT NULL,
  organizer_account_id VARCHAR(36) NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL

  PRIMARY KEY (id),
  FOREIGN KEY (organizer_account_id)
);
