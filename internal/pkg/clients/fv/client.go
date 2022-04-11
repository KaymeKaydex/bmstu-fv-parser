package fv

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/config"
)

const getWorkingOutPath = "/branch/ajax/event/day"

type Client struct {
	ctx        context.Context
	httpClient *http.Client
}

func New(ctx context.Context) *Client {
	return &Client{
		ctx:        ctx,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) WithTransport(transport *http.Transport) {
	c.httpClient.Transport = transport
}

func (c *Client) GetWorkingOut(req RequestGetWorkingOut) (*ResponseGetWorkingOut, error) {
	cfg := config.FromContext(c.ctx).FVConfig

	url := url.URL{
		Scheme: cfg.Protocol,
		Host:   cfg.SiteAddress,
		Path:   getWorkingOutPath,
	}

	fmt.Println(url.String())

	//req := http.NewRequest(http.MethodPost)
	//c.httpClient.Do()

	return nil, nil
}
