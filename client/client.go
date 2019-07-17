package client

import (
	"fmt"

	"github.com/evalphobia/httpwrapper/request"
)

// Client is http client for Face++ API.
type Client struct {
	Option
	APIKey    string
	APISecret string
}

func New() *Client {
	return &Client{}
}

func (c *Client) SetAPIKeys(apiKey, apiSecret string) {
	c.APIKey = apiKey
	c.APISecret = apiSecret
}

func (c *Client) SetOption(opt Option) {
	c.Option = opt
}

// CallPOST sends POST request to `url` with `params` and set reqponse to `result`
func (c *Client) CallPOST(path string, params, result interface{}) (err error) {
	opt := c.Option
	url := fmt.Sprintf("%s%s", opt.getBaseURL(), path)

	resp, err := request.POST(url, request.Option{
		Payload: parameter{
			Credential: Credential{c.APIKey, c.APISecret},
			Extra:      params,
		},
		PayloadType: request.PayloadTypeFORM,
		Retry:       opt.Retry,
		Debug:       opt.Debug,
		UserAgent:   opt.getUserAgent(),
		Timeout:     opt.getTimeout(),
	})
	if err != nil {
		return err
	}

	err = resp.JSON(result)
	return err
}

type parameter struct {
	Credential `url:",squash"`
	Extra      interface{} `url:",squash"`
}
