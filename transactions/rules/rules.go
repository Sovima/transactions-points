package rules

import (
	"math"
)

// The values are given in cents
var REQUIREMENT1 = map[string]int {"tim_hortons": 2500, "subway": 2500, "sportcheck" : 7500,}
var REQUIREMENT2 = map[string]int {"tim_hortons": 2500, "subway": 0, "sportcheck" : 7500,}
var REQUIREMENT3 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 7500,}
var REQUIREMENT4 = map[string]int {"tim_hortons": 1000, "subway": 1000, "sportcheck" : 2500,}
var REQUIREMENT5 = map[string]int {"tim_hortons": 1000, "subway": 0, "sportcheck" : 2500,}
var REQUIREMENT6 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 2000,}
var REQUIREMENT7 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}


func CollectRewards( transactions map[string]int ) []Rule {
	var output []Rule

	// fmt.Println("Starting values")

	// fmt.Println(transactions)

	// Check how many rules1 can be applied

	rule1 := new(Rule1)
	transactions, output = countRules(rule1, 1, transactions, output)
	
	
	// Check how many 2 * rules4 can be applied
	rule4 := new(Rule4)
	transactions, output = countRules(rule4, 2, transactions, output)

	// Check how many rules2 can be applied
	rule2 := new(Rule2)
	transactions, output = countRules(rule2, 1, transactions, output)

	// Check how many rules4 can be applied
	transactions, output = countRules(rule4, 1, transactions, output)

	// Check how many 3*rules6 can be applied
	rule6 := new(Rule6)
	transactions, output = countRules(rule6, 3, transactions, output)

	// Check how many rules3 can be applied
	rule3 := new(Rule3)
	transactions, output = countRules(rule3, 1, transactions, output)

	// Check how many rules6 can be applied
	transactions, output = countRules(rule6, 1, transactions, output)


	// Check how many rules7 can be applied
	rule7 := new(Rule7)
	total_remaining_points := transactions["tim_hortons"] / 100 + transactions["subway"] / 100 + transactions["sportcheck"] / 100

	for i := 0; i < total_remaining_points; i++ {
		output = append(output, rule7)
	}

	// fmt.Println("Remaining values")

	// fmt.Println(transactions)
	return output
}


func countRules(ruleToAdd Rule, group int, transactions map[string]int, output []Rule) (map[string]int, []Rule) {
	ruleRequirements := ruleToAdd.GetRequirements()
	max_with_zero := int(math.Max(float64(transactions["sportcheck"]), math.Max(float64(transactions["tim_hortons"]), float64(transactions["subway"]))))
	max_sportcheck := max_with_zero
	max_tims := max_with_zero
	max_subway := max_with_zero
	if ruleRequirements["sportcheck"] > 0 {
		max_sportcheck = transactions["sportcheck"] / ruleRequirements["sportcheck"] / group
	}
	if ruleRequirements["tim_hortons"] > 0 {
		max_tims = transactions["tim_hortons"] / ruleRequirements["tim_hortons"] / group
	}
	if ruleRequirements["subway"] > 0 {
		max_subway = transactions["subway"] / ruleRequirements["subway"] / group
	}

	rule_count := int(math.Min(float64(max_sportcheck), math.Min(float64(max_tims), float64(max_subway))))

	for i := 0; i < rule_count; i++ {
		for j := 0; j < group; j++ {
			output = append(output, ruleToAdd)
		}
		transactions["sportcheck"] -= group * ruleRequirements["sportcheck"]
		transactions["tim_hortons"] -= group * ruleRequirements["tim_hortons"]
		transactions["subway"] -= group * ruleRequirements["subway"]
	}

	// fmt.Println("Remaining transactions ", transactions)


	return transactions, output
}



type Rule interface {
	GetRequirements() map[string] int
	ApplyRule() int
}


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
	return REQUIREMENT1
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