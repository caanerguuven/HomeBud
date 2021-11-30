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

//CreateExpenseDefinition function
func CreateExpenseDefinition(w http.ResponseWriter, r *http.Request) {
	_cur := &modelOf.ExpenseDefinition{}
	utils.ParseBody(r, _cur)
	result := _cur.CreateExpenseDefinition()
	response, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	utils.WriteResponse(w, http.StatusOK, response)
}

//DeleteExpenseDefinition function
func DeleteExpenseDefinition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	expenseDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(expenseDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	expenseDefinition := modelOf.DeleteExpenseDefinition(id)
	response, _ := json.Marshal(expenseDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetAllExpenseDefinitions function
func GetAllExpenseDefinitions(w http.ResponseWriter, r *http.Request) {
	expenseDefinitions := modelOf.GetAllExpenseDefinitions()
	response, _ := json.Marshal(expenseDefinitions)
	utils.WriteResponse(w, http.StatusOK, response)
}

//GetExpenseDefinitionByID function
func GetExpenseDefinitionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	expenseDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(expenseDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	expenseDefinition, _ := modelOf.GetExpenseDefinitionByID(id)
	response, _ := json.Marshal(&expenseDefinition)
	utils.WriteResponse(w, http.StatusOK, response)
}

//UpdateExpenseDefinition function
func UpdateExpenseDefinition(w http.ResponseWriter, r *http.Request) {
	var expenseDefinitionDto = &modelOf.ExpenseDefinition{}
	utils.ParseBody(r, expenseDefinitionDto)
	vars := mux.Vars(r)
	expenseDefinitionID := vars["Id"]
	id, err := strconv.ParseInt(expenseDefinitionID, 0, 0)
	if err != nil {
		fmt.Println("int parse error")
		return
	}

	expenseDefinition, db := modelOf.GetExpenseDefinitionByID(id)
	if expenseDefinitionDto.Name != "" {
		expenseDefinition.Name = expenseDefinitionDto.Name
	}

	if expenseDefinitionDto.ExpenseDefinitionGroupID != nil {
		expenseDefinition.ExpenseDefinitionGroupID = expenseDefinitionDto.ExpenseDefinitionGroupID
	}

	if expenseDefinitionDto.Active != nil {
		expenseDefinition.Active = expenseDefinitionDto.Active
	}

	db.Save(&expenseDefinition)
	response, _ := json.Marshal(&expenseDefinition)

	utils.WriteResponse(w, http.StatusOK, response)
}
