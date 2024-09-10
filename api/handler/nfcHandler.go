package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"viseh-api/database/query"
	"viseh-api/services/jwt"
	nfcService "viseh-api/services/nfc"
	"viseh-api/services/validate"
	"viseh-api/types"
)

// UpdateNFC update NFC chip data in the database
func UpdateNFC(w http.ResponseWriter, r *http.Request) {
	var nfc types.Chip
	var nfcEncrypted types.ChipEncryptedV2

	//Get id from request path
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleUser, uid) {
		return
	}
	//decode request body to user type
	err = json.NewDecoder(r.Body).Decode(&nfc)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !validate.NfcIsValid(nfc) {
		BadRequestHandler(w)
		return
	}

	user, err := query.GetUser(context.Background(), uid)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	nfc.Data.MEDID = user.MedID

	nfcEncrypted.Version = 2 //hardcoded version in case of later iterations of data format
	nfcEncrypted = nfcService.CreateEncryptedV2(nfc)

	nfcEncryptedString, err := json.Marshal(nfcEncrypted)

	//query database to update nfc data
	err = query.UpdateNfc(context.Background(), uid, nfcEncryptedString)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("nfc updated"))
		return
	}
}

// GetNfc get NFC chip data from the database
func GetNfc(w http.ResponseWriter, r *http.Request) {
	var nfcEncrypted types.ChipEncryptedV2

	//Get id from request path
	uid, err := strconv.Atoi(r.PathValue("uid"))
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}
	if !jwt.Handler(w, r, jwt.RoleUser, uid) {
		return
	}
	//query database for nfc data
	nfcUser, err := query.GetNfcWithId(context.Background(), uid)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	err = json.Unmarshal(nfcUser.Data, &nfcEncrypted)
	if err != nil {
		return
	}
	//Set content type
	w.Header().Set("Content-Type", "application/json")
	//Return user data to api request
	err = json.NewEncoder(w).Encode(nfcEncrypted)
	if err != nil {
		return
	}

}

// GetNfcWithMedID get NFC chip data from the database
func GetNfcWithMedID(w http.ResponseWriter, r *http.Request) {
	if !jwt.Handler(w, r, jwt.RoleAmbulance, 0) {
		return
	}

	var nfc types.Chip
	var nfcEncrypted types.ChipEncryptedV2

	//Get id from request path
	medID := r.PathValue("med_id")

	//query database for nfc data
	nfcEncryptedString, err := query.GetNfcWithMedID(context.Background(), medID)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	err = json.Unmarshal(nfcEncryptedString.Data, &nfcEncrypted)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	//Set content type
	w.Header().Set("Content-Type", "application/json")
	//decrypt data
	nfc, err = nfcService.CreateDecryptedV2(nfcEncrypted)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	//Return user data to api request
	err = json.NewEncoder(w).Encode(nfc)
	if err != nil {
		return
	}

}
