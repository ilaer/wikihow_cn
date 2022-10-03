package cmd

type WikiHow struct {
	Width             float32    `ini:"gui_width"`
	Height            float32    `ini:"gui_height"`
	TableHeader       []string   `json:"table_headers"`
	TableData         [][]string `json:"table_data"`
	TableColumnWidths []float32  `json:"table_column_widths"`
	TableColumnHeight float32    `json:"table_column_height"`
}
