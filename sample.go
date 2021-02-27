import (
	"database/sql"
	"log"
	"os"
	"time"
)

func badGetTimeOfDay() {
	now := time.Now()
	if now.Hour >= 0 && now.Hour < 6 {
		return "Night"
	}
	if now.Hour >= 6 && now.Hour < 12 {
		return "Morning"
	}
	if now.Hour >= 12 && now.Hour < 18 {
		return "Afternoon"
	}
	return "Evening"
}

func betterGetTimeOfDay(now Time) {
	if now.Hour >= 0 && now.Hour < 6 {
		return "Night"
	}
	if now.Hour >= 6 && now.Hour < 12 {
		return "Morning"
	}
	if now.Hour >= 12 && now.Hour < 18 {
		return "Afternoon"
	}
	return "Evening"
}

var db *sql.DB

func badWriteUserName(id int) {
	var name string
	err := db.Query("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	f, err = os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(name)
	if err != nil {
		log.Fatal(err)
	}
}

type myDB struct {
	db *sql.DB
}

func newMyDB(db *sql.DB) *UserRepo {
	return &myDB{
		db: db,
	}
}

func (db *myDB) betterWriteUserName(id int, file string) {
	var name string
	err := db.Query("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	f, err = os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(name)
	if err != nil {
		log.Fatal(err)
	}
}
