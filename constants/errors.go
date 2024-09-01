// Copyright (c) nano Author and TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package constants

import (
	"github.com/jxncyjq/lib_stardust/core/errors"
)

// Errors that can occur during message handling.
var (
	ErrBindingNotFound                = errors.New("binding for this user was not found in etcd", 100)
	ErrBrokenPipe                     = errors.New("broken low-level pipe", 101)
	ErrBufferExceed                   = errors.New("session send buffer exceed", 102)
	ErrChangeDictionaryWhileRunning   = errors.New("you shouldn't change the dictionary while the app is already running", 103)
	ErrChangeRouteWhileRunning        = errors.New("you shouldn't change routes while app is already running", 104)
	ErrCloseClosedGroup               = errors.New("close closed group", 105)
	ErrCloseClosedSession             = errors.New("close closed session", 106)
	ErrClosedGroup                    = errors.New("group closed", 107)
	ErrEmptyUID                       = errors.New("empty uid", 108)
	ErrEtcdGrantLeaseTimeout          = errors.New("timed out waiting for etcd lease grant", 108)
	ErrEtcdLeaseNotFound              = errors.New("etcd lease not found in group", 109)
	ErrFrontSessionCantPushToFront    = errors.New("frontend session can't push to front", 110)
	ErrFrontendTypeNotSpecified       = errors.New("for using SendPushToUsers from a backend server you have to specify a valid frontendType", 111)
	ErrGroupAlreadyExists             = errors.New("group already exists", 112)
	ErrGroupNotFound                  = errors.New("group not found", 113)
	ErrIllegalUID                     = errors.New("illegal uid", 114)
	ErrInvalidCertificates            = errors.New("certificates must be exactly two", 115)
	ErrInvalidSpanCarrier             = errors.New("tracing: invalid span carrier", 116)
	ErrKickingUsers                   = errors.New("failed to kick users, check array with failed uids", 117)
	ErrMemberAlreadyExists            = errors.New("member already exists in group", 118)
	ErrMemberNotFound                 = errors.New("member not found in the group", 119)
	ErrMemoryTTLNotFound              = errors.New("memory group TTL not found", 120)
	ErrMetricNotKnown                 = errors.New("the provided metric does not exist", 121)
	ErrNatsMessagesBufferSizeZero     = errors.New("pitaya.buffer.cluster.rpc.server.nats.messages cant be zero", 122)
	ErrNatsNoRequestTimeout           = errors.New("pitaya.cluster.rpc.client.nats.requesttimeout cant be empty", 123)
	ErrNatsPushBufferSizeZero         = errors.New("pitaya.buffer.cluster.rpc.server.nats.push cant be zero", 124)
	ErrNilCondition                   = errors.New("pitaya/timer: nil condition", 125)
	ErrNoBindingStorageModule         = errors.New("for sending remote pushes or using unique session module while using grpc you need to pass it a BindingStorage", 126)
	ErrNoConnectionToServer           = errors.New("rpc client has no connection to the chosen server", 127)
	ErrNoContextFound                 = errors.New("no context found", 128)
	ErrNoNatsConnectionString         = errors.New("you have to provide a nats url", 129)
	ErrNoServerTypeChosenForRPC       = errors.New("no server type chosen for sending RPC, send a full route in the format server.service.component", 130)
	ErrNoServerWithID                 = errors.New("can't find any server with the provided ID", 131)
	ErrNoServersAvailableOfType       = errors.New("no servers available of this type", 132)
	ErrNoUIDBind                      = errors.New("you have to bind an UID to the session to do that", 133)
	ErrNonsenseRPC                    = errors.New("you are making a rpc that may be processed locally, either specify a different server type or specify a server id", 134)
	ErrNotImplemented                 = errors.New("method not implemented", 135)
	ErrNotifyOnRequest                = errors.New("tried to notify a request route", 136)
	ErrOnCloseBackend                 = errors.New("onclose callbacks are not allowed on backend servers", 137)
	ErrProtodescriptor                = errors.New("failed to get protobuf message descriptor", 138)
	ErrPushingToUsers                 = errors.New("failed to push message to users, check array with failed uids", 139)
	ErrRPCClientNotInitialized        = errors.New("RPC client is not running", 140)
	ErrRPCJobAlreadyRegistered        = errors.New("rpc job was already registered", 141)
	ErrRPCLocal                       = errors.New("RPC must be to a different server type", 142)
	ErrRPCServerNotInitialized        = errors.New("RPC server is not running", 143)
	ErrReplyShouldBeNotNull           = errors.New("reply must not be null", 144)
	ErrReplyShouldBePtr               = errors.New("reply must be a pointer", 145)
	ErrRequestOnNotify                = errors.New("tried to request a notify route", 146)
	ErrRouterNotInitialized           = errors.New("router is not initialized", 147)
	ErrServerNotFound                 = errors.New("server not found", 148)
	ErrServiceDiscoveryNotInitialized = errors.New("service discovery client is not initialized", 149)
	ErrSessionAlreadyBound            = errors.New("session is already bound to an uid", 150)
	ErrSessionDuplication             = errors.New("session exists in the current group", 151)
	ErrSessionNotFound                = errors.New("session not found", 152)
	ErrSessionOnNotify                = errors.New("current session working on notify mode", 153)
	ErrTimeoutTerminatingBinaryModule = errors.New("timeout waiting to binary module to die", 154)
	ErrWrongValueType                 = errors.New("protobuf: convert on wrong type value", 155)
	ErrRateLimitExceeded              = errors.New("rate limit exceeded", 156)
	ErrReceivedMsgSmallerThanExpected = errors.New("received less data than expected, EOF?", 157)
	ErrReceivedMsgBiggerThanExpected  = errors.New("received more data than expected", 158)
	ErrConnectionClosed               = errors.New("client connection closed", 159)
)
