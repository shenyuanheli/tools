/**
  @author: 1043193460@qq.com
  @date: 2022/12/13 14:32
  @note:
**/
package net

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
}

func (c *Client) NewRequest() *Request {
	r := &Request{
		timeout: 30,
		headers: map[string]string{},
		cookies: map[string]string{},
	}
	return r
}

// Debug model
func (c *Client) Debug(v bool) *Request {
	r := c.NewRequest()
	return r.Debug(v)
}

func (c *Client) Jar(v http.CookieJar) *Request {
	r := c.NewRequest()
	return r.Jar(v)
}

func (c *Client) DisableKeepAlives(v bool) *Request {
	r := c.NewRequest()
	return r.DisableKeepAlives(v)
}

func (c *Client) CheckRedirect(v func(req *http.Request, via []*http.Request) error) *Request {
	r := c.NewRequest()
	return r.CheckRedirect(v)
}

func (c *Client) TLSClient(v *tls.Config) *Request {
	r := c.NewRequest()
	return r.SetTLSClient(v)
}

func (c *Client) SetTLSClient(v *tls.Config) *Request {
	r := c.NewRequest()
	return r.SetTLSClient(v)
}

func (c *Client) SetHeaders(headers map[string]string) *Request {
	r := c.NewRequest()
	return r.SetHeaders(headers)
}

func (c *Client) SetCookies(cookies map[string]string) *Request {
	r := c.NewRequest()
	return r.SetCookies(cookies)
}

func (c *Client) JSON() *Request {
	r := c.NewRequest()
	return r.JSON()
}

func (c *Client) Proxy(v func(*http.Request) (*url.URL, error)) *Request {
	r := c.NewRequest()
	return r.Proxy(v)
}

func (c *Client) SetTimeout(d time.Duration) *Request {
	r := c.NewRequest()
	return r.SetTimeout(d)
}

func (c *Client) Transport(v *http.Transport) *Request {
	r := c.NewRequest()
	return r.Transport(v)
}

// Get is a get http request
func (c *Client) Get(url string, data ...interface{}) (string, error) {
	r := c.NewRequest()
	return r.Get(url, data...)
}

func (c *Client) Post(url string, data ...interface{}) (string, error) {
	r := c.NewRequest()
	return r.Post(url, data...)
}

// Put is a put http request
func (c *Client) Put(url string, data ...interface{}) (string, error) {
	r := c.NewRequest()
	return r.Put(url, data...)
}

// Delete is a delete http request
func (c *Client) Delete(url string, data ...interface{}) (string, error) {
	r := c.NewRequest()
	return r.Delete(url, data...)
}

// UploadFilePath File upload via file path
func (c *Client) UploadFilePath(url, filename, filePath string) (string, error) {
	r := c.NewRequest()
	return r.UploadFilePath(url, filename, filePath)
}
