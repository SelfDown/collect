package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
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
		//设置标题样式
		err = f.SetCellStyle(sheetName, "A1", "A1", titleStyle)
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
		for j, nameField := range sheet.Fields {
			colName, error := excelize.ColumnNumberToName(j + 1)
			if error != nil {
				template.LogData(error)
			}
			cellName := colName + "2"
			f.SetCellValue(sheetName, cellName, nameField.Name)
			f.SetCellStyle(sheetName, cellName, cellName, nameStyle)
			if nameField.Width > 0 {
				f.SetColWidth(sheetName, colName, colName, nameField.Width)
			}

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
				if utils.IsRenderVar(nameField.Field) {
					value = utils.RenderVar(nameField.Field, dataItem)
				} else {
					value = utils.RenderTplData(nameField.FieldTpl, dataItem)
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
