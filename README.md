# Benchmark Test for MongoDB and JSONB (pg)
This benchmark test compares the performance of MongoDB and JSONB when inserting, updating, reading and deleting documents.

### Test Setup
The test was performed on a machine with the following specifications:

* CPU: intel core i3 9100f
* RAM: 16GB DDR4
* Storage: 256GB SSD
* OS: Fedora 37

### Test Results
The test results are shown in the table below:

| Test Name       | Number of Documents | Time (ns/op)         | Memory (B/op) | Allocs/op |
|-----------------|---------------------|----------------------|----------------|-----------|
| BenchmarkMongo1 | 1                   | 0.03156 ns/op        | 0 B/op         | 0         |
| BenchmarkMongo100 | 100                | 0.4614 ns/op         | 0 B/op         | 0         |
| BenchmarkMongo200 | 200                | 0.9224 ns/op         | 0 B/op         | 0         |
| BenchmarkMongo500 | 500                | 2076665187 ns/op    | 15726880 B/op  | 251912    |
| BenchmarkJSONB1 | 1                   | 0.01976 ns/op        | 0 B/op         | 0         |
| BenchmarkJSONB100 | 100                | 204185849 ns/op      | 136963 B/op    | 2044      |
| BenchmarkJSONB200 | 200                | 2617058157 ns/op     | 1292336 B/op   | 20323     |
| BenchmarkJSONB500 | 500                | 6457528499 ns/op     | 3113936 B/op   | 50623     |

* MongoDB tests took about 0.03 seconds for 1 document, 0.46 seconds for 100 documents, and 0.92 seconds for 200 documents. 
* JSONB tests took about 0.019 seconds for 1 document, 0.204 seconds for 100 documents, and 2.61 seconds for 200 documents.

1. This is just one way to measure the performance of these technologies, and the results may vary depending on the specific use case and hardware being used. Other factors such as the complexity of the queries, the amount of data being stored and retrieved, and the number of concurrent requests being made can also affect performance.
2. Maybe JSONB is more performant when working with big data, but MongoDB is more performant when dealing with small data. It depends on your use case and the amount of data that you need to manage.






## OLD TEST(MariaDB vs Postgres) RESULT:
### mariadb
hard 136 mb
mem 335 mb

### postgresql
hard 80
mem 262

### posgres first: 20s
BenchmarkPostgres50-4   1000000000  0.3320      ns/op  0       B/op  0        allocs/op
BenchmarkPostgres100-4  6128        172249      ns/op  23      B/op  0        allocs/op
BenchmarkPostgres200-4  1           1545513570  ns/op  256448  B/op  6906     allocs/op

### postgres second: 8s
BenchmarkPostgres50-4   1000000000  0.3794      ns/op  0       B/op  0        allocs/op
BenchmarkPostgres100-4  21          51429318    ns/op  6835    B/op  171      allocs/op
BenchmarkPostgres200-4  1           1406782021  ns/op  256752  B/op  6909     allocs/op

mem after 276

### mariadb first: 18s
BenchmarkPostgres50-4   1000000000  0.3135      ns/op  0       B/op  0        allocs/op
BenchmarkPostgres100-4  4622        220736      ns/op  24      B/op  0        allocs/op
BenchmarkPostgres200-4  1           1264586084  ns/op  213152  B/op  6875     allocs/op

### mariadb second: 15s
BenchmarkPostgres50-4   1000000000  0.3150      ns/op  0       B/op  0        allocs/op
BenchmarkPostgres100-4  271         3899507     ns/op  422     B/op  12       allocs/op
BenchmarkPostgres200-4  1           1576092967  ns/op  212048  B/op  6866     allocs/op

mem after 276