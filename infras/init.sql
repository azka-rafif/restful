CREATE TABLE run (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  run_time TEXT NOT NULL,
  run_kilometers REAL NOT NULL,
  run_city TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  created_by TEXT NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_by TEXT NOT NULL,
  deleted_at TIMESTAMP DEFAULT NULL,
  deleted_by TEXT DEFAULT NULL
);