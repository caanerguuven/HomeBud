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

//CreateIncomingGroupDefinition function
func CreateIncomingGroupDefinition(w http.ResponseWriter, r *http.Request) {
	_indg := &modelOf.IncomingGroupDefinition{}
	utils.ParseBody(r, _indg)
	result := _indg.CreateIncomingGroupDefinition()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteIncomingGroupDefinition function
func DeleteIncomingGroupDefinition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incomingGroupDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(incomingGroupDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	incomingGroupDefinition := modelOf.DeleteIncomingGroupDefinition(id)
	response, _ := json.Marshal(incomingGroupDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllIncomingGroupDefinitions function
func GetAllIncomingGroupDefinitions(w http.ResponseWriter, r *http.Request) {
	incomingGroupDefinitions := modelOf.GetAllIncomingGroupDefinitions()
	response, _ := json.Marshal(incomingGroupDefinitions)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetIncomingGroupDefinitionByID function
func GetIncomingGroupDefinitionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incomingGroupDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(incomingGroupDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	incomingGroupDefinition, _ := modelOf.GetIncomingGroupDefinitionByID(id)
	response, _ := json.Marshal(&incomingGroupDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateIncomingGroupDefinition function
func UpdateIncomingGroupDefinition(w http.ResponseWriter, r *http.Request) {
	var incomingGroupDefinitionDto = &modelOf.IncomingGroupDefinition{}
	utils.ParseBody(r, incomingGroupDefinitionDto)
	vars := mux.Vars(r)
	incomingGroupDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(incomingGroupDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	incomingGroupDefinition, db := modelOf.GetIncomingGroupDefinitionByID(id)
	if incomingGroupDefinitionDto.Name != "" {
		incomingGroupDefinition.Name = incomingGroupDefinitionDto.Name
	}

	if incomingGroupDefinitionDto.Active != nil {
		incomingGroupDefinition.Active = incomingGroupDefinitionDto.Active
	}

	db.Save(&incomingGroupDefinition)
	response, _ := json.Marshal(&incomingGroupDefinition)

	utils.WriteResponse(w, http.StatusOK, response)
}
