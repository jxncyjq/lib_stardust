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

package packet

import "errors"

// Type represents the network packet's type such as: handshake and so on.
type Type byte

const (
	_ Type = iota
	// Handshake represents a handshake: request(client) <====> handshake response(server) 客户端请求服务器
	Handshake = 0x01

	// HandshakeAck represents a handshake ack from client to server  服务器返回客户端
	HandshakeAck = 0x02

	// Heartbeat represents a heartbeat 客户端服务器心跳包
	Heartbeat = 0x03

	// Data represents a common data packet  公共数据包
	Data = 0x04

	// Kick represents a kick off packet  断开连接包
	Kick = 0x05 // disconnect message from server
)

// ErrWrongPomeloPacketType represents a wrong packet type.
var ErrWrongPomeloPacketType = errors.New("wrong packet type")

// ErrInvalidPomeloHeader represents an invalid header
var ErrInvalidPomeloHeader = errors.New("invalid header")
