package main

import (
	"Project1/config"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type note struct {
	text string
	id   int
}

type Conn struct {
	ConnTest *sql.DB
}

func (db *Conn) Connect(user string, password string, dbname string, port int, host string) (connect string, err error) {
	defer func() { err = errors.Wrap(err, "Conn.Connect") }()
	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	u := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", host, port),
		Path:   "test",
	}
	// v := url.Values{}
	v, err := url.ParseQuery("sslmode=disable")
	if err != nil {
		err = errors.Wrap(err, "Failed ParseQuery")
		return
	}
	u.RawQuery = v.Encode()
	u.User = url.UserPassword(user, password)
	// connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%d", user, password, dbname, port)
	fmt.Println(u.String())
	db.ConnTest, err = sql.Open("postgres", u.String())
	if err != nil {
		err = errors.Wrap(err, "Failed connection to DB")
		return
	}
	connect = u.String()
	return
}

func (db *Conn) Close() (err error) {
	defer func() { err = errors.Wrap(err, "Conn.Close") }()
	err = db.ConnTest.Close()
	if err != nil {
		err = errors.Wrap(err, "Failed close DB")
		return
	}
	return
}

func selectDB(db *sql.DB) {
	rows, err := db.Query("select * from notes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	notes := []note{}

	for rows.Next() {
		n := note{}
		err := rows.Scan(&n.id, &n.text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		notes = append(notes, n)
	}
	for _, n := range notes {
		fmt.Println(n.id, n.text)
	}
}

func insertDB(a string, db *sql.DB) {
	_, err := db.Exec("insert into notes (text) values ($1)", a) // вместо "a" пишем нужный текст
	if err != nil {
		panic(err)
	}
}

func deleteDB(b int, db *sql.DB) {
	_, err := db.Exec("delete from notes where id = $1", b) //где "b" нужное значение
	if err != nil {
		panic(err)
	}
}

func main() {
	conf := new(config.Config)
	config.ParseConfigFile("config.hcl", &conf)
	fmt.Println(conf.Data.DataDB)
	fmt.Println(conf.Data.DataPassword)
	fmt.Println(conf.Data.DataUser)
	// var conTest Conn
	// _, err := conTest.Connect("ivan_user", "password", "test", 5432, "localhost")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer func() {
	// 	err = conTest.Close()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }()
	// var query string = ""
	// var id int = 0
	// fmt.Println(" \n &param select - выводит таблицу в терминал\n &param insert - запрашивает текст для ввода(без пробелов(испульзуйте _)) и вставляет в таблицу \n &param delete - запрашивает номер строки и удаляет её")
	// for {
	// 	fmt.Println("\n Write SQL query")
	// 	fmt.Fscan(os.Stdin, &query)
	// 	if query != "" {
	// 		switch query {
	// 		case "select":
	// 			selectDB(conTest.ConnTest)
	// 		case "insert":
	// 			fmt.Println("Write text to insert ")
	// 			scanner := bufio.NewScanner(os.Stdin)
	// 			if scanner.Scan() {
	// 				line := scanner.Text()
	// 				insertDB(line, conTest.ConnTest)
	// 			}
	// 		case "delete":
	// 			id = 0
	// 			fmt.Println("Write id where from need to delete")
	// 			fmt.Fscan(os.Stdin, &id)
	// 			deleteDB(id, conTest.ConnTest)
	// 		case "close DB":
	// 			return
	// 		case "quit":
	// 			os.Exit(0)
	// 		default:
	// 			fmt.Println("Invalid sytax")
	// 		}
	// 	}
	// }
}
