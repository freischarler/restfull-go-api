// internal/data/post_repository.go
package data

import (
	"context"
	"strings"
	"time"

	"github.com/martinpaz/restfulapi/pkg/post"
)

// PostRepository manages the operations with the database that
// correspond to the post model.
type PostRepository struct {
	Data *Data
}

// GetAll returns all posts.
func (pr *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	q := `
	SELECT recipe_name, recipe_type, user_id , ingredients, description, thumbnail, likes, created_at, updated_at
		FROM posts;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		var array_ingredients, array_description string

		rows.Scan(&p.Recipe_name, &p.Recipe_type, &p.UserID, &array_ingredients, &array_description, &p.Thumbnail, &p.Likes, &p.CreatedAt, &p.UpdatedAt)

		p.Ingredients = strings.Split(array_ingredients, "$$$")
		p.Description = strings.Split(array_description, "$$$")

		posts = append(posts, p)
	}

	return posts, nil
}

// GetOne returns one post by id.
func (pr *PostRepository) GetOne(ctx context.Context, id uint) (post.Post, error) {
	q := `
	SELECT id, recipe_name, recipe_type, user_id , ingredients, description, thumbnail, likes, created_at, updated_at
		FROM posts WHERE id = $1;
	`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p post.Post
	var array_ingredients, array_description string
	//var likes_count int

	err := row.Scan(&p.ID, &p.Recipe_name, &p.Recipe_type, &p.UserID, &array_ingredients, &array_description, &p.Thumbnail, &p.Likes, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return post.Post{}, err
	}

	/*fmt.Printf("count:%d\n", likes_count)
	if likes_count == 0 {
		p.Likes = 0
	}*/

	p.Ingredients = strings.Split(array_ingredients, "$$$")
	p.Description = strings.Split(array_description, "$$$")

	return p, nil
}

// GetByUser returns all user posts.
func (pr *PostRepository) GetByUser(ctx context.Context, userID uint) ([]post.Post, error) {
	q := `
	SELECT id, body, user_id, created_at, updated_at
		FROM posts
		WHERE user_id = $1;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.UserID, &p.Recipe_name, &p.Description, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

// GetByUser returns all user posts.
func (pr *PostRepository) GetByType(ctx context.Context, recipeTYPE string) ([]post.Post, error) {
	q := `
	SELECT recipe_name, recipe_type, user_id , ingredients, description, thumbnail, likes, created_at, updated_at
		FROM posts
		WHERE recipe_type = $1;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q, recipeTYPE)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		var array_ingredients, array_description string

		rows.Scan(&p.Recipe_name, &p.Recipe_type, &p.UserID, &array_ingredients, &array_description, &p.Thumbnail, &p.Likes, &p.CreatedAt, &p.UpdatedAt)

		p.Ingredients = strings.Split(array_ingredients, "$$$")
		p.Description = strings.Split(array_description, "$$$")

		posts = append(posts, p)
	}

	return posts, nil
}

// Create adds a new post.
func (pr *PostRepository) Create(ctx context.Context, p *post.Post) error {
	q := `
	INSERT INTO posts (recipe_name, recipe_type, user_id, ingredients, description, thumbnail, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	if p.Thumbnail == "" {
		p.Thumbnail = "https://food.unl.edu/newsletters/images/mise-en-plase.jpg"
	}

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Recipe_name, p.Recipe_type, p.UserID, strings.Join(p.Ingredients, "$$$"), strings.Join(p.Description, "$$$"), p.Thumbnail, time.Now(), time.Now())

	err = row.Scan(&p.UserID)
	if err != nil {
		return err
	}

	return nil
}

// Create adds a new like.
func (pr *PostRepository) UpdateLike(ctx context.Context, id uint) error {
	q := `
	UPDATE posts
		SET likes=likes+1
		WHERE id=$1;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
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

// Update updates a post by id.
func (pr *PostRepository) Update(ctx context.Context, id uint, p post.Post) error {
	q := `
	UPDATE posts set body=$1, updated_at=$2
		WHERE id=$3;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Description, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post by id.
func (pr *PostRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM posts WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
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

// GetRank returns all posts.
func (pr *PostRepository) GetRank(ctx context.Context) ([]post.Post, error) {
	q := `
	SELECT recipe_name, recipe_type, user_id , thumbnail, likes, created_at, updated_at
	FROM posts
	ORDER BY likes DESC
	LIMIT 10
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post

		rows.Scan(&p.Recipe_name, &p.Recipe_type, &p.UserID, &p.Thumbnail, &p.Likes, &p.CreatedAt, &p.UpdatedAt)

		posts = append(posts, p)
	}

	return posts, nil
}

/*SELECT recipe_name, recipe_type, likes, id
FROM posts
ORDER BY likes DESC
LIMIT 10*/
