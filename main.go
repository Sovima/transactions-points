// This is the starting point of the webapp
// It will be handling all of the calls from/to the frontend
package main

import (
	"html/template"
	"fmt"
	"net/http"
	"display/transactions"
)


var points int = 0
var newTransactions transactions.Transactions = transactions.Transactions{}



func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even treat them like variables in Go)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)
	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        transactionsJson := r.FormValue("transactions")

		newTransactions.Populate(transactionsJson)
		newTransactions.GetRules()
		points = newTransactions.ApplyRules()
		fmt.Println(newTransactions)



        tmpl.Execute(w, struct{ Total_points int; Individual_points int}{points, 0})
    })


	http.HandleFunc("/check-one", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        transactionsJson := r.FormValue("transaction-num")

		oneTransactionRewards := newTransactions.GetRewardsPerTransaction(transactionsJson)
        tmpl.Execute(w, struct{ Total_points int; Individual_points int}{points, oneTransactionRewards})
    })


	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}