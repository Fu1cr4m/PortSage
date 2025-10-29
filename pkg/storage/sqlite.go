    package storage

    import (
        "database/sql"
        _ "github.com/mattn/go-sqlite3"
    )

    const schema = `
CREATE TABLE IF NOT EXISTS audit (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ts TEXT,
  type TEXT,
  instance_id TEXT,
  class TEXT,
  location TEXT,
  action TEXT
);`

    type DB struct{ *sql.DB }

    func Open(path string) (*DB, error) {
        db, err := sql.Open("sqlite3", path)
        if err != nil { return nil, err }
        if _, err := db.Exec(schema); err != nil { return nil, err }
        return &DB{db}, nil
    }
