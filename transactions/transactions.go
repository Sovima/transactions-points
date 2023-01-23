// This is the package that collects all of the transactions
// and calculates the total rewards as well as the 
// rewards for the individual transactions
package transactions

import (
	"fmt"
	"encoding/json"
	"display/transactions/rules"
)


// Create a struct for individual transactions
// Follows the input format
type Transaction struct {
	amountCents float64
    date string
    merchantCode string
}


// Create a struct that collects all of the individual transactions
// as well as the full list of rules to be applied to those transactions
type Transactions struct {
	individualTransactions map[string]Transaction
	totalValues map[string] int
	rewards []rules.Rule

}


func (t *Transactions) Populate(collectedTransactions string) {
	// Populate the Transactions struct with the information from 
	// collectedTransactions
	// input: 
	//     collectedTransactions string : all of the monthly transactions 
	//         in json string format. For example, 
	//         {"T01": {"date": "2021-05-01", "merchant_code" : "sportcheck", "amount_cents": 6123}}
	// return:
	//     No return but the fields of the Transactions struct get updated
	//     to reflect the input list of transactions


	// Convert the json string into map[string]Transaction
	t.individualTransactions = convertTransactions(collectedTransactions)

	// Keep the running totals for each merchant_code
	t.totalValues = map[string]int{"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}

	// Go through each transaction and update the running total
	// for the corresponding total
	for _, value := range t.individualTransactions {
		t.totalValues[value.merchantCode] += int(value.amountCents)
	}

	fmt.Printf("total Tim Hortons: %v\n", t.totalValues["tim_hortons"])
	fmt.Printf("total Subway: %v\n", t.totalValues["subway"])
	fmt.Printf("total SportCheck: %v\n", t.totalValues["sportcheck"])
}


func (t *Transactions) GetRules() {
	// Collect the rewards for the collected transactions
	// input: 
	//     no input but fields from the passed Transactions class are used to 
	//     get running totals for each merchant
	// return:
	//     No return but the rewards field of the Transactions struct get updated
	//     to reflect the full list of reward rules to be applied 
	//     to the input transactions

	t.rewards = rules.CollectRewards(t.totalValues)
}

func (t *Transactions) ApplyRules() int {
	// Collect the reward points from each reward rule
	// input: 
	//     no input but rewards field from the Transactions class is used to 
	//     update the total number of rewards
	// return:
	//     int : the total number of points collected with the given transactions

	points := 0

	// Go through each reward and collect the reward points from each
	for _, reward := range t.rewards {
		points += reward.ApplyRule()
	}

	return points
}


func (t *Transactions) GetRewardsPerTransaction(transaction string) int {
	// Get reward points for individual transaction
	// input: 
	//     transaction string : ID of the desired transaction
	//     Additionally, the fields of Transactions class are used 
	//     to get the information about a transaction based on the input ID
	// return:
	//     int : the total number of points collected in a single transaction

	// Get the transaction from the map of all monthly transactions
	oneTransaction := t.individualTransactions[transaction]

	// Create a map to have the running total for a single transaction
	// note that two of the three fields will be 0 since one transaction 
	// can contain at most one merchant
	oneTransactionSpending := map[string]int{"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}

	// Update the amount spent for the merchant in a single transaction
	oneTransactionSpending[oneTransaction.merchantCode] += int(oneTransaction.amountCents)

	// Collect the rewards for the single transaction
	rewards := rules.CollectRewards(oneTransactionSpending)

	// Walk through each reward rule and accumulate the 
	// obtained points
	var points int
	for _, reward := range rewards {
		points += reward.ApplyRule()
	}

	return points
}



func convertTransactions(collectedTransactions string) map[string]Transaction{
	// Convert json string containing transaction information into
	// map[string]Transaction (a map from transaction ID to the corresponding transaction info)
	// input: 
	//     collectedTransactions string : json string containing information for 
	//                                    all transactions
	// return:
	//     map[string]Transaction : a map from transaction ID to the corresponding transaction info


	// Define the structure of the output data
	data := map[string]Transaction{}

	// Define the initial structure of the input
	var first_input map[string]interface{}

	// give the initial string structure of map[string]interface{}
	jsonErr := json.Unmarshal([]byte(collectedTransactions), &first_input)

	// check for any errors
	if jsonErr != nil {
		fmt.Print("ERROR JSON:")
        fmt.Print(jsonErr.Error())
	}

	// Go through every field in the initial input and 
	// convert it into Transaction struct
	for key, transaction := range first_input {
		curr_transaction := transaction.(map[string]interface{})
		data[key] = Transaction{curr_transaction["amount_cents"].(float64), 
								curr_transaction["date"].(string), 
								curr_transaction["merchant_code"].(string)}
	}
    return data
}