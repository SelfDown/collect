package collect

import (
	common "collect.mod/src/collect/common"
	config "collect.mod/src/collect/config"
	utils "collect.mod/src/collect/utils"
	"github.com/xuri/excelize/v2"
)

type Excel2Data struct {
	BaseHandler
}

func (uf *Excel2Data) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *TemplateService) *common.Result {

	f, error := excelize.OpenReader(ts.File)
	if error != nil {
		return common.NotOk("文件转换失败,支持XLAM/XLSM/XLSX/XLTM/XLTX格式，详情:" + error.Error())
	}
	defer ts.File.Close()
	dataList := make([]map[string]interface{}, 0)
	for index, sheet := range template.ExcelConfigData.Sheets {
		sheetName := "Sheet" + utils.Strval(index+1)
		rows, err := f.GetRows(sheetName)
		if err != nil {
			template.LogData(err.Error())
			return common.NotOk(err.Error())
		}

		for index, row := range rows {
			if index < 2 { // 跳过前面2行标题
				continue
			}
			dataItem := make(map[string]interface{})
			for j, nameField := range sheet.Fields {
				field := utils.GetRenderVarName(nameField.Field)
				if j < len(row) {
					dataItem[field] = row[j]
				} else {
					dataItem[field] = nil
				}
			}
			dataList = append(dataList, dataItem)
		}

	}
	r := common.Ok(dataList, "处理参数成功")
	return r
}
