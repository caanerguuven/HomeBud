package controllers

import (
	"encoding/json"
	modelOf "homebud/models"
	utils "homebud/utils"
	"net/http"
)

//GetMonths function
func GetMonths(w http.ResponseWriter, r *http.Request) {
	result := modelOf.GetMonths()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//GetYears function
func GetYears(w http.ResponseWriter, r *http.Request) {
	result := modelOf.GetYears()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}
