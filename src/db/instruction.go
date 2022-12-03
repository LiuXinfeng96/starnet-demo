package db

const INSTRUCTION_TABLE_NAME = "instruction"

type InstructionType int32

const (
	OPERATION InstructionType = iota + 1

	AUTOMATIC
)

const (
	OPERATION_STR = "操作规避"

	AUTOMATIC_STR = "自主规避"
)

var InstructionTypeName = map[InstructionType]string{
	OPERATION: OPERATION_STR,
	AUTOMATIC: AUTOMATIC_STR,
}

var InstructionTypeValue = map[string]InstructionType{
	OPERATION_STR: OPERATION,
	AUTOMATIC_STR: AUTOMATIC,
}

type InstructionExecState int32

const (
	NOTEXEC InstructionExecState = iota + 1
	INEXEC
	EXECSUCCESS
	EXECFAIL
)

const (
	NOTEXEC_STR = "未执行"

	INEXEC_STR = "执行中"

	EXECSUCCESS_STR = "执行成功"

	EXECFAIL_STR = "执行失败"
)

var ExecStateName = map[InstructionExecState]string{
	NOTEXEC:     NOTEXEC_STR,
	INEXEC:      INEXEC_STR,
	EXECSUCCESS: EXECSUCCESS_STR,
	EXECFAIL:    EXECFAIL_STR,
}

var ExecStateValue = map[string]InstructionExecState{
	NOTEXEC_STR:     NOTEXEC,
	INEXEC_STR:      INEXEC,
	EXECSUCCESS_STR: EXECSUCCESS,
	EXECFAIL_STR:    EXECFAIL,
}

type Instruction struct {
	GeneralField
	BlockChainField
	InstructionId       string `gorm:"index"`
	Type                InstructionType
	InstructionContent  string
	InstructionSource   string
	ExecInstructionTime int64
	GenInstructionTime  int64
	DebrisId            string
	DebrisName          string
	SatelliteId         string `gorm:"index"`
	SatelliteName       string `gorm:"index"`
	ExecState           InstructionExecState
}

func (i *Instruction) TableName() string {
	return INSTRUCTION_TABLE_NAME
}

func init() {
	instruction := new(Instruction)
	TableSlice = append(TableSlice, &instruction)
}
