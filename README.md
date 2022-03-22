# CaseStudyGolang
Persist data from csv file and access the data by the rest full api service

Hello every one, 
Please clone the project and run => `go run .\main.go`
Also build the project => `go build .\main.go`  for creating main.exe

It provides you with a rest api for fetching the data per Id, Id is a Guid 

## Sample url
http://localhost:8384/promotions/10121f45-45bf-4350-9547-d2961eda4329

For enhancing the process of csv file, I use bulk insert also chunk the data, which chunked by 1000. This chunkSize is an editable variable.
I could improve the speed of preparing the date and inserting 200,000 less than one minute. In a first approach which I read the csv file by **"record, err := csvLines.Read()"** and insert in Mysql per record it took more the 30 minutes.

## Benchmark

| Chunk Size    | time        | description  |
| ------------- |:-----------:| -----:|
| 1             | 30 minutes  | per read from csv having single insert |
| 500           | 1.05 minutes| using bulk insert |
| 1000          | 45 seconds  | using bulk insert |

 I think if I execute each chunked data in a separate gorutiue may be it gives a better result but I don't enough time to test it. I was looking for a mechanism in Golang equivalent to `Parallel.ForEach` in **c#**
 
 
## Features, tools and principals that use in this code
1. Goroutines
2. GoCron
3. GORM
4. Bulk insert for enhance the performance of insert
5. Chunk the data which obtain from csv file.
6. Graceful shutdown for my Webapi by gorutine and context: if exit the project by this command `ctrl + c`
7. Create layering such as data layer and csv service
8. Using transaction in my bulk insert
9. AutoMigrate mechanism
