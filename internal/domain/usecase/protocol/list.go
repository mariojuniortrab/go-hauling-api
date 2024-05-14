package protocol_usecase

type List struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	OrderBy   string `json:"orderBy"`
	OrderType string `json:"orderType"`
	Q         string `json:"q"`
}
