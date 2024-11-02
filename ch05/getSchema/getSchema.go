package main

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	flag.Parse()
	if flag.NArg() != 5 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}

	hostname := flag.Arg(0)
	p := flag.Arg(1)
	user := flag.Arg(2)
	pass := flag.Arg(3)
	database := flag.Arg(4)

	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number:", err)
		return
	}
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, hostname, port, database)
	// 데이터베이스에 연결한다.
	dbpool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		fmt.Println("pgxpool.New():", err)
		return
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(),
		`SELECT "datname" 
        FROM "pg_database" 
        WHERE datistemplate = false;`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Print("Scan", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()

	query := `
    SELECT table_name
    FROM information_schema.tables
    WHERE table_schema = 'public'
    ORDER BY table_name;
    `
    rows, err = dbpool.Query(context.Background(), query)
    if err != nil {
        fmt.Println("Query", err)
        return
    }

    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if err != nil {
            fmt.Println("Scan", err)
            return
        }
        fmt.Print("+T", name)
    }
    defer rows.Close()
}
