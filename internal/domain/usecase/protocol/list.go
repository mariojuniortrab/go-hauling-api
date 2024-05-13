package protocol_usecase

type List struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Q     string `json:"q"`
}
