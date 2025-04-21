package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connectin pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))
	`

	// Use the Exec() method on the embedded connection pool to execute the statement.
	// This method returns a sql.Result type, which contains some basic
	// information about happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method on the result toget the ID of newly inserted in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has type int64, so we convert it on an int type before returning
	return int(id), nil
}

// Return a specific snippet base on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `
		SELECT id, title, content, created, ecpires FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND id = ?
	`

	// This returns a pointer to a sql.Row object which holds the result from the database.
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroes Snippet struct.
	s := &Snippet{}

	// Use row.Scan() to copy the values from each field in sql.Row to the corresponding field in the Snippet struct.
	// Notice that the arguments to row.Scan() are *pointers* to the place you want to copy the data into,
	// and the number or arguments must be exactly the same as the number of column returned by your statement.
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// If the query returns no rows, than row.Scan() will be return a sql.ErrorNoRows.
		// We use the errors.Is() function check for that error specifivally, and return own ErrorNoRecord error instead
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// If everything went OK the return the Snippet object
	return s, nil
}

// Return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
