package domains

import ()

type Record struct {
	Keyval string
	Valval int
}

type Entry struct {
	Remote_addr        string
	Time_local         string
	Host               string
	Request            string
	Status             string
	Body_bytes_sent    string
	Http_referer       string
	Http_user_agent    string
	Geoip_country_code string
	Uid_set            string
	Uid_got            string
}

type KeywordObj struct {
	
	Keyword string
	Cl int
	Lang string
		
}