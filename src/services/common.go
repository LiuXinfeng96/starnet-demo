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
	SortType    SortType
	SearchIndex []string
	SearchInput string
}

func (s *Server) InsertOneObjertToDB(object interface{}) error {
	if err := s.gormDb.Create(object).Error; err != nil {
		s.sulog.Infof("insert one object to db failed, err:[%s], object: [%+v]\n",
			err.Error(), object)
		return err
	}
	return nil
}

func (s *Server) InsertObjectsToDB(objects []interface{}) error {
	if err := s.gormDb.CreateInBatches(objects, DEFAULT_BATCHES_SIZE).Error; err != nil {
		s.sulog.Infof("insert objects to db failed, err:[%s]\n",
			err.Error())
		return err
	}
	return nil
}

func (s *Server) QueryObjectsWithPage(params *QueryObjectsParams) (*sql.Rows, error) {
	switch params.SortType {
	case SORTTYPE_TIME:
		offset := (params.Page - 1) * params.PageSize

		querySub := s.gormDb.Model(params.ModelStruct).Select("id").
			Order("last_time desc").Limit(int(params.PageSize)).Offset(int(offset))

		if len(params.SearchIndex) != 0 && len(params.SearchInput) != 0 {
			for i, v := range params.SearchIndex {
				if i == 0 {
					querySub = querySub.Where(v+" LIKE ?", "%"+params.SearchInput+"%")
					continue
				}
				querySub = querySub.Or(v+" LIKE ?", "%"+params.SearchInput+"%")
			}
		}

		return s.gormDb.Model(params.ModelStruct).Joins("inner join (?) as t2 using(id)", querySub).Rows()
	default:
		err := errors.New("the sort type does not exist")
		s.sulog.Infof("query objects with page from db failed, err:[%s]\n",
			err.Error())
		return nil, err
	}
}

func (s *Server) QueryObjectsWithPageSC(params *QueryObjectsParams) (*sql.Rows, error) {
	switch params.SortType {
	case SORTTYPE_TIME:
		offset := (params.Page - 1) * params.PageSize

		querySub := s.gormDb.Model(params.ModelStruct).Select("id").
			Where("chain_id = ?", s.satelliteChainId).
			Order("last_time desc").Limit(int(params.PageSize)).Offset(int(offset))

		if len(params.SearchIndex) != 0 && len(params.SearchInput) != 0 {
			for i, v := range params.SearchIndex {
				if i == 0 {
					querySub = querySub.Where(v+" LIKE ?", "%"+params.SearchInput+"%")
					continue
				}
				querySub = querySub.Or(v+" LIKE ?", "%"+params.SearchInput+"%")
			}
		}

		return s.gormDb.Model(params.ModelStruct).Joins("inner join (?) as t2 using(id)", querySub).Rows()
	default:
		err := errors.New("the sort type does not exist")
		s.sulog.Infof("query objects with page from db failed, err:[%s]\n",
			err.Error())
		return nil, err
	}
}

func (s *Server) ScanRows(rows *sql.Rows, object interface{}) error {
	err := s.gormDb.ScanRows(rows, object)
	if err != nil {
		s.sulog.Infof("scan sql rows failed, err:[%s]\n",
			err.Error())
		return err
	}
	return nil
}

func (s *Server) QueryObjectById(modelStruct interface{},
	id int32) error {

	if err := s.gormDb.Model(modelStruct).
		Where("id = ?", id).First(modelStruct).Error; err != nil {

		s.sulog.Infof("query object by id failed, err:[%s], object:[%+v]\n",
			err.Error(), modelStruct)
		return err
	}
	return nil
}

func (s *Server) QueryObjectByCondition(modelStruct interface{},
	searchIndex, searchInput string) error {

	if err := s.gormDb.Model(modelStruct).
		Where(searchIndex+" = ?", searchInput).First(modelStruct).Error; err != nil {

		s.sulog.Infof("query object by [%s] failed, err:[%s], object:[%+v]\n",
			searchIndex, err.Error(), modelStruct)
		return err
	}
	return nil
}

func (s *Server) QueryObjectsByCondition(modelStruct interface{},
	searchIndex, searchInput string) (*sql.Rows, error) {

	sqlRows, err := s.gormDb.Model(modelStruct).
		Where(searchIndex+" = ?", searchInput).
		Order("last_time desc").Rows()

	if err != nil {
		s.sulog.Infof("query objects by [%s] failed, err:[%s], object:[%+v]\n",
			searchIndex, err.Error(), modelStruct)
		return nil, err
	}

	return sqlRows, nil
}
