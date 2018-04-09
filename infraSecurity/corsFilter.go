package infraSecurity

import "github.com/rs/cors"

func CorsFilter() *cors.Cors {
	corsFilter := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods : []string{"POST", "GET", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders : []string{"Access-Control-Allow-Headers", "accept, origin, X-Requested-With, Content-Type, X-Codingpedia, Authorization"},
		MaxAge : 172800,
	})

	return corsFilter
}