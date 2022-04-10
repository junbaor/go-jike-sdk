package jike

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go-jike-sdk/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BaseURL = "https://api.ruguoapp.com"
const HeaderAccessToken = "x-jike-access-token"
const HeaderRefreshToken = "x-jike-access-token"

type request struct {
	Debug      bool
	HTTPMethod string
	HTTPPath   string
	Params     map[string]string
	Input      interface{}
	Output     interface{}
	Context    context.Context
	req        *http.Request
}

type Option interface {
	apply(map[string]string)
}

var DeviceId = "4653BFCE-9D54-471C-809C-422AC240DA7B"
var IDFV = "5C5FC6BB-F6E6-4689-BB5A-E88763186C55"
var httpHeader = make(map[string]string)

func init() {
	httpHeader["Host"] = "api.ruguoapp.com"
	httpHeader["Cookie"] = "abtest_info={}; abtest_info.sig=TpMSLxutJSIb6SX-RcpsEJ9rvBM"
	httpHeader["king-card-status"] = "unavailable"
	httpHeader["client-request-id"] = uuid.NewString()
	httpHeader["user-agent"] = "jike/7.18.1 (com.ruguoapp.jike; build:1901; iOS 14.7.0) Alamofire/5.4.3"
	httpHeader["x-jike-device-properties"] = "{\"idfv\":\"" + IDFV + "\"}"
	httpHeader["app-buildno"] = "1901"
	httpHeader["x-jike-device-id"] = DeviceId
	httpHeader["os"] = "ios"
	httpHeader["manufacturer"] = "Apple"
	httpHeader["bundleid"] = "com.ruguoapp.jike"
	httpHeader["accept-language"] = "zh-Hans-CN;q=1.0"
	httpHeader["support-h265"] = "true"
	httpHeader["model"] = "iPhone10] = 1"
	httpHeader["app-permissions"] = "14"
	httpHeader["accept"] = "*/*"
	httpHeader["app-version"] = "7.18.1"
	httpHeader["wificonnected"] = "true"
	httpHeader["os-version"] = "Version 14.7 (Build 18G5033e)"
}

func (r *request) send(j *Jike) error {
	err := r.doSend(j)
	if err != nil {
		var response *utils.ErrorResponse
		if errors.As(err, &response) && response.StatusCode == 401 {
			err := refreshToken(j)
			if err != nil {
				return err
			}
			return r.doSend(j)
		}
	}
	return err
}

func refreshToken(j *Jike) error {
	req := &request{
		Debug:      j.UserService.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `/app_auth_tokens.refresh`,
		Context:    context.Background(),
	}
	return req.doSend(j)
}

func (r *request) doSend(j *Jike) error {
	var c = j.client
	if c == nil {
		return errors.New("miss http client")
	}
	if r.Context == nil {
		return errors.New("miss context.Context")
	}

	var err = r.buildRequest()
	if err != nil {
		return err
	}

	r.req.Header.Set(HeaderAccessToken, getAccessToken(j.areaCode, j.phone))
	r.req.Header.Set(HeaderRefreshToken, getRefreshToken(j.areaCode, j.phone))

	if r.Debug {
		utils.ShowInformation("Jike Request", r.req)
	}

	resp, err := c.Do(r.req)
	if err != nil {
		return errors.WithStack(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if r.Debug {
		utils.ShowInformation("Jike Response", string(body))
	}

	if err := r.handleError(resp, body); err != nil {
		return err
	}

	// set access token
	at := resp.Header.Get(HeaderAccessToken)
	if at != "" {
		err = setAccessToken(j.areaCode, j.phone, at)
		if err != nil {
			return err
		}
	}

	// set refresh token
	rt := resp.Header.Get(HeaderAccessToken)
	if rt != "" {
		err = setRefreshToken(j.areaCode, j.phone, rt)
		if err != nil {
			return err
		}
	}

	err = json.Unmarshal(body, r.Output)
	return errors.WithStack(err)
}

func (r *request) handleError(resp *http.Response, body []byte) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	e := &utils.ErrorResponse{}
	err := json.Unmarshal(body, e)
	if err != nil {
		e.StatusCode = resp.StatusCode
		return errors.WithStack(err)
	}
	e.StatusCode = resp.StatusCode
	return errors.WithStack(e)
}

func (r *request) buildRequest() error {
	if r.HTTPMethod == "" {
		r.HTTPMethod = http.MethodGet
	}

	var (
		req    *http.Request
		err    error
		uri    = BaseURL + r.HTTPPath
		params = buildQuery(r.Params)
	)

	if r.HTTPMethod == http.MethodPost && r.Input != nil {
		payload, err := json.Marshal(r.Input)
		if err != nil {
			return errors.WithStack(err)
		}
		req, err = http.NewRequestWithContext(r.Context, r.HTTPMethod, uri, bytes.NewReader(payload))
		if err != nil {
			return errors.WithStack(err)
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(r.Context, r.HTTPMethod, uri, nil)
		if err != nil {
			return errors.WithStack(err)
		}
		req.URL.RawQuery = params.Encode()
	}

	for k, v := range httpHeader {
		req.Header.Set(k, v)
	}

	r.req = req
	return nil
}

func buildQuery(input map[string]string) url.Values {
	q := url.Values{}
	if len(input) == 0 {
		return q
	}

	for k, v := range input {
		if v != "" {
			q.Add(k, v)
		}
	}

	return q
}

type Jike struct {
	debug        bool
	UserService  *UsersService
	TopicService *TopicService
	client       *http.Client
	areaCode     string
	phone        string
}

func NewJike(areaCode, phone string) *Jike {
	return NewJikeWithClient(areaCode, phone, &http.Client{})
}

func NewJikeWithClient(areaCode, phone string, client *http.Client) *Jike {
	f := &Jike{
		debug:    utils.IsDebug(),
		client:   client,
		areaCode: areaCode,
		phone:    phone,
	}
	f.initService(f)
	return f
}

func (f *Jike) initService(jike *Jike) {
	f.UserService = NewUsersService(jike, f.client, f.debug)
	f.TopicService = NewTopicService(jike, f.client, f.debug)
}
