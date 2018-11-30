package main

import (
  "log"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/iande/budgegator/init"
)


func initDB(db *sql.DB) {
  sqlStmt := `
    create table categories (
      id integer not null primary key,
      name text,
      notes text,
      parent_category_id integer
    );
    create table accounts (
      id integer not null primary key,
      name text,
      notes text,
      account_type_id integer,
      budgeted boolean,
      active boolean
    );
    create table account_types (
      id integer not null primary key,
      name text,
      recommend_budgeted boolean
    );
    create table payees (
      id integer not null primary key,
      name text,
      last_category_id integer
    );
    create table transactions (
      id integer not null primary key,
      date text,
      memo text,
      amount integer,
      account_id integer,
      category_id integer,
      payee_id integer,
      parent_transaction_id integer,
      status text
    );
  `
  db.Exec(sqlStmt)
}

func main() {
  db, err := sql.Open("sqlite3", "./budget.db")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  initDB(db)

  log.Println("Hello world")
}
