// This is the package that contains all of the 
// logic for collecting rules and 
package rules

import (
	"math"
)

// The following variables store 
// minimum requirements for the corresponding rule
// For example REQUIREMENT1 stores requirements for rule1
// That is, $25 tims, $25 subway, and $75 sportcheck
// Note that all amounts are stored in cents
var REQUIREMENT1 = map[string]int {"tim_hortons": 2500, "subway": 2500, "sportcheck" : 7500,}
var REQUIREMENT2 = map[string]int {"tim_hortons": 2500, "subway": 0, "sportcheck" : 7500,}
var REQUIREMENT3 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 7500,}
var REQUIREMENT4 = map[string]int {"tim_hortons": 1000, "subway": 1000, "sportcheck" : 2500,}
var REQUIREMENT5 = map[string]int {"tim_hortons": 1000, "subway": 0, "sportcheck" : 2500,}
var REQUIREMENT6 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 2000,}
var REQUIREMENT7 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}
var REQUIREMENT8 = map[string]int {"tim_hortons": 1000, "subway": 1000, "sportcheck" : 6500,}


func CollectRewards( transactions map[string]int ) []Rule {
	// this function implements the logic for collecting the
	// maximum number of reward points for the input transactions
	// input: 
	//     transactions map[string]int : the total amounts spent for each merchant
	// return:
	//     []Rule : the list of rules that will be applied to the 
	//              provided transactions

	// Define the output variable
	var output []Rule

	// Check how many rule1s can be applied
	rule1 := new(Rule1)
	// Append the right number of rule1 to the output
	// and adjust the remaining amounts spent for each merchant
	transactions, output = countRules(rule1, 1, transactions, output)
	
	
	// Check how many 2 * rules4 can be applied
	rule4 := new(Rule4)
	// Append the right number of 2 * rule4 to the output
	// and adjust the remaining amounts spent for each merchant
	transactions, output = countRules(rule4, 2, transactions, output)


	// Check how many rules2 can be applied
	rule2 := new(Rule2)
	// Append the right number of rule2 to the output
	// and adjust the remaining amounts spent for each merchant
	transactions, output = countRules(rule2, 1, transactions, output)


	// Check how many rules8 can be applied
	rule8 := new(Rule8)
	// Append the right number of rule8 to the output
	// and adjust the remaining amounts spent for each merchant
	transactions, output = countRules(rule8, 1, transactions, output)


	// Check how many 3*rules6 can be applied
	rule6 := new(Rule6)
	// Append the right number of rule6 to the output
	// and adjust the remaining amounts spent for each merchant
	transactions, output = countRules(rule6, 3, transactions, output)
	

	// Check how many rules4 can be applied
	// Append the right number of rule4 to the output
	// and adjust the remaining amounts spent for each merchant

	transactions, output = countRules(rule4, 1, transactions, output)


	// Check how many rules6 can be applied

	transactions, output = countRules(rule6, 1, transactions, output)


	// Check how many rules7 can be applied
	rule7 := new(Rule7)

	// Collect the total remaining values between all merchants
	// and add rule 7 to every remaining dollar
	total_remaining_points := transactions["tim_hortons"] / 100 + transactions["subway"] / 100 + transactions["sportcheck"] / 100

	// Add rule 7 to every remaining dollar
	for i := 0; i < total_remaining_points; i++ {
		output = append(output, rule7)
	}

	return output
}


