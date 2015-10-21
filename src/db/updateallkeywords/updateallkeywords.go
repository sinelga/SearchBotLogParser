package updateallkeywords

import (
	
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"domains"
	"fmt"
	"log"
	"os"
	"time"
)

func UpdateAll(db sql.DB, locale string, themes string, keywords []domains.KeywordObj) {

	fmt.Println("Need Update ", len(keywords), "keywords")

	tx, err := db.Begin()
	if err != nil {

		log.Fatal(err)
		os.Exit(1)
	}

	timenow := int64(time.Now().UnixNano()) / 1000000

	for _, keywordObj := range keywords {

		//		fmt.Println(phrase)
		stmt, err := tx.Prepare("update keywords set Hits=Hits+1,Updated=? where Keyword=?")
		if err != nil {
			log.Fatal(err)

		}
		defer stmt.Close()
		if _, err = stmt.Exec(timenow, keywordObj.Keyword); err != nil {
			log.Fatal(err)
			os.Exit(1)

		}

	}

	tx.Commit()

}
