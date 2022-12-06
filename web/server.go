package web

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Page struct {
	tmpl *template.Template
	db   *sql.DB
}

type Result struct {
	Count int
}

func Setup() {
	db, err := sql.Open("mysql", databaseConnectString())
	if err != nil {
		panic(err)
	}

	query := `CREATE TABLE IF NOT EXISTS clicks (
		id int not null primary key auto_increment,
		timestamp datetime not null default NOW()
	)`
	res, err := db.ExecContext(context.Background(), query)
	if err != nil {
		panic(err)
	}
	fmt.Print(res)
}

func Serve() {
	db, err := sql.Open("mysql", databaseConnectString())
	if err != nil {
		panic(err)
	}

	p := &Page{
		tmpl: template.Must(template.ParseFiles("web/layout.html")),
		db:   db,
	}

	http.HandleFunc("/", p.root)
	http.HandleFunc("/increment", p.increment)
	http.ListenAndServe(":3000", nil)
}

func (p *Page) root(w http.ResponseWriter, r *http.Request) {
	row := p.db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM clicks")
	var result Result
	row.Scan(&result.Count)
	p.tmpl.Execute(w, result)
}

func (p *Page) increment(w http.ResponseWriter, r *http.Request) {
	_, err := p.db.Exec("INSERT INTO clicks VALUES()")
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", 302)
}

func databaseConnectString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_NAME"))
}
