/*
 * Hello Service
 *
 * A hello world service.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	HelloApiService := openapi.NewHelloApiService()
	HelloApiController := openapi.NewHelloApiController(HelloApiService)

	router := openapi.NewRouter(HelloApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}