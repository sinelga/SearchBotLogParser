package insertnewkeywords

import (
	
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"domains"
	"fmt"
	"log"
	"os"
	"time"
)

func InsertAll(db sql.DB, locale string, themes string, keywords []domains.KeywordObj) {

	fmt.Println("Need Insert ", len(keywords), "keywords")

	tx, err := db.Begin()
	if err != nil {

		log.Fatal(err)
		os.Exit(1)
	}


	timenow := int64(time.Now().UnixNano()) / 1000000
	for _, keywordObj := range keywords {

		stmt, err := tx.Prepare("insert into keywords('Locale','Themes','Keyword','Created','Updated','Cl','Lang')  values(?,?,?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)

		}
		defer stmt.Close()
		if _, err = stmt.Exec(locale, themes, keywordObj.Keyword, timenow, timenow, keywordObj.Cl, keywordObj.Lang); err != nil {
			log.Fatal(err)
			os.Exit(1)

		}

	}

	tx.Commit()

}
