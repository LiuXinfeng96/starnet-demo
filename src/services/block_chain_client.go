package services

import (
	"encoding/json"
	"errors"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"sync"
	"time"

	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"github.com/golang/groupcache/lru"
)

type SdkPool struct {
	cache *lru.Cache

	mutex sync.Mutex
}

func NewSdkPool(maxEntries int) *SdkPool {
	var sdkPool SdkPool
	sdkPool.cache = lru.New(maxEntries)
	sdkPool.cache.OnEvicted = func(key lru.Key, value interface{}) {
		client, ok := value.(sdk.ChainClient)
		if !ok {
			return
		}
		client.Stop()
	}

	return &sdkPool
}

func (s *SdkPool) Add(key string, client *sdk.ChainClient) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Add(key, client)
}

func (s *SdkPool) Get(key string) (client *sdk.ChainClient, ok bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	v, ok := s.cache.Get(key)

	client, ok = v.(*sdk.ChainClient)

	return
}

func (s *SdkPool) Remove(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Remove(key)
}

func (s *SdkPool) RemoveOldest() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.RemoveOldest()
}

func (s *SdkPool) Len() int {
	return s.cache.Len()
}

func (s *SdkPool) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Clear()
}

func (s *Server) InitChainClient() {

	masterName := s.config.BCConfig[0].UserName
	user := new(db.User)
	err := s.QueryObjectByCondition(user, "user_name", masterName)
	if err != nil {
		materUser := &db.User{
			UserName:     masterName,
			UserRole:     db.CONTROL,
			UserPwd:      s.config.BCConfig[0].UserPwd,
			UserNickName: "主链用户",
			UserPhoneNum: "",
			UserEmail:    "",
		}
		err := s.InsertOneObjertToDB(materUser)
		if err != nil {
			panic(err)
		}
	}

	execName := s.config.BCConfig[1].UserName
	user = new(db.User)
	err = s.QueryObjectByCondition(user, "user_name", execName)
	if err != nil {
		execUser := &db.User{
			UserName:     execName,
			UserRole:     db.EXEC,
			UserPwd:      s.config.BCConfig[1].UserPwd,
			UserNickName: "星座链用户",
			UserPhoneNum: "",
			UserEmail:    "",
		}
		err := s.InsertOneObjertToDB(execUser)
		if err != nil {
			panic(err)
		}
	}

	masterClient, err := sdk.NewChainClient(
		sdk.WithConfPath(s.config.BCConfig[0].SdkConfigPath),
		sdk.WithChainClientLogger(s.sulog),
	)
	if err != nil {
		panic(err)
	}

	execClient, err := sdk.NewChainClient(
		sdk.WithConfPath(s.config.BCConfig[1].SdkConfigPath),
		sdk.WithChainClientLogger(s.sulog),
	)
	if err != nil {
		panic(err)
	}

	s.sdkPool.Add(masterName+s.GetMasterChainId(), masterClient)
	s.sdkPool.Add(masterName+s.GetExecChainId(), execClient)

	s.sdkPool.Add(execName+s.GetExecChainId(), execClient)
	s.sdkPool.Add(execName+s.GetMasterChainId(), masterClient)

}

func (s *Server) GetSdkClient(userName string) (*sdk.ChainClient, error) {
	client, ok := s.sdkPool.Get(userName)
	if !ok {
		return nil, errors.New("the chain client not found")
	}

	return client, nil
}

func (s *Server) GetMasterChainUserName() string {
	return s.config.BCConfig[0].UserName
}

func (s *Server) GetExecChainUserName() string {
	return s.config.BCConfig[1].UserName
}

func (s *Server) SendTxToBlockChain(contractName, funcName string, client *sdk.ChainClient,
	kvs []*common.KeyValuePair, model db.ModelStruct, blockChainField *db.BlockChainField) {
	i := 0
	for {
		resp, err := client.InvokeContract(contractName,
			funcName, "", kvs, -1, true)
		if err != nil {
			s.GetSuLogger().Warnf("send tx to the chain failed: [%s]\n", err.Error())
			if i < s.retryTime {
				i++
				s.GetSuLogger().Warnf("wait [%d]ms\n", s.retryInterval/time.Millisecond)
				time.Sleep(s.retryInterval)
				s.GetSuLogger().Warnf("retry send tx, time: [%d]\n", i)
				continue
			}
			return
		}

		if resp.Code != common.TxStatusCode_SUCCESS {
			s.GetSuLogger().Warnf("invoke contract failed: tx status code: [%d], msg: [%s]\n",
				resp.Code, resp.Message)
			return
		}

		if blockChainField != nil {
			*blockChainField, err = getBlockChainFiledFromResp(resp.ContractResult)
			if err != nil {
				s.GetSuLogger().Warnf("get block chain info from contract resp failed, err: [%s], result: [%+v]",
					err.Error(), resp.ContractResult)
				return
			}
		}

		if model != nil {
			s.UpdateObject(model)
		}

		return
	}

}

func getBlockChainFiledFromResp(cr *common.ContractResult) (db.BlockChainField, error) {
	var resp models.ContractResp
	err := json.Unmarshal(cr.Result, &resp)
	if err != nil {
		return db.BlockChainField{}, err
	}
	var bcFiled db.BlockChainField
	bcFiled.TxId = resp.TxId
	bcFiled.BlockHeight = resp.BlockHeight
	bcFiled.ChainTime = time.Now().Unix()
	return bcFiled, nil
}
