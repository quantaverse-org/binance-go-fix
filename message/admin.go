package message

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
	"time"
)

const (
	TagRefSeqNum           Tag = 45
	TagText                Tag = 58
	TagRawDataLength       Tag = 95
	TagRawData             Tag = 96
	TagEncryptMethod       Tag = 98
	TagHeartBtInt          Tag = 108
	TagTestReqID           Tag = 112
	TagResetSeqNumFlag     Tag = 141
	TagHeadline            Tag = 148
	TagRefTagID            Tag = 371
	TagRefMsgType          Tag = 372
	TagSessionRejectReason Tag = 373
	TagUsername            Tag = 553
	TagDropCopyFlag        Tag = 9406
	TagErrorCode           Tag = 25016
	TagMessageHandling     Tag = 25035
	TagResponseMode        Tag = 25036
	TagUUID                Tag = 25037
)

type EncryptMethod string

const (
	EncryptMethodNone EncryptMethod = "0"
)

type MessageHandling string

const (
	MessageHandlingUnordered  MessageHandling = "1"
	MessageHandlingSequential MessageHandling = "2"
)

type ResponseMode string

const (
	ResponseModeEverything ResponseMode = "1"
	ResponseModeOnlyAcks   ResponseMode = "2"
)

type SessionRejectReason string

const (
	SessionRejectReasonInvalidTagNumber               SessionRejectReason = "0"
	SessionRejectReasonRequiredTagMissing             SessionRejectReason = "1"
	SessionRejectReasonTagNotDefinedForMessageType    SessionRejectReason = "2"
	SessionRejectReasonUndefinedTag                   SessionRejectReason = "3"
	SessionRejectReasonValueIsIncorrect               SessionRejectReason = "5"
	SessionRejectReasonIncorrectDataFormatForValue    SessionRejectReason = "6"
	SessionRejectReasonSignatureProblem               SessionRejectReason = "8"
	SessionRejectReasonSendingTimeAccuracyProblem     SessionRejectReason = "10"
	SessionRejectReasonXMLValidationError             SessionRejectReason = "12"
	SessionRejectReasonTagAppearsMoreThanOnce         SessionRejectReason = "13"
	SessionRejectReasonTagSpecifiedOutOfRequiredOrder SessionRejectReason = "14"
	SessionRejectReasonRepeatingGroupFieldsOutOfOrder SessionRejectReason = "15"
	SessionRejectReasonIncorrectNumInGroupCount       SessionRejectReason = "16"
	SessionRejectReasonOther                          SessionRejectReason = "99"
)

type Heartbeat struct {
	TestReqID string
}

func NewHeartbeat(testReqID string) *Heartbeat {
	return &Heartbeat{TestReqID: testReqID}
}

func (h *Heartbeat) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newAdminBuilder(MsgTypeHeartbeat, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime)
	if h.TestReqID != "" {
		builder.SetField(TagTestReqID, h.TestReqID)
	}
	return builder.Build()
}

func (h *Heartbeat) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeHeartbeat); err != nil {
		return err
	}
	h.TestReqID = optionalString(m, TagTestReqID)
	return nil
}

type TestRequest struct {
	TestReqID string
}

func NewTestRequest(testReqID string) *TestRequest {
	return &TestRequest{TestReqID: testReqID}
}

func (r *TestRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	return newAdminBuilder(MsgTypeTestRequest, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagTestReqID, r.TestReqID).
		Build()
}

func (r *TestRequest) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeTestRequest); err != nil {
		return err
	}
	testReqID, err := requiredString(m, TagTestReqID)
	if err != nil {
		return err
	}
	r.TestReqID = testReqID
	return nil
}

type Reject struct {
	RefSeqNum           *uint32
	RefTagID            *Tag
	RefMsgType          string
	SessionRejectReason *SessionRejectReason
	ErrorCode           *int64
	Text                string
}

func NewReject() *Reject {
	return new(Reject)
}

func (r *Reject) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeReject); err != nil {
		return err
	}

	refSeqNum, err := optionalSeqNum(m, TagRefSeqNum)
	if err != nil {
		return err
	}
	refTagID, err := optionalTag(m, TagRefTagID)
	if err != nil {
		return err
	}
	sessionRejectReason, err := optionalSessionRejectReason(m, TagSessionRejectReason)
	if err != nil {
		return err
	}
	errorCode, err := optionalInt(m, TagErrorCode)
	if err != nil {
		return err
	}

	r.RefSeqNum = refSeqNum
	r.RefTagID = refTagID
	r.RefMsgType = optionalString(m, TagRefMsgType)
	r.SessionRejectReason = sessionRejectReason
	r.ErrorCode = errorCode
	r.Text = optionalString(m, TagText)
	return nil
}

