// pkg/post/post.go
package post

import "time"

type Post struct {
	ID          uint      `json:"id,omitempty"`
	Recipe_name string    `json:"recipe_name,omitempty"`
	Recipe_type string    `json:"recipe_type,omitempty"`
	Thumbnail   string    `json:"thumbnail,omitempty"`
	Ingredients []string  `json:"ingredients,omitempty"`
	Description []string  `json:"description,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	Likes       uint      `json:"likes,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
