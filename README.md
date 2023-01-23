# Reward points

## Table of contents
- [Reward points](#reward-points)
  * [Walkthrough](#walkthrough)
  * [Project Description](#project-description)
  * [Usage](#usage)
  * [Adding a new rule](#adding-a-new-rule)

## Walkthrough

For the video walk through the functionalities of the website, go to [video walkthrough](https://youtu.be/S1-GgV5g3O8)

## Project Description

This module implements a web server that runs on localhost, collects transactions, and allows to track reward points for all transactoins as well as individually. 

The rules by which reward points are awarded are defined in the `transactions/rules/rules.go` file. In this case, the greedy approach is used. That is, each rule is applied as many times as possible with the following priorities:
1. Rule 1
2. 2x Rule 4
3. Rule 2
3. Rule4 + 2x Rule6 (Rule8)
4. 3x Rule 6
5. Rule 4
7. Rule 6
8. Rule 7

Note that rule 5 and rule3 are not present as more points can be earned by fully substituting them with other available rules. 

The rules package is used in `transactions/transactions.go` which implements the pipeline for collecting transactions, collecting rules, and applying the rules to collect reward points.

Lastly, `SofyaMalashchenko.go` contains implementation of the server that runs on local host and combines above mentioned packages and files and gives an easy-to-use web application. 

For further information on the specific files and functions, please view the comments in `rules.go`, `transactions.go`, and `SofyaMalashchenko.go`.

## Usage
## Adding a new rule

Current implementation of the rule structs allows for an easy addition of new rules. 

To add a new rule, create a new rule struct that has GetRequirements and ApplyRule methods as well as update `collectRewards( transactions map[string]int ) []Rule` function in rules.go to add the new rule. These implementations should be added in the `rules.go` file.


For example, see rule 8:
```
var REQUIREMENT8 = map[string]int {"tim_hortons": 1000, "subway": 1000, "sportcheck" : 6500,}

// An additional rule that is a combination of rule6 * 2 and rule4
type Rule8 struct {}

func (r *Rule8) GetRequirements() map[string] int {
	return REQUIREMENT8
}

func (r *Rule8) ApplyRule() int {
	return 300
}
```
