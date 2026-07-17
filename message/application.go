package message

import (
	"time"
)

const (
	TagClOrdID                                 Tag = 11
	TagCumQty                                  Tag = 14
	TagCurrency                                Tag = 15
	TagExecID                                  Tag = 17
	TagExecInst                                Tag = 18
	TagLastPx                                  Tag = 31
	TagLastQty                                 Tag = 32
	TagOrderID                                 Tag = 37
	TagOrderQty                                Tag = 38
	TagOrdStatus                               Tag = 39
	TagOrdType                                 Tag = 40
	TagOrigClOrdID                             Tag = 41
	TagPrice                                   Tag = 44
	TagSide                                    Tag = 54
	TagSymbol                                  Tag = 55
	TagTimeInForce                             Tag = 59
	TagTransactTime                            Tag = 60
	TagListID                                  Tag = 66
	TagAllocID                                 Tag = 70
	TagNoOrders                                Tag = 73
	TagOrdRejReason                            Tag = 103
	TagMaxFloor                                Tag = 111
	TagNoMiscFees                              Tag = 136
	TagMiscFeeAmt                              Tag = 137
	TagMiscFeeCurr                             Tag = 138
	TagMiscFeeType                             Tag = 139
	TagNoRelatedSym                            Tag = 146
	TagExecType                                Tag = 150
	TagLeavesQty                               Tag = 151
	TagCashOrderQty                            Tag = 152
	TagPegOffsetValue                          Tag = 211
	TagMDReqID                                 Tag = 262
	TagSubscriptionRequestType                 Tag = 263
	TagMarketDepth                             Tag = 264
	TagAggregatedBook                          Tag = 266
	TagNoMDEntryTypes                          Tag = 267
	TagNoMDEntries                             Tag = 268
	TagMDEntryType                             Tag = 269
	TagMDEntryPx                               Tag = 270
	TagMDEntrySize                             Tag = 271
	TagMDUpdateAction                          Tag = 279
	TagMDReqRejReason                          Tag = 281
	TagInstrumentReqID                         Tag = 320
	TagListStatusType                          Tag = 429
	TagListOrderStatus                         Tag = 431
	TagCxlRejResponseTo                        Tag = 434
	TagMassCancelRequestType                   Tag = 530
	TagMassCancelResponse                      Tag = 531
	TagMassCancelRejectReason                  Tag = 532
	TagTotalAffectedOrders                     Tag = 533
	TagInstrumentListRequestType               Tag = 559
	TagMinTradeVol                             Tag = 562
	TagMatchType                               Tag = 574
	TagWorkingIndicator                        Tag = 636
	TagPegMoveType                             Tag = 835
	TagPegOffsetType                           Tag = 836
	TagPeggedPrice                             Tag = 839
	TagTargetStrategy                          Tag = 847
	TagMinPriceIncrement                       Tag = 969
	TagTradeID                                 Tag = 1003
	TagAggressorIndicator                      Tag = 1057
	TagPegPriceType                            Tag = 1094
	TagTriggerType                             Tag = 1100
	TagTriggerAction                           Tag = 1101
	TagTriggerPrice                            Tag = 1102
	TagTriggerPriceType                        Tag = 1107
	TagTriggerPriceDirection                   Tag = 1109
	TagMaxTradeVol                             Tag = 1140
	TagContingencyType                         Tag = 1385
	TagListRejectReason                        Tag = 1386
	TagAggressorSide                           Tag = 2446
	TagStartPriceRange                         Tag = 2551
	TagEndPriceRange                           Tag = 2552
	TagReqID                                   Tag = 6136
	TagStrategyID                              Tag = 7940
	TagSelfTradePreventionMode                 Tag = 25001
	TagCancelRestrictions                      Tag = 25002
	TagNoLimitIndicators                       Tag = 25003
	TagLimitType                               Tag = 25004
	TagLimitCount                              Tag = 25005
	TagLimitMax                                Tag = 25006
	TagLimitResetInterval                      Tag = 25007
	TagLimitResetIntervalResolution            Tag = 25008
	TagTriggerTrailingDeltaBips                Tag = 25009
	TagNoListTriggeringInstructions            Tag = 25010
	TagListTriggerType                         Tag = 25011
	TagListTriggerTriggerIndex                 Tag = 25012
	TagListTriggerAction                       Tag = 25013
	TagClListID                                Tag = 25014
	TagOrigClListID                            Tag = 25015
	TagCumQuoteQty                             Tag = 25017
	TagOrderCreationTime                       Tag = 25018
	TagWorkingFloor                            Tag = 25021
	TagTrailingTime                            Tag = 25022
	TagWorkingTime                             Tag = 25023
	TagPreventedMatchID                        Tag = 25024
	TagPreventedExecutionPrice                 Tag = 25025
	TagPreventedExecutionQty                   Tag = 25026
	TagTradeGroupID                            Tag = 25027
	TagCounterSymbol                           Tag = 25028
	TagCounterOrderID                          Tag = 25029
	TagPreventedQty                            Tag = 25030
	TagLastPreventedQty                        Tag = 25031
	TagSOR                                     Tag = 25032
	TagOrderCancelRequestAndNewOrderSingleMode Tag = 25033
	TagCancelClOrdID                           Tag = 25034
	TagOrderRateLimitExceededMode              Tag = 25038
	TagMinQtyIncrement                         Tag = 25039
	TagMarketMinTradeVol                       Tag = 25040
	TagMarketMaxTradeVol                       Tag = 25041
	TagMarketMinQtyIncrement                   Tag = 25042
	TagFirstBookUpdateID                       Tag = 25043
	TagLastBookUpdateID                        Tag = 25044
	TagOPO                                     Tag = 25046
	TagExpiryReason                            Tag = 25056
)

