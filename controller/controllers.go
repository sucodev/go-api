package controller

import (
	"github.com/sucodev/go-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func IntializeControllers() {
	logger = config.GetLogger("controller")
	db = config.GetSqlite()
}
