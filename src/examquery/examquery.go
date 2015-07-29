package examquery

import (
	
	"strings"
)


type Querys map[string]struct{}

type Entryquerys struct {
	QueryEntry Querys
}

func NewQuerys() *Entryquerys {

	return &Entryquerys{make(Querys)}

}

type Keywords map[string]struct{}

type Entrykeywords struct {
	KeywordsEntry Keywords 
}

func NewKeywords() *Entrykeywords {
	
	return &Entrykeywords{make(Keywords)}
	
}


func (newQuerys *Entryquerys) Exam(query string) {

	if len(strings.Split(query, " ")) > 1 {

		newQuerys.QueryEntry[strings.ToLower(query)] = struct{}{}

	}


}


func (newKeywords *Entrykeywords) Exam(query string) {

	keywordsarr :=strings.Split(query, " ")

	if len(keywordsarr) <= 1 {

		newKeywords.KeywordsEntry [strings.ToLower(query)] = struct{}{}

	} else {
		
		for _,keyword := range keywordsarr {
			
			newKeywords.KeywordsEntry [strings.ToLower(keyword)] = struct{}{}
					
		}
				
		
	}


}



