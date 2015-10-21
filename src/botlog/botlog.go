package main

import (
	"bytes"
//	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	_ "github.com/mxk/go-sqlite/sqlite3"
	"database/sql"
	"db/getallkeywords"
	"db/getallphrases"
	"db/insertnewkeywords"
	"db/insertnewphrases"
	"db/updateallkeywords"
	"db/updateallphrases"
	"domains"
	"entryHandler"
	"examkeyword"
	//		"fmt"
	"github.com/satyrius/gonx"
	"io"
	"log"
	"os"
	//	"sort"
	"sort_map_by_value"
	"strconv"
	"time"
)

func main() {

	file, err := os.Open("/home/juno/workspace/SearchGoogleLogParser/logscollector/rawallaccess.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	format := `$remote_addr $remote_user [$time_local] $host "$request" "$request_time" "$upstream_response_time" $status $body_bytes_sent "$http_referer" "$http_user_agent" "$http_x_forwarded_for" "$uid_got" "$uid_set" "$geoip_country_code"`

	reader := gonx.NewReader(file, format)

	var entryArr []domains.Entry

	newHitsEntry := entryHandler.NewEmptyHitsEntry()

	bing_newHitsEntry := entryHandler.NewEmptyHitsEntry()

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		remote_addr, _ := rec.Field("remote_addr")
		time_local, _ := rec.Field("time_local")
		host, _ := rec.Field("host")
		request, _ := rec.Field("request")
		status, _ := rec.Field("status")
		body_bytes_sent, _ := rec.Field("body_bytes_sent")
		http_referer, _ := rec.Field("http_referer")
		http_user_agent, _ := rec.Field("http_user_agent")
		geoip_country_code, _ := rec.Field("geoip_country_code")
		uid_set, _ := rec.Field("uid_set")
		uid_got, _ := rec.Field("uid_got")

		entry := domains.Entry{

			Remote_addr:        remote_addr,
			Time_local:         time_local,
			Host:               host,
			Request:            request,
			Status:             status,
			Body_bytes_sent:    body_bytes_sent,
			Http_referer:       http_referer,
			Http_user_agent:    http_user_agent,
			Geoip_country_code: geoip_country_code,
			Uid_set:            uid_set,
			Uid_got:            uid_got,
		}

		entryArr = append(entryArr, entry)

	}

	allsearchengines := make(map[string]int)

	for _, entry := range entryArr {

		searchengine := (*newHitsEntry).Handle(entry, "www.google")
		(*bing_newHitsEntry).Handle(entry, "www.bing")

		if searchengine != "" {

			_, ok := allsearchengines[searchengine]

			if ok {
				i := allsearchengines[searchengine]
				allsearchengines[searchengine] = i + 1
			} else {

				allsearchengines[searchengine] = 0

			}
		}

	}

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	alloldphrases := getallphrases.GetAll(*db, "fi_FI", "porno")
	alloldkeywords := getallkeywords.GetAll(*db, "fi_FI", "porno")

	var phrasestoUpdate []string
	var phrasestoInsert []string

	var keywodstoUpdate []domains.KeywordObj
	var keywodstoInsert []domains.KeywordObj

	for key := range entryHandler.NewQueryEntry.QueryEntry {

		_, ok := alloldphrases[key]

		if ok {

			phrasestoUpdate = append(phrasestoUpdate, key)

		} else {

			phrasestoInsert = append(phrasestoInsert, key)
		}

	}

	for key := range entryHandler.NewKewordsEntry.KeywordsEntry {

		keywordObj := examkeyword.ExamKeyword(key)

		if keywordObj.Keyword != "" && len(keywordObj.Keyword) > 2 {

			_, ok := alloldkeywords[keywordObj.Keyword]

			if ok {

				keywodstoUpdate = append(keywodstoUpdate, keywordObj)

			} else {

				keywodstoInsert = append(keywodstoInsert, keywordObj)

			}

		}

	}

	updateallphrases.UpdateAll(*db, "fi_FI", "porno", phrasestoUpdate)
	insertnewphrases.InsertAll(*db, "fi_FI", "porno", phrasestoInsert)
	updateallkeywords.UpdateAll(*db, "fi_FI", "porno", keywodstoUpdate)
	insertnewkeywords.InsertAll(*db, "fi_FI", "porno", keywodstoInsert)

	var g_total int = 0
	var b_total int = 0
	var s_total int = 0
	var buffer bytes.Buffer

	currenttime := time.Now().Local()
	buffer.WriteString("\n" + currenttime.Format("2006-01-02 15:04:05 +0300") + "\n\n")
	buffer.WriteString("google -------------------------\n\n")


	google_records := sort_map_by_value.SortByValue(newHitsEntry.Domhits)
	
	for _, record := range google_records {

		buffer.WriteString(record.Keyval + " " + strconv.Itoa(record.Valval) + "\n")
		g_total = g_total + record.Valval

	}
	
	

	buffer.WriteString("-------------------------\n")
	buffer.WriteString("g_total " + strconv.Itoa(len(newHitsEntry.Domhits)) + " " + strconv.Itoa(g_total) + "\n\n")

	buffer.WriteString("bing -------------------------\n\n")

	bing_records := sort_map_by_value.SortByValue(bing_newHitsEntry.Domhits)

	for _, record := range bing_records {

		buffer.WriteString(record.Keyval + " " + strconv.Itoa(record.Valval) + "\n")
		b_total = b_total + record.Valval

	}

	buffer.WriteString("-------------------------\n")
	buffer.WriteString("b_total " + strconv.Itoa(len(bing_newHitsEntry.Domhits)) + " " + strconv.Itoa(b_total) + "\n\n")

	searchengines_record := sort_map_by_value.SortByValue(allsearchengines)

	buffer.WriteString("\nSearchengines-----------------------\n\n")
	for _, record := range searchengines_record {

		buffer.WriteString(record.Keyval + " " + strconv.Itoa(record.Valval) + "\n")
		s_total = s_total + record.Valval

	}

	buffer.WriteString("\n-------------------------\n")
	buffer.WriteString("s_total " + strconv.Itoa(len(allsearchengines)) + " " + strconv.Itoa(s_total) + "\n\n")

	file, err = os.OpenFile("stat.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(buffer.String()); err != nil {
		panic(err)
	}

}
