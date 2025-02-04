package note_model

import (
	"database/sql"
	"fmt"
)

func AddNote(db *sql.DB, userId int64, text string) error {
	query := "INSERT INTO notes (user_id, text) VALUES (?, ?)"
	_, err := db.Exec(query, userId, text)
	if err != nil {
		return fmt.Errorf("Error adding note: %v", err)
	}
	return nil
}
