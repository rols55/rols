CREATE TABLE IF NOT EXISTS group_invites (
  user_id INTEGER NOT NULL,
  group_id INTEGER NOT NULL,
  state TEXT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, group_id)
);