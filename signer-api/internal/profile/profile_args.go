package profile

type Args struct {
	FirstName string `json:"first-name"url:"first-name"`
	LastName  string `json:"last-name"url:"last-name"`
	Limit     int    `json:"limit"url:"limit"`
}
