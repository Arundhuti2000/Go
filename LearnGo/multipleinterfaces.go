package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed != true{
		cost:= len(e.body)*5
		return cost
	} else {
		return len(e.body)*2
	}
}

func (e email) format() string {
	if e.isSubscribed == true{
		return fmt.Sprintf("'%s' | Subscribed", e.body)
	} else {
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
