package main

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

/*
CREATE TABLE `tickets` (
`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
`subdomain_id` int(11) NOT NULL,
`subject` varchar(255) NOT NULL DEFAULT ”,
`state` varchar(255) NOT NULL DEFAULT 'open',
PRIMARY KEY (`id`)
)

CREATE TABLE "tickets" (

	id int check (id > 0) NOT NULL ,
	  subdomain_id int NOT NULL,
	  subject varchar(255) NOT NULL DEFAULT '',
	  state varchar(255) NOT NULL DEFAULT 'open',
	  PRIMARY KEY (id)
	  )
*/
var multip = 10000

func bBenchmarkSimplest(b *testing.B) {
	b.ResetTimer()
	db, err := sql.Open("mysql", os.Getenv("MARIADB")+"/test")
	if err != nil {
		log.Fatal(err)
	}
	for n := 0; n < b.N*multip; n++ {
		rows, err := db.Query("select id, subject, state from tickets where subdomain_id = ? and (state = ? or state = ?) limit 1", 1, "open", "spam")
		if err != nil {
			log.Fatalln(err)
		}
		var id int
		var subject, state string
		for rows.Next() {
			err := rows.Scan(&id, &subject, &state)
			//log.Println(id, subject, state)
			if err != nil {
				log.Fatalln(err)
			}
		}
		rows.Close()
	}
	db.Close()
}

func BenchmarkSimplestPostgre(b *testing.B) {
	b.ResetTimer()
	db, err := sql.Open("postgres", "host=192.168.122.111 user=sql password=test dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	for n := 0; n < b.N*multip; n++ {
		rows, err := db.Query("select id, subject, state from TICKETS limit 1")
		if err != nil {
			log.Fatalln(err)
		}
		var id int
		var subject, state string
		for rows.Next() {
			err := rows.Scan(&id, &subject, &state)
			//log.Println(id, subject, state)
			if err != nil {
				log.Fatalln(err)
			}
		}
		rows.Close()
	}
	db.Close()
}
