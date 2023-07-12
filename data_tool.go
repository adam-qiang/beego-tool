/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-07 00:43:38
 */

package tool

import (
	"bytes"
	"encoding/csv"
	"fmt"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/tealeg/xlsx/v3"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

// ExportCsv 数据导出CSV
// @param ctx *beegoContext.Context
// @param title []string
// @param dataList [][]string
// @param fileName string
// @return error
func ExportCsv(ctx *beegoContext.Context, title []string, dataList [][]string, fileName string) error {
	//处理基础数据
	fileNameGbk, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(fileName)), simplifiedchinese.GBK.NewEncoder()))
	fileName = string(fileNameGbk)
	for k, col := range title {
		t, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(col)), simplifiedchinese.GBK.NewEncoder()))
		title[k] = string(t)
	}

	//设置头信息
	ctx.Output.Header("Content-Type", "application/csv")
	ctx.Output.Header("Content-Disposition", "attachment; filename="+string(fileName)+".csv")

	//写入数据流
	writer := csv.NewWriter(ctx.ResponseWriter)
	err := writer.Write(title)
	if err != nil {
		return err
	}
	for _, items := range dataList {
		for key, item := range items {
			itemArr, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(item)), simplifiedchinese.GBK.NewEncoder()))
			items[key] = string(itemArr)
		}
		err := writer.Write(items)
		if err != nil {
			return err
		}
		writer.Flush()
	}
	return nil
}

// ExportExcel 数据导出excel
// @param w http.ResponseWriter
// @param r *http.Request
// @param titleList []string
// @param dataList [][]interface{}
// @param fileName string
func ExportExcel(w http.ResponseWriter, r *http.Request, titleList []string, dataList [][]interface{}, fileName string) {
	// 生成一个新的文件
	file := xlsx.NewFile()

	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")

	// 插入表头
	titleRow := sheet.AddRow()
	titleRow.WriteSlice(titleList, -1)
	//单独单元格赋值同时设置单元格样式等
	//for _, v := range titleList {
	//	cell := titleRow.AddCell()
	//	cell.Value = v
	//	//cell.GetStyle().Font.Color = "00FF0000"
	//}

	// 插入内容
	for _, data := range dataList {
		row := sheet.AddRow()
		row.WriteSlice(data, -1)
	}
	fileName = fmt.Sprintf("%s.xlsx", fileName)
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content := bytes.NewReader(buffer.Bytes())
	http.ServeContent(w, r, fileName, time.Now(), content)
}
