package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/godror/godror"
)

type Angestellter struct {
	id                int
	Name              string
	Einstellungsdatum time.Time
}

func main() {
	db, err := sql.Open("godror", "app/app@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=172.17.0.1)(PORT=1521)))(CONNECT_DATA=(SID=ORCLCDB)))")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(0)
	defer db.Close()

	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true})
	stmt, err := tx.Prepare("INSERT INTO EMP (EMPNO, ENAME, HIREDATE) VALUES(:id1, :name2, :date3)")
	if err != nil {
		log.Fatalf("Fehler bei Statement: %v", err)
	}
	for i := 10; i < 100; i++ {
		_, err = stmt.Exec(
			sql.Named("id1", i),
			sql.Named("name2", "JOE"),
			sql.Named("date3", time.Now()))

		if err != nil {
			log.Fatalf("Fehler bei Exec: %v", err)
		}
	}
	defer stmt.Close()
	tx.Commit()

}
