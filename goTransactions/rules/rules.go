package rules

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"strings"
)


type Rule interface {
	ApplyRule()
}


// Rule1
type Rule1 struct {
	requirements map[string]int
	rewards int
}


func (r *Rule1) ApplyRule() int {
	return r.rewards
}


// Rule2
type Rule2 struct {
	requirements map[string]int
	rewards int
}


func (r *Rule2) ApplyRule() int {
	return r.rewards
}

// Rule3
type Rule3 struct {
	requirements map[string]int
	rewards int
}


func (r *Rule3) ApplyRule() int {
	return r.rewards
}

// Rule4
type Rule4 struct {
	requirements map[string]int
	rewards int
}


func (r *Rule4) ApplyRule() int {
	return r.rewards
}


// Rule5
type Rule5 struct {
	requirements map[string]int
	rewards int
}


func (r *Rule5) ApplyRule() int {
	return r.rewards
}


// Rule6
type Rule6 struct {
	requirements map[string]int
	rewards int
}


func (r *Rule6) ApplyRule() int {
	return r.rewards
}

