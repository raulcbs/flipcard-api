package response

type CardResponse struct {
	ID     int    `json:"id"`
	Front  string `json:"front"`
	Back   string `json:"back"`
	DeckID uint   `json:"deck_id"`
}
