# Reward points


The whole module can be found in `goTransactions`

The rules are defined in the rules.py file


In this case, the greedy approach is used where each rule is applied as many times as possible in the following order:
1. Rule 1
2. 2x Rule 4
3. Rule 2
4. 3x Rule 6
5. Rule 4
7. Rule 6
8. Rule 7

Note that rule 5 is not present as it is redundant compared to rule 6. 

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
