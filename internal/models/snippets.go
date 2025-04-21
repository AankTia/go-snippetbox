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
		SELECT id, title, content, created, expires FROM snippets
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
	stmt := `
		SELECT id, title, content, created, expires from snippets
		WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10
	`

	// Use the Query() method on the connection pool to execute our SQL statement.
	// This returns a sql.Rows() resultset containing the result of our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is always properly closed before the Latest() method returns.
	// This defer statement should come *affter* you check for an error from the Query() method.
	// Otherwise, of Query() returns an error, you'll get a panic trying to close a nil resultset.
	defer rows.Close()

	// Initializing an empty slice to holed the Snippet structs.
	snippets := []*Snippet{}

	// Use rows.Next to iterate through the rows in resultset.
	// This prepare the first (and then each subsequent) row to be acted on by the row.Scan() method.
	// If iteration over all the rows completes then the resultset automatically closes itself
	// and free-up the underlying database connection.
	for rows.Next() {
		// Create a pointer to a new zeroed Snippet struct.
		s := &Snippet{}

		// Use rows.Scan() to copy the values from each field in the row to the new Snippet object.
		// The arguments to row.Scan() must be pointers to the place tou want to copy the data into,
		// and the number of arguments must be exactly the same as the number of columns returned by your statement.
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		// Append it to the slice snippets.
		snippets = append(snippets, s)
	}

	// When the rows.Next loop has finished we call rows.Err() to retrieve any error that was encountered during the iteration.
	// It's important to call diss - don't assume that a successfull iteration was completed over the whole resultset.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippets slice.
	return snippets, nil
}
