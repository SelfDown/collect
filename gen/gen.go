package main

import (
	serviceImp "collect.mod/src/collect/service_imp"
	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:        "model_gen",
		FieldNullable:  true,
		FieldCoverable: true,
		Mode:           gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

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
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormDB) // reuse your gorm db

	//// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.UserAccount{})
	//
	//// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.UserAccount{})

	// Generate the code
	g.ApplyBasic(
		g.GenerateModel("user_account"),
		//g.GenerateModel("sys_projects"),
		//g.GenerateModel("server_instance"),
	)
	g.Execute()
}
