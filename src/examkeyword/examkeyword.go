package examkeyword

import (
//	"fmt"
	"regexp"
	"domains"
)

func ExamKeyword(keyword string)  domains.KeywordObj{


	var keywordObj domains.KeywordObj

	rp := regexp.MustCompile(`[[:punct:]]`)
	rp2 := regexp.MustCompile(`[[:alpha:]]`)
	rp3 := regexp.MustCompile(`[äöøæ]`)

	if rp.MatchString(keyword) {

	} else {

		if rp2.MatchString(keyword) {

			if rp3.MatchString(keyword) {

//				fmt.Println("FIN", keyword)
				
				keywordObj = domains.KeywordObj{
					Keyword: keyword,
					Cl: 1,
					Lang: "UND",
				}
				

			} else {

//				fmt.Println("Cl", keyword)
					keywordObj = domains.KeywordObj{
					Keyword: keyword,
					Cl: 0,
					Lang: "UND",
				}
				

			}

		}

	}
return keywordObj

}
