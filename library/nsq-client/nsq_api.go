package nsq_client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type deadlinedConn struct {
	Timeout time.Duration
	net.Conn
}

func (c *deadlinedConn) Read(b []byte) (n int, err error) {
	if err = c.Conn.SetReadDeadline(time.Now().Add(c.Timeout));err != nil{
		return 0,err
	}else {
		return c.Conn.Read(b)
	}
}

func (c *deadlinedConn) Write(b []byte) (n int, err error) {
	if err = c.Conn.SetWriteDeadline(time.Now().Add(c.Timeout));err != nil {
		return 0,err
	} else {
		return c.Conn.Write(b)
	}
}

func NewDeadlineTransport(timeout time.Duration) *http.Transport {
	transport := &http.Transport{
		DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, timeout)
			if err != nil {
				return nil, err
			}
			return &deadlinedConn{timeout, c}, nil
		},
	}
	return transport
}

type wrappedResp struct {
	Status     string      `json:"status_txt"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

// stores the result in the value pointed to by ret(must be a pointer)
func apiRequestNegotiateV1(method string, endpoint string, body io.Reader, ret interface{}) error {
	httpclient := &http.Client{Transport: NewDeadlineTransport(2 * time.Second)}
	req, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.nsq; version=1.0")

	resp, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %s %q", resp.Status, respBody)
	}

	if len(respBody) == 0 {
		respBody = []byte("{}")
	}

	return json.Unmarshal(respBody, ret)
}
