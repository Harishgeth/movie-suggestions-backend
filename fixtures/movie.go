package fixtures

import "movie-suggestions-api/dtos"

func DeadpoolResult() []dtos.Movie {
	return []dtos.Movie{
		dtos.Movie{
			Name:   "Deadpool (2016)",
			Rating: "8.0 based on 699,445 user ratings",
		},
		dtos.Movie{
			Name:   "Deadpool (2013) (TV Series)",
			Rating: "5.6 based on 244 user ratings",
		},
	}
}

func JustLikeHeavenResult() []dtos.Movie {
	return []dtos.Movie{
		dtos.Movie{
			Name:   "Just (2013) (Short)",
			Rating: "3.1 based on 9 user ratings",
		},
		dtos.Movie{
			Name:   "Just (2007) (Short)",
			Rating: "5.6 based on 244 user ratings",
		},
	}
}
