package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"net/http"
	"viseh-api/types"
)

func RegisterMedEndpoints() {
	api.Get("/med/users/{id}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Get medical user by uid.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.MedUser]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Get("/med/users").
		HasQueryParameter("page", rest.QueryParam{
			Description: "current page index",
			Type:        "integer",
		}).
		HasQueryParameter("items", rest.QueryParam{
			Description: "amount of items",
			Type:        "integer",
		}).
		HasDescription("Get list of medical users.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.MedUsers]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Put("/med/users/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Update medical user.").
		HasRequestModel(rest.ModelOf[types.MedUser]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Get("/med/users/search").
		HasQueryParameter("email", rest.QueryParam{
			Description: "user email",
			Type:        "string",
		}).
		HasDescription("Search for medical user.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.MedUsers]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Patch("/med/users/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Patch medical user.").
		HasRequestModel(rest.ModelOf[patchMedUser]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Post("/med/register").
		HasDescription("Register.").
		HasRequestModel(rest.ModelOf[types.MedUser]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Post("/med/login").
		HasDescription("Login.").
		HasRequestModel(rest.ModelOf[types.UserLogin]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())
}
