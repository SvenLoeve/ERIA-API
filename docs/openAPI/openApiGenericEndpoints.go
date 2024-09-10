package main

import (
	"github.com/a-h/rest"
	"net/http"
)

func RegisterGenericEndpoints() {

	api.Get("/mock/{count}").
		HasPathParameter("count", rest.PathParam{
			Description: "amount of users to mock",
		}).
		HasDescription("Generate mock users.").
		HasResponseModel(http.StatusNotFound, rest.ModelOf[string]())

	api.Get("/med/mock/{count}").
		HasPathParameter("count", rest.PathParam{
			Description: "amount of medical users to mock",
		}).
		HasDescription("Generate mock medical users.").
		HasResponseModel(http.StatusNotFound, rest.ModelOf[string]())

	api.Get("/privacy").
		HasDescription("Privacy policy.").
		HasResponseModel(http.StatusCreated, rest.ModelOf[string]())

	api.Get("/dataProtection").
		HasDescription("Data protection policy.").
		HasResponseModel(http.StatusNotFound, rest.ModelOf[string]())
}
