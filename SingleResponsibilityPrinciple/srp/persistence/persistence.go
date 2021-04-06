package persistence

import (
	"github.com/danielgom/golang-desing-patterns/SingleResponsibilityPrinciple/srp"
	"io/ioutil"
	"strings"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *srp.Journal, fileName string) {
	err := ioutil.WriteFile(fileName, []byte(strings.Join(j.Entries, p.LineSeparator)), 0644)
	if err != nil {
		panic(err)
	}
}
