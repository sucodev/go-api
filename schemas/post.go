package schemas

import (
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
