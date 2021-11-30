package main

import (
	"fmt"
	"homebud/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World")
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", dashboardHandler)

	handleAccountController(myRouter)
	handleCurrencyController(myRouter)
	handleEnumController(myRouter)
	handleExpenseDefinition(myRouter)
	handleExpenseGroupDefinition(myRouter)
	handleIncomingDefinition(myRouter)
	handleIncomingGroupDefinition(myRouter)
	handleMonthlyExpenseRecord(myRouter)
	handleMonthlyIncomingRecord(myRouter)

	http.Handle("/", myRouter)
	http.ListenAndServe(":8080", nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome the HomeBud Project")
}

func handleAccountController(myRouter *mux.Router) {
	myRouter.HandleFunc("/accounts", controllers.GetAllAccounts).Methods("GET")
	myRouter.HandleFunc("/accounts/{id}", controllers.GetAccountByID).Methods("GET")
	myRouter.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	myRouter.HandleFunc("/accounts", controllers.UpdateAccount).Methods("PUT")
	myRouter.HandleFunc("/accounts/{id}", controllers.DeleteAccount).Methods("DELETE")
}

func handleCurrencyController(myRouter *mux.Router) {
	myRouter.HandleFunc("/currencies", controllers.GetAllCurrencies).Methods("GET")
	myRouter.HandleFunc("/currencies/{id}", controllers.GetCurrencyByID).Methods("GET")
	myRouter.HandleFunc("/currencies", controllers.CreateCurrency).Methods("POST")
	myRouter.HandleFunc("/currencies", controllers.UpdateCurrency).Methods("PUT")
	myRouter.HandleFunc("/currencies/{id}", controllers.DeleteCurrency).Methods("DELETE")
}

func handleEnumController(myRouter *mux.Router) {
	myRouter.HandleFunc("/enums/months", controllers.GetMonths).Methods("GET")
	myRouter.HandleFunc("/enums/years", controllers.GetYears).Methods("GET")
}

func handleExpenseDefinition(myRouter *mux.Router) {
	myRouter.HandleFunc("/expenseDefinitions", controllers.GetAllExpenseDefinitions).Methods("GET")
	myRouter.HandleFunc("/expenseDefinitions/{id}", controllers.GetExpenseDefinitionByID).Methods("GET")
	myRouter.HandleFunc("/expenseDefinitions", controllers.CreateExpenseDefinition).Methods("POST")
	myRouter.HandleFunc("/expenseDefinitions", controllers.UpdateExpenseDefinition).Methods("PUT")
	myRouter.HandleFunc("/expenseDefinitions/{id}", controllers.DeleteExpenseDefinition).Methods("DELETE")
}

func handleExpenseGroupDefinition(myRouter *mux.Router) {
	myRouter.HandleFunc("/expenseGroupDefinitions", controllers.GetAllExpenseGroupDefinitions).Methods("GET")
	myRouter.HandleFunc("/expenseGroupDefinitions/{id}", controllers.GetExpenseGroupDefinitionByID).Methods("GET")
	myRouter.HandleFunc("/expenseGroupDefinitions", controllers.CreateExpenseGroupDefinition).Methods("POST")
	myRouter.HandleFunc("/expenseGroupDefinitions", controllers.UpdateExpenseGroupDefinition).Methods("PUT")
	myRouter.HandleFunc("/expenseGroupDefinitions/{id}", controllers.DeleteExpenseGroupDefinition).Methods("DELETE")
}

func handleIncomingDefinition(myRouter *mux.Router) {
	myRouter.HandleFunc("/incomingDefinitions", controllers.GetAllIncomingDefinitions).Methods("GET")
	myRouter.HandleFunc("/incomingDefinitions/{id}", controllers.GetIncomingDefinitionByID).Methods("GET")
	myRouter.HandleFunc("/incomingDefinitions", controllers.CreateIncomingDefinition).Methods("POST")
	myRouter.HandleFunc("/incomingDefinitions", controllers.UpdateIncomingDefinition).Methods("PUT")
	myRouter.HandleFunc("/incomingDefinitions/{id}", controllers.DeleteIncomingDefinition).Methods("DELETE")
}

func handleIncomingGroupDefinition(myRouter *mux.Router) {
	myRouter.HandleFunc("/incomingGroupDefinitions", controllers.GetAllIncomingGroupDefinitions).Methods("GET")
	myRouter.HandleFunc("/incomingGroupDefinitions/{id}", controllers.GetIncomingGroupDefinitionByID).Methods("GET")
	myRouter.HandleFunc("/incomingGroupDefinitions", controllers.CreateIncomingGroupDefinition).Methods("POST")
	myRouter.HandleFunc("/incomingGroupDefinitions", controllers.UpdateIncomingGroupDefinition).Methods("PUT")
	myRouter.HandleFunc("/incomingGroupDefinitions/{id}", controllers.DeleteIncomingGroupDefinition).Methods("DELETE")
}

func handleMonthlyExpenseRecord(myRouter *mux.Router) {
	myRouter.HandleFunc("/monthlyExpenseRecords", controllers.GetAllMonthlyExpenseRecords).Methods("GET")
	myRouter.HandleFunc("/monthlyExpenseRecords/{id}", controllers.GetMonthlyExpenseRecordByID).Methods("GET")
	myRouter.HandleFunc("/monthlyExpenseRecords", controllers.CreateMonthlyExpenseRecord).Methods("POST")
	myRouter.HandleFunc("/monthlyExpenseRecords", controllers.UpdateMonthlyExpenseRecord).Methods("PUT")
	myRouter.HandleFunc("/monthlyExpenseRecords/{id}", controllers.DeleteMonthlyExpenseRecord).Methods("DELETE")
}

func handleMonthlyIncomingRecord(myRouter *mux.Router) {
	myRouter.HandleFunc("/monthlyIncomingRecords", controllers.GetAllMonthlyIncomingRecords).Methods("GET")
	myRouter.HandleFunc("/monthlyIncomingRecords/{id}", controllers.GetMonthlyIncomingRecordByID).Methods("GET")
	myRouter.HandleFunc("/monthlyIncomingRecords", controllers.CreateMonthlyIncomingRecord).Methods("POST")
	myRouter.HandleFunc("/monthlyIncomingRecords", controllers.UpdateMonthlyIncomingRecord).Methods("PUT")
	myRouter.HandleFunc("/monthlyIncomingRecords/{id}", controllers.DeleteMonthlyIncomingRecord).Methods("DELETE")
}
