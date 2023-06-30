package test

import (
	"fmt"
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"testing"
)

func tt() {
	// oracle://user:password@127.0.0.1:1521/service
	url := oracle.BuildUrl("172.26.0.113", 1521, "wgtest2", "MOONGOD_DBA", "kdll=3ldls---d##sdsSSd0-2--3", nil)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})

	if err != nil {
		// panic error or log error info
	}

	t := db.Exec(`select * from "BCS"."BCS_AUTHORIZE_LOG"`)
	fmt.Println(t.Rows())

	// do somethings
}
func Test_oracle(t *testing.T) {
	tt()
}
