CREATE TABLE IF NOT EXISTS messages (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  sender_id INTEGER NOT NULL,
  reciver_id INTEGER NOT NULL,
  message_text TEXT,
  timestamp DATETIME,
  is_read BOOLEAN DEFAULT FALSE,
  group_id INTEGER
);