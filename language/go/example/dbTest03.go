package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func year2021_count(db *sql.DB, err error) {
    var result string
    err = db.QueryRow("SELECT sum(population) AS population FROM population WHERE year = 2021").Scan(&result)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
}

func year_count(db *sql.DB, dong string) {
    var year string
    var population string
    var querySet string
    if dong == "" {
        fmt.Println("year_count() dong is nil")
        querySet = "SELECT year, sum(population) AS population FROM population GROUP BY year ORDER BY year"
    } else {
        fmt.Println("year_count() dong is not nil")
        querySet = "SELECT year, sum(population) AS population FROM population WHERE dong = '" + dong + "' GROUP BY year ORDER BY year"
    }
    rows, err := db.Query(querySet) 
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&year, &population)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(year, population)
    }
}

func year2021_sex(db *sql.DB, dong string) {
    var sex string
    var population string
    var per string
    var querySet string
    if dong == "" {
        fmt.Println("year2021_sex() dong is nil")
        querySet = "SELECT sex, sum(population) AS population, round(sum(population)/(select sum(population) from population where year = 2021)*100, 1) AS per FROM population WHERE year = 2021 GROUP BY sex"
    } else {
        fmt.Println("year2021_sex() dong is not nil")
        querySet = "SELECT sex, sum(population) AS population, round(sum(population)/(select sum(population) from population where year = 2021 AND dong = '" + dong + "')*100, 1) AS per FROM population WHERE year = 2021 AND dong = '" + dong + "' GROUP BY sex"
    }
    rows, err := db.Query(querySet)  
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&sex, &population, &per)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(sex, population, per)
    }
}

func year_age(db *sql.DB, dong string) {
    var year string
    var age string
    var population string
    var querySet string
    if dong == "" {
        fmt.Println("year_age() dong is nil")
        querySet = "SELECT year, age, sum(population) AS population FROM population GROUP BY year, age ORDER BY year, age"
    } else {
        fmt.Println("year_age() dong is not nil")
        querySet = "SELECT year, age, sum(population) AS population FROM population WHERE dong = '" + dong + "' GROUP BY year, age ORDER BY year, age"
    }
    rows, err := db.Query(querySet)  
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    for rows.Next() {
        err := rows.Scan(&year, &age, &population)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(year, age, population)
    }
}

func year_item(db *sql.DB, dong string) {
    var year string
    var item string
    var population string
    var querySet string
    if dong == "" {
        fmt.Println("year_item() dong is nil")
        querySet = "SELECT year, item, sum(population) AS population FROM item GROUP BY year, item ORDER BY year, item"
    } else {
        fmt.Println("year_item() dong is not nil")
        querySet = "SELECT year, item, sum(population) AS population FROM item WHERE dong = '" + dong + "' GROUP BY year, item ORDER BY year, item"
    }
    rows, err := db.Query(querySet)  
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    for rows.Next() {
        err := rows.Scan(&year, &item, &population)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(year, item, population)
    }
}

func main() {

    var dong string

    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "id:pw@tcp(addr:port)/dbName")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // year2021_count()
    year2021_count(db, err)
    fmt.Println("=============== year2021_count()")

    // year_count()
    year_count(db, dong)
    fmt.Println("=============== year_count(), dong nil")
    year_count(db, "거제면")
    fmt.Println("=============== year_count(), dong 거제면")
    
    // year2021_sex()
    year2021_sex(db, "")
    fmt.Println("=============== year2021_sex(), dong nil")
    year2021_sex(db, "거제면")
    fmt.Println("=============== year2021_sex(), dong 거제면")
    
    // year_age()
    year_age(db, dong)
    fmt.Println("=============== year_age(), dong nil")
    year_age(db, "거제면")
    fmt.Println("=============== year_age(), dong 거제면")
    
    // year_item()
    year_item(db, dong)
    fmt.Println("=============== year_item(), dong nil")
    year_item(db, "거제면")
    fmt.Println("=============== year_item(), dong 거제면")
    
}
