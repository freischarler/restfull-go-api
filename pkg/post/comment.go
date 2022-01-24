// pkg/post/post.go
package post

import "time"

type Comment struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Comment   string    `json:"recipe_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
