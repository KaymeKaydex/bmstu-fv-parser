package fv

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/config"
)

// Обязательные постоянные заголовки
const (
	HeaderContentType string = "application/x-www-form-urlencoded; charset=UTF-8"
	HeaderUserAgent   string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Safari/605.1.15"
	// HeaderCookie todo: сделать автогенерацию сессий на каждый запрос
	HeaderCookie string = "sputnik_session=1650302481186|2; _ga=GA1.2.1957892335.1644173282; _gid=GA1.2.1043741576.1650302481; PHPSESSID=o7gik5iupctes1imnt0jdabu72"
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

	log.Info("generated url ", url.String())

	rawReq := &RequestGetWorkingOut{
		Id:            14,
		Date:          time.Date(2022, 4, 18, 0, 0, 0, 0, time.Local),
		SecurityLSKey: "cfeff17599c13002d5685ca2c4fe25e5",
	}

	rawReqBody := []byte(rawReq.String())

	reqToFV, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(rawReqBody))
	if err != nil {
		return nil, err
	}

	// подсовываем заголовки в запрос, чтобы обойти проверки сервера
	reqToFV.Header.Set("Content-Type", HeaderContentType)
	reqToFV.Header.Set("Host", cfg.SiteAddress)
	reqToFV.Header.Set("User-Agent", HeaderUserAgent)
	reqToFV.Header.Set("Cookie", HeaderCookie)

	r, err := c.httpClient.Do(reqToFV)
	if err != nil {
		return nil, err
	}

	bts, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var resp ResponseGetWorkingOut

	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
