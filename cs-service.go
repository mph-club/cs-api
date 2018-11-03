package main

import (
	"csportal-server/server"
	"fmt"
	"log"
	"os"
)

func main() {
	region := os.Getenv("AWS_REGION")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")

	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json", region, userPoolID)
	jwk := server.GetJWK(jwkURL)

	log.Println(jwk)

	server.CreateAndListen()
}
