package main

import "fmt"

// Dependency Inversion principle
// High level modules should not depend on low level modules
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type RelationInfo struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-Level modules (It stores maybe to a database)
type Relationships struct {
	relations []RelationInfo
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, RelationInfo{parent, Parent, child})
	r.relations = append(r.relations, RelationInfo{child, Child, parent})
}

// High-Level module, please check dip file for solution

type Research struct {
	// Break DIP
	// The reason being that we are making the research "module" depending on "Relationships" module which is a low level
	// module

	// If something changes on the Low-Level module Relationships then the code breaks
	relationships *Relationships
}

func (r *Research) Investigate() {
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("Jhon has a child called", rel.to.name)
		}
	}
}

func main() {

	parent := Person{"John"}
	child := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Solution:")
	// Solution
	research2 := Research2{&relationships}
	research2.Investigate()

}
