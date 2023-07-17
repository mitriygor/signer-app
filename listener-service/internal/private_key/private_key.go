package private_key

type PrivateKey struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Secret string `json:"secret"`
}
