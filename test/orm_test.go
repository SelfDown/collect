package test

import (
	"bytes"
	serviceImp "collect.mod/src/collect/service_imp"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	//"github.com/go-sql-driver/mysql"
	"log"
	"testing"

	model "collect.mod/model"
	modelGen "collect.mod/model_gen"
)

func SnakeCase(name string) string {

	var b bytes.Buffer
	var lastUnderscore bool
	ln := len(name)
	if ln == 0 {
		return ""
	}
	b.WriteRune(unicode.ToLower(rune(name[0])))
	for i := 1; i < ln; i++ {
		r := rune(name[i])
		nextIsLower := false
		if i < ln-1 {
			n := rune(name[i+1])
			nextIsLower = unicode.IsLower(n) && unicode.IsLetter(n)
		}
		if unicode.IsUpper(r) {
			if !lastUnderscore && nextIsLower {
				b.WriteRune('_')
				lastUnderscore = true
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
			lastUnderscore = false
		}
	}
	return b.String()
}
func get_datasource() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev_user:!QAZ2wsx#EDC4rfv2022@{xx[,*.]}.com2022@.com@tcp(172.26.0.13:3306)/devops?charset=utf8")
	if err != nil {
		log.Fatal("数据库打开出现了问题：", err)
		return nil, err
	}
	db.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return nil, err
	}
	return db, err

}
func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return camel
}
func Test_orm_2(t *testing.T) {
	user_account := model.GetModel("user_account")
	u := reflect.ValueOf(user_account)
	column := u.FieldByName("Statu")
	switch column.Kind() {
	case reflect.Int32:
		break
	}
	fmt.Printf("%#v\n", column.Kind().String())
	fmt.Printf("%#v\n", column.Type())
	fmt.Printf("%#v", u)
}
func Test_orm(t *testing.T) {
	ab := toCamelCase("abc_i")

	fmt.Printf(ab)

	//db, _ := get_datasource()
	//gormDB, _ := gorm.Open(mysql.New(mysql.Config{
	//	Conn: db,
	//}), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true, // 使用单数表名
	//	},
	//})
	base := serviceImp.BaseHandler{}
	gormDB := base.GetGormDb()
	var value int32
	value = 1
	user := model.UserAccount{Userid: "5", Username: "z", Statu: &value, RoleID: "1", Userpwd: "1"}
	//u := model.GetModel("user_account")
	//utils.SetDataValue("Userid", "2", &u)
	//utils.SetDataValue("Username", "1", &u)
	//utils.SetDataValue("RoleId", "1", &u)
	//utils.SetDataValue("Userpwd", "1", &u)
	//reflect.ValueOf(model)
	//fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", user)
	modelGen.SetDefault(gormDB)
	column, _ := modelGen.UserAccount.GetFieldByName("username")
	//u := modelGen.UserAccount
	//err := u.WithContext(ctx).Create(&user)

	fmt.Printf("%#v", column)

	//gormDB.Create(user)

}
