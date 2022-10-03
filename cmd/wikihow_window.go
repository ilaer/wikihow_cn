package cmd

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"golang.org/x/image/colornames"
	"image/color"
	"wikihow_cn/crawler"
)

func (m *WikiHow) MainWindow() {
	ap := app.New()
	mw := ap.NewWindow("WikiHow")

	// 最外围的盒子,上下结构

	// 二. 下面的table部分

	// 1. 数据加载提示栏
	label1 := widget.NewLabel("搜索结果 : 0")
	label1.Move(fyne.NewPos(50, 35))
	label1.Resize(fyne.NewSize(500, 30))

	// 2. 数据展示表格

	//创建表格
	table1 := widget.NewTable(nil, nil, nil)

	// 获取表格的行数和列数
	table1.Length = func() (int, int) {
		return len(m.TableData), len(m.TableData[0])
	}

	// 创建表格单元
	table1.CreateCell = func() fyne.CanvasObject {
		item := widget.NewLabel("Template")
		item.Wrapping = fyne.TextWrapBreak
		item.Resize(fyne.Size{
			Width:  192,
			Height: m.TableColumnHeight,
		})
		return item
	}

	// 更新表格单元的数据
	table1.UpdateCell = func(cell widget.TableCellID, item fyne.CanvasObject) {

		val := m.TableData[cell.Row][cell.Col]
		switch cell.Col {
		case 0:
			if len(val) > 38 {
				item.(*widget.Label).SetText(fmt.Sprintf("%s", val[:38]))
			} else {
				item.(*widget.Label).SetText(fmt.Sprintf("%s", val))
			}
		case 1:

			item.(*widget.Label).SetText(fmt.Sprintf("%s", val))

		case 2:
			if len(val) > 38 {
				item.(*widget.Label).SetText(fmt.Sprintf("%s", val[:38]))
			} else {
				item.(*widget.Label).SetText(fmt.Sprintf("%s", val))
			}
		}

	}

	table1.OnSelected = func(id widget.TableCellID) {
		//出发选中单元格时,判断是否有对应的数据,是则复制到剪切板
		if id.Row < len(m.TableData) && id.Col < len(m.TableData[0]) {

			clipboard.WriteAll(m.TableData[id.Row][id.Col])

		}

	}

	table1.Move(fyne.NewPos(30, 80))
	table1.Resize(fyne.NewSize(960, 600))

	for idx, width := range m.TableColumnWidths {
		table1.SetColumnWidth(idx, width)
	}

	// 一. 上面的选择部分

	// 1. 搜索关键词输入
	kwEntry := widget.NewEntry()
	//kwEntry.Resize(fyne.NewSize(800, 600))
	kwEntry.SetPlaceHolder("請輸入")
	kwEntry.TextStyle = fyne.TextStyle{Bold: true}
	//kwEntry.Wrapping=fyne.TextTruncate
	kwEntry.Move(fyne.NewPos(350, 20))
	kwEntry.Resize(fyne.NewSize(400, 30))

	// 4. 搜索按钮
	submitBTN := widget.NewButton("查询", func() {
		kw := kwEntry.Text
		datas, _ := crawler.Search(kw)
		m.TableData = m.TableData[:1]
		m.TableData = append(m.TableData, datas...)
		table1.Refresh()
		label1.SetText(fmt.Sprintf("搜索结果 : %d", len(datas)))
	})
	submitBTN.Move(fyne.NewPos(800, 20))
	submitBTN.Resize(fyne.NewSize(80, 30))

	// 交互组件和数据展示组件的分割线
	Line1 := canvas.NewLine(color.Color(colornames.Black))
	Line1.StrokeWidth = 2

	Line1.Position1 = fyne.NewPos(30, 25)
	Line1.Position2 = fyne.NewPos(m.Width-30, 25)

	// 布局
	mw.SetContent(
		container.NewVBox(
			container.NewWithoutLayout(
				kwEntry,
				submitBTN,
			),
			container.NewWithoutLayout(
				Line1),
			container.NewWithoutLayout(

				label1,
				table1,
			),
		),
	)

	mw.Resize(fyne.NewSize(m.Width, m.Height))

	mw.ShowAndRun()
}
