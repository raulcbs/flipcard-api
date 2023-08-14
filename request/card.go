package request

type CreateCardRequest struct {
	Front  string `json:"front"`
	Back   string `json:"back"`
	DeckID uint   `json:"deck_id"`
}
