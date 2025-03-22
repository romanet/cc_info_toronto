package db_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

type ProgramSchedulerRecord struct {
	CentreId        int    `db:"centre_id"`
	WeekStartDate   string `db:"week_start_date"`
	Sport           string `db:"sport"`
	ProgramCategory string `db:"program_category"`
	Age             string `db:"age"`
	DayOfWeek       string `db:"day_of_week"`
	ProgramTime     string `db:"program_time"`
	Status          string `db:"status"`
	Comment         string `db:"comment"`
	Link            string `db:"link"`
}

type CentreRecord struct {
	ID       int     `db:"id"`
	Name     string  `db:"name"`
	Address  string  `db:"address"`
	Phone    string  `db:"phone"`
	District string  `db:"district"`
	X        string  `db:"x"`
	Y        string  `db:"y"`
	Wifi     int     `db:"wifi"`
	Lng      float64 `db:"lng"`
	Lat      float64 `db:"lat"`
}

var (
	DbSchema = `CREATE TABLE IF NOT EXISTS program_scheduler_records (
		centre_id INTEGER,
		week_start_date TEXT,
		sport TEXT,
		program_category TEXT,
		age TEXT,
		day_of_week TEXT,
		program_time TEXT,
		status TEXT,
		comment TEXT,
		link TEXT
	);
	`
	DbSchemaCentres = `CREATE  TABLE IF NOT EXISTS centres (
		id INTEGER,
		name TEXT,
		address TEXT,
		phone TEXT,
		district TEXT,
		x TEXT,
		y TEXT,
		wifi INTEGER,
		lng TEXT,
		lat TEXT
	);

	`
	DbInsert = `INSERT INTO 
	program_scheduler_records (centre_id, week_start_date, sport, program_category, age, day_of_week, program_time, status, comment,link)
	values                   (:centre_id,:week_start_date,:sport,:program_category,:age,:day_of_week,:program_time,:status,:comment,:link)
	`

	DbInsertCentres = `INSERT INTO 
	centres (id,name,address,phone,district,x,y,wifi,lng,lat)
	values (:id,:name,:address,:phone,:district,:x,:y,:wifi,:lng,:lat)
	`
)

type Db struct {
	Conn *sqlx.DB
}

func NewDb(dbPath string) Db {
	mydb, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return Db{
		Conn: mydb,
	}
}

func (db Db) InitDB() {

	db.Conn.Exec(DbSchema)
	db.Conn.Exec(DbSchemaCentres)

	db.Conn.Exec("truncate table centres")
	db.Conn.Exec("truncate table program_scheduler_records")
}

func (db Db) AddCentresRecords(tmp []CentreRecord) {
	_, err := db.Conn.NamedExec(DbInsertCentres, tmp)
	if err != nil {
		panic(err)
	}
}

func (db Db) AddProgramSchedulerRecords(tmp []ProgramSchedulerRecord) {
	for i := range len(tmp) {
		_, err := db.Conn.NamedExec(DbInsert, tmp[i])
		if err != nil {
			logger.Println("ERROR:", DbInsert)
			vv, _ := json.Marshal(tmp[i])
			logger.Println("ERROR:", i, ")", string(vv))
			panic(err)
		}
	}

}

func (db Db) Save() {
	err := db.Conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
