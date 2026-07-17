package message

type MsgType string

const (
	MsgTypeHeartbeat                           MsgType = "0"
	MsgTypeTestRequest                         MsgType = "1"
	MsgTypeReject                              MsgType = "3"
	MsgTypeLogout                              MsgType = "5"
	MsgTypeExecutionReport                     MsgType = "8"
	MsgTypeOrderCancelReject                   MsgType = "9"
	MsgTypeLogon                               MsgType = "A"
	MsgTypeNewOrderSingle                      MsgType = "D"
	MsgTypeNewOrderList                        MsgType = "E"
	MsgTypeOrderCancelRequest                  MsgType = "F"
	MsgTypeListStatus                          MsgType = "N"
	MsgTypeOrderMassCancelRequest              MsgType = "q"
	MsgTypeOrderMassCancelReport               MsgType = "r"
	MsgTypeOrderCancelRequestAndNewOrderSingle MsgType = "XCN"
	MsgTypeLimitQuery                          MsgType = "XLQ"
	MsgTypeLimitResponse                       MsgType = "XLR"
	MsgTypeNews                                MsgType = "B"
	MsgTypeInstrumentListRequest               MsgType = "x"
	MsgTypeInstrumentList                      MsgType = "y"
	MsgTypeMarketDataRequest                   MsgType = "V"
	MsgTypeMarketDataRequestReject             MsgType = "Y"
	MsgTypeMarketDataSnapshot                  MsgType = "W"
	MsgTypeMarketDataIncrementalRefresh        MsgType = "X"
	MsgTypeOrderAmendKeepPriorityRequest       MsgType = "XAK"
	MsgTypeOrderAmendReject                    MsgType = "XAR"
)

func (m MsgType) IsValid() bool {
	switch m {
	case MsgTypeHeartbeat,
		MsgTypeTestRequest,
		MsgTypeReject,
		MsgTypeLogout,
		MsgTypeExecutionReport,
		MsgTypeOrderCancelReject,
		MsgTypeLogon,
		MsgTypeNewOrderSingle,
		MsgTypeNewOrderList,
		MsgTypeOrderCancelRequest,
		MsgTypeListStatus,
		MsgTypeOrderMassCancelRequest,
		MsgTypeOrderMassCancelReport,
		MsgTypeOrderCancelRequestAndNewOrderSingle,
		MsgTypeLimitQuery,
		MsgTypeLimitResponse,
		MsgTypeNews,
		MsgTypeInstrumentListRequest,
		MsgTypeInstrumentList,
		MsgTypeMarketDataRequest,
		MsgTypeMarketDataRequestReject,
		MsgTypeMarketDataSnapshot,
		MsgTypeMarketDataIncrementalRefresh,
		MsgTypeOrderAmendKeepPriorityRequest,
		MsgTypeOrderAmendReject:
		return true
	default:
		return false
	}
}
