**Table Of Contents**
- [XRAW (Raw Query ORM) Description](#XRAW-Raw-Query-ORM-Description)
- [Benchmarking vs Other ORMs](#Benchmarking-vs-Other-ORMs)
- [Support Database](#Support-Database)
- [Installation](#Installation)
- [Features](#Features)
- [How To Use](#How-To-Use)
  - [Configure the Host](#Configure-the-Host)
  - [Init New Engine](#Init-New-Engine)
  - [Init the models](#Init-the-models)
  - [Create New SQL Select Query](#Create-New-SQL-Select-Query)
    - [Get All Data](#Get-All-Data)
      - [With SQL Raw](#With-SQL-Raw)
      - [WITH Query Builder](#WITH-Query-Builder)
    - [Get Multiple Result Data with Where Condition](#Get-Multiple-Result-Data-with-Where-Condition)
    - [Get Single Result Data with Where Condition](#Get-Single-Result-Data-with-Where-Condition)
  - [Create, Update, Delete Query](#Create-Update-Delete-Query)
    - [Insert](#Insert)
      - [Single Insert](#Single-Insert)
      - [Multiple Insert](#Multiple-Insert)
    - [Update](#Update)
    - [Delete](#Delete)
# XRAW (Raw Query ORM) Description
Raw Query ORM is a Query Builder as light as raw query and as easy as ORM

# Benchmarking vs Other ORMs
source : https://github.com/kihamo/orm-benchmark

command : ``` orm-benchmark -orm=all (-multi=1 default) ```

```bash
Reports: 

  2000 times - Insert
 ---  xraw:     2.07s      1034107 ns/op     480 B/op      7 allocs/op ---
       orm:     2.61s      1307297 ns/op    1433 B/op     37 allocs/op
       raw:     2.79s      1397290 ns/op     568 B/op     14 allocs/op
       qbs:     3.12s      1562002 ns/op    4319 B/op    105 allocs/op
      gorp:     3.82s      1909134 ns/op    1405 B/op     31 allocs/op
      hood:     3.98s      1991849 ns/op   10736 B/op    156 allocs/op
      xorm:     4.18s      2091037 ns/op    2584 B/op     69 allocs/op
      modl:     4.47s      2232501 ns/op    1333 B/op     30 allocs/op
      gorm:     5.70s      2851877 ns/op    7717 B/op    150 allocs/op

   500 times - MultiInsert 100 row
  --- xraw:     1.00s      1991180 ns/op   42711 B/op    205 allocs/op ---
       orm:     1.21s      2428474 ns/op  104592 B/op   1629 allocs/op
       raw:     1.91s      3824940 ns/op  110983 B/op   1110 allocs/op
      xorm:     2.78s      5566887 ns/op  230481 B/op   4962 allocs/op
      gorp:     Not support multi insert
      hood:     Not support multi insert
       qbs:     Not support multi insert
      modl:     Not support multi insert
      gorm:     Not support multi insert

  2000 times - Update
       raw:     1.08s       541486 ns/op     632 B/op     16 allocs/op
       orm:     1.12s       559463 ns/op    1392 B/op     38 allocs/op
---   xraw:     1.35s       675182 ns/op     288 B/op      5 allocs/op ---
      gorp:     2.55s      1273046 ns/op    1552 B/op     37 allocs/op
      xorm:     2.61s      1305718 ns/op    2896 B/op    107 allocs/op
      modl:     2.81s      1405387 ns/op    1504 B/op     38 allocs/op
       qbs:     3.77s      1882914 ns/op    4313 B/op    105 allocs/op
      hood:     6.14s      3068596 ns/op   10731 B/op    156 allocs/op
      gorm:     8.30s      4150627 ns/op   18622 B/op    385 allocs/op

  4000 times - Read
       orm:     2.28s       570828 ns/op    2640 B/op     95 allocs/op
       raw:     2.53s       632325 ns/op    1432 B/op     37 allocs/op
       qbs:     2.89s       721450 ns/op    6360 B/op    176 allocs/op
 ---  xraw:     4.67s      1167800 ns/op    1792 B/op     40 allocs/op ---
      hood:     5.11s      1277683 ns/op    4016 B/op     48 allocs/op
      xorm:     5.43s      1358713 ns/op    9981 B/op    267 allocs/op
      modl:     5.58s      1394672 ns/op    1873 B/op     45 allocs/op
      gorp:     5.59s      1396913 ns/op    1872 B/op     52 allocs/op
      gorm:     5.99s      1496423 ns/op   12156 B/op    239 allocs/op

  2000 times - MultiRead limit 100
      modl:     1.76s       878631 ns/op   49864 B/op   1721 allocs/op
       raw:     1.76s       881717 ns/op   34704 B/op   1320 allocs/op
       orm:     1.92s       961051 ns/op   88198 B/op   4483 allocs/op
      gorp:     2.15s      1072788 ns/op   63673 B/op   1909 allocs/op
       qbs:     2.21s      1104192 ns/op  165632 B/op   6428 allocs/op
---   xraw:     3.04s      1520843 ns/op   40440 B/op   1536 allocs/op ---
      hood:     4.44s      2217598 ns/op  136053 B/op   6358 allocs/op
      xorm:     4.86s      2430348 ns/op  178372 B/op   7890 allocs/op
      gorm:     5.31s      2654807 ns/op  255421 B/op   6225 allocs/op
```

# Support Database
| No   | Database   |
| :--- | :--------- |
| 1    | MySQL      |
| 2    | Postgres   |
| 3    | SQL Server |

# Installation
```go
go get github.com/radityaapratamaa/xraw
```

import to your project

```go
import "github.com/radityaapratamaa/xraw"
```
# Features
| Feature       | Using                                               | Description                                                                                                      |
| :------------ | :-------------------------------------------------- | :--------------------------------------------------------------------------------------------------------------- |
| Select        | Select(cols ...string)                              | Specify the column will be query                                                                                 |
| SelectSum     | SelectSumn(col string)                              | Specify the single column to be summarize                                                                        |
| SelectAverage | SelectAverage(col string)                           | Specify the single column to average the value                                                                   |
| SelectMax     | SelectMax(col string)                               | Specify the single column to get max the value                                                                   |
| SelectMin     | SelectMin(col string)                               | Specify the single column to get min the value                                                                   |
| SelectCount   | SelectCount(col string)                             | Specify the single column to get total data that column                                                          |
| Where         | Where(col string, value interface{}, opt ...string) | set the condition, ex: Where("id", 1, ">") -> **"WHERE id > 1"** Where("name", "test") => **"WHERE name = 'test'"** |
| WhereIn       |                                                     |
| WhereNotIn    |                                                     |
| WhereLike     |                                                     |
| Or            |                                                     |
| OrIn          |                                                     |
| OrNotIn       |                                                     |
| OrLike        |                                                     |
| GroupBy       |                                                     |
| Join          |                                                     |
| Limit         |                                                     |
| OrderBy       |                                                     |
| Asc           |                                                     |
| Desc          |                                                     |

# How To Use
## Configure the Host
```go
// Mandatory DBConfig
dbConfig := &xraw.DbConfig{
    Host: "localhost",
    Driver: "(mysql | postgres | sqlserver)",
    Username: "your_username",
    DbName: "your_dbName",
}
```

All Property DBConfig
```go
dbConfig := &xraw.DbConfig{
    Host: "your_host", //mandatory
    Driver: "DB Driver", //mandatory
    Username: "dbUsername", //mandatory
    DbName:"database_name", //mandatory
    Password: "dbPass",
    DbScheme: "db Scheme", // for postgres scheme, default is "Public" Scheme
    Port: "port", //default 3306 (mysql), 5432 (postgres)
    Protocol: "db Protocol", //default is "tcp"
    DbInstance: "db Instance", // for sqlserver driver if necessary
}
```
## Init New Engine
```go
// Please make sure the variable "dbConfig" pass by reference (pointer)
db, err := xraw.New(dbConfig)
if err != nil {
    log.Fatalln("Cannot Connet to database")
}
log.Println("Success Connect to Database")
```

## Init the models
```sql
    -- We Have a table with this structure
    CREATE TABLE Student (
        id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NULL,
        address TEXT NULL,
        is_active BOOLEAN NULL,
        birth_date DATE NULL, 
    )
```

```go
    // Init the models (struct name MUST BE SAME with table name)
    type Student struct {
        // db tag, filled based on column name
        Id int `db:"id"`
        Name string 
        Address string
        // IsActive is boolean field
        IsActive bool `db:"is_active"`
        // BirthDate is "Date" data type column field
        BirthDate string `db:"birth_date"`
    }

    // init student struct to variable
    var studentList []Student
    var student Student
```

## Create New SQL Select Query
### Get All Data
#### With SQL Raw
```go
    if err := db.SQLRaw(`SELECT name, address, birth_date FROM Student [JOIN ..... ON .....] 
    [WHERE ......... [ORDER BY ...] [LIMIT ...]]`).Get(&studentList); err != nil {
        log.Fatalln(err.Error())
    }
    log.Println("result is, ", studentList)
```
#### WITH Query Builder
```go
    // Get All Students data
    if err := db.Get(&studentList); err != nil {
        log.Fatalln(err.Error())
    }
    log.Println("result is, ", studentList)
```
```sql
    --  it will generate : 
    SELECT * FROM student
```
### Get Multiple Result Data with Where Condition
```go
    // Get Specific Data
    if err := db.Select("name, address, birth_date").Where("is_active", true).Get(&studentList); err != nil {
        log.Fatalln(err.Error())
    }
    log.Println("result is, ", studentList)
```
```sql
-- it will generate: (prepared Statement)
SELECT name, address, birth_date FROM student WHERE is_active = ?
```
---
```go
    // Get Specific Data (other example)
    if err := db.Select("name", "address", "birth_date").
    Where("is_active", 1).WhereLike("name", "%Lorem%").
    Get(&studentList); err != nil {
        log.Fatalln(err.Error())
    }
    log.Println("result is, ", studentList)
```
```sql
-- it will generate: 
SELECT name, address, birth_date FROM student WHERE is_active = ? AND name LIKE ?
```
### Get Single Result Data with Where Condition
```go
    // Get Specific Data (single Result)
    if err := db.Select("name, address, birth_date").Where("id", 1).Get(&student); err != nil {
        log.Fatalln(err.Error())
    }
    log.Println("result is, ", student)
```
```sql
-- it will generate: 
SELECT name, address, birth_date FROM student WHERE id = ?
```

## Create, Update, Delete Query
### Insert
#### Single Insert
```go
    dtStudent := Student{
        Name: "test",
        Address: "test",
        IsActive: 1,
        BirthDate: "2010-01-01",
    }

    affected, err := db.Insert(&dtStudent)
    if err != nil {
        log.Fatalln("Error When Insert")
    }

    if affected > 0 {
        log.Println("Success Insert")
    }
```
```sql
    -- it will generate : (mysql)
    INSERT INTO Student (name, address, is_active, birth_date) VALUES (?,?,?,?)
    -- prepared Values :
    -- ('test', 'test', 1, '2010-01-01')
```
#### Multiple Insert
```go
    dtStudents := []Student{
        Student{
            Name: "test",
            Address: "test",
            IsActive: 1,
            BirthDate: "2010-01-01",
        },
        Student{
            Name: "test2",
            Address: "test2",
            IsActive: 1,
            BirthDate: "2010-01-02",
        },
        Student{
            Name: "test3",
            Address: "test3",
            IsActive: 1,
            BirthDate: "2010-01-03",
        },
    }

    affected, err := db.Insert(&dtStudent)
    if err != nil {
        log.Fatalln("Error When Insert")
    }

    if affected > 0 {
        log.Println("Success Insert")
    }
```
```sql
    -- it will generate : (mysql)
    INSERT INTO Student (name, address, is_active, birth_date) VALUES (?,?,?,?)
    -- prepared Values :
    -- 1. ('test', 'test', 1, '2010-01-01')
    -- 2. ('test2', 'test2', 1, '2010-01-02')
    -- 3. ('test3', 'test3', 1, '2010-01-03')
```

### Update
```go
    dtStudent := Student{
        Name: "change",
        Address: "change",
        IsActive: 1,
        BirthDate: "2010-01-10",
    }

    affected, err := db.Where("id", 1).Update(&dtStudent)
    if err != nil {
        log.Fatalln("Error When Update")
    }

    if affected > 0 {
        log.Println("Success Update")
    }
```
```sql
    -- it will generate : (mysql)
    UPDATE Student SET name = ?, address = ?, is_active = ?, birth_date = ? WHERE id = ?
    -- prepared Values :
    -- ('change', 'change', 1, '2010-01-10', 1)
```
### Delete
```go
    affected, err := db.Where("id", 1).Delete(&Student{})
    if err != nil {
        log.Fatalln("Error When Delete")
    }

    if affected > 0 {
        log.Println("Success Delete")
    }
```
```sql
    -- it will generate : (mysql)
    DELETE FROM Student WHERE id = ?
    -- prepared Values :
    -- (1)
```


