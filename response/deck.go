package response

import "github.com/raulcbs/flipcard-api/model"

type DeckResponse struct {
	ID    uint         `json:"id"`
	Name  string       `json:"name"`
	Cards []model.Card `json:"cards"`
}
