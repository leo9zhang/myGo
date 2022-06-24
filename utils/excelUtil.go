// Package utils TODO
/*
 * @Author: leospzhang leospzhang@tencent.com
 * @Date: 2022-06-24 14:10:01
 * @LastEditors: leospzhang leospzhang@tencent.com
 * @LastEditTime: 2022-06-24 14:21:53
 * @FilePath: /gotool/utils/excelUtil.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// ExportExcel TODO
func ExportExcel(url string) {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "helloworld")
	f.SetCellValue("Sheet2", "B2", 100)
	f.SetActiveSheet(index)

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Errorf(err.Error())
	}
}
