CREATE TABLE IF NOT EXISTS followers (
  follower TEXT NOT NULL,
  followed TEXT NOT NULL,
  allowed BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (follower) REFERENCES users(uuid) ON DELETE CASCADE,
  FOREIGN KEY (followed) REFERENCES users(uuid) ON DELETE CASCADE,
  CONSTRAINT unique_followers UNIQUE (follower, followed)
);