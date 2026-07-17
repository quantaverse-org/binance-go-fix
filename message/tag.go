package message

// Tag identifies a FIX field by its numeric tag.
type Tag uint32

const (
	// TagBeginString is BeginString (8): always FIX.4.4 and the first field.
	TagBeginString = 8
	// TagBodyLength is BodyLength (9): message length in bytes and the second field.
	TagBodyLength = 9
	// TagCheckSum is CheckSum (10): three-character checksum at the end of the message.
	TagCheckSum = 10
	// TagMsgSeqNum is MsgSeqNum (34): monotonically increasing message sequence number.
	TagMsgSeqNum = 34
	// TagMsgType is MsgType (35): message type and the third field in the message.
	TagMsgType = 35
	// TagSenderCompID is SenderCompID (49): unique session identifier for the sender.
	TagSenderCompID = 49
	// TagTargetCompID is TargetCompID (56): identifier for this TCP connection.
	TagTargetCompID = 56
	// TagSendingTime is SendingTime (52): UTC timestamp when the message was sent.
	TagSendingTime = 52
	// TagRecvWindow is RecvWindow (25000): request validity window after SendingTime in milliseconds.
	TagRecvWindow = 25000
)
