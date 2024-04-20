CREATE TABLE IF NOT EXISTS notifications (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  sender_uuid TEXT,
  reciver_uuid TEXT,
  notification_text TEXT,
  notification_type TEXT,
  group_id INTEGER,
  timestamp DATETIME,
  is_read BOOLEAN DEFAULT FALSE
);