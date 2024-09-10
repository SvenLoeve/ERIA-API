package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"net/http"
	"viseh-api/types"
)

func RegisterNfcEndpoints() {

	api.Get("/nfc/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Get nfc data by uid.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.ChipEncryptedV2]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Post("/nfc/{uid}").
		HasPathParameter("uid", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Get user by uid.").
		HasRequestModel(rest.ModelOf[types.Chip]()).
		HasResponseModel(http.StatusCreated, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Get("/med/nfc/{med_id}").
		HasPathParameter("med_id", rest.PathParam{
			Description: "medical id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Get nfc data by medical id.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.Chip]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())
}
