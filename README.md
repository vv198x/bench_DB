#RESULT:
###mariadb
######hard 136 mb
######mem 335 mb

###postgresql
######hard 80
######mem 262

###posgres first: 20s
######BenchmarkPostgres50-4   1000000000  0.3320      ns/op  0       B/op  0        allocs/op
######BenchmarkPostgres100-4  6128        172249      ns/op  23      B/op  0        allocs/op
######BenchmarkPostgres200-4  1           1545513570  ns/op  256448  B/op  6906     allocs/op

###postgres second: 8s
######BenchmarkPostgres50-4   1000000000  0.3794      ns/op  0       B/op  0        allocs/op
######BenchmarkPostgres100-4  21          51429318    ns/op  6835    B/op  171      allocs/op
######BenchmarkPostgres200-4  1           1406782021  ns/op  256752  B/op  6909     allocs/op

###mem after 276

###mariadb first: 18s
######BenchmarkPostgres50-4   1000000000  0.3135      ns/op  0       B/op  0        allocs/op
######BenchmarkPostgres100-4  4622        220736      ns/op  24      B/op  0        allocs/op
######BenchmarkPostgres200-4  1           1264586084  ns/op  213152  B/op  6875     allocs/op

###mariadb second: 15s
######BenchmarkPostgres50-4   1000000000  0.3150      ns/op  0       B/op  0        allocs/op
######BenchmarkPostgres100-4  271         3899507     ns/op  422     B/op  12       allocs/op
######BenchmarkPostgres200-4  1           1576092967  ns/op  212048  B/op  6866     allocs/op

###mem after 276
