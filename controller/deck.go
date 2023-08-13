package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raulcbs/flipcard-api/helper"
	"github.com/raulcbs/flipcard-api/model"
)

type addDeck struct {
	name string
}

func AddDeck(context *gin.Context) {
	var input addDeck
	var nmUserDeck model.NmUserDeck
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deck := model.Deck{
		Name: input.name,
	}

	savedDeck, err := deck.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nmUserDeck.UserID = user.ID
	nmUserDeck.DeckID = savedDeck.ID

	_, err = deck.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedDeck})
}

func GetAllDecks(context *gin.Context) {
	var deck model.Deck

	decks, err := deck.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": decks})
}
