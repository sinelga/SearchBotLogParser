package insertnewphrases

import (
	
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func InsertAll(db sql.DB, locale string, themes string, phrases []string) {

	fmt.Println("Need Insert ", len(phrases), "phrases")

	tx, err := db.Begin()
	if err != nil {

		log.Fatal(err)
		os.Exit(1)
	}

	//	timenow := int64(time.Now().Unix())*1000

	timenow := int64(time.Now().UnixNano()) / 1000000
	for _, phrase := range phrases {

//		fmt.Println("Insert ", phrase)
		stmt, err := tx.Prepare("insert into phrases('Locale','Themes','Phrase','Created','Updated')  values(?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)

		}
		defer stmt.Close()
		if _, err = stmt.Exec(locale, themes, phrase, timenow, timenow); err != nil {
			log.Fatal(err)
			os.Exit(1)

		}

	}

	tx.Commit()

}
