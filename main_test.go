package main

import (
	"context"
	"encoding/json"
	"github.com/go-pg/pg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
	"time"
)

func ConnectMongoDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(`mongodb://root:rootpassword@192.168.122.161:27017/`))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func benchmarkMongoDB(b *testing.B, lim int) {
	client := ConnectMongoDB()
	defer client.Disconnect(context.TODO())

	collection := client.Database("test").Collection("benchmark")
	var result bson.M

	for i := 0; i < lim; i++ {
		collection.InsertOne(context.TODO(), bson.M{
			"name":  "test",
			"value": "benchmark",
			"attachments": []bson.M{
				{
					"name": "attachment1.jpg",
					"type": "image/jpeg",
					"data": "base64 encoded image data",
				},
				{
					"name": "attachment2.pdf",
					"type": "application/pdf",
					"data": "base64 encoded pdf data",
				},
			},
			"array": []bson.M{
				{
					"name":  "item1",
					"value": "value1",
				},
				{
					"name":  "item2",
					"value": "value2",
				},
			},
		})

		collection.UpdateOne(context.TODO(), bson.M{"name": "test"}, bson.M{"$set": bson.M{"value": "updated"}})

		collection.UpdateOne(context.TODO(),
			bson.M{"name": "test"},
			bson.M{
				"$push": bson.M{
					"array": bson.M{"name": "upd", "value": "updated_value"},
				},
			})

		collection.FindOne(context.TODO(), bson.M{"name": "test"}).Decode(&result)

		collection.DeleteOne(context.TODO(), bson.M{"name": "test"})
	}

}

func BenchmarkMongo1(b *testing.B) {
	benchmarkMongoDB(b, 1)
}
func BenchmarkMongo100(b *testing.B) {
	benchmarkMongoDB(b, 100)
}
func BenchmarkMongo200(b *testing.B) {
	benchmarkMongoDB(b, 200)
}
func BenchmarkMongo500(b *testing.B) {
	benchmarkMongoDB(b, 500)
}

func BenchmarkHold(b *testing.B) {
	time.Sleep(time.Second * 2)
}

type Benchmark struct {
	ID          int
	Name_       string
	Value_      string
	Attachments json.RawMessage
	Array_      json.RawMessage
}

func init() {

	connectPG().Exec(`CREATE TABLE IF NOT EXISTS benchmarks(
    id SERIAL PRIMARY KEY,
    name_ VARCHAR(255) NOT NULL,
    value_ VARCHAR(255) NOT NULL,
    attachments JSONB NOT NULL,
    array_ JSONB NOT NULL
);`)
}
func connectPG() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     "192.168.122.161:5432",
		User:     "postgres",
		Password: "password",
		Database: "database",
	})
}

func benchmarkJSONB(b *testing.B, lim int) {
	db := connectPG()
	defer db.Close()
	var result Benchmark

	for i := 0; i < lim; i++ {

		db.Model(&Benchmark{
			Name_:       "test",
			Value_:      "benchmark",
			Attachments: json.RawMessage(`[{"name":"attachment1.jpg","type":"image/jpeg","data":"base64 encoded image data"},{"name":"attachment2.pdf","type":"application/pdf","data":"base64 encoded pdf data"}]`),
			Array_:      json.RawMessage(`[{"name":"item1","value":"value1"},{"name":"item2","value":"value2"}]`),
		}).Insert()

		db.Model(&Benchmark{}).Set("value = ?", "updated").Where("name = ?", "test").Update()

		db.Exec(`UPDATE benchmarks SET array_ = jsonb_set(array_, '{1,name}', '"upd"');`)
		//db.Model(&Benchmark{}).Set(`array_ = jsonb_set(array_, '{1,name}', 'upd')`).Update()

		db.Model(&result).Where("name = ?", "test").Select()

		db.Model(&result).Where("name = ?", "test").Delete()

	}

}

func BenchmarkJSONB1(b *testing.B) {
	benchmarkJSONB(b, 1)
}
func BenchmarkJSONB100(b *testing.B) {
	benchmarkJSONB(b, 100)
}
func BenchmarkJSONB200(b *testing.B) {
	benchmarkJSONB(b, 200)
}
func BenchmarkJSONB500(b *testing.B) {
	benchmarkJSONB(b, 500)
}

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


func benchmarkMongoDB(b *testing.B, lim int) {
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
