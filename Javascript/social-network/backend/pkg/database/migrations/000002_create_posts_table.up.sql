CREATE TABLE IF NOT EXISTS posts (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id STRING NOT NULL,
  author TEXT, 
  title TEXT,
  text TEXT,
  image TEXT,
  privacy TEXT,
  followers TEXT,
  creation_date TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(uuid) ON DELETE CASCADE
);