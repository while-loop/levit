package repo

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateConnection(host, user, password, db string) (*gorm.DB, error) {
	return gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, db),
	)
}
