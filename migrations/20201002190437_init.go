package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20201002190437, Down20201002190437)
}

func Up20201002190437(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create table "user"
		(
			id   serial not null
				constraint user_pk
					primary key,
			name varchar(255)
		);
		
		create table "book"
		(
			id      serial not null
				constraint book_pk
					primary key,
			user_id integer,
			name    varchar(255)
		);
		
		create table "author"
		(
			id   serial not null
				constraint author_pk
					primary key,
			name varchar(255)
		);
		
		create table "book_author_link"
		(
			id        serial not null
				constraint table_name_pk
					primary key,
			book_id   integer,
			author_id integer
		);
	`)
	return err
}

func Down20201002190437(tx *sql.Tx) error {
	_, err := tx.Exec(`
		drop table "user";
		drop table "book";
		drop table "author";
		drop table "book_author_link";
	`)
	return err
}