type Side string

const (
	SideBuy  Side = "1"
	SideSell Side = "2"
)

type OrdType string

const (
	OrdTypeMarket    OrdType = "1"
	OrdTypeLimit     OrdType = "2"
	OrdTypeStop      OrdType = "3"
	OrdTypeStopLimit OrdType = "4"
	OrdTypePegged    OrdType = "P"
)

type ExecInst string

const (
	ExecInstParticipateDontInitiate ExecInst = "6"
)

type TimeInForce string

const (
	TimeInForceGoodTillCancel    TimeInForce = "1"
	TimeInForceImmediateOrCancel TimeInForce = "3"
	TimeInForceFillOrKill        TimeInForce = "4"
)

type SelfTradePreventionMode string

const (
	SelfTradePreventionModeNone        SelfTradePreventionMode = "1"
	SelfTradePreventionModeExpireTaker SelfTradePreventionMode = "2"
	SelfTradePreventionModeExpireMaker SelfTradePreventionMode = "3"
	SelfTradePreventionModeExpireBoth  SelfTradePreventionMode = "4"
	SelfTradePreventionModeDecrement   SelfTradePreventionMode = "5"
	SelfTradePreventionModeTransfer    SelfTradePreventionMode = "6"
)

type PegPriceType string

const (
	PegPriceTypeMarketPeg  PegPriceType = "4"
	PegPriceTypePrimaryPeg PegPriceType = "5"
)

type PegMoveType string

const (
	PegMoveTypeFixed PegMoveType = "1"
)

type PegOffsetType string

const (
	PegOffsetTypePriceTier PegOffsetType = "3"
)

type TriggerType string

const (
	TriggerTypePriceMovement TriggerType = "4"
)

type TriggerAction string

const (
	TriggerActionActivate TriggerAction = "1"
)

type TriggerPriceType string

const (
	TriggerPriceTypeLastTrade TriggerPriceType = "2"
)

type TriggerPriceDirection string

const (
	TriggerPriceDirectionUp   TriggerPriceDirection = "U"
	TriggerPriceDirectionDown TriggerPriceDirection = "D"
)

type ExecType string

const (
	ExecTypeNew      ExecType = "0"
	ExecTypeCanceled ExecType = "4"
	ExecTypeReplaced ExecType = "5"
	ExecTypeRejected ExecType = "8"
	ExecTypeTrade    ExecType = "F"
	ExecTypeExpired  ExecType = "C"
)

type OrdStatus string

const (
	OrdStatusNew             OrdStatus = "0"
	OrdStatusPartiallyFilled OrdStatus = "1"
	OrdStatusFilled          OrdStatus = "2"
	OrdStatusCanceled        OrdStatus = "4"
	OrdStatusPendingCancel   OrdStatus = "6"
	OrdStatusRejected        OrdStatus = "8"
	OrdStatusPendingNew      OrdStatus = "A"
	OrdStatusExpired         OrdStatus = "C"
)

type MiscFeeType string

const (
	MiscFeeTypeExchangeFees MiscFeeType = "4"
)

type CancelRestrictions string

const (
	CancelRestrictionsOnlyNew             CancelRestrictions = "1"
	CancelRestrictionsOnlyPartiallyFilled CancelRestrictions = "2"
)

type CxlRejResponseTo string

const (
	CxlRejResponseToOrderCancelRequest CxlRejResponseTo = "1"
)

type OrderCancelRequestAndNewOrderMode string

const (
	OrderCancelRequestAndNewOrderModeStopOnFailure OrderCancelRequestAndNewOrderMode = "1"
	OrderCancelRequestAndNewOrderModeAllowFailure  OrderCancelRequestAndNewOrderMode = "2"
)

type OrderRateLimitExceededMode string

const (
	OrderRateLimitExceededModeDoNothing  OrderRateLimitExceededMode = "1"
	OrderRateLimitExceededModeCancelOnly OrderRateLimitExceededMode = "2"
)

type MassCancelRequestType string

const (
	MassCancelRequestTypeCancelSymbolOrders MassCancelRequestType = "1"
)

type MassCancelResponse string

const (
	MassCancelResponseCancelRequestRejected MassCancelResponse = "0"
	MassCancelResponseCancelSymbolOrders    MassCancelResponse = "1"
)

type MassCancelRejectReason string

const (
	MassCancelRejectReasonOther MassCancelRejectReason = "99"
)

type ContingencyType string

const (
	ContingencyTypeOneCancelsTheOther  ContingencyType = "1"
	ContingencyTypeOneTriggersTheOther ContingencyType = "2"
)

type ListStatusType string

const (
	ListStatusTypeResponse    ListStatusType = "2"
	ListStatusTypeExecStarted ListStatusType = "4"
	ListStatusTypeAllDone     ListStatusType = "5"
	ListStatusTypeUpdated     ListStatusType = "100"
)

type ListOrderStatus string

const (
	ListOrderStatusExecuting ListOrderStatus = "3"
	ListOrderStatusAllDone   ListOrderStatus = "6"
	ListOrderStatusReject    ListOrderStatus = "7"
)

type ListRejectReason string

const (
	ListRejectReasonOther ListRejectReason = "99"
)

type OrdRejReason string

const (
	OrdRejReasonOther OrdRejReason = "99"
)

type ListTriggerType string

const (
	ListTriggerTypeActivated       ListTriggerType = "1"
	ListTriggerTypePartiallyFilled ListTriggerType = "2"
	ListTriggerTypeFilled          ListTriggerType = "3"
)

type ListTriggerAction string

