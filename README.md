# CaseStudyGolang
Persist data from csv file and access the data by the rest full api service

Hello every one, 
Please clone the project and run => go run .\main.go
Also build the project => go build .\main.go  for creating main.exe

It provides you with a rest api for fetching the data per Id, Id is a Guid 
http://localhost:8384/promotions/10121f45-45bf-4350-9547-d2961eda4329

For enhancing the process of csv file, I use bulk insert also chunk the data, which chunked by 1000. This chunkSize is an editable variable.
I could improve the speed of preparing the date and inserting 200,000 less than one minute. In a first approach which I read the csv file by "record, err := csvLines.Read()" and insert in Mysql per record it took more the 30 minutes. I think if I execute each chunked data in a separate gorutiue may be it gives a better result but I don't enough time to test it. I was looking for a mechanism in golang equivalent to Parallel.ForEach in c# 
 
------------------------------------------------------------------------------------
My name is Mehrdad Salimian. This is my first time Golang project. My time is very limited. But I made an effort to dedicate two days for implementing this project. I am .net developer and I have had 14 years’ experience for designing and developing application. I tried to use some features, tools and principals in this project such as:
1- Goroutines
2- goCron
3- GORM
4- Bulk insert for enhance the performance of insert
5- Chunk the data which obtain from csv file.
6- Graceful shutdown for my Webapi by gorutine and context
7- Create layering such as data layer and csv service
8- Using transaction in my bulk insert
9- AutoMigrate mechanism



