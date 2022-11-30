package services

import (
	"database/sql"
	"errors"
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
	SearchIndex []string
	SearchInput string
}

func (s *Server) InsertOneObjertToDB(object interface{}) error {
	if err := s.gormDb.Create(object).Error; err != nil {
		s.sulog.Infof("insert one object to db failed, err:[%s], object: [%+v]",
			err.Error(), object)
		return err
	}
	return nil
}

func (s *Server) InsertObjectsToDB(objects []interface{}) error {
	if err := s.gormDb.CreateInBatches(objects, DEFAULT_BATCHES_SIZE).Error; err != nil {
		s.sulog.Infof("insert objects to db failed, err:[%s]",
			err.Error())
		return err
	}
	return nil
}

func (s *Server) QueryObjectsWithPage(params *QueryObjectsParams) (*sql.Rows, error) {
	switch params.SortType {
	case SORTTYPE_TIME_STR:
		offset := (params.Page - 1) * params.PageSize

		querySub := s.gormDb.Model(params.ModelStruct).Select("last_time").
			Order("last_time desc").Limit(1).Offset(int(offset))

		if len(params.SearchIndex) != 0 && len(params.SearchInput) != 0 {
			for i, v := range params.SearchIndex {
				if i == 0 {
					querySub = querySub.Where(v+" LIKE ?", "%"+params.SearchInput+"%")
					continue
				}
				querySub = querySub.Or(v+" LIKE ?", "%"+params.SearchInput+"%")
			}
		}

		return s.gormDb.Model(params.ModelStruct).Where("last_time <= (?)", querySub).
			Order("last_time desc").Limit(int(params.PageSize)).Rows()
	default:
		err := errors.New("the sort type does not exist")
		s.sulog.Infof("query objects with page from db failed, err:[%s]",
			err.Error())
		return nil, err
	}
}

func (s *Server) ScanRows(rows *sql.Rows, object interface{}) error {
	err := s.gormDb.ScanRows(rows, object)
	if err != nil {
		s.sulog.Infof("scan sql rows failed, err:[%s]",
			err.Error())
		return err
	}
	return nil
}