const (
	ListTriggerActionRelease ListTriggerAction = "1"
	ListTriggerActionCancel  ListTriggerAction = "2"
)

type LimitType string

const (
	LimitTypeOrder   LimitType = "1"
	LimitTypeMessage LimitType = "2"
)

type LimitResetIntervalResolution string

const (
	LimitResetIntervalResolutionSecond LimitResetIntervalResolution = "s"
	LimitResetIntervalResolutionMinute LimitResetIntervalResolution = "m"
	LimitResetIntervalResolutionHour   LimitResetIntervalResolution = "h"
	LimitResetIntervalResolutionDay    LimitResetIntervalResolution = "d"
)

type InstrumentListRequestType string

const (
	InstrumentListRequestTypeSingleInstrument InstrumentListRequestType = "0"
	InstrumentListRequestTypeAllInstruments   InstrumentListRequestType = "4"
)

type SubscriptionRequestType string

const (
	SubscriptionRequestTypeSubscribe   SubscriptionRequestType = "1"
	SubscriptionRequestTypeUnsubscribe SubscriptionRequestType = "2"
)

type MDEntryType string

const (
	MDEntryTypeBid   MDEntryType = "0"
	MDEntryTypeOffer MDEntryType = "1"
	MDEntryTypeTrade MDEntryType = "2"
)

type MDReqRejReason string

const (
	MDReqRejReasonDuplicateMDReqID     MDReqRejReason = "1"
	MDReqRejReasonTooManySubscriptions MDReqRejReason = "2"
)

type MDUpdateAction string

const (
	MDUpdateActionNew    MDUpdateAction = "0"
	MDUpdateActionChange MDUpdateAction = "1"
	MDUpdateActionDelete MDUpdateAction = "2"
)

type OrderFields struct {
	ClOrdID                  string
	OrderQty                 string
	OrdType                  OrdType
	ExecInst                 ExecInst
	Price                    string
	Side                     Side
	Symbol                   string
	TimeInForce              TimeInForce
	MaxFloor                 string
	CashOrderQty             string
	TargetStrategy           string
	StrategyID               string
	SelfTradePreventionMode  SelfTradePreventionMode
	PegOffsetValue           string
	PegPriceType             PegPriceType
	PegMoveType              PegMoveType
	PegOffsetType            PegOffsetType
	TriggerType              TriggerType
	TriggerAction            TriggerAction
	TriggerPrice             string
	TriggerPriceType         TriggerPriceType
	TriggerPriceDirection    TriggerPriceDirection
	TriggerTrailingDeltaBips string
	SOR                      *bool
}

type NewOrderSingle struct {
	OrderFields
}

func NewNewOrderSingle(clOrdID string, ordType OrdType, side Side, symbol string) *NewOrderSingle {
	return &NewOrderSingle{
		OrderFields: OrderFields{
			ClOrdID: clOrdID,
			OrdType: ordType,
			Side:    side,
			Symbol:  symbol,
		},
	}
}

func (r *NewOrderSingle) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeNewOrderSingle, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime)
	setOrderFields(builder, r.OrderFields, true)
	return builder.Build()
}

type ExecutionReport struct {
	ExecID                   string
	ClOrdID                  string
	OrigClOrdID              string
	OrderID                  string
	OrderQty                 string
	OrdType                  OrdType
	Side                     Side
	Symbol                   string
	ExecInst                 ExecInst
	Price                    string
	TimeInForce              TimeInForce
	TransactTime             string
	OrderCreationTime        string
	MaxFloor                 string
	ListID                   string
	CashOrderQty             string
	TargetStrategy           string
	StrategyID               string
	SelfTradePreventionMode  SelfTradePreventionMode
	ExecType                 ExecType
	CumQty                   string
	LeavesQty                string
	CumQuoteQty              string
	AggressorIndicator       string
	TradeID                  string
	LastPx                   string
	LastQty                  string
	OrdStatus                OrdStatus
	AllocID                  string
	MatchType                string
	WorkingFloor             string
	TrailingTime             string
	WorkingIndicator         string
	WorkingTime              string
	PreventedMatchID         string
	PreventedExecutionPrice  string
	PreventedExecutionQty    string
	TradeGroupID             string
	CounterSymbol            string
	CounterOrderID           string
	PreventedQty             string
	LastPreventedQty         string
	SOR                      string
	ErrorCode                string
	Text                     string
	NoMiscFees               string
	TriggerType              TriggerType
	TriggerAction            TriggerAction
	TriggerPrice             string
	TriggerPriceType         TriggerPriceType
	TriggerPriceDirection    TriggerPriceDirection
	TriggerTrailingDeltaBips string
	PegOffsetValue           string
	PegPriceType             PegPriceType
	PegMoveType              PegMoveType
	PegOffsetType            PegOffsetType
	PeggedPrice              string
	ExpiryReason             string
}

func NewExecutionReport() *ExecutionReport {
	return new(ExecutionReport)
}

func (r *ExecutionReport) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeExecutionReport); err != nil {
		return err
	}

	ordType, err := requiredString(m, TagOrdType)
	if err != nil {
		return err
	}
	side, err := requiredString(m, TagSide)
	if err != nil {
		return err
	}
	symbol, err := requiredString(m, TagSymbol)
	if err != nil {
		return err
	}
	execType, err := requiredString(m, TagExecType)
	if err != nil {
		return err
	}
	cumQty, err := requiredString(m, TagCumQty)
	if err != nil {
		return err
	}
	lastQty, err := requiredString(m, TagLastQty)
	if err != nil {
		return err
	}
	ordStatus, err := requiredString(m, TagOrdStatus)
	if err != nil {
		return err
	}

	r.ExecID = optionalString(m, TagExecID)
	r.ClOrdID = optionalString(m, TagClOrdID)
	r.OrigClOrdID = optionalString(m, TagOrigClOrdID)
	r.OrderID = optionalString(m, TagOrderID)
	r.OrderQty = optionalString(m, TagOrderQty)
	r.OrdType = OrdType(ordType)
	r.Side = Side(side)
	r.Symbol = symbol
	r.ExecInst = ExecInst(optionalString(m, TagExecInst))
	r.Price = optionalString(m, TagPrice)
	r.TimeInForce = TimeInForce(optionalString(m, TagTimeInForce))
	r.TransactTime = optionalString(m, TagTransactTime)
	r.OrderCreationTime = optionalString(m, TagOrderCreationTime)
	r.MaxFloor = optionalString(m, TagMaxFloor)
	r.ListID = optionalString(m, TagListID)
	r.CashOrderQty = optionalString(m, TagCashOrderQty)
	r.TargetStrategy = optionalString(m, TagTargetStrategy)
	r.StrategyID = optionalString(m, TagStrategyID)
	r.SelfTradePreventionMode = SelfTradePreventionMode(optionalString(m, TagSelfTradePreventionMode))
	r.ExecType = ExecType(execType)
	r.CumQty = cumQty
	r.LeavesQty = optionalString(m, TagLeavesQty)
	r.CumQuoteQty = optionalString(m, TagCumQuoteQty)
	r.AggressorIndicator = optionalString(m, TagAggressorIndicator)
	r.TradeID = optionalString(m, TagTradeID)
	r.LastPx = optionalString(m, TagLastPx)
	r.LastQty = lastQty
	r.OrdStatus = OrdStatus(ordStatus)
	r.AllocID = optionalString(m, TagAllocID)
	r.MatchType = optionalString(m, TagMatchType)
	r.WorkingFloor = optionalString(m, TagWorkingFloor)
	r.TrailingTime = optionalString(m, TagTrailingTime)
	r.WorkingIndicator = optionalString(m, TagWorkingIndicator)
	r.WorkingTime = optionalString(m, TagWorkingTime)
	r.PreventedMatchID = optionalString(m, TagPreventedMatchID)
	r.PreventedExecutionPrice = optionalString(m, TagPreventedExecutionPrice)
	r.PreventedExecutionQty = optionalString(m, TagPreventedExecutionQty)
	r.TradeGroupID = optionalString(m, TagTradeGroupID)
	r.CounterSymbol = optionalString(m, TagCounterSymbol)
	r.CounterOrderID = optionalString(m, TagCounterOrderID)
	r.PreventedQty = optionalString(m, TagPreventedQty)
	r.LastPreventedQty = optionalString(m, TagLastPreventedQty)
	r.SOR = optionalString(m, TagSOR)
	r.ErrorCode = optionalString(m, TagErrorCode)
	r.Text = optionalString(m, TagText)
	r.NoMiscFees = optionalString(m, TagNoMiscFees)
	r.TriggerType = TriggerType(optionalString(m, TagTriggerType))
	r.TriggerAction = TriggerAction(optionalString(m, TagTriggerAction))
	r.TriggerPrice = optionalString(m, TagTriggerPrice)
	r.TriggerPriceType = TriggerPriceType(optionalString(m, TagTriggerPriceType))
	r.TriggerPriceDirection = TriggerPriceDirection(optionalString(m, TagTriggerPriceDirection))
	r.TriggerTrailingDeltaBips = optionalString(m, TagTriggerTrailingDeltaBips)
	r.PegOffsetValue = optionalString(m, TagPegOffsetValue)
	r.PegPriceType = PegPriceType(optionalString(m, TagPegPriceType))
	r.PegMoveType = PegMoveType(optionalString(m, TagPegMoveType))
	r.PegOffsetType = PegOffsetType(optionalString(m, TagPegOffsetType))
	r.PeggedPrice = optionalString(m, TagPeggedPrice)
	r.ExpiryReason = optionalString(m, TagExpiryReason)
	return nil
}

func (r *ExecutionReport) Error() string {
	if r == nil {
		return "order rejected"
	}
	return rejectErrorMessage(
		"order rejected",
		errorPart("clOrdID", r.ClOrdID),
		errorPart("origClOrdID", r.OrigClOrdID),
		errorPart("orderID", r.OrderID),
		errorPart("symbol", r.Symbol),
		errorPart("execType", string(r.ExecType)),
		errorPart("ordStatus", string(r.OrdStatus)),
		errorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type OrderCancelRequest struct {
	ClOrdID            string
	OrigClOrdID        string
	OrderID            string
	OrigClListID       string
	ListID             string
	Symbol             string
	CancelRestrictions CancelRestrictions
}

func NewOrderCancelRequest(clOrdID string, symbol string) *OrderCancelRequest {
	return &OrderCancelRequest{ClOrdID: clOrdID, Symbol: symbol}
}

func (r *OrderCancelRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeOrderCancelRequest, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagClOrdID, r.ClOrdID).
		SetField(TagSymbol, r.Symbol)
	setStringField(builder, TagOrigClOrdID, r.OrigClOrdID)
	setStringField(builder, TagOrderID, r.OrderID)
	setStringField(builder, TagOrigClListID, r.OrigClListID)
	setStringField(builder, TagListID, r.ListID)
	setStringField(builder, TagCancelRestrictions, string(r.CancelRestrictions))
	return builder.Build()
}

type OrderCancelReject struct {
	ClOrdID            string
	OrigClOrdID        string
	OrderID            string
	OrigClListID       string
	ListID             string
	Symbol             string
	CancelRestrictions CancelRestrictions
	CxlRejResponseTo   CxlRejResponseTo
	ErrorCode          string
	Text               string
}

func NewOrderCancelReject() *OrderCancelReject {
	return new(OrderCancelReject)
}

func (r *OrderCancelReject) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeOrderCancelReject); err != nil {
		return err
	}

	clOrdID, err := requiredString(m, TagClOrdID)
	if err != nil {
		return err
	}
	symbol, err := requiredString(m, TagSymbol)
	if err != nil {
		return err
	}
	cxlRejResponseTo, err := requiredString(m, TagCxlRejResponseTo)
	if err != nil {
		return err
	}
	errorCode, err := requiredString(m, TagErrorCode)
	if err != nil {
		return err
	}
	text, err := requiredString(m, TagText)
	if err != nil {
		return err
	}

	r.ClOrdID = clOrdID
	r.OrigClOrdID = optionalString(m, TagOrigClOrdID)
	r.OrderID = optionalString(m, TagOrderID)
	r.OrigClListID = optionalString(m, TagOrigClListID)
	r.ListID = optionalString(m, TagListID)
	r.Symbol = symbol
	r.CancelRestrictions = CancelRestrictions(optionalString(m, TagCancelRestrictions))
	r.CxlRejResponseTo = CxlRejResponseTo(cxlRejResponseTo)
	r.ErrorCode = errorCode
	r.Text = text
	return nil
}

