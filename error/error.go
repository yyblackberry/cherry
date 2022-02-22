package cherryError

import (
	"errors"
	"fmt"
)

func Error(text string) error {
	return errors.New(text)
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a))
}

// session
var (
	SessionMemberNotFound    = Error("member not found in the group")
	SessionClosedGroup       = Error("group is closed")
	SessionDuplication       = Error("session has existed in the current group")
	SessionNotFoundInContext = Error("session not found in context")
)

// route
var (
	RouteFieldCantEmpty = Error("route field can not be empty")
	RouteInvalid        = Error("invalid route")
)

// packet
var (
	PacketWrongType              = Error("wrong packet type")
	PacketSizeExceed             = Error("codec: packet size exceed")
	PacketConnectClosed          = Error("client connection closed")
	PacketInvalidHeader          = Error("invalid header")
	PacketMsgSmallerThanExpected = Error("received less data than expected, EOF?")
)

// message
var (
	MessageWrongType     = Error("wrong message type")
	MessageInvalid       = Error("invalid message")
	MessageRouteNotFound = Error("route info not found in dictionary")
)

var (
	ProtobufWrongValueType = Error("convert on wrong type value")
)

var (
	DiscoveryGetMemberListIsEmpty = Error("get member list is empty.")
)

// cluster
var (
	ClusterRPCClientIsStop = Error("rpc client is stop")
	ClusterNoImplement     = Error("no implement")
)