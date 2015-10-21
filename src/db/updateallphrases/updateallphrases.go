package updateallphrases

import (
	
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func UpdateAll(db sql.DB, locale string, themes string, phrases []string) {

	fmt.Println("Need Update ", len(phrases), "phrases")

	tx, err := db.Begin()
	if err != nil {

		log.Fatal(err)
		os.Exit(1)
	}

	//	timenow := int64(time.Now().Unix())*1000

	timenow := int64(time.Now().UnixNano()) / 1000000

	for _, phrase := range phrases {

		fmt.Println(phrase)
		stmt, err := tx.Prepare("update phrases set Hits=Hits+1,Updated=? where phrase=?")
		if err != nil {
			log.Fatal(err)

		}
		defer stmt.Close()
		if _, err = stmt.Exec(timenow, phrase); err != nil {
			log.Fatal(err)
			os.Exit(1)

		}

	}

	tx.Commit()

}