func (r *OrderCancelReject) Error() string {
	if r == nil {
		return "order cancel reject"
	}
	return rejectErrorMessage(
		"order cancel reject",
		errorPart("clOrdID", r.ClOrdID),
		errorPart("origClOrdID", r.OrigClOrdID),
		errorPart("orderID", r.OrderID),
		errorPart("symbol", r.Symbol),
		errorPart("responseTo", string(r.CxlRejResponseTo)),
		errorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type OrderCancelRequestAndNewOrderSingle struct {
	Mode                       OrderCancelRequestAndNewOrderMode
	OrderRateLimitExceededMode OrderRateLimitExceededMode
	OrderID                    string
	CancelClOrdID              string
	OrigClOrdID                string
	CancelRestrictions         CancelRestrictions
	OrderFields
}

func NewOrderCancelRequestAndNewOrderSingle(mode OrderCancelRequestAndNewOrderMode, clOrdID string, ordType OrdType, side Side, symbol string) *OrderCancelRequestAndNewOrderSingle {
	return &OrderCancelRequestAndNewOrderSingle{
		Mode: mode,
		OrderFields: OrderFields{
			ClOrdID: clOrdID,
			OrdType: ordType,
			Side:    side,
			Symbol:  symbol,
		},
	}
}

func (r *OrderCancelRequestAndNewOrderSingle) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeOrderCancelRequestAndNewOrderSingle, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagOrderCancelRequestAndNewOrderSingleMode, string(r.Mode))
	setStringField(builder, TagOrderRateLimitExceededMode, string(r.OrderRateLimitExceededMode))
	setStringField(builder, TagOrderID, r.OrderID)
	setStringField(builder, TagCancelClOrdID, r.CancelClOrdID)
	setStringField(builder, TagOrigClOrdID, r.OrigClOrdID)
	setStringField(builder, TagCancelRestrictions, string(r.CancelRestrictions))
	setOrderFields(builder, r.OrderFields, false)
	return builder.Build()
}

type OrderMassCancelRequest struct {
	ClOrdID               string
	Symbol                string
	MassCancelRequestType MassCancelRequestType
}

func NewOrderMassCancelRequest(clOrdID string, symbol string) *OrderMassCancelRequest {
	return &OrderMassCancelRequest{
		ClOrdID:               clOrdID,
		Symbol:                symbol,
		MassCancelRequestType: MassCancelRequestTypeCancelSymbolOrders,
	}
}

func (r *OrderMassCancelRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	return newApplicationBuilder(MsgTypeOrderMassCancelRequest, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagClOrdID, r.ClOrdID).
		SetField(TagSymbol, r.Symbol).
		SetField(TagMassCancelRequestType, string(r.MassCancelRequestType)).
		Build()
}

type OrderMassCancelReport struct {
	Symbol                 string
	ClOrdID                string
	MassCancelRequestType  MassCancelRequestType
	MassCancelResponse     MassCancelResponse
	MassCancelRejectReason MassCancelRejectReason
	TotalAffectedOrders    string
	ErrorCode              string
	Text                   string
}

func NewOrderMassCancelReport() *OrderMassCancelReport {
	return new(OrderMassCancelReport)
}

func (r *OrderMassCancelReport) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeOrderMassCancelReport); err != nil {
		return err
	}

	symbol, err := requiredString(m, TagSymbol)
	if err != nil {
		return err
	}
	clOrdID, err := requiredString(m, TagClOrdID)
	if err != nil {
		return err
	}
	requestType, err := requiredString(m, TagMassCancelRequestType)
	if err != nil {
		return err
	}
	response, err := requiredString(m, TagMassCancelResponse)
	if err != nil {
		return err
	}

	r.Symbol = symbol
	r.ClOrdID = clOrdID
	r.MassCancelRequestType = MassCancelRequestType(requestType)
	r.MassCancelResponse = MassCancelResponse(response)
	r.MassCancelRejectReason = MassCancelRejectReason(optionalString(m, TagMassCancelRejectReason))
	r.TotalAffectedOrders = optionalString(m, TagTotalAffectedOrders)
	r.ErrorCode = optionalString(m, TagErrorCode)
	r.Text = optionalString(m, TagText)
	return nil
}

