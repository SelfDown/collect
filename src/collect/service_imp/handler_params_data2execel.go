package collect

import (
	common "collect/src/collect/common"
	config "collect/src/collect/config"
	utils "collect/src/collect/utils"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type Data2Excel struct {
	BaseHandler
}

func (uf *Data2Excel) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {
	params := template.GetParams()
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.

	for index, sheet := range template.ExcelConfigData.Sheets {
		sheetName := "Sheet" + utils.Strval(index+1)
		_, err := f.NewSheet(sheetName)
		if err != nil {
			return common.NotOk(err.Error())
		}
		title := utils.RenderTplData(sheet.TitleTpl, params)
		f.SetCellRichText(sheetName, "A1", []excelize.RichTextRun{
			{
				Text: title.(string),
			},
		})
		titleStyle, err := f.NewStyle(&sheet.TitleStyle)
		if err != nil {
			template.LogData(err.Error())
		}
		nameStyle, err := f.NewStyle(&sheet.NameStyle)
		if err != nil {
			template.LogData(err.Error())
		}

		//设置行高
		if err := f.SetRowHeight(sheetName, 1, sheet.TitleHeight); err != nil {
			template.LogData(err.Error())
		}
		// 合并头
		colName, _ := excelize.ColumnNumberToName(len(sheet.Fields))
		if err := f.MergeCell(sheetName, "A1", colName+"1"); err != nil {
			template.LogData(err.Error())
		}

		contentStyle, err := f.NewStyle(&sheet.ContentStyle)
		if err != nil {
			template.LogData(err.Error())
		}

		if err != nil {
			template.LogData(err.Error())
		}
		for j, nameField := range sheet.Fields {
			colName, error := excelize.ColumnNumberToName(j + 1)
			if error != nil {
				template.LogData(error)
			}
			cellName := colName + "2"
			// 设置所有列只能文本
			err = f.SetColStyle(sheetName, colName, contentStyle)
			// 设置表头
			f.SetCellValue(sheetName, cellName, nameField.Name)
			// 设置表头样式
			f.SetCellStyle(sheetName, cellName, cellName, nameStyle)
			if nameField.Width > 0 {
				f.SetColWidth(sheetName, colName, colName, nameField.Width)
			}

		}

		//设置标题样式
		err = f.SetCellStyle(sheetName, "A1", "A1", titleStyle)
		if err != nil {
			template.LogData(err.Error())
		}
		// 冻结第一列和第二行
		f.SetPanes(sheetName, &excelize.Panes{
			Freeze: true,
			Split:  false,
			XSplit: 1,
			YSplit: 2,
		})
		dataList, _ := utils.RenderVarToArrMap(sheet.Data, params)
		for dataRow, dataItem := range dataList {
			for j, nameField := range sheet.Fields {
				colName, error := excelize.ColumnNumberToName(j + 1)
				if error != nil {
					template.LogData(error)
				}
				row := dataRow + 3
				cellName := colName + utils.Strval(row)
				var value interface{}
				// 优先模板
				if !utils.IsValueEmpty(nameField.Template) {
					value = utils.RenderTplData(nameField.TemplateTpl, dataItem)
				} else {
					value = utils.RenderVar(nameField.Field, dataItem)
				}
				// 渲染模板值
				f.SetCellStr(sheetName, cellName, utils.Strval(value))

			}

		}

	}

	f.SetActiveSheet(0)
	savePath := utils.RenderVar(handlerParam.Path, params).(string)
	// 如果目录不存在则创建目录
	dir := utils.ParentDirName(savePath)
	if !utils.IsPathExist(dir) {
		if err := utils.CreateDirs(dir); err != nil {
			return common.NotOk(err.Error())
		}
	}

	if err := f.SaveAs(savePath); err != nil {
		common.NotOk(err.Error())
	}
	r := common.Ok(nil, "处理参数成功")
	return r
}
