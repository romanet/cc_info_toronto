package conf

type Conf struct {
	DEFAULT_DB_FILE                   string
	CENTERS_LISTING_URL               string
	CENTER_SPORT_PROGRAMS_TEMPLATE    string
	CENTER_SPORT_WEEK_DAILY_SCHEDULER string
}

var GlobConf = Conf{
	DEFAULT_DB_FILE:                   "cc-info-toronto.db",
	CENTERS_LISTING_URL:               "https://www.toronto.ca/data/parks/live/locations/centres.json",
	CENTER_SPORT_PROGRAMS_TEMPLATE:    "https://www.toronto.ca/data/parks/live/locations/{{ .ID }}/sports/info.json",
	CENTER_SPORT_WEEK_DAILY_SCHEDULER: "https://www.toronto.ca/data/parks/live/locations/{{ .CentreId }}/sports/{{ .JsonFile }}",
}
