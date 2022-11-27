package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
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

func benchmarkMariaDB(b *testing.B, lim int) {
	b.ResetTimer()
	db, err := sql.Open("mysql", os.Getenv("MARIADB")+"/test")
	if err != nil {
		log.Fatal(err)
	}
	for n := 1; n <= lim; n++ {
		st := fmt.Sprintf("INSERT into tickets(id, subdomain_id,subject,state) VALUES (%v,%v,'1','1')", n, n)
		_, err := db.Exec(st)
		rows, err := db.Query("select id, subject, state from tickets")
		if err != nil {
			fmt.Println(err)
		}
		var id int
		var subject, state string
		for rows.Next() {
			err := rows.Scan(&id, &subject, &state)
			st = fmt.Sprintf("DELETE FROM tickets WHERE id = %v ", id)
			_, err = db.Exec(st)
			if err != nil {
				fmt.Println(err)

			}
		}
		rows.Close()
	}
	db.Close()
}

/*
func benchmarkSimplestPostgres(b *testing.B, lim int) {
	b.ResetTimer()
	db, err := sql.Open("postgres", "host=192.168.122.111 user=sql password=test dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	for n := 1; n <= lim; n++ {
		st := fmt.Sprintf("INSERT into tickets(id, subdomain_id,subject,state) VALUES (%v,%v,'1','1')", n, n)
		_, err := db.Exec(st)
		rows, err := db.Query("select id, subject, state from tickets")
		if err != nil {
			fmt.Println(err)
		}
		var id int
		var subject, state string
		for rows.Next() {
			err := rows.Scan(&id, &subject, &state)
			db.Exec("DELETE from tickets WHERE id=$1", id)
			if err != nil {
				log.Fatalln(err)
			}
		}
		rows.Close()
	}
	db.Close()
}
func BenchmarkPostgres50(b *testing.B) {
	benchmarkSimplestPostgres(b, 50)
}
func BenchmarkPostgres100(b *testing.B) {
	benchmarkSimplestPostgres(b, 100)
}
func BenchmarkPostgres200(b *testing.B) {
	benchmarkSimplestPostgres(b, 200)
}
*/

func BenchmarkPostgres50(b *testing.B) {
	benchmarkMariaDB(b, 50)
}
func BenchmarkPostgres100(b *testing.B) {
	benchmarkMariaDB(b, 100)
}
func BenchmarkPostgres200(b *testing.B) {
	benchmarkMariaDB(b, 200)
}