func countRules(ruleToAdd Rule, group int, transactions map[string]int, output []Rule) (map[string]int, []Rule) {
	// This function calculates the maximum number of 
	// times a given rule can be applied to the remaining transactions
	// input: 
	//     ruleToAdd Rule : the type of rule we wish to add 
	//     group int : the number of times a rule is grouped
	//                 For example, if set to 3 and ruleToAdd is rule6
	//                 we will only apply rule6 in groups of 3
	//     transactions map[string]int : the total amounts spent for each merchant
	//     output []Rule : the list of so far collected rules
	// return:
	//     map[string]int : the updated remaining values spent on each merchant
	//     []Rule : the updated list of rules to be applied to the list of transactions


	// Get the requirements for the added rule
	ruleRequirements := ruleToAdd.GetRequirements()

	// In case any of the required values is 0 we want to aviod division by zero
	max_with_zero := int(math.Max(float64(transactions["sportcheck"]), math.Max(float64(transactions["tim_hortons"]), float64(transactions["subway"]))))

	// These values will store the maximum number of times a given rule can be 
	// applied to the specific merchant
	max_sportcheck := max_with_zero
	max_tims := max_with_zero
	max_subway := max_with_zero

	// Now, if the restrictions are not 0, we can divide by them and update 
	// maximum number of times a given rule can be applied to the specific merchant
	if ruleRequirements["sportcheck"] > 0 {
		max_sportcheck = transactions["sportcheck"] / ruleRequirements["sportcheck"] / group
	}
	if ruleRequirements["tim_hortons"] > 0 {
		max_tims = transactions["tim_hortons"] / ruleRequirements["tim_hortons"] / group
	}
	if ruleRequirements["subway"] > 0 {
		max_subway = transactions["subway"] / ruleRequirements["subway"] / group
	}


	// Note that the count for a given rule
	// is bounded above by the smallest number of times that the given rule 
	// can be applied to specific merchant. Hence we are taking the minimum
	rule_count := int(math.Min(float64(max_sportcheck), math.Min(float64(max_tims), float64(max_subway))))

	// Append the given rule to the output list rule_count of times
	// and update the remaining transaction values for each merchant
	for i := 0; i < rule_count; i++ {
		for j := 0; j < group; j++ {
			output = append(output, ruleToAdd)
		}
		transactions["sportcheck"] -= group * ruleRequirements["sportcheck"]
		transactions["tim_hortons"] -= group * ruleRequirements["tim_hortons"]
		transactions["subway"] -= group * ruleRequirements["subway"]
	}

	return transactions, output
}

// This interface allows us to pass different rules in
// the countRules function 
// It lists the methods that each of the concrete classes
// must implement
type Rule interface {
	GetRequirements() map[string] int
	ApplyRule() int
}

// Below you will see individual rules that implement
// the abstract Rule class

// Rule1
type Rule1 struct {}


func (r *Rule1) GetRequirements() map[string] int {
	return REQUIREMENT1
}


func (r *Rule1) ApplyRule() int {
	return 500
}


// Rule2
type Rule2 struct {}


func (r *Rule2) GetRequirements() map[string] int {
	return REQUIREMENT2
}


func (r *Rule2) ApplyRule() int {
	return 300
}


// Rule3
type Rule3 struct {}


func (r *Rule3) GetRequirements() map[string] int {
	return REQUIREMENT3
}


func (r *Rule3) ApplyRule() int {
	return 200
}


// Rule4
type Rule4 struct {}


func (r *Rule4) GetRequirements() map[string] int {
	return REQUIREMENT4
}


func (r *Rule4) ApplyRule() int {
	return 150
}


// Rule5
type Rule5 struct {}


func (r *Rule5) GetRequirements() map[string] int {
	return REQUIREMENT5
}


func (r *Rule5) ApplyRule() int {
	return 75
}


// Rule6
type Rule6 struct {}


func (r *Rule6) GetRequirements() map[string] int {
	return REQUIREMENT6
}


func (r *Rule6) ApplyRule() int {
	return 75
}

// Rule7
type Rule7 struct {}

func (r *Rule7) GetRequirements() map[string] int {
	return REQUIREMENT7
}

func (r *Rule7) ApplyRule() int {
	return 1
}


// Rule8
// An additional rule that is a combination of rule6 * 2 and rule4
type Rule8 struct {}

func (r *Rule8) GetRequirements() map[string] int {
	return REQUIREMENT8
}

func (r *Rule8) ApplyRule() int {
	return 300
}

