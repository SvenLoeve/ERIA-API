package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"net/http"
	"viseh-api/types"
)

func RegisterUserEndpoints() {
	api.Get("/users/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Get user by uid.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.User]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Get("/users").
		HasQueryParameter("page", rest.QueryParam{
			Description: "current page index",
			Type:        "integer",
		}).
		HasQueryParameter("items", rest.QueryParam{
			Description: "amount of items",
			Type:        "integer",
		}).
		HasDescription("Get list of users.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.Users]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Put("/users/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Update user.").
		HasRequestModel(rest.ModelOf[types.User]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Get("/users/search").
		HasQueryParameter("email", rest.QueryParam{
			Description: "user email",
			Type:        "string",
		}).
		HasDescription("Search for user.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.Users]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Patch("/users/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Patch user.").
		HasRequestModel(rest.ModelOf[patchUser]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Post("/register").
		HasDescription("Register.").
		HasRequestModel(rest.ModelOf[types.FullUser]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Post("/login").
		HasDescription("Login.").
		HasRequestModel(rest.ModelOf[types.UserLogin]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())
}