func (r *OrderMassCancelReport) Error() string {
	if r == nil {
		return "order mass cancel reject"
	}
	return rejectErrorMessage(
		"order mass cancel reject",
		errorPart("clOrdID", r.ClOrdID),
		errorPart("symbol", r.Symbol),
		errorPart("response", string(r.MassCancelResponse)),
		errorPart("reason", string(r.MassCancelRejectReason)),
		errorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type ListTriggeringInstruction struct {
	ListTriggerType         ListTriggerType
	ListTriggerTriggerIndex string
	ListTriggerAction       ListTriggerAction
}

type NewOrderListOrder struct {
	OrderFields
	ListTriggeringInstructions []ListTriggeringInstruction
}

type NewOrderList struct {
	ClListID        string
	ContingencyType ContingencyType
	OPO             *bool
	Orders          []NewOrderListOrder
}

func NewNewOrderList(clListID string, orders []NewOrderListOrder) *NewOrderList {
	return &NewOrderList{ClListID: clListID, Orders: orders}
}

func (r *NewOrderList) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeNewOrderList, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagClListID, r.ClListID)
	setStringField(builder, TagContingencyType, string(r.ContingencyType))
	setBoolField(builder, TagOPO, r.OPO)
	if len(r.Orders) > 0 {
		builder.SetField(TagNoOrders, FormatUint(uint64(len(r.Orders))))
	}
	for _, order := range r.Orders {
		addOrderFields(builder, order.OrderFields, false)
		if len(order.ListTriggeringInstructions) > 0 {
			builder.AddField(TagNoListTriggeringInstructions, FormatUint(uint64(len(order.ListTriggeringInstructions))))
		}
		for _, instruction := range order.ListTriggeringInstructions {
			addStringField(builder, TagListTriggerType, string(instruction.ListTriggerType))
			addStringField(builder, TagListTriggerTriggerIndex, instruction.ListTriggerTriggerIndex)
			addStringField(builder, TagListTriggerAction, string(instruction.ListTriggerAction))
		}
	}
	return builder.Build()
}

type ListStatus struct {
	Symbol           string
	ListID           string
	ClListID         string
	OrigClListID     string
	ContingencyType  ContingencyType
	ListStatusType   ListStatusType
	ListOrderStatus  ListOrderStatus
	ListRejectReason ListRejectReason
	OrdRejReason     OrdRejReason
	TransactTime     string
	ErrorCode        string
	Text             string
	NoOrders         string
}

func NewListStatus() *ListStatus {
	return new(ListStatus)
}

func (r *ListStatus) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeListStatus); err != nil {
		return err
	}

	listStatusType, err := requiredString(m, TagListStatusType)
	if err != nil {
		return err
	}
	listOrderStatus, err := requiredString(m, TagListOrderStatus)
	if err != nil {
		return err
	}

	r.Symbol = optionalString(m, TagSymbol)
	r.ListID = optionalString(m, TagListID)
	r.ClListID = optionalString(m, TagClListID)
	r.OrigClListID = optionalString(m, TagOrigClListID)
	r.ContingencyType = ContingencyType(optionalString(m, TagContingencyType))
	r.ListStatusType = ListStatusType(listStatusType)
	r.ListOrderStatus = ListOrderStatus(listOrderStatus)
	r.ListRejectReason = ListRejectReason(optionalString(m, TagListRejectReason))
	r.OrdRejReason = OrdRejReason(optionalString(m, TagOrdRejReason))
	r.TransactTime = optionalString(m, TagTransactTime)
	r.ErrorCode = optionalString(m, TagErrorCode)
	r.Text = optionalString(m, TagText)
	r.NoOrders = optionalString(m, TagNoOrders)
	return nil
}

func (r *ListStatus) Error() string {
	if r == nil {
		return "order list reject"
	}
	return rejectErrorMessage(
		"order list reject",
		errorPart("clListID", r.ClListID),
		errorPart("origClListID", r.OrigClListID),
		errorPart("listID", r.ListID),
		errorPart("symbol", r.Symbol),
		errorPart("status", string(r.ListOrderStatus)),
		errorPart("reason", string(r.ListRejectReason)),
		errorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type OrderAmendKeepPriorityRequest struct {
	ClOrdID     string
	OrigClOrdID string
	OrderID     string
	Symbol      string
	OrderQty    string
}

func NewOrderAmendKeepPriorityRequest(clOrdID string, symbol string) *OrderAmendKeepPriorityRequest {
	return &OrderAmendKeepPriorityRequest{ClOrdID: clOrdID, Symbol: symbol}
}

func (r *OrderAmendKeepPriorityRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeOrderAmendKeepPriorityRequest, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagClOrdID, r.ClOrdID).
		SetField(TagSymbol, r.Symbol)
	setStringField(builder, TagOrigClOrdID, r.OrigClOrdID)
	setStringField(builder, TagOrderID, r.OrderID)
	setStringField(builder, TagOrderQty, r.OrderQty)
	return builder.Build()
}

type OrderAmendReject struct {
	ClOrdID     string
	OrigClOrdID string
	OrderID     string
	Symbol      string
	OrderQty    string
	ErrorCode   string
	Text        string
}

func NewOrderAmendReject() *OrderAmendReject {
	return new(OrderAmendReject)
}

func (r *OrderAmendReject) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeOrderAmendReject); err != nil {
		return err
	}

	clOrdID, err := requiredString(m, TagClOrdID)
	if err != nil {
		return err
	}
	symbol, err := requiredString(m, TagSymbol)
	if err != nil {
		return err
	}
	orderQty, err := requiredString(m, TagOrderQty)
	if err != nil {
		return err
	}
	errorCode, err := requiredString(m, TagErrorCode)
	if err != nil {
		return err
	}
	text, err := requiredString(m, TagText)
	if err != nil {
		return err
	}

	r.ClOrdID = clOrdID
	r.OrigClOrdID = optionalString(m, TagOrigClOrdID)
	r.OrderID = optionalString(m, TagOrderID)
	r.Symbol = symbol
	r.OrderQty = orderQty
	r.ErrorCode = errorCode
	r.Text = text
	return nil
}

