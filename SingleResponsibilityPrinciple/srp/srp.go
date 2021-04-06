package srp

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	// Map may be a better example
	Entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.Entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.Entries = append(j.Entries, entry)
	return entryCount
}

// RemoveEntry returns true if the change could be done, otherwise, false is thrown if
// there is an error
func (j *Journal) RemoveEntry(index int) bool {
	if index+1 >= len(j.Entries) {
		return false
	}
	j.Entries[index-1] = j.Entries[len(j.Entries)-1]
	j.Entries = j.Entries[:len(j.Entries)-1]
	return true
}

// RemoveEntryOrdered returns true if the change could be done, otherwise, false is thrown if
// there is an error, RemoveEntryOrdered modifies the entries slice in order to get an ordered slice in the end
func (j *Journal) RemoveEntryOrdered(index int) bool {
	if index+1 >= len(j.Entries) {
		return false
	}
	j.Entries = append(j.Entries[0:index-1], j.Entries[index:]...)
	return true
}

// Separation of concerns , Journal should not be in charge of also saving and loading the file
// Breaking the single responsibility principle, separate component or separate package (see persistence file)

// **********************************************************

/*
type Persistence struct {
	lineSeparator string
}

*/

var LineSeparator = "\n"

func SaveToFile(j *Journal, fileName string) {
	_ = ioutil.WriteFile(fileName, []byte(strings.Join(j.Entries, LineSeparator)), 0644)
}

func (j *Journal) Save(fileName string) {
	_ = ioutil.WriteFile(fileName, []byte(j.String()), 0644)
}

func (j *Journal) Load() {
	panic("Implementation")
}

func (j *Journal) LoadFromWeb() {
	panic("Implementation")
}

// **********************************************************
