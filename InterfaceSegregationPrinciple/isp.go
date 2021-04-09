package main

import (
	"fmt"
)

type Printer interface {
	Print(Document)
}

type Scanner interface {
	Scan(Document)
}

type Fax interface {
	Fax(Document)
}

// OldPrinter can only print
type OldPrinter struct {
}

func (o *OldPrinter) Print(document Document) {
	fmt.Println(document)
}

// NewFax can only scan
type NewFax struct {
}

func (n *NewFax) Fax(document Document) {
	fmt.Println(document)
}

// NewPrinter can either print or scan
type NewPrinter struct {
}

func (n *NewPrinter) Print(document Document) {
	fmt.Println(document)
}

func (n *NewPrinter) Scan(document Document) {
	fmt.Println("Thisscan")
	fmt.Println(document)
}

// We can combine two interfaces into a single interface such as io.ReadWriter

type MultiFunctionalDevice interface {
	Scanner
	Fax
}

type NewEraDevice struct {
	scanner Scanner
	fax     Fax
}

func (n *NewEraDevice) Scan(document Document) {
	fmt.Println("scan")
	n.scanner.Scan(document)
}

func (n *NewEraDevice) Fax(document Document) {
	n.fax.Fax(document)
}
