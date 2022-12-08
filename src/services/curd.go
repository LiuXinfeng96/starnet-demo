package services

import (
	"database/sql"
	"errors"
	"starnet-demo/src/db"
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
	QueryMap    map[string]string
}

type QueryLatestObjectsParams struct {
	ModelStruct db.ModelStruct
	Page        int32
	PageSize    int32
	SortType    SortType
	SearchIndex []string
	SearchInput string
	GroupIndex  string
}

func (s *Server) InsertOneObjertToDB(object interface{}) error {
	if err := s.gormDb.Create(object).Error; err != nil {
		s.sulog.Infof("insert one object to db failed, err:[%s], object: [%+v]\n",
			err.Error(), object)
		return err
	}
	return nil
}

func (s *Server) InsertObjectsToDB(objects interface{}) error {
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
			Limit(int(params.PageSize)).Offset(int(offset))

		if len(params.SearchIndex) != 0 && len(params.SearchInput) != 0 {
			for i, v := range params.SearchIndex {
				if i == 0 {
					querySub = querySub.Where(v+" LIKE ?", "%"+params.SearchInput+"%")
					continue
				}
				querySub = querySub.Or(v+" LIKE ?", "%"+params.SearchInput+"%")
			}
		}

		if len(params.QueryMap) != 0 {
			for k, v := range params.QueryMap {
				querySub = querySub.Where(k+" = ?", v)
			}
		}

		return s.gormDb.Model(params.ModelStruct).Order("last_time desc").
			Joins("inner join (?) as t2 using(id)", querySub).Rows()
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
			Where("chain_id = ?", s.GetExecChainId()).
			Limit(int(params.PageSize)).Offset(int(offset))

		if len(params.SearchIndex) != 0 && len(params.SearchInput) != 0 {
			for i, v := range params.SearchIndex {
				if i == 0 {
					querySub = querySub.Where(v+" LIKE ?", "%"+params.SearchInput+"%")
					continue
				}
				querySub = querySub.Or(v+" LIKE ?", "%"+params.SearchInput+"%")
			}
		}

		if len(params.QueryMap) != 0 {
			for k, v := range params.QueryMap {
				querySub = querySub.Where(k+" = ?", v)
			}
		}

		return s.gormDb.Model(params.ModelStruct).Order("last_time desc").
			Joins("inner join (?) as t2 using(id)", querySub).Rows()
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

func (s *Server) QueryLatestObjectsWithPage(params *QueryLatestObjectsParams) (*sql.Rows, error) {
	switch params.SortType {
	case SORTTYPE_TIME:
		offset := (params.Page - 1) * params.PageSize

		groupSub := s.gormDb.Model(params.ModelStruct).Select(params.GroupIndex + ",MAX(last_time) as latest_time").
			Group(params.GroupIndex)

		querySub := s.gormDb.Model(params.ModelStruct).Select("id").Order("last_time desc").
			Joins("inner join (?) as t2 on t2."+params.GroupIndex+
				" = "+params.ModelStruct.TableName()+"."+params.GroupIndex+
				" AND t2.latest_time = "+params.ModelStruct.TableName()+".last_time",
				groupSub).Limit(int(params.PageSize)).Offset(int(offset))

		querySub = s.gormDb.Model(params.ModelStruct).Order("last_time desc").
			Joins("inner join (?) as t2 using(id)", querySub)

		if len(params.SearchIndex) != 0 && len(params.SearchInput) != 0 {
			for i, v := range params.SearchIndex {
				if i == 0 {
					querySub = querySub.Where(v+" LIKE ?", "%"+params.SearchInput+"%")
					continue
				}
				querySub = querySub.Or(v+" LIKE ?", "%"+params.SearchInput+"%")
			}
		}
		return querySub.Rows()
	default:
		err := errors.New("the sort type does not exist")
		s.sulog.Infof("query objects with page from db failed, err:[%s]\n",
			err.Error())
		return nil, err
	}
}
