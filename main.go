package main

import (
	"os"
	"path/filepath"
	"wikihow_cn/cmd"
)

func main() {
	wikiHow := &cmd.WikiHow{}
	wikiHow.Width = 1024
	wikiHow.Height = 768
	wikiHow.TableHeader = []string{"Title", "View", "URL"}
	wikiHow.TableData = append(wikiHow.TableData, wikiHow.TableHeader)

	// table column width
	wikiHow.TableColumnWidths = []float32{300, 300, 350}
	// table column height
	wikiHow.TableColumnHeight = 30

	currentPath, _ := os.Getwd()
	os.Setenv("FYNE_FONT", filepath.Join(currentPath, "wqy-microhei.ttc"))
	wikiHow.MainWindow()
	os.Unsetenv("FYNE_FONT")
}
