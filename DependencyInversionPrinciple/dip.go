package main

import (
	"fmt"
)

// Solution for the DIP problem

type RelationshipBrowser interface {
	FindAllChildrenOf(string) []*Person
}

type Research2 struct {
	// Depend on abstractions
	browser RelationshipBrowser
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Research2) Investigate() {

	// If we were to change something, then it needs to be done in the low level module without affecting the high level
	// module, in this case we would need to change FindAllChildrenOf method from Relationships instead of Investigate
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}
