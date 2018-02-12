package repo

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/while-loop/levit/common/test"
	"gopkg.in/ory-am/dockertest.v3"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	if !test.HasDockerInstalled() {
		os.Exit(m.Run())
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	pool.MaxWait = time.Minute * 2
	resource, err := pool.Run("mysql", "8.0.3", []string{
		"MYSQL_ROOT_PASSWORD=pass",
		"MYSQL_DATABASE=test",
		"MYSQL_USER=user",
		"MYSQL_PASSWORD=pass",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = CreateConnection(fmt.Sprintf("localhost:%s", resource.GetPort("3306/tcp")), "user", "pass", "test")
		if err != nil {
			return err
		}
		return db.DB().Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestMySqlImpl(t *testing.T) {
	if !test.HasDockerInstalled() {
		t.Skip("skipping mysql test, docker not installed")
	}

	testImpl(t, NewMySql(db))
}