func (r *Reject) Error() string {
	if r == nil {
		return "session reject"
	}
	return rejectErrorMessage(
		"session reject",
		uint32PtrErrorPart("refSeqNum", r.RefSeqNum),
		tagPtrErrorPart("refTagID", r.RefTagID),
		errorPart("refMsgType", r.RefMsgType),
		sessionRejectReasonPtrErrorPart("reason", r.SessionRejectReason),
		int64PtrErrorPart("errorCode", r.ErrorCode),
		errorPart("text", r.Text),
	)
}

type LogonRequest struct {
	EncryptMethod   EncryptMethod
	HeartBtInt      int64
	PrivateKey      ed25519.PrivateKey
	ResetSeqNumFlag bool
	Username        string
	MessageHandling MessageHandling
	ResponseMode    *ResponseMode
	DropCopyFlag    *bool
}

func NewLogonRequest(username string, privateKey ed25519.PrivateKey, heartBtInt int64, messageHandling MessageHandling) *LogonRequest {
	return &LogonRequest{
		EncryptMethod:   EncryptMethodNone,
		HeartBtInt:      heartBtInt,
		PrivateKey:      privateKey,
		ResetSeqNumFlag: true,
		Username:        username,
		MessageHandling: messageHandling,
	}
}

func (r *LogonRequest) WithResponseMode(responseMode ResponseMode) *LogonRequest {
	r.ResponseMode = &responseMode
	return r
}

func (r *LogonRequest) WithDropCopyFlag(dropCopyFlag bool) *LogonRequest {
	r.DropCopyFlag = &dropCopyFlag
	return r
}

func (r *LogonRequest) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	sendingTimeValue := FormatTimestampMs(sendingTime)
	rawData, err := SignLogonRawData(r.PrivateKey, senderCompID, targetCompID, seqNum, sendingTimeValue)
	if err != nil {
		return nil, err
	}

	builder := newAdminBuilder(MsgTypeLogon, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime).
		SetField(TagEncryptMethod, string(r.EncryptMethod)).
		SetField(TagHeartBtInt, FormatInt(r.HeartBtInt)).
		SetField(TagRawDataLength, FormatUint(uint64(len(rawData)))).
		SetField(TagRawData, rawData).
		SetField(TagResetSeqNumFlag, FormatBool(r.ResetSeqNumFlag)).
		SetField(TagUsername, r.Username).
		SetField(TagMessageHandling, string(r.MessageHandling))

	if r.ResponseMode != nil {
		builder.SetField(TagResponseMode, string(*r.ResponseMode))
	}
	if r.DropCopyFlag != nil {
		builder.SetField(TagDropCopyFlag, FormatBool(*r.DropCopyFlag))
	}
	return builder.Build()
}

func SignLogonRawData(privateKey ed25519.PrivateKey, senderCompID string, targetCompID string, seqNum uint32, sendingTime string) (string, error) {
	if len(privateKey) != ed25519.PrivateKeySize {
		return "", fmt.Errorf("invalid Ed25519 private key size: %d", len(privateKey))
	}

	payload := LogonSignaturePayload(senderCompID, targetCompID, seqNum, sendingTime)
	signature := ed25519.Sign(privateKey, []byte(payload))
	return base64.StdEncoding.EncodeToString(signature), nil
}

func LogonSignaturePayload(senderCompID string, targetCompID string, seqNum uint32, sendingTime string) string {
	return strings.Join([]string{
		string(MsgTypeLogon),
		senderCompID,
		targetCompID,
		FormatUint(uint64(seqNum)),
		sendingTime,
	}, string(SOH))
}

func ParseLogonPrivateKeyPEM(data []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	ed25519PrivateKey, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not Ed25519")
	}
	return ed25519PrivateKey, nil
}

type LogonResponse struct {
	EncryptMethod EncryptMethod
	HeartBtInt    int64
	UUID          string
}

func NewLogonResponse(encryptMethod EncryptMethod, heartBtInt int64, uuid string) *LogonResponse {
	return &LogonResponse{
		EncryptMethod: encryptMethod,
		HeartBtInt:    heartBtInt,
		UUID:          uuid,
	}
}

func (r *LogonResponse) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeLogon); err != nil {
		return err
	}

	encryptMethod, err := requiredEncryptMethod(m, TagEncryptMethod)
	if err != nil {
		return err
	}
	heartBtInt, err := requiredInt(m, TagHeartBtInt)
	if err != nil {
		return err
	}
	uuid, err := requiredString(m, TagUUID)
	if err != nil {
		return err
	}

	r.EncryptMethod = encryptMethod
	r.HeartBtInt = heartBtInt
	r.UUID = uuid
	return nil
}

type Logout struct {
	Text string
}

func NewLogout(text string) *Logout {
	return &Logout{Text: text}
}

func (l *Logout) ToMessage(senderCompID string, targetCompID string, seqNum uint32, sendingTime time.Time) (*Message, error) {
	builder := newAdminBuilder(MsgTypeLogout, senderCompID, targetCompID, seqNum).
		SendingTime(sendingTime)
	if l.Text != "" {
		builder.SetField(TagText, l.Text)
	}
	return builder.Build()
}

func (l *Logout) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeLogout); err != nil {
		return err
	}
	l.Text = optionalString(m, TagText)
	return nil
}

type News struct {
	Headline string
}

func NewNews(headline string) *News {
	return &News{Headline: headline}
}

func (n *News) FromMessage(m *Message) error {
	if err := requireMsgType(m, MsgTypeNews); err != nil {
		return err
	}
	headline, err := requiredString(m, TagHeadline)
	if err != nil {
		return err
	}
	n.Headline = headline
	return nil
}

func newAdminBuilder(msgType MsgType, senderCompID string, targetCompID string, seqNum uint32) *MsgBuilder {
	return NewMsgBuilder().
		MsgType(msgType).
		SenderCompID(senderCompID).
		TargetCompID(targetCompID).
		MsgSeqNum(seqNum)
}

func requireMsgType(m *Message, want MsgType) error {
	got, err := m.GetRequiredField(TagMsgType)
	if err != nil {
		return err
	}
	if got != string(want) {
		return fmt.Errorf("invalid message type: %s", got)
	}
	return nil
}

func requiredString(m *Message, tag Tag) (string, error) {
	value, err := m.GetRequiredField(tag)
	if err != nil {
		return "", fmt.Errorf("missing field: %d", tag)
	}
	return value, nil
}

func optionalString(m *Message, tag Tag) string {
	value, _ := m.GetField(tag)
	return value
}

func rejectErrorMessage(title string, parts ...string) string {
	filtered := make([]string, 0, len(parts)+1)
	filtered = append(filtered, title)
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	return strings.Join(filtered, ", ")
}

func errorPart(name string, value string) string {
	if value == "" {
		return ""
	}
	return name + "=" + value
}

func uint32PtrErrorPart(name string, value *uint32) string {
	if value == nil {
		return ""
	}
	return name + "=" + FormatUint(uint64(*value))
}

func int64PtrErrorPart(name string, value *int64) string {
	if value == nil {
		return ""
	}
	return name + "=" + FormatInt(*value)
}

func tagPtrErrorPart(name string, value *Tag) string {
	if value == nil {
		return ""
	}
	return name + "=" + FormatTag(*value)
}

func sessionRejectReasonPtrErrorPart(name string, value *SessionRejectReason) string {
	if value == nil {
		return ""
	}
	return name + "=" + string(*value)
}

func requiredInt(m *Message, tag Tag) (int64, error) {
	value, err := requiredString(m, tag)
	if err != nil {
		return 0, err
	}
	parsed, err := ParseInt(value)
	if err != nil {
		return 0, fmt.Errorf("parse field %d: %w", tag, err)
	}
	return parsed, nil
}

func requiredEncryptMethod(m *Message, tag Tag) (EncryptMethod, error) {
	value, err := requiredString(m, tag)
	if err != nil {
		return "", err
	}
	return EncryptMethod(value), nil
}

func optionalInt(m *Message, tag Tag) (*int64, error) {
	value, ok := m.GetField(tag)
	if !ok {
		return nil, nil
	}
	parsed, err := ParseInt(value)
	if err != nil {
		return nil, fmt.Errorf("parse field %d: %w", tag, err)
	}
	return &parsed, nil
}

func optionalSeqNum(m *Message, tag Tag) (*uint32, error) {
	value, ok := m.GetField(tag)
	if !ok {
		return nil, nil
	}
	parsed, err := ParseUint(value)
	if err != nil {
		return nil, fmt.Errorf("parse field %d: %w", tag, err)
	}
	if parsed > uint64(^uint32(0)) {
		return nil, fmt.Errorf("parse field %d: value %d overflows uint32", tag, parsed)
	}
	seqNum := uint32(parsed)
	return &seqNum, nil
}

func optionalTag(m *Message, tag Tag) (*Tag, error) {
	value, ok := m.GetField(tag)
	if !ok {
		return nil, nil
	}
	parsed, err := ParseTag(value)
	if err != nil {
		return nil, fmt.Errorf("parse field %d: %w", tag, err)
	}
	return &parsed, nil
}

func optionalSessionRejectReason(m *Message, tag Tag) (*SessionRejectReason, error) {
	value, ok := m.GetField(tag)
	if !ok {
		return nil, nil
	}
	reason := SessionRejectReason(value)
	return &reason, nil
}

var _ Request = (*Heartbeat)(nil)
var _ Response = (*Heartbeat)(nil)
var _ Request = (*TestRequest)(nil)
var _ Response = (*TestRequest)(nil)
var _ Response = (*Reject)(nil)
var _ error = (*Reject)(nil)
var _ Request = (*LogonRequest)(nil)
var _ Response = (*LogonResponse)(nil)
var _ Request = (*Logout)(nil)
var _ Response = (*Logout)(nil)
var _ Response = (*News)(nil)
