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

To run the server, run 
```
go run SofyaMalashchenko.go
```
from the root project folder (the folder where this `README.md` is located). 
Then, in browser open `http://localhost:8080/`. This will take you to the starting point of the web app. 

In the form, insert the list of transactions in json string format and press the submit button. For example:
```
{
    "T01": {"date": "2021-05-01", "merchant_code" : "sportcheck", "amount_cents": 3500},
    "T02": {"date": "2021-05-02", "merchant_code" : "sportcheck", "amount_cents": 8700},
    "T03": {"date": "2021-05-03", "merchant_code" : "tim_hortons", "amount_cents": 323},
    "T04": {"date": "2021-05-04", "merchant_code" : "tim_hortons", "amount_cents": 1267},
    "T05": {"date": "2021-05-05", "merchant_code" : "tim_hortons", "amount_cents": 2116},
    "T06": {"date": "2021-05-06", "merchant_code" : "tim_hortons", "amount_cents": 2211},
    "T07": {"date": "2021-05-07", "merchant_code" : "subway", "amount_cents": 1853},
    "T08": {"date": "2021-05-08", "merchant_code" : "subway", "amount_cents": 2153},
    "T09": {"date": "2021-05-09", "merchant_code" : "sportcheck", "amount_cents": 7326},
    "T10": {"date": "2021-05-10", "merchant_code" : "tim_hortons", "amount_cents": 1321}
}
```

The output will be the total number of reward points collected from these transactions. Also, a new form will apear under the first form. After this, there are two options:
1. In the newly appeared form, submit a transaction ID to view the number of reward points collected from that transaction. The value of reward points will appear under the form. Otherwise, no response will be produced. 
2. In the original form, submit another list of transactions in the same format. 

Repeat option 1 or option 2 to view more information about the current or new transactions. 

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
