package dtos

// Movie ID, Title, Description, Categories, Image, Actor, Current
type Movie struct {
	MovieID     string   `bson:"movie_id" json:"movie_id"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	Categories  []string `bson:"categories" json:"categories"`
	Image       string   `bson:"image" json:"image"`
	Actor       []string `bson:"actors" json:"actors"`
	CRR         float64  `bson:"current_recommended_rate" json:"current_recommended_rate"`
	// Rating string `json:"rating"`
}

// type ScrollDataCaptured struct {
// 	UserID           string    `json:"user_id"`
// 	Timestamp        time.Time `json:"timestamp"`
// 	DurationOfScroll int       `json:"duration_of_scroll"`
// 	PostID           string    `json:"post_id"`
// }

type PaginationSpecifics struct {
	Page  int `json:"page"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

// type ScrollDataShow struct {
// 	MovieID     string   `bson:"movie_id" json:"movie_id"`
// 	Title       string   `bson:"title" json:"title"`
// 	Description string   `bson:"description" json:"description"`
// 	Categories  []string `bson:"categories" json:"categories"`
// 	Image       string   `bson:"image" json:"image"`
// 	Actor       []string `bson:"actors" json:"actors"`
// 	CRR         float64  `bson:"current_recommended_rate" json:"current_recommended_rate"`
// }
