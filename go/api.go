/*
 * Hello Service
 *
 * A hello world service.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// HelloApiRouter defines the required methods for binding the api requests to a responses for the HelloApi
// The HelloApiRouter implementation should parse necessary information from the http request,
// pass the data to a HelloApiServicer to perform the required actions, then write the service results to the http response.
type HelloApiRouter interface { 
	HelloHello(http.ResponseWriter, *http.Request)
	HelloOpenapiJson(http.ResponseWriter, *http.Request)
}


// HelloApiServicer defines the api actions for the HelloApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type HelloApiServicer interface { 
	HelloHello(context.Context, string) (ImplResponse, error)
	HelloOpenapiJson(context.Context) (ImplResponse, error)
}
