package db

const BCIDENTITY_TABLE_NAME = "block_chain_identity"

type ChainRoleType int32

const (
	CHAIN_ADMIN ChainRoleType = iota + 1
	CHAIN_CLIENT
	CHAIN_LIGHT
	CHAIN_CONSENSUS
	CHAIN_COMMON
)

const (
	CHAIN_ADMIN_STR = "admin"

	CHAIN_CLIENT_STR = "client"

	CHAIN_LIGHT_STR = "light"

	CHAIN_CONSENSUS_STR = "consensus"

	CHAIN_COMMON_STR = "common"
)

var ChainRoleTypeName = map[ChainRoleType]string{
	CHAIN_ADMIN:     CHAIN_ADMIN_STR,
	CHAIN_CLIENT:    CHAIN_CLIENT_STR,
	CHAIN_LIGHT:     CHAIN_LIGHT_STR,
	CHAIN_CONSENSUS: CHAIN_CONSENSUS_STR,
	CHAIN_COMMON:    CHAIN_COMMON_STR,
}

var ChainRoleTypeValue = map[string]ChainRoleType{
	CHAIN_ADMIN_STR:     CHAIN_ADMIN,
	CHAIN_CLIENT_STR:    CHAIN_CLIENT,
	CHAIN_LIGHT_STR:     CHAIN_LIGHT,
	CHAIN_CONSENSUS_STR: CHAIN_CONSENSUS,
	CHAIN_COMMON_STR:    CHAIN_COMMON,
}

type CertUsageType int32

const (
	SIGN CertUsageType = iota + 1

	TLS

	TLS_SIGN

	TLS_ENC
)

const (
	SIGN_STR = "sign"

	TLS_STR = "tls"

	TLS_SIGN_STR = "tls-sign"

	TLS_ENC_STR = "tls-enc"
)

var CertUsageTypeName = map[CertUsageType]string{
	SIGN:     SIGN_STR,
	TLS:      TLS_STR,
	TLS_SIGN: TLS_SIGN_STR,
	TLS_ENC:  TLS_ENC_STR,
}

var CertUsageTypeValue = map[string]CertUsageType{
	SIGN_STR:     SIGN,
	TLS_STR:      TLS,
	TLS_SIGN_STR: TLS_SIGN,
	TLS_ENC_STR:  TLS_ENC,
}

type BlockChainIdentity struct {
	GeneralField
	UserId          int32
	ChainId         string
	ChainOrg        string
	ChainRole       ChainRoleType
	ChainCertUsage  CertUsageType
	IdentityContent string `gorm:"type:longtext"`
	PrivateKey      string `gorm:"type:longtext"`
}

func (b *BlockChainIdentity) TableName() string {
	return BCIDENTITY_TABLE_NAME
}

func init() {
	blockChainIdentity := new(BlockChainIdentity)
	TableSlice = append(TableSlice, &blockChainIdentity)
}
