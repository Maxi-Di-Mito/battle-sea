package utils

import "encoding/json"

func ToLog(obj interface{}) string {
	data, _ := json.Marshal(&obj)

	return string(data)
}

type CellValue string

const (
	CELLVALUE_WATER   CellValue = "water"
	CELLVALUE_UNKNOWN CellValue = "unknown"
	CELLVALUE_BOAT    CellValue = "boat"
)
