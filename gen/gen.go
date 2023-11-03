package main

import (
	serviceImp "collect.mod/src/collect/service_imp"
	utils "collect.mod/src/collect/utils"
	"fmt"
	"gorm.io/gen"
	"io/ioutil"
	"strings"
)

func fixModel() {
	directory := "./model" // The current directory

	files, err := ioutil.ReadDir(directory) //read the files from the directory
	if err != nil {
		fmt.Println("error reading directory:", err) //print error if directory is not read properly
		return
	}
	addTable := "package model\n\nfunc addTable() {\n"
	index := 0
	for _, file := range files {
		name := file.Name() //print the files from the directory
		if !strings.HasSuffix(name, ".gen.go") {
			continue
		}
		index += 1
		fmt.Println(name)
		filePath := directory + "/" + name
		content, _ := ioutil.ReadFile(filePath)
		contentData := utils.Strval(content)
		tableName := strings.ReplaceAll(name, ".gen.go", "")
		modelName := utils.ToSchemaName(tableName)
		primaryKeyList := make([]string, 0)
		varName := "table" + modelName
		addTable += "\n\t//" + utils.Strval(index) + " " + tableName
		addTable += "\n\t" + varName + " := " + modelName + "{}"
		addTable += "\n\tmodelMap[\"" + tableName + "\"] = " + varName
		addTable += "\n\tprimaryKeyMap[\"" + tableName + "\"] = " + varName + ".PrimaryKey()"
		for _, line := range strings.Split(contentData, "\n") {
			if strings.Contains(line, "primaryKey") {
				key := strings.Split(line, "column:")[1]
				key = strings.Split(key, ";")[0]
				key = strings.TrimSpace("\"" + key + "\"")
				primaryKeyList = append(primaryKeyList, key)
			}

		}

		contentData += "\n" +
			"func (*" + modelName + ") PrimaryKey() []string " +
			"{\n\t" +
			"return []string{" + strings.Join(primaryKeyList, ",") + "}\n" +
			"}"
		ioutil.WriteFile(filePath, []byte(contentData), 0777)
	}
	addTable += "\n}"
	ioutil.WriteFile(directory+"/add_table.go", []byte(addTable), 0777)
}

func main1() {
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
	tableList, _ := gormDB.Migrator().GetTables()
	for _, item := range tableList {
		fmt.Println(item)
		g.ApplyBasic(g.GenerateModel(item))
	}
	//g.ApplyBasic(
	//g.GenerateAllTable(),
	//g.GenerateModel("user_account"),
	//g.GenerateModel("sys_projects"),
	//g.GenerateModel("server_instance"),
	//)
	g.Execute()
	fixModel()
}
