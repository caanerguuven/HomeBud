package controllers

import (
	"encoding/json"
	"fmt"
	modelOf "homebud/models"
	utils "homebud/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateIncomingDefinition function
func CreateIncomingDefinition(w http.ResponseWriter, r *http.Request) {
	_ind := &modelOf.IncomingDefinition{}
	utils.ParseBody(r, _ind)
	result := _ind.CreateIncomingDefinition()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteIncomingDefinition function
func DeleteIncomingDefinition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incomingDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(incomingDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	incomingDefinition := modelOf.DeleteIncomingDefinition(id)
	response, _ := json.Marshal(incomingDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllIncomingDefinitions function
func GetAllIncomingDefinitions(w http.ResponseWriter, r *http.Request) {
	incomingDefinitions := modelOf.GetAllIncomingDefinitions()
	response, _ := json.Marshal(incomingDefinitions)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetIncomingDefinitionByID function
func GetIncomingDefinitionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incomingDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(incomingDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	incomingDefinition, _ := modelOf.GetIncomingDefinitionByID(id)
	response, _ := json.Marshal(&incomingDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateIncomingDefinition function
func UpdateIncomingDefinition(w http.ResponseWriter, r *http.Request) {
	var incomingDefinitionDto = &modelOf.IncomingDefinition{}
	utils.ParseBody(r, incomingDefinitionDto)
	vars := mux.Vars(r)
	incomingDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(incomingDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	incomingDefinition, db := modelOf.GetIncomingDefinitionByID(id)
	if incomingDefinitionDto.Name != "" {
		incomingDefinition.Name = incomingDefinitionDto.Name
	}

	if incomingDefinitionDto.IncomingDefinitionGroupID != nil {
		incomingDefinition.IncomingDefinitionGroupID = incomingDefinitionDto.IncomingDefinitionGroupID
	}

	if incomingDefinitionDto.Active != nil {
		incomingDefinition.Active = incomingDefinitionDto.Active
	}

	db.Save(&incomingDefinition)
	response, _ := json.Marshal(&incomingDefinition)

	utils.WriteResponse(w, http.StatusOK, response)
}
