CREATE TABLE IF NOT EXISTS comments (
  id INTEGER NOT NULL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  post_id INTEGER NOT NULL,
  title TEXT,
  text TEXT,
  image Text,
  creation_date TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(uuid) ON DELETE CASCADE,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
);