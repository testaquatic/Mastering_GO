package restdb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres 접속을 위한 정보
var (
	Hostname = "localhost"
	Port     = 5432
	Username = "postgres"
	Password = "pass"
	Database = "restapi"
)

type PgPool struct {
	*pgxpool.Pool
}

// 사용자 정보
type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// Postgres 서버와 연결한다.
func ConnectPostgres() *PgPool {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", Username, Password, Hostname, Port, Database)
	dbpool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &PgPool{dbpool}
}

func (pgpool PgPool) Close() {
	pgpool.Pool.Close()

}

func (pgpool PgPool) DeleteUser(ID int) bool {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return false
	}

	t := pgpool.FindUserID(ID)
	if t.ID == 0 {
		log.Println("User", ID, "does not exist.")
		return false
	}

	// 트랜잭션을 시작한다.
	tx, err := pgpool.Begin(context.Background())
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}
	_, err = tx.Exec(context.Background(), "DELETE FROM users WHERE id=$1", ID)
	if err != nil {
		log.Println("DeleteUser:", err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Println("DeleteUser:", err)
		}
		return false
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}

	return true
}

func (pgpool PgPool) ListAllUsers() []User {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return nil
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Println("ListAllUsers:", err)
		return nil
	}
	defer rows.Close()

	all := []User{}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Admin, &user.Active)
		if err != nil {
			log.Println("ListAllUsers:", err)
			continue
		}
		all = append(all, user)
	}

	log.Println("All:", all)
	return all
}

func (pgpool PgPool) IsUserValid(u User) bool {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return false
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users WHERE username=$1", u.Username)
	if err != nil {
		log.Println("IsUserValid:", err)
		return false
	}
	defer rows.Close()

	temp := User{}
	if rows.Next() {
		// 같은 이름을 갖는 여러 사용자가 있더라도 하나의 레코드만 사용한다.
		err = rows.Scan(&temp.ID, &temp.Username, &temp.Password, &temp.LastLogin, &temp.Admin, &temp.Active)
		if err != nil {
			log.Println("IsUserValid:", err)
			return false
		}
	}
	if u.Username == temp.Username && u.Password == temp.Password {
		return true
	}

	return false
}

func (pgpool PgPool) InsertUser(u User) bool {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return false
	}

	if pgpool.IsUserValid(u) {
		log.Println("User", u.Username, "already exists!")
		return false
	}

	// 트랜잭션을 시작한다.
	tx, err := pgpool.Begin(context.Background())
	if err != nil {
		log.Println("InsertUser:", err)
		return false
	}

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO users (username, password, lastlogin, admin, active) VALUES ($1, $2, $3, $4, $5)",
		u.Username, u.Password, u.LastLogin, u.Admin, u.Active,
	)
	if err != nil {
		log.Println("InsertUser:", err)
		err := tx.Rollback(context.Background())
		if err != nil {
			log.Println("InsertUser:", err)
		}
		return false
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("InsertUser:", err)
		return false
	}

	return true
}

func (pgpool PgPool) ListLogged() []User {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return nil
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users WHERE active = 1")
	if err != nil {
		log.Println("ListLogged:", err)
		return nil
	}
	defer rows.Close()

	all := []User{}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Admin, &user.Active)
		if err != nil {
			log.Println("ListLogged:", err)
			continue
		}
		all = append(all, user)
	}

	log.Println("All:", all)
	return all
}

func (pgpool PgPool) FindUserID(ID int) User {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return User{}
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users WHERE id=$1", ID)
	if err != nil {
		log.Println("FindUserID:", err)
		return User{}
	}
	defer rows.Close()

	user := User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Admin, &user.Active)
		if err != nil {
			log.Println("FindUserID:", err)
			return User{}
		}
	}

	return user
}

func (pgpool PgPool) FindUserUsername(username string) User {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return User{}
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		log.Println("FindUserUsername:", err)
		return User{}
	}
	defer rows.Close()

	user := User{}
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Admin, &user.Active)
		if err != nil {
			log.Println("FindUserUsername:", err)
			return User{}
		}
		log.Println("Found user:", user)
	}

	return user
}

func (pgpool PgPool) ReturnLoggedUsers() []User {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return nil
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users WHERE active = 1")
	if err != nil {
		log.Println("ReturnLoggedUsers:", err)
		return nil
	}
	defer rows.Close()

	all := []User{}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.LastLogin, &user.Admin, &user.Active)
		if err != nil {
			log.Println("ReturnLoggedUsers:", err)
			continue
		}
		all = append(all, user)
	}

	log.Println("Logged in:", all)

	return all
}

func (pgpool PgPool) IsUserAdmin(u User) bool {
	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return false
	}

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM users WHERE username=$1", u.Username)
	if err != nil {
		log.Println("IsUserAdmin:", err)
		return false
	}
	defer rows.Close()

	temp := User{}
	if rows.Next() {
		// 같은 이름을 갖는 여러 사용자가 있더라도 하나의 레코드만 사용한다.
		err = rows.Scan(&temp.ID, &temp.Username, &temp.Password, &temp.LastLogin, &temp.Admin, &temp.Active)
		if err != nil {
			log.Println("IsUserAdmin:", err)
			return false
		}
	}
	if u.Username == temp.Username && u.Password == temp.Password && temp.Admin == 1 {
		return true
	}

	return false
}

func (pgpool PgPool) UpdateUser(u User) bool {
	log.Println("Updating user:", u)

	if pgpool.Pool == nil {
		log.Println("Cannot connect to PostgreSQL!")
		// nil에 대해서 Close를 수행하는 것은 맞지 않는 것 같다.
		// pgpool.Pool.Close()
		return false
	}

	// 트랜잭션을 시작한다.
	tx, err := pgpool.Begin(context.Background())
	if err != nil {
		log.Println("UpdateUser:", err)
		return false
	}

	res, err := tx.Exec(
		context.Background(),
		"UPDATE users SET username=$1, password=$2, admin=$3, active=$4 WHERE id=$5",
		u.Username, u.Password, u.Admin, u.Active, u.ID,
	)
	if err != nil {
		log.Println("UpdateUser:", err)
		err := tx.Rollback(context.Background())
		if err != nil {
			log.Println("UpdateUser:", err)
		}
		return false
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("UpdateUser:", err)
		return false
	}

	affected := res.RowsAffected()
	if affected == 0 {
		log.Println("RowsAffected() failed")
		return false
	}
	log.Println("Affected:", affected)

	return true
}
