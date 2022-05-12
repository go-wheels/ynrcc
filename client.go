package ynrcc

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"

	"github.com/go-wheels/ynrcc/sm2util"
)

var ErrInvalidSign = errors.New("invalid sign")

type Options struct {
	GatewayURL string
	Cert       string
	PriKey     string
	BankPubKey string
	MerID      string
	TemID      string
}

type Client struct {
	httpClient http.Client
	gatewayURL string
	cert       string
	merID      string
	temID      string
	priKey     *sm2.PrivateKey
	bankPubKey *sm2.PublicKey
}

func NewClient(options Options) (client *Client, err error) {
	client = &Client{
		gatewayURL: options.GatewayURL,
		cert:       options.Cert,
		merID:      options.MerID,
		temID:      options.TemID,
	}
	client.priKey, err = x509.ReadPrivateKeyFromHex(options.PriKey)
	if err != nil {
		client = nil
		return
	}
	client.bankPubKey, err = x509.ReadPublicKeyFromHex(options.BankPubKey)
	if err != nil {
		client = nil
		return
	}
	return
}

func (c *Client) MerID() string { return c.merID }

func (c *Client) TemID() string { return c.temID }

func (c *Client) Execute(req, res any) (err error) {
	reqBody, err := json.Marshal(request{Request: req})
	if err != nil {
		return
	}
	httpReq, err := http.NewRequest(http.MethodPost, c.gatewayURL, bytes.NewReader(reqBody))
	if err != nil {
		return
	}
	reqSign, err := c.sign(reqBody)
	if err != nil {
		return
	}
	httpReq.Header.Set("ynrcc-sign", reqSign)
	httpReq.Header.Set("ynrcc-cert", c.cert)

	httpRes, err := c.httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpRes.Body.Close()

	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return
	}
	resWrap := &response{Response: res}
	err = json.Unmarshal(resBody, resWrap)
	if err != nil {
		return
	}
	resSign := httpRes.Header.Get("ynrcc-sign")
	if !c.verify(resBody, resSign) {
		err = ErrInvalidSign
	}
	return
}

func (c *Client) ReadNotify(httpReq *http.Request) (req *PayResultMerNotifyRequest, err error) {
	defer httpReq.Body.Close()
	reqBody, err := io.ReadAll(httpReq.Body)
	if err != nil {
		return
	}
	req = &PayResultMerNotifyRequest{}
	reqWrap := &request{Request: req}
	err = json.Unmarshal(reqBody, reqWrap)
	if err != nil {
		req = nil
		return
	}
	reqSign := httpReq.Header.Get("ynrcc-sign")
	if !c.verify(reqBody, reqSign) {
		err = ErrInvalidSign
	}
	return
}

func (c *Client) verify(data []byte, sign string) (ok bool) {
	digest := sha256Sum(data)
	r, s, ok := decodeSign(sign)
	if !ok {
		return
	}
	ok = sm2.Verify(c.bankPubKey, digest, r, s)
	return
}

func (c *Client) sign(data []byte) (sign string, err error) {
	digest := sha256Sum(data)
	r, s, err := sm2util.Sign(c.priKey, digest, nil)
	if err != nil {
		return
	}
	sign = encodeSign(r, s)
	return
}

func encodeSign(r, s *big.Int) (sign string) {
	sign = fmt.Sprintf("%X%X", r, s)
	return
}

func decodeSign(sign string) (r, s *big.Int, ok bool) {
	if len(sign) != 128 {
		return
	}
	r, ok2 := new(big.Int).SetString(sign[:64], 16)
	s, ok3 := new(big.Int).SetString(sign[64:], 16)
	ok = ok2 && ok3
	if !ok {
		r, s = nil, nil
	}
	return
}

func sha256Sum(data []byte) (digest []byte) {
	hash := sha256.New()
	hash.Write(data)
	digest = hash.Sum(nil)
	return
}
