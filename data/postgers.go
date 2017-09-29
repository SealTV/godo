package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER           = "postgres"
	DB_PASSWORD       = "postgres"
	DB_NAME           = "todo_database"
	DB_SSL_MODE       = "disable"
	DB_DEFAULT_SCHEMA = "main_schema"
)

func IniDB() *sql.DB {
	db, err := sql.Open("postgres", getConnectionSating())
	if err != nil {
		log.Fatal(err)
	}

	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}

	_, err = db.Exec(fmt.Sprintf("SET search_path TO %s;", DB_DEFAULT_SCHEMA))
	if err != nil {
		log.Fatal(err)
	}

	todo := Todo{
		UserId:      1,
		IsActive:    true,
		Description: "some description",
		ListId:      1,
		Title:       "Som title",
	}

	todo = AddTodo(db, todo)
	todo.Title = "Another title 2"
	UpdateTodo(db, todo)
	DeleteTodo(db, todo)
	todo.Id = 5
	DeleteTodo(db, todo)
	return db
}

func getConnectionSating() string {
	dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=%s", DB_USER, DB_NAME, DB_SSL_MODE)
	return dbInfo
}

func migrate(db *sql.DB) {
	sql, err := db.Prepare(`
    create sequence todos_id_seq;
	create sequence users_id_seq;
	create sequence lists_id_seq;

	create table todos
	(
		id serial not null
			constraint todos_pkey
				primary key,
		title varchar(256) not null,
		description varchar(2048),
		list_id integer,
		is_active boolean default true,
		user_id integer,
		date_create date default now() not null
	);

	create unique index todos_id_uindex
		on todos (id);

	create index todos_user_list
		on todos (user_id, list_id);

	create index todos_user
		on todos (user_id);

	comment on table todos is 'main table';

	create table users
	(
		id serial not null
			constraint users_pkey
				primary key,
		login varchar(128) not null,
		password varchar(256) not null,
		email varchar(256) not null,
		register_date timestamp default now() not null
	);

	create unique index users_id_uindex
		on users (id);

	create unique index users_login_uindex
		on users (login);

	create unique index users_email_uindex
		on users (email);

	comment on table users is 'Users table';

	alter table todos
		add constraint todos_users_id_fk
			foreign key (user_id) references users
	;
	create table lists
	(
		id serial not null
			constraint lists_pkey
				primary key,
		name varchar(256) not null,
		user_id integer not null
			constraint lists_users_id_fk
				references users
	);

	create unique index lists_id_uindex
		on lists (id);

	create index lists_user_id_index
		on lists (user_id);

	alter table todos
		add constraint todos_lists_id_fk
			foreign key (list_id) references lists;
    `)
	_, err = sql.Exec()
	// выходим, если будут ошибки с SQL запросом выше
	if err != nil {
		log.Fatal(err)
	}
}
