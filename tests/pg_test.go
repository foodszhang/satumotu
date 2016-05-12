package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := sql.Open("postgres", "user=pretty password=pppp dbname=pretty sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	rows, err := db.Query(`select id, name , source from product where id > $1 limit $2`, 50, 50)
	defer rows.Close()
	t.Log(rows)
	for rows.Next() {
		var id int
		var name string
		var source string
		err = rows.Scan(&id, &name, &source)
		t.Log(id, name, source)

	}
}
