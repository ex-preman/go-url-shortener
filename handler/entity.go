package handler

type EncodePayload struct {
	URL string `json:"url"`
}

type EncodeResponse struct {
	ShortenURL string `json:"shorten_url"`
}
