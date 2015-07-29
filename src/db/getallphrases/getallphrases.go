package getallphrases

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"log"
)

func GetAll(db sql.DB, locale string, themes string) map[string]struct{} {

	alloldphrases := make(map[string]struct{})

	sqlstr := "select phrase from phrases where locale='" + locale + "' and themes='" + themes + "'"

	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var phrase string
		rows.Scan(&phrase)
		alloldphrases[phrase] = struct{}{}

	}
	rows.Close()
	
	return alloldphrases 

}
