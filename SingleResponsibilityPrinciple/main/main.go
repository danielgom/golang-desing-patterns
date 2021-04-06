package main

import (
	"fmt"
	"github.com/danielgom/golang-desing-patterns/SingleResponsibilityPrinciple/srp"
	"github.com/danielgom/golang-desing-patterns/SingleResponsibilityPrinciple/srp/persistence"
)

func main() {
	journal := srp.Journal{Entries: []string{}}
	fmt.Println(journal.Entries)
	journal.AddEntry("first")
	journal.AddEntry("second")
	journal.AddEntry("third")
	journal.AddEntry("fourth")
	journal.AddEntry("fifth")
	journal.AddEntry("sixth")
	journal.AddEntry("seventh")
	journal.AddEntry("eighth")
	fmt.Printf("%#v\n", journal.Entries)
	journal.RemoveEntryOrdered(5)
	fmt.Printf("%#v\n", journal.Entries)
	journal.RemoveEntry(5)
	fmt.Printf("%#v\n", journal.Entries)

	// We either create another struct and call the method or simply call another function that is not a part of the
	// Journal methods

	// Struct and method approach
	per := persistence.Persistence{LineSeparator: "\n"}
	per.SaveToFile(&journal, "this_file.txt")

	// Struct function approach srp is the package the function is in, not a method
	srp.SaveToFile(&journal, "this_file.txt")

}
