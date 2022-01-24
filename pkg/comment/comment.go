// pkg/post/post.go
package comment

import "time"

type Comment struct {
	ID        uint      `json:"id,omitempty"`
	Recipe_id uint      `json:"recipe_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
