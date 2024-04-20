CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  uuid TEXT UNIQUE,
  username TEXT NOT NULL,
  firstname TEXT NOT NULL,
  lastname TEXT,
  sex TEXT,
  birthday TEXT,
  email TEXT,
  public BOOLEAN DEFAULT 0,
  nickname TEXT DEFAULT "",
  aboutme TEXT DEFAULT "",
  image TEXT DEFAULT "defaultImage.png",
  password TEXT
);