package data

import (
	"context"
	"time"

	"github.com/martinpaz/restfulapi/pkg/comment"
)

// CommentRepository manages the operations with the database that
// correspond to the user model.
type CommentRepository struct {
	Data *Data
}

// Create adds a new user.
func (cr *CommentRepository) Create(ctx context.Context, c *comment.Comment) error {
	q := `
	INSERT INTO comments (recipe_id, name, comment, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	stmt, err := cr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, c.Recipe_id, c.Name, c.Comment, time.Now())

	err = row.Scan(&c.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetOne returns one user by id.
func (cr *CommentRepository) GetOne(ctx context.Context, id uint) ([]comment.Comment, error) {
	q := `
	SELECT id, name, comment, created_at
		FROM comments WHERE recipe_id = $1;
	`

	/*row := cr.Data.DB.QueryRowContext(ctx, q, id)

	var c comment.Comment
	err := row.Scan(&c.ID, &c.Name, &c.Comment, &c.CreatedAt)
	if err != nil {
		return comment.Comment{}, err
	}*/

	rows, err := cr.Data.DB.QueryContext(ctx, q, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []comment.Comment
	for rows.Next() {
		var c comment.Comment
		rows.Scan(&c.ID, &c.Name, &c.Comment, &c.CreatedAt)
		comments = append(comments, c)
	}

	return comments, nil
}

// Delete removes a user by id.
func (cr *CommentRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM comments WHERE id=$1;`

	stmt, err := cr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
