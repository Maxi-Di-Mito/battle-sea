package utils

import "encoding/json"

func ToLog(obj interface{}) string {
	data, _ := json.Marshal(&obj)

	return string(data)
}

type CellType string

const (
	CELLTYPE_WATER CellType = "water"
	CELLTYPE_BOAT  CellType = "boat"
)
