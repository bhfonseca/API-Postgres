package models

import "api-postgesql/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`select * from todos where id=$1, id`, id)

	err = row.Scan(&todo.ID, todo.Title, todo.Description, todo.Done)

	return
}
