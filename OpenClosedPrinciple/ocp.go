package main

/*
	Open-Close principle basically we need to provide something our user can implement rather than modifying the code already
	created, Open for extension, closed for modification
*/

type Specification interface {
	IsSatisfied(*Product) bool
}

// For size comparison
type SizeSpecification struct {
	size Size
}

// For color comparison
type ColorSpecification struct {
	color Color
}

/*
	What if we want a color an size comparison=, instead of writing a function such as "BetterFilterProductsSizeAndColor"
	we are going to create a struct which is going to extend functionality of the Specification interface already created
*/

type AndSpecification struct {
	first  Specification
	second Specification
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// We need to satisfy the interface

func (a *AndSpecification) IsSatisfied(product *Product) bool {
	return a.first.IsSatisfied(product) && a.second.IsSatisfied(product)
}

func BetterFilterProducts(prs []Product, specification Specification) []Product {
	r := make([]Product, 0)
	for i, v := range prs {
		if specification.IsSatisfied(&v) {
			r = append(r, prs[i])
		}
	}
	return r
}
