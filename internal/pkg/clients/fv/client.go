package fv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
	"moul.io/http2curl"

	"github.com/KaymeKaydex/bmstu-fv-parser.git/internal/app/config"
)

// Обязательные постоянные заголовки
const (
	HeaderContentType string = "application/x-www-form-urlencoded; charset=UTF-8"
	HeaderUserAgent   string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Safari/605.1.15"
	// HeaderCookie todo: сделать автогенерацию сессий на каждый запрос
	HeaderCookie string = "sputnik_session=1652722300905|2; _ga=GA1.2.1421662848.1652722301; _gid=GA1.2.193746774.1652722301; _gat_gtag_UA_54621671_2=1; sp_test=1; PHPSESSID=918kkgh22urvoa3j8t65qj061n"
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

	rawReqBody := []byte(req.String())

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

	command, _ := http2curl.GetCurlCommand(reqToFV)
	fmt.Println(command)

	bts, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	log.Debug(string(bts))

	var resp ResponseGetWorkingOut

	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
