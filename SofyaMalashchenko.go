// This is the starting point of the webapp
// It will be handling all of the calls from/to the frontend
package main

import (
	"html/template"
	"net/http"
	"display/transactions"
)

// Use points and newTransactions variables to 
// store information about input transactions and 
// total rewards points
var points int = 0
var newTransactions transactions.Transactions = transactions.Transactions{}



func main() {
	// The frontend consists of a single page, the index page
	// which is updated as more input is recieved
	tmpl := template.Must(template.ParseFiles("index.html"))


	// Add a handler to load the index page with the get request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// index page is loaded through a get request
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }
    })

	// Add a handler to recieve the response from the first 
	// form and update the total amount of reward points
	// for all transactions
	http.HandleFunc("/calculate-total", func(w http.ResponseWriter, r *http.Request) {

		// the form output is sent through a post request
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

		// Extract all of the transactions from the 
		// transactions form 
        transactionsJson := r.FormValue("transactions")

		// Populate the newTransactions variable with the
		// input transactions
		newTransactions.Populate(transactionsJson)

		// Get the rules for the input transactions
		// For the algorithm to determine the rules
		// check transactions/rules/rules.go

		newTransactions.GetRules()

		// Apply the rules to calculate the reward points
		points = newTransactions.ApplyRules()

		// Load the index page with the updated number of total points
		// Note that individual_points are set to 0 as we have 
		// not yet selected the transaction id to calculate individual points for
        tmpl.Execute(w, struct{ Total_points int; Individual_points int}{points, 0})
    })

	// Add a handler to recieve the response from the second form
	// (the one that accepts individual transaction IDs) 
	// and update the number of reward points for a single transaction
	http.HandleFunc("/check-one", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

		// Get the transaction ID from the form
        transactionsJson := r.FormValue("transaction-num")

		// Find the individual transaction in the newTransactions,
		// calculate individual points, and update Individual_points
		// variable used in index.html
		oneTransactionRewards := newTransactions.GetRewardsPerTransaction(transactionsJson)
        tmpl.Execute(w, struct{ Total_points int; Individual_points int}{points, oneTransactionRewards})
    })


	// Now that we have defined the server, we Listen and Serve any of the 
	// requests
	http.ListenAndServe(":8080", nil)
}