package dtos

// Movie ID, Title, Description, Categories, Image, Actor, Current
type Movie struct {
	MovieID     string   `json:"movie_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
	Image       string   `json:"image"`
	Actor       []string `json:"actors"`
	CRR         float64  `json:"current_recommended_rate"`
	// Rating string `json:"rating"`
}
