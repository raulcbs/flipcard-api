package migration

import (
	"github.com/raulcbs/flipcard-api/database"
	"github.com/raulcbs/flipcard-api/model"
)

func Migration() {
	// * create tables and columns
	// ! the order it's important
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Deck{})
	database.Database.AutoMigrate(&model.Card{})
	database.Database.AutoMigrate(&model.Word{})
	database.Database.AutoMigrate(&model.NmUserDeck{})
}
