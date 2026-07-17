package message

import (
	"strings"
	"testing"
	"time"
)

func TestLogonRequestToMessage(t *testing.T) {
	privateKey, err := ParseLogonPrivateKeyPEM([]byte(examplePrivateKeyPEM))
	if err != nil {
		t.Fatalf("ParseLogonPrivateKeyPEM() error = %v", err)
	}

	message, err := NewLogonRequest("api-key", privateKey, 30, MessageHandlingSequential).
		WithResponseMode(ResponseModeOnlyAcks).
		WithDropCopyFlag(true).
		ToMessage("EXAMPLE", "SPOT", 1, exampleSendingTime())
	if err != nil {
		t.Fatalf("ToMessage() error = %v", err)
	}

	assertField(t, message, TagMsgType, string(MsgTypeLogon))
	assertField(t, message, TagSenderCompID, "EXAMPLE")
	assertField(t, message, TagTargetCompID, "SPOT")
	assertField(t, message, TagMsgSeqNum, "1")
	assertField(t, message, TagSendingTime, "20240627-11:17:25.223")
	assertField(t, message, TagEncryptMethod, "0")
	assertField(t, message, TagHeartBtInt, "30")
	assertField(t, message, TagRawDataLength, "88")
	assertField(t, message, TagRawData, exampleLogonSignature)
	assertField(t, message, TagResetSeqNumFlag, "Y")
	assertField(t, message, TagUsername, "api-key")
	assertField(t, message, TagMessageHandling, "2")
	assertField(t, message, TagResponseMode, "2")
	assertField(t, message, TagDropCopyFlag, "Y")
}

func TestSignLogonRawData(t *testing.T) {
	privateKey, err := ParseLogonPrivateKeyPEM([]byte(examplePrivateKeyPEM))
	if err != nil {
		t.Fatalf("ParseLogonPrivateKeyPEM() error = %v", err)
	}

	payload := LogonSignaturePayload("EXAMPLE", "SPOT", 1, "20240627-11:17:25.223")
	wantPayload := strings.Join([]string{"A", "EXAMPLE", "SPOT", "1", "20240627-11:17:25.223"}, string(SOH))
	if payload != wantPayload {
		t.Fatalf("LogonSignaturePayload() = %q, want %q", payload, wantPayload)
	}

	rawData, err := SignLogonRawData(privateKey, "EXAMPLE", "SPOT", 1, "20240627-11:17:25.223")
	if err != nil {
		t.Fatalf("SignLogonRawData() error = %v", err)
	}
	if rawData != exampleLogonSignature {
		t.Fatalf("SignLogonRawData() = %q, want %q", rawData, exampleLogonSignature)
	}
}

func TestLogonRequestToMessageInvalidPrivateKey(t *testing.T) {
	_, err := NewLogonRequest("api-key", nil, 30, MessageHandlingSequential).
		ToMessage("EXAMPLE", "SPOT", 1, exampleSendingTime())
	if err == nil {
		t.Fatal("ToMessage() error = nil, want error")
	}
}

func TestRejectFromMessage(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=3|45=7|371=35|372=D|373=1|25016=-1102|58=Missing field|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	reject := new(Reject)
	if err := reject.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}
	if reject.RefSeqNum == nil || *reject.RefSeqNum != 7 {
		t.Fatalf("RefSeqNum = %v, want %d", reject.RefSeqNum, 7)
	}
	if reject.RefTagID == nil || *reject.RefTagID != TagMsgType {
		t.Fatalf("RefTagID = %v, want %d", reject.RefTagID, TagMsgType)
	}
	if reject.RefMsgType != string(MsgTypeNewOrderSingle) {
		t.Fatalf("RefMsgType = %q, want %q", reject.RefMsgType, string(MsgTypeNewOrderSingle))
	}
	if reject.SessionRejectReason == nil || *reject.SessionRejectReason != SessionRejectReasonRequiredTagMissing {
		t.Fatalf("SessionRejectReason = %v, want %s", reject.SessionRejectReason, SessionRejectReasonRequiredTagMissing)
	}
	if reject.ErrorCode == nil || *reject.ErrorCode != -1102 {
		t.Fatalf("ErrorCode = %v, want %d", reject.ErrorCode, -1102)
	}
	if reject.Text != "Missing field" {
		t.Fatalf("Text = %q, want %q", reject.Text, "Missing field")
	}
}

func TestNewsFromMessage(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=B|148=Reconnect soon|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	news := new(News)
	if err := news.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}
	if news.Headline != "Reconnect soon" {
		t.Fatalf("Headline = %q, want %q", news.Headline, "Reconnect soon")
	}
}

func assertField(t *testing.T, message *Message, tag Tag, want string) {
	t.Helper()

	got, err := message.GetRequiredField(tag)
	if err != nil {
		t.Fatalf("GetRequiredField(%d) error = %v", tag, err)
	}
	if got != want {
		t.Fatalf("field %d = %q, want %q", tag, got, want)
	}
}

func withSOH(message string) string {
	return strings.ReplaceAll(message, "|", string(SOH))
}

func exampleSendingTime() time.Time {
	return time.Date(2024, time.June, 27, 11, 17, 25, 223000000, time.UTC)
}

const examplePrivateKeyPEM = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIIJEYWtGBrhACmb9Dvy+qa8WEf0lQOl1s4CLIAB9m89u
-----END PRIVATE KEY-----`

const exampleLogonSignature = "4MHXelVVcpkdwuLbl6n73HQUXUf1dse2PCgT1DYqW9w8AVZ1RACFGM+5UdlGPrQHrgtS3CvsRURC1oj73j8gCA=="
