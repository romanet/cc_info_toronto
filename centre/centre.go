package centre

//module  cc_info_toronto/center

import (
	"encoding/json"
	"log"
	"os"

	"cc-info-toronto.org/conf"
	"cc-info-toronto.org/db_utils"
	"cc-info-toronto.org/http_utils"
)

var (
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
)

type Centre struct {
	ID       int     `json:"ID"`
	Name     string  `json:"Name"`
	Address  string  `json:"Address"`
	Phone    string  `json:"Phone"`
	District string  `json:"District"`
	X        string  `json:"X"`
	Y        string  `json:"Y"`
	Wifi     int     `json:"wifi"`
	Lng      float64 `json:"lng,string,omitempty"`
	Lat      float64 `json:"lat,string,omitempty"`
}

type Centres struct {
	All []Centre `json:"all"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveCentresToDb(db db_utils.Db, centres []db_utils.CentreRecord) {
	db.AddCentresRecords(centres)
}

func GetCenters(ce chan Centre, db db_utils.Db) {
	respData, err := http_utils.ReadURL(conf.GlobConf.CENTERS_LISTING_URL)

	if err != nil {
		panic(err)
	}

	var f Centres
	err = json.Unmarshal(respData, &f)
	check(err)
	var ceRecords []db_utils.CentreRecord

	for _, v := range f.All {
		tmp := (*db_utils.CentreRecord)(&v)

		ceRecords = append(ceRecords, *tmp)
		ce <- v

		//if i > 4 {			break		} // for debigging
	}
	logger.Println("SaveCentresToDb - start")
	SaveCentresToDb(db, ceRecords)
	logger.Println("SaveCentresToDb - end")
	close(ce)

}
