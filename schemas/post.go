package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Description string
	// Categories  []string
	Role    string
	Company string
	Link    string
	Public  bool
}

type PostResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeleteAt    time.Time `json:"deleteAt,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Link        string    `json:"link"`
	Public      bool      `json:"public"`
}
