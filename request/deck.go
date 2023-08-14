package request

import "github.com/raulcbs/flipcard-api/model"

type CreateDeckRequest struct {
	Name  string       `json:"name"`
	Cards []model.Card `json:"cards"`
}
