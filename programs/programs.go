package programs

import (
	"encoding/json"
	"log"
	"strings"
	"text/template"


	"cc-info-toronto.org/centre"
	"cc-info-toronto.org/conf"
	"cc-info-toronto.org/http_utils"
)

/* var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.LstdFlags|log.Lshortfile)
	//logger = log.New(os.Stdout, "logger: ", log.LstdFlags|log.Lshortfile)
) */

type CentreInfo struct {
	CentreId int
	Title    string `json:"title"`
	Weeks    []struct {
		CentreId      int
		WeekStartDate string `json:"title"`
		JsonFile      string `json:"json"`
		HasPrograms   bool   `json:"hasPrograms,string"`
	} `json:"weeks"`
}

type WeeklyPrograms struct {
	Programs []struct {
		CentreId        int
		ProgramCategory string               `json:"program"`
		Days            []DayOfWeeklyProgram `json:"days"`
	} `json:"programs"`
}

type ProgramSchedulerRecordJson struct {
	CentreId        int
	WeekStartDate   string
	Sport           string
	ProgramCategory string
	Age             string
	DayOfWeek       string `json:"day"`
	ProgramTime     string `json:"title"` //json.RawMessage
	Status          string `json:"status"`
	Comment         string `json:"comment"`
	Link            string `json:"link"`
}

type DayOfWeeklyProgram struct {
	CentreId        int
	WeekStartDate   string
	ProgramCategory string `json:"program"`
	//DayOfWeek             string    `json:"day"`
	Sport   string                       `json:"title"`
	Age     string                       `json:"age"`
	Status  string                       `json:"status"`
	Comment string                       `json:"comment"`
	Link    string                       `json:"link"`
	Times   []ProgramSchedulerRecordJson `json:"times"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
 

func GetPrograms(ce chan centre.Centre, pr chan ProgramSchedulerRecordJson) {
	//logger.SetOutput(os.Stdout)

	tmpl, err := template.New("url").Parse(conf.GlobConf.CENTER_SPORT_PROGRAMS_TEMPLATE)
	check(err)
	tmpl_prog, err := template.New("url").Parse(conf.GlobConf.CENTER_SPORT_WEEK_DAILY_SCHEDULER)
	check(err)
	for centre := range ce {
		url := new(strings.Builder)
		err = tmpl.Execute(url, centre)
		check(err)

		respData, err := http_utils.ReadURL(url.String())

		check(err)

		//logger.Println("respData=", string(respData))
 

		var f CentreInfo
		err = json.Unmarshal(respData, &f)
		if err != nil {
			//log.Println("ERROR http_utils.ReadURL(\"", url, "\")")
			//log.Println("ERROR respData=", string(respData))
			//log.Println("ERROR utf8Data=", string(utf8Data))
			log.Println("Skipping centre", centre.ID, "due to error in json.Unmarshal, URL:", url.String())
			continue
		}
		//check(err)

		//logger.Println("respData=", f)

		for week_indx := range f.Weeks {
			f.Weeks[week_indx].CentreId = centre.ID
			//logger.Println("Week:", f.Weeks[week_indx])
			url_prog := new(strings.Builder)
			err = tmpl_prog.Execute(url_prog, f.Weeks[week_indx])
			check(err)
			if f.Weeks[week_indx].HasPrograms /* == "true"*/ {
				//logger.Println("http_utils.ReadURL(\"", url_prog.String(), "\")")
				respData, err := http_utils.ReadURL(url_prog.String())
				check(err)
				//logger.Println("respData=", string(respData))
				//utf8Data, err := UTF16BytesToUTF8(respData)
				//check(err)
				var wps WeeklyPrograms
				err = json.Unmarshal(respData, &wps)
				check(err)
				for wps_indx := range wps.Programs {
					wps.Programs[wps_indx].CentreId = f.Weeks[week_indx].CentreId
					//ww, _ := json.Marshal(wps.Programs[wps_indx])
					//logger.Println("WeeklyProgram=", string(ww))
					for days_indx := range wps.Programs[wps_indx].Days {
						wps.Programs[wps_indx].Days[days_indx].CentreId = wps.Programs[wps_indx].CentreId
						wps.Programs[wps_indx].Days[days_indx].WeekStartDate = f.Weeks[week_indx].WeekStartDate
						wps.Programs[wps_indx].Days[days_indx].ProgramCategory = wps.Programs[wps_indx].ProgramCategory
						for i := range wps.Programs[wps_indx].Days[days_indx].Times {
							wps.Programs[wps_indx].Days[days_indx].Times[i].CentreId = wps.Programs[wps_indx].Days[days_indx].CentreId
							wps.Programs[wps_indx].Days[days_indx].Times[i].Sport = wps.Programs[wps_indx].Days[days_indx].Sport
							wps.Programs[wps_indx].Days[days_indx].Times[i].ProgramCategory = wps.Programs[wps_indx].Days[days_indx].ProgramCategory
							wps.Programs[wps_indx].Days[days_indx].Times[i].Age = wps.Programs[wps_indx].Days[days_indx].Age
							wps.Programs[wps_indx].Days[days_indx].Times[i].WeekStartDate = wps.Programs[wps_indx].Days[days_indx].WeekStartDate

							pr <- wps.Programs[wps_indx].Days[days_indx].Times[i]

						}

					}
				}

			}

		}
	}
	close(pr)

}
