package main

// Interface segregation principle, we should not put too much to a single interface
// better way to solve it is to break up the big interface into smaller interfaces

type Document struct {
}

type Machine interface {
	Print(Document)
	Fax(Document)
	Scan(Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(document Document) {
}

func (m *MultiFunctionPrinter) Fax(document Document) {
}

func (m *MultiFunctionPrinter) Scan(document Document) {
}

// Here we have a small problem, Old printers do not have the functionality of either Scan or Fax, but we are implementing
// those anyway and panic them or making those methods deprecated may not be the best solution, check the isp file
// for the solution
type OldFashionPrinter struct {
}

func (o *OldFashionPrinter) Print(document Document) {
	// ok
}

func (o *OldFashionPrinter) Fax(document Document) {
	panic("Operation not supported")
}

func (o *OldFashionPrinter) Scan(document Document) {
	panic("Operation not supported")
}

func main() {

	// We can use the old printer as a Printer and thats it
	document := Document{}
	var p Printer = &OldPrinter{}
	p.Print(document)

	var device MultiFunctionalDevice = &NewEraDevice{
		scanner: &NewPrinter{},
		fax:     &NewFax{},
	}
	device.Scan(document)

}
