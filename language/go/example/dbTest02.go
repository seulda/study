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

  
    // 1. SELECT
    var name01 string
    err = db.QueryRow("SELECT count(*) FROM board").Scan(&name01)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("count : ", name01)

  
    // 2. INSERT
    one := "goID"
    two := "goPW"
    three := "title by golang test02"
    four := "content by golang test02"
    result, err := db.Exec("INSERT INTO board(id, pw, b_title, b_content) VALUES (?, ?, ?, ?)", one, two, three, four)
    if err != nil {
        log.Fatal(err)
    }
    // sql.Result.RowsAffected() 체크
    // Result 객체로부터 갱신된 레코드수(RowsAffected())와 새로 추가된 Id (LastInsertId())를 구할 수 있음
    n, err := result.RowsAffected()
    fmt.Println("insert result >> count : ", n, "  error: ", err)
    if n == 1 {
        fmt.Println("@@ congratulations. 1 row inserted @@")
    }

  
    // 3. UPDATE
    // Prepared Statement 생성
    stmt, err := db.Prepare("UPDATE board SET b_content=? WHERE b_num=?")
    checkError(err)
    defer stmt.Close()
    // Prepared Statement 실행, Placeholder 파라미터 순서대로 전달
    _, err = stmt.Exec("Golang update test : content", 5)
    checkError(err)
  
  
    // 4. DELETE
    result, err := db.Exec("DELETE FROM board WHERE b_num=5")
    if err != nil {
        log.Fatal(err)
    }
    // sql.Result.RowsAffected() 체크
    n, err := result.RowsAffected()
    fmt.Println("delete result >> count : ", n, "  error: ", err)
    if n == 1 {
        fmt.Println("@@ congratulations. 1 row deleted @@")
    }
  
  
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
