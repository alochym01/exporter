package storage

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

type RedFishClient struct {
	User    string
	Pass    string
	Timeout time.Duration
}

// Get ....
func (rf RedFishClient) Get(url string) ([]byte, error) {
	// Make a http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Add header accept application/json
	req.Header.Add("Accept", `application/json`)

	// Set username/password in a http request
	req.SetBasicAuth(rf.User, rf.Pass)

	// http.Transport
	transport := http.DefaultTransport.(*http.Transport).Clone()
	// transport.IdleConnTimeout = 60 * time.Second
	// transport.MaxIdleConns = 10
	// transport.MaxConnsPerHost = 10
	// transport.MaxIdleConnsPerHost = 10
	// // Disable SSL check
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// http.Client setup
	c := http.Client{
		Transport: transport,
		Timeout:   rf.Timeout,
	}

	// Make http Request
	res, err := c.Do(req)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}

	// Close http.Request connection
	defer res.Body.Close()

	// read the whole body into a []bytes
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewClient return a RedFishClient struct
func NewClient(user, pass string, t time.Duration) RedFishClient {
	return RedFishClient{
		User:    user,
		Pass:    pass,
		Timeout: t * time.Second,
	}
}
