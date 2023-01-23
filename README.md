# Reward points

## Project description

This module implements a web server that runs on localhost, collects transactions, and allows to track reward points for all transactoins as well as individually. 

The rules by which reward points are awarded are defined in the `transactions/rules/rules.go` file. In this case, the greedy approach is used. That is, each rule is applied as many times as possible with the following priorities:
1. Rule 1
2. 2x Rule 4
3. Rule 2
3. Rule4 + 2x Rule6
4. 3x Rule 6
5. Rule 4
7. Rule 6
8. Rule 7

Note that rule 5 and rule3 are not present as more points can be earned by fully substituting them with other available rules. 

The rules package is used in `transactions/transactions.go` which implements the pipeline for collecting transactions, collecting rules, and applying the rules to collect reward points.

Lastly, `SofyaMalashchenko.go`


## Walkthrough

For the video walk through the functionalities of the website, go to [video walkthrough](https://youtu.be/S1-GgV5g3O8)

## Adding a new rule

To add a new rule, create a new rule struct that has GetRequirements and ApplyRule methods. For example, see rule 7:
```
var REQUIREMENT7 = map[string]int {"tim_hortons": 0, "subway": 0, "sportcheck" : 0,}

type Rule7 struct {}


func (r *Rule7) GetRequirements() map[string] int {
	return REQUIREMENT7
}


func (r *Rule7) ApplyRule() int {
	return 1
}
```
