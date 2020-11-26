package request

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/anacoondaaa/ncloud-sdk-go/oauth"
)

func encryptHmacSha256(message, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))
	sha := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sha)
}

// NewRequest is http request with oauth
func NewRequest(accessKey string, secretKey string, method string, url string, params map[string]string) ([]byte, *http.Response, error) {
	c := oauth.NewConsumer(accessKey, secretKey, method, url)

	for k, v := range params {
		c.AdditionalParams[k] = v
	}

	reqURL, err := c.GetRequestUrl()
	if err != nil {
		return nil, nil, err
	}
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)

	req, err := http.NewRequest(method, reqURL, nil)
	if err != nil {
		return nil, nil, err
	}

	message := method + " " + url + "\n" + timestamp + "\n" + accessKey

	apigwSignature := encryptHmacSha256(message, secretKey)

	req.Header.Set("x-ncp-iam-access-key", accessKey)
	req.Header.Set("x-ncp-apigw-signature-v2", apigwSignature)
	req.Header.Set("x-ncp-apigw-timestamp", timestamp)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes, resp, nil
}