func (r *OrderAmendReject) Error() string {
	if r == nil {
		return "order amend reject"
	}
	return rejectErrorMessage(
		"order amend reject",
		errorPart("clOrdID", r.ClOrdID),
		errorPart("origClOrdID", r.OrigClOrdID),
		errorPart("orderID", r.OrderID),
		errorPart("symbol", r.Symbol),
		errorPart("orderQty", r.OrderQty),
		errorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type LimitQuery struct {
	ReqID string
}

func NewLimitQuery(reqID string) *LimitQuery {
	return &LimitQuery{ReqID: reqID}
}

func (r *LimitQuery) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	return newApplicationBuilder(MsgTypeLimitQuery, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagReqID, r.ReqID).
		Build()
}

type LimitResponse struct {
	ReqID             string
	NoLimitIndicators string
}

func NewLimitResponse() *LimitResponse {
	return new(LimitResponse)
}

func (r *LimitResponse) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeLimitResponse); err != nil {
		return err
	}

	reqID, err := requiredString(m, TagReqID)
	if err != nil {
		return err
	}
	noLimitIndicators, err := requiredString(m, TagNoLimitIndicators)
	if err != nil {
		return err
	}

	r.ReqID = reqID
	r.NoLimitIndicators = noLimitIndicators
	return nil
}

type InstrumentListRequest struct {
	InstrumentReqID           string
	InstrumentListRequestType InstrumentListRequestType
	Symbol                    string
}

func NewInstrumentListRequest(instrumentReqID string, requestType InstrumentListRequestType) *InstrumentListRequest {
	return &InstrumentListRequest{
		InstrumentReqID:           instrumentReqID,
		InstrumentListRequestType: requestType,
	}
}

func (r *InstrumentListRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeInstrumentListRequest, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagInstrumentReqID, r.InstrumentReqID).
		SetField(TagInstrumentListRequestType, string(r.InstrumentListRequestType))
	setStringField(builder, TagSymbol, r.Symbol)
	return builder.Build()
}

type InstrumentList struct {
	InstrumentReqID string
	NoRelatedSym    string
}

func NewInstrumentList() *InstrumentList {
	return new(InstrumentList)
}

func (r *InstrumentList) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeInstrumentList); err != nil {
		return err
	}

	instrumentReqID, err := requiredString(m, TagInstrumentReqID)
	if err != nil {
		return err
	}
	noRelatedSym, err := requiredString(m, TagNoRelatedSym)
	if err != nil {
		return err
	}

	r.InstrumentReqID = instrumentReqID
	r.NoRelatedSym = noRelatedSym
	return nil
}

type MarketDataRequest struct {
	MDReqID                 string
	SubscriptionRequestType SubscriptionRequestType
	MarketDepth             string
	AggregatedBook          *bool
	Symbols                 []string
	MDEntryTypes            []MDEntryType
}

func NewMarketDataRequest(mdReqID string, subscriptionRequestType SubscriptionRequestType) *MarketDataRequest {
	return &MarketDataRequest{
		MDReqID:                 mdReqID,
		SubscriptionRequestType: subscriptionRequestType,
	}
}

func (r *MarketDataRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newApplicationBuilder(MsgTypeMarketDataRequest, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagMDReqID, r.MDReqID).
		SetField(TagSubscriptionRequestType, string(r.SubscriptionRequestType))
	setStringField(builder, TagMarketDepth, r.MarketDepth)
	setBoolField(builder, TagAggregatedBook, r.AggregatedBook)
	if len(r.Symbols) > 0 {
		builder.SetField(TagNoRelatedSym, FormatUint(uint64(len(r.Symbols))))
		for _, symbol := range r.Symbols {
			builder.AddField(TagSymbol, symbol)
		}
	}
	if len(r.MDEntryTypes) > 0 {
		builder.SetField(TagNoMDEntryTypes, FormatUint(uint64(len(r.MDEntryTypes))))
		for _, entryType := range r.MDEntryTypes {
			builder.AddField(TagMDEntryType, string(entryType))
		}
	}
	return builder.Build()
}

type MarketDataRequestReject struct {
	MDReqID        string
	MDReqRejReason MDReqRejReason
	ErrorCode      string
	Text           string
}

func NewMarketDataRequestReject() *MarketDataRequestReject {
	return new(MarketDataRequestReject)
}

func (r *MarketDataRequestReject) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeMarketDataRequestReject); err != nil {
		return err
	}

	mdReqID, err := requiredString(m, TagMDReqID)
	if err != nil {
		return err
	}

	r.MDReqID = mdReqID
	r.MDReqRejReason = MDReqRejReason(optionalString(m, TagMDReqRejReason))
	r.ErrorCode = optionalString(m, TagErrorCode)
	r.Text = optionalString(m, TagText)
	return nil
}

func (r *MarketDataRequestReject) Error() string {
	if r == nil {
		return "market data request reject"
	}
	return rejectErrorMessage(
		"market data request reject",
		errorPart("mdReqID", r.MDReqID),
		errorPart("reason", string(r.MDReqRejReason)),
		errorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type MarketDataSnapshot struct {
	MDReqID          string
	Symbol           string
	LastBookUpdateID string
	NoMDEntries      string
}

func NewMarketDataSnapshot() *MarketDataSnapshot {
	return new(MarketDataSnapshot)
}

func (r *MarketDataSnapshot) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeMarketDataSnapshot); err != nil {
		return err
	}

	mdReqID, err := requiredString(m, TagMDReqID)
	if err != nil {
		return err
	}
	symbol, err := requiredString(m, TagSymbol)
	if err != nil {
		return err
	}
	noMDEntries, err := requiredString(m, TagNoMDEntries)
	if err != nil {
		return err
	}

	r.MDReqID = mdReqID
	r.Symbol = symbol
	r.LastBookUpdateID = optionalString(m, TagLastBookUpdateID)
	r.NoMDEntries = noMDEntries
	return nil
}

