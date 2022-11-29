package services

import (
	"errors"

	"gorm.io/gorm"
)

const DEFAULT_BATCHES_SIZE = 1000

type SortType int32

const (
	SORTTYPE_TIME SortType = iota + 1
)

const (
	SORTTYPE_TIME_STR = "最新时间"
)

var SortTypeName = map[SortType]string{
	SORTTYPE_TIME: SORTTYPE_TIME_STR,
}

var SortTypeValue = map[string]SortType{
	SORTTYPE_TIME_STR: SORTTYPE_TIME,
}

type QueryObjectsParams struct {
	ModelStruct interface{}
	Page        int32
	PageSize    int32
	SortType    string
	Condition   string
	Db          *gorm.DB
}

func InsertOneObjertToDB(object interface{}, db *gorm.DB) error {
	if err := db.Create(object).Error; err != nil {
		return err
	}
	return nil
}

func InsertObjectsToDB(objects []interface{}, db *gorm.DB) error {
	if err := db.CreateInBatches(objects, DEFAULT_BATCHES_SIZE).Error; err != nil {
		return err
	}
	return nil
}

func QueryObjectsWithPage(params *QueryObjectsParams) (*gorm.Rows, error) {
	switch params.SortType {
	case SORTTYPE_TIME_STR:
		params.Db.Model(params.ModelStruct).
	default:
		return nil, errors.New("the sort type does not exist")
	}
}
