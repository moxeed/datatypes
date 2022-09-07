# GORM Data Types

## Date

```go
import "gorm.io/datatypes"

type UserWithDate struct {
	gorm.Model
	Name string
	Date datatypes.Date
}

user := UserWithDate{Name: "jinzhu", Date: datatypes.Date(time.Now())}
DB.Create(&user)
// INSERT INTO `user_with_dates` (`name`,`date`) VALUES ("jinzhu","2020-07-17 00:00:00")

DB.First(&result, "name = ? AND date = ?", "jinzhu", datatypes.Date(curTime))
// SELECT * FROM user_with_dates WHERE name = "jinzhu" AND date = "2020-07-17 00:00:00" ORDER BY `user_with_dates`.`id` LIMIT 1
```

## Time

MySQL, PostgreSQL, SQLite, SQLServer are supported.

Time with nanoseconds is supported for some databases which support for time with fractional second scale.

```go
import "gorm.io/datatypes"

type UserWithTime struct {
    gorm.Model
    Name string
    Time datatypes.Time
}

user := UserWithTime{Name: "jinzhu", Time: datatypes.NewTime(1, 2, 3, 0)}
DB.Create(&user)
// INSERT INTO `user_with_times` (`name`,`time`) VALUES ("jinzhu","01:02:03")

DB.First(&result, "name = ? AND time = ?", "jinzhu", datatypes.NewTime(1, 2, 3, 0))
// SELECT * FROM user_with_times WHERE name = "jinzhu" AND time = "01:02:03" ORDER BY `user_with_times`.`id` LIMIT 1
```

NOTE: If the current using database is SQLite, the field column type is defined as `TEXT` type
when GORM AutoMigrate because SQLite doesn't have time type.
