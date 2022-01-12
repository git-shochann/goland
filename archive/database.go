package archive

// コードの中では使わないけど, ビルドして一緒にコンパイルしないと使えないのでこうしておく
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// コネクションを確立する
var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func Db() {
	// create table
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age  INT
	)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	// insert
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Yua", 22)
	if err != nil {
		log.Fatalln(err)
	}

	// update
	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 5, "Sho")
	if err != nil {
		log.Fatalln(err)
	}

	// multiple select
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person // Person型のスライスを作成
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age) // rows.Scan()はpointerで渡した型に取得したdataをマッピングするのはなぜ？
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p) // ppのスライスにpを追加して新しいスライスを作成
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// single select
	cmd = "SELECT * FROM person where age = ?"
	row := DbConnection.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age) // どういうこと？
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// delete
	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Sho")
	if err != nil {
		log.Fatalln(err)
	}
}