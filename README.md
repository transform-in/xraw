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

command : ``` orm-benchmark -orm=xorm,xraw (-multi=1 default) ```

```bash
Reports: 

  2000 times - Insert
       raw:     1.97s       984686 ns/op     568 B/op     14 allocs/op
      xraw:     2.90s      1451076 ns/op     480 B/op      7 allocs/op
      xorm:     3.51s      1753535 ns/op    2584 B/op     69 allocs/op

   500 times - MultiInsert 100 row
      xraw:     1.17s      2333593 ns/op   42711 B/op    205 allocs/op
       raw:     1.91s      3817893 ns/op  110983 B/op   1110 allocs/op
      xorm:     2.05s      4103739 ns/op  230517 B/op   4962 allocs/op

  2000 times - Update
      xraw:     1.10s       550760 ns/op     288 B/op      5 allocs/op
       raw:     1.11s       554043 ns/op     632 B/op     16 allocs/op
      xorm:     2.25s      1122649 ns/op    2897 B/op    107 allocs/op

  4000 times - Read
       raw:     2.04s       510274 ns/op    1432 B/op     37 allocs/op
      xraw:     4.24s      1059379 ns/op    1793 B/op     40 allocs/op
      xorm:     5.85s      1462664 ns/op    9981 B/op    267 allocs/op

  2000 times - MultiRead limit 100
       raw:     1.47s       734600 ns/op   34704 B/op   1320 allocs/op
      xraw:     2.90s      1452457 ns/op   40440 B/op   1536 allocs/op
      xorm:     4.11s      2052743 ns/op  178377 B/op   7890 allocs/op
```

command: ``` orm-benchmark -orm=xorm,xraw,raw -multi=10 ```

```
Reports: 

 20000 times - Insert
       raw:    22.77s      1138388 ns/op     568 B/op     14 allocs/op
      xraw:    23.99s      1199305 ns/op     480 B/op      7 allocs/op
      xorm:    38.69s      1934604 ns/op    2577 B/op     69 allocs/op

  5000 times - MultiInsert 100 row
      xraw:    13.06s      2612550 ns/op   42632 B/op    205 allocs/op
       raw:    20.18s      4035083 ns/op  110905 B/op   1110 allocs/op
      xorm:    23.67s      4733712 ns/op  230398 B/op   4962 allocs/op

 20000 times - Update
       raw:    11.58s       579139 ns/op     632 B/op     16 allocs/op
      xraw:    12.02s       601225 ns/op     288 B/op      5 allocs/op
      xorm:    25.18s      1258879 ns/op    2896 B/op    107 allocs/op

 40000 times - Read
       raw:    26.16s       654035 ns/op    1432 B/op     37 allocs/op
      xraw:    51.59s      1289778 ns/op    1792 B/op     40 allocs/op
      xorm:    56.53s      1413237 ns/op    9978 B/op    267 allocs/op

 20000 times - MultiRead limit 100
       raw:    15.03s       751372 ns/op   34705 B/op   1320 allocs/op
      xraw:    32.58s      1629140 ns/op   40440 B/op   1536 allocs/op
      xorm:    41.83s      2091260 ns/op  178378 B/op   7890 allocs/op
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


