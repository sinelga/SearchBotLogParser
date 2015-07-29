package entryHandler

import (
	"domains"
	"examquery"
	//	"fmt"
	"net/url"
	"strings"
)

type Hits map[string]int

type Entrydomain struct {
	Domhits Hits
}

func NewEmptyHitsEntry() *Entrydomain {
	return &Entrydomain{make(Hits)}
}

var domain string

var NewQueryEntry = examquery.NewQuerys()
var NewKewordsEntry = examquery.NewKeywords()

func (hitsentry *Entrydomain) Handle(entry domains.Entry,bot string) string {

	var searchengine string = ""
	var host string

	if entry.Http_referer != "-" {

		u, err := url.Parse(entry.Http_referer)
		if err != nil {
			panic(err)
		}
		host = u.Host

		if host != entry.Host && entry.Geoip_country_code == "FI" && entry.Status == "200" {
			searchengine = host
		}

		m, _ := url.ParseQuery(u.RawQuery)

		searchstr := loopForBestQ(m, "q")

		if searchstr == "" {

			searchstr = loopForBestQ(m, "p")

		}
		if searchstr == "" {

			searchstr = loopForBestQ(m, "query")

		}

		if len(searchstr) > 2 {

			(*NewQueryEntry).Exam(searchstr)
			(*NewKewordsEntry).Exam(searchstr) 

		}

	}

	//	fmt.Println("Host->",host)

	//	if strings.HasPrefix(entry.Http_referer, "http://www.google") && entry.Geoip_country_code == "FI" && (strings.Index(entry.Request, ".html") > 0 || strings.HasPrefix(entry.Request, "GET / ")) {
	//	if strings.Index(entry.Http_referer, "www.google") >-1 &&  entry.Geoip_country_code == "FI" && entry.Status=="200" {
	if strings.HasPrefix(host, bot) && entry.Geoip_country_code == "FI" && entry.Status == "200" {

		//	if strings.HasPrefix(entry.Http_referer, "http://www.bing.com") && entry.Geoip_country_code=="FI" && (strings.Index(entry.Request,".html") > 0 || strings.HasPrefix(entry.Request,"GET / ")) {

		hostsplit := strings.Split(entry.Host, ".")

		if len(hostsplit) > 2 {

			domain = hostsplit[len(hostsplit)-2] + "." + hostsplit[len(hostsplit)-1]
		} else {

			domain = entry.Host

		}

		value, ok := hitsentry.Domhits[domain]

		if !ok {

			hitsentry.Domhits[domain] = 0

		} else {

			hitsentry.Domhits[domain] = value + 1

		}

	}
	return searchengine
}

func loopForBestQ(allRawQuery map[string][]string, search_key string) string {

	var searchstr string

	for key, value := range allRawQuery {

		if key == search_key {

			if len(value[0]) > 2 {

				searchstr = value[0]

			}

		}

	}

	return searchstr
}
