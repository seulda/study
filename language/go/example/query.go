package main
 
import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "id:pw@tcp(addr:port)/dbName")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
     
    // 하나의 Row를 갖는 SQL 쿼리
    var name01 string
    err = db.QueryRow("SELECT sum(population) AS population FROM population WHERE year = 2021").Scan(&name01)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name01)

    // 복수 Row를 갖는 SQL 쿼리
    var id int
    var name02 string
    rows, err := db.Query("SELECT year, sum(population) AS population FROM population GROUP BY year ORDER BY year")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
 
    for rows.Next() {
        err := rows.Scan(&id, &name02)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name02)
    }
}
