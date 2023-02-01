package job

import (
	"fmt"
	"github.com/DogFoodingCN/workC/models"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
	"time"
)

func SaveJobs(c *gin.Context) {
	var jobs []models.Job
	if err := c.ShouldBindJSON(&jobs); err != nil {
		fmt.Println(err)
		return
	}

	fileName := "Jobs_" + time.Now().Format("2006-01-02") + ".xlsx"
	_, err := os.Stat(fileName)
	if err == nil {
		// 存在 打开 计数
		fmt.Println("append")
		err = appendFile(fileName, jobs)
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "append success",
			"data": nil,
		})
		return
	}

	if os.IsNotExist(err) {
		fmt.Println("create")
		err = createFile(fileName, jobs)
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  err,
				"data": nil,
			})
			return
		}
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "create success",
			"data": nil,
		})
		return
	}
}

func createFile(fileName string, data []models.Job) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 插入
	if err := insertExcel(f, data, 0); err != nil {
		return err
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}

func appendFile(fileName string, data []models.Job) error {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 读取行数
	rows, err := f.GetRows("Sheet1")
	fmt.Println("rows:", len(rows))
	err = insertExcel(f, data, len(rows))
	if err != nil {
		return err
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}

// 从 start 位置开始插入数据
func insertExcel(f *excelize.File, data []models.Job, start int) error {

	_ = f.SetColWidth("Sheet1", "A", "A", 20)
	_ = f.SetColWidth("Sheet1", "B", "B", 35)
	_ = f.SetColWidth("Sheet1", "C", "C", 15)
	_ = f.SetColWidth("Sheet1", "F", "F", 20)
	_ = f.SetColWidth("Sheet1", "G", "H", 60)

	var tmp []string
	for idx, job := range data {
		row := idx + 1 + start
		// 加表头
		if start == 0 && idx == 0 {
			cell, err := excelize.CoordinatesToCellName(1, row)
			if err != nil {
				return err
			}
			tmp = append(tmp, "公司")
			tmp = append(tmp, "岗位")
			tmp = append(tmp, "薪资")
			tmp = append(tmp, "经验")
			tmp = append(tmp, "学历")
			tmp = append(tmp, "地区")
			tmp = append(tmp, "要求")
			tmp = append(tmp, "福利")
			tmp = append(tmp, "公司概况")
			if err := f.SetSheetRow("Sheet1", cell, &tmp); err != nil {
				return err
			}
			tmp = nil
			row++
		}

		// 添加数据
		{
			cell, err := excelize.CoordinatesToCellName(1, row)
			if err != nil {
				return err
			}
			tmp = append(tmp, job.Company.Name)
			tmp = append(tmp, job.Base.Name)
			tmp = append(tmp, job.Base.Salary)
			tmp = append(tmp, job.Base.Tags[0])
			tmp = append(tmp, job.Base.Tags[1])
			tmp = append(tmp, job.Base.Area)
			tmp = append(tmp, strings.Join(job.Requirement.Tags, ","))
			tmp = append(tmp, job.Company.Desc)
			tmp = append(tmp, strings.Join(job.Company.Tags, ","))
			// 设置行数据
			if err := f.SetSheetRow("Sheet1", cell, &tmp); err != nil {
				return err
			}

			// 超链接样式
			style, err := f.NewStyle(&excelize.Style{
				Font: &excelize.Font{Color: "#1265BE"},
			})
			if err != nil {
				return err
			}

			// 公司超链接
			{
				cell, err = excelize.CoordinatesToCellName(1, row)
				if err != nil {
					return err
				}
				if err := f.SetCellHyperLink("Sheet1",
					cell, job.Company.Link, "External",
					excelize.HyperlinkOpts{
						Display: &job.Company.Link,
						Tooltip: &job.Company.Name,
					}); err != nil {
					return err
				}
				_ = f.SetCellStyle("Sheet1", cell, cell, style)
			}

			// 岗位超链接
			{
				cell, err = excelize.CoordinatesToCellName(2, row)
				if err != nil {
					return err
				}
				if err := f.SetCellHyperLink("Sheet1",
					cell, job.Base.Link, "External",
					excelize.HyperlinkOpts{
						Display: &job.Base.Link,
						Tooltip: &job.Base.Name,
					}); err != nil {
					return err
				}
				err = f.SetCellStyle("Sheet1", cell, cell, style)
			}
		}

		tmp = nil

	}
	return nil
}