type MarketDataIncrementalRefresh struct {
	MDReqID     string
	NoMDEntries string
}

func NewMarketDataIncrementalRefresh() *MarketDataIncrementalRefresh {
	return new(MarketDataIncrementalRefresh)
}

func (r *MarketDataIncrementalRefresh) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeMarketDataIncrementalRefresh); err != nil {
		return err
	}

	mdReqID, err := requiredString(m, TagMDReqID)
	if err != nil {
		return err
	}
	noMDEntries, err := requiredString(m, TagNoMDEntries)
	if err != nil {
		return err
	}

	r.MDReqID = mdReqID
	r.NoMDEntries = noMDEntries
	return nil
}

func newApplicationBuilder(msgType MsgType, senderCompID string, targetCompID string, seqNum uint32) *MsgBuilder {
	return NewMsgBuilder().
		MsgType(msgType).
		SenderCompID(senderCompID).
		TargetCompID(targetCompID).
		MsgSeqNum(seqNum)
}

func setOrderFields(builder *MsgBuilder, order OrderFields, includeSOR bool) {
	writeOrderFields(builder.SetField, order, includeSOR)
}

func addOrderFields(builder *MsgBuilder, order OrderFields, includeSOR bool) {
	writeOrderFields(builder.AddField, order, includeSOR)
}

func writeOrderFields(write func(Tag, string) *MsgBuilder, order OrderFields, includeSOR bool) {
	writeString(write, TagClOrdID, order.ClOrdID)
	writeString(write, TagOrderQty, order.OrderQty)
	writeString(write, TagOrdType, string(order.OrdType))
	writeString(write, TagExecInst, string(order.ExecInst))
	writeString(write, TagPrice, order.Price)
	writeString(write, TagSide, string(order.Side))
	writeString(write, TagSymbol, order.Symbol)
	writeString(write, TagTimeInForce, string(order.TimeInForce))
	writeString(write, TagMaxFloor, order.MaxFloor)
	writeString(write, TagCashOrderQty, order.CashOrderQty)
	writeString(write, TagTargetStrategy, order.TargetStrategy)
	writeString(write, TagStrategyID, order.StrategyID)
	writeString(write, TagSelfTradePreventionMode, string(order.SelfTradePreventionMode))
	writeString(write, TagPegOffsetValue, order.PegOffsetValue)
	writeString(write, TagPegPriceType, string(order.PegPriceType))
	writeString(write, TagPegMoveType, string(order.PegMoveType))
	writeString(write, TagPegOffsetType, string(order.PegOffsetType))
	writeString(write, TagTriggerType, string(order.TriggerType))
	writeString(write, TagTriggerAction, string(order.TriggerAction))
	writeString(write, TagTriggerPrice, order.TriggerPrice)
	writeString(write, TagTriggerPriceType, string(order.TriggerPriceType))
	writeString(write, TagTriggerPriceDirection, string(order.TriggerPriceDirection))
	writeString(write, TagTriggerTrailingDeltaBips, order.TriggerTrailingDeltaBips)
	if includeSOR && order.SOR != nil {
		write(TagSOR, FormatBool(*order.SOR))
	}
}

func setStringField(builder *MsgBuilder, tag Tag, value string) {
	writeString(builder.SetField, tag, value)
}

func addStringField(builder *MsgBuilder, tag Tag, value string) {
	writeString(builder.AddField, tag, value)
}

func writeString(write func(Tag, string) *MsgBuilder, tag Tag, value string) {
	if value != "" {
		write(tag, value)
	}
}

func setBoolField(builder *MsgBuilder, tag Tag, value *bool) {
	if value != nil {
		builder.SetField(tag, FormatBool(*value))
	}
}

var _ Request = (*NewOrderSingle)(nil)
var _ Response = (*ExecutionReport)(nil)
var _ error = (*ExecutionReport)(nil)
var _ Request = (*OrderCancelRequest)(nil)
var _ Response = (*OrderCancelReject)(nil)
var _ error = (*OrderCancelReject)(nil)
var _ Request = (*OrderCancelRequestAndNewOrderSingle)(nil)
var _ Request = (*OrderMassCancelRequest)(nil)
var _ Response = (*OrderMassCancelReport)(nil)
var _ error = (*OrderMassCancelReport)(nil)
var _ Request = (*NewOrderList)(nil)
var _ Response = (*ListStatus)(nil)
var _ error = (*ListStatus)(nil)
var _ Request = (*OrderAmendKeepPriorityRequest)(nil)
var _ Response = (*OrderAmendReject)(nil)
var _ error = (*OrderAmendReject)(nil)
var _ Request = (*LimitQuery)(nil)
var _ Response = (*LimitResponse)(nil)
var _ Request = (*InstrumentListRequest)(nil)
var _ Response = (*InstrumentList)(nil)
var _ Request = (*MarketDataRequest)(nil)
var _ Response = (*MarketDataRequestReject)(nil)
var _ error = (*MarketDataRequestReject)(nil)
var _ Response = (*MarketDataSnapshot)(nil)
var _ Response = (*MarketDataIncrementalRefresh)(nil)
