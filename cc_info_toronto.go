package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"cc-info-toronto.org/centre"
	"cc-info-toronto.org/conf"
	"cc-info-toronto.org/db_utils"
	"cc-info-toronto.org/programs"
)

var (
	//buf    bytes.Buffer
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
)

func main() {
	defaultDbFilePath, _ := filepath.Abs(conf.GlobConf.DEFAULT_DB_FILE)
	var dbPath = flag.String("db-path", defaultDbFilePath, "path to DuckDB file")
	flag.Parse()
	absDbPath, _ := filepath.Abs(*dbPath)
	logger.Println("Database path", absDbPath)

	db := db_utils.NewDb(absDbPath)
	db.InitDB()

	var ce = make(chan centre.Centre)
	var pr = make(chan programs.ProgramSchedulerRecordJson)
	go centre.GetCenters(ce, db)
	go programs.GetPrograms(ce, pr)

	var PrArray []db_utils.ProgramSchedulerRecord
	for v := range pr {
		tmp := (*db_utils.ProgramSchedulerRecord)(&v)
		PrArray = append(PrArray, *tmp)
	}
	logger.Println("Before Saving Programs")
	db.AddProgramSchedulerRecords(PrArray)
	logger.Println("Saved Programs")
	logger.Println("Done")

}
