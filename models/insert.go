package models

import	(
	"github.com/franciscoleal/api_go/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.CreateConnection()
	
	return
}