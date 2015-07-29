package domain_hits

import (
//	"domains"

)
type Hits map[string]int


type Entrydomain struct {
	Domhits Hits
}


func NewEmptyEntry() *Entrydomain {
	return &Entrydomain{make(Hits)}
}

func (entry *Entrydomain) Hits(domain string) {

	value, ok := entry.Domhits[domain]
	if !ok {

		entry.Domhits[domain] = 0

	} else {

		entry.Domhits[domain] = value + 1

	}

}