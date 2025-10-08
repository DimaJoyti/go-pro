module github.com/DimaJoyti/go-pro/services/api-gateway

go 1.22

require (
	github.com/DimaJoyti/go-pro/services/shared v0.0.0
	github.com/golang-jwt/jwt/v5 v5.3.0
)

require github.com/google/uuid v1.5.0 // indirect

replace github.com/DimaJoyti/go-pro/services/shared => ../shared
