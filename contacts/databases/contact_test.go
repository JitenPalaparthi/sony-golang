package database

import (
	"os/exec"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	// step-0: make sure all required commands are available
	// step-1 : DB -> Create a docker db instance
	// step-2: connect your application to the db
	// step-3: insert a record
	// step-4: do test
	// step-5: delete db (docker)instance

	//docker run -d -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=contactsdb postgres
	// err := TearUp()
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Fail()
	// }

	// dsn := "host=localhost user=postgres password=postgres dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// db, err := GetConnection(dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// contactdb := Contact{DB: db}
	// contact := new(models.Contact)
	// contact.Name = "Jiten"
	// contact.Email = "Jitenp@outlook.Com"
	// contact.Status = "Active"
	// var t1 int64 = 111111111
	// contact.Lastmodified = t1

	// err = contactdb.Create(contact)
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Fail()
	// }

	// c1, err := contactdb.Get("1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	t.Fail()
	// }
	// contact.ID = 1
	// if c1 != contact {
	// 	fmt.Println(c1, " and", contact, "are not same")
	// 	t.Fail()
	// }

	// // db has been created

}

func TearUp() error {
	cmd := exec.Command("docker", "info")
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	if !strings.Contains(string(output), "Server Version") {
		return err
	}
	cname := "pg1"
	cmd = exec.Command("docker", "run", "-d", "--name", cname, "-p", "5432:5432", "-e", "POSTGRES_USER=postgres", "-e", "POSTGRES_PASSWORD=postgres", "-e", "POSTGRES_DB=contactsdb", "postgres")

	output, err = cmd.Output()
	if err != nil {
		return err
	}
	if len(string(output)) != 64 {
		return err
	}
	return nil
}

func TearDown() error {
	cmd := exec.Command("docker", "rm", "-f", "pg1")
	_, err := cmd.Output()
	return err
}

// https://github.com/axw/gocov
// https://github.com/matm/gocov-html

// https://github.com/matm/gocov-html

// alias testcases="sed -n 's/func.*\(Test.*\)(.*/\1/p' | xargs | sed 's/ /|/g'"

// go test -v -run $(cat coordinator_test.go | testcases)
