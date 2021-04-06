package main

import (
	"fmt"
)

// OCP
// Open for extension, closed for modification
// Specification enterprise pattern

type Color int
type Size int

const (
	green Color = iota
	blue
)

const (
	small Size = iota
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

// Code breaks the Open-Closed principle, check ocp.go file
//***********************************************************************************

// First requirement, FilterByColor
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {

	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// Second requirement, FilterBySize
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// Third requirement FilterByColorAndSize
func (f *Filter) FilterByColorAndSize(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

//***********************************************************************************

func main() {

	apple := Product{
		name:  "Apple",
		color: green,
		size:  small,
	}

	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}
	products := []Product{apple, tree, house}
	fmt.Println("Green products (old):")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Println(v.name + " - is green")
	}

	colorSpec := ColorSpecification{color: green}
	fmt.Println("Green products (new):")
	for _, v := range BetterFilterProducts(products, &colorSpec) {
		fmt.Println(v.name + " - is green")
	}

	sizeSpec := SizeSpecification{size: large}
	fmt.Println("Large products:")
	for _, v := range BetterFilterProducts(products, &sizeSpec) {
		fmt.Println(v.name + " - is large")
	}

	// Example of the double specification created
	color := ColorSpecification{color: green}
	size := SizeSpecification{size: large}
	andSpec := AndSpecification{&color, &size}
	fmt.Println("Large and green products: ")
	for _, v := range BetterFilterProducts(products, &andSpec) {
		fmt.Println(v.name + " - is large and green")
	}

}
