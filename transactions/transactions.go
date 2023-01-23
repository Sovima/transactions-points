package transactions

import (
	"fmt"
	// "io/ioutil"
	// "os"
	"encoding/json"
	// "strings"
	"display/transactions/rules"
)


type Transaction struct {
	amountCents float64
    date string
    merchantCode string
}


type Transactions struct {
	individualTransactions map[string]Transaction

	totalValues map[string] int

	rewards []rules.Rule

}

func (t *Transactions) Populate(collectedTransactions string) {
	t.individualTransactions = collectTransactions(collectedTransactions)
	t.totalValues = map[string]int{"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}

	for _, value := range t.individualTransactions {
		t.totalValues[value.merchantCode] += int(value.amountCents)
		// fmt.Printf("Amount cents %f with merchant %s\n", value.amountCents, value.merchantCode)

	}

	fmt.Printf("total Tim Hortons: %v\n", t.totalValues["tim_hortons"])
	fmt.Printf("total Subway: %v\n", t.totalValues["subway"])
	fmt.Printf("total SportCheck: %v\n", t.totalValues["sportcheck"])
}


func (t *Transactions) GetRules() {
	t.rewards = rules.CollectRewards(t.totalValues)
	// At this point we must decide which rules to apply
}

func (t *Transactions) ApplyRules() int {
	points := 0
	for _, reward := range t.rewards {
		points += reward.ApplyRule()
	}

	return points
	
}


func (t *Transactions) GetRewardsPerTransaction(transaction string) int {
	oneTransaction := t.individualTransactions[transaction]

	oneTransactionSpending := map[string]int{"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}
	oneTransactionSpending[oneTransaction.merchantCode] += int(oneTransaction.amountCents)

	fmt.Println(oneTransactionSpending)
	rewards := rules.CollectRewards(oneTransactionSpending)
	var points int
	for _, reward := range rewards {
		points += reward.ApplyRule()
		fmt.Println("Getting points" , points)
	}

	return points

}



func collectTransactions(collectedTransactions string) map[string]Transaction {
	// Read from the transactions file and 
	// convert the input into a pre-defined interface
	// Note that this assumes the dictionary structure
	// defined in the question


	// Open jsonTransactions
	// jsonTransactions, err := os.Open("transactions.json")
	// // handle possible errors
	// if err != nil {
	// 	fmt.Println(err)
	// }


	// fmt.Println("Successfully Opened users.json")
	// // To make sure the file gets closed in the end
	// defer jsonTransactions.Close()


	// byteTransactions, _ := ioutil.ReadAll(jsonTransactions)

	data := map[string]Transaction{}
	var test map[string]interface{}
	jsonErr := json.Unmarshal([]byte(collectedTransactions), &test)

	if jsonErr != nil {
		fmt.Print("ERROR JSON:")
        fmt.Print(jsonErr.Error())
	}

	for key, transaction := range test {
		// fmt.Println(key,transaction)
		curr_transaction := transaction.(map[string]interface{})
		data[key] = Transaction{curr_transaction["amount_cents"].(float64), 
								curr_transaction["date"].(string), 
								curr_transaction["merchant_code"].(string)}
	}

	// fmt.Println(data)

    return data
}


// func main() {
// 	monthlyTransactions := Transactions{}
// 	monthlyTransactions.Populate() // Collect the transactions
// 	monthlyTransactions.GetRules() // Collect the rules
// 	monthlyTransactions.ApplyRules() // Apply collected rules
// 	// monthlyTransactions.getRewardsPerTransaction("T09") // Apply collected rules
// }