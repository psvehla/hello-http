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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// HelloApiController binds http requests to an api service and writes the service results to the http response
type HelloApiController struct {
	service      HelloApiServicer
	errorHandler ErrorHandler
}

// HelloApiOption for how the controller is set up.
type HelloApiOption func(*HelloApiController)

// WithHelloApiErrorHandler inject ErrorHandler into controller
func WithHelloApiErrorHandler(h ErrorHandler) HelloApiOption {
	return func(c *HelloApiController) {
		c.errorHandler = h
	}
}

// NewHelloApiController creates a default api controller
func NewHelloApiController(s HelloApiServicer, opts ...HelloApiOption) Router {
	controller := &HelloApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the HelloApiController
func (c *HelloApiController) Routes() Routes {
	return Routes{
		{
			"HelloHello",
			strings.ToUpper("Get"),
			"/hello/{name}",
			c.HelloHello,
		},
		{
			"HelloOpenapiJson",
			strings.ToUpper("Get"),
			"/openapi.json",
			c.HelloOpenapiJson,
		},
	}
}

// HelloHello - hello hello
func (c *HelloApiController) HelloHello(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nameParam := params["name"]

	result, err := c.service.HelloHello(r.Context(), nameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// HelloOpenapiJson - Download ./gen/http/openapi.json
func (c *HelloApiController) HelloOpenapiJson(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.HelloOpenapiJson(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
