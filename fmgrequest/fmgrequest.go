package fmgrequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/poroping/fortimanager-devicedb-sdk-go/config"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func Create(c *config.Config, r *models.FmgCmdbRequest) (*models.FmgCmdbResponse, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	log.Printf("[DEBUG] Body: %s", string(body))
	if err != nil {
		return nil, err
	}

	response := &models.FmgCmdbResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during CREATE/UPDATE %s", err)
		return nil, err
	}

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func CreateNoSession(c *config.Config, r *models.FmgCmdbRequest) (*models.FmgCmdbResponseNoSession, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	log.Printf("[DEBUG] Body: %s", string(body))
	if err != nil {
		return nil, err
	}

	response := &models.FmgCmdbResponseNoSession{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during CREATE/UPDATE %s", err)
		return nil, err
	}

	err = fortiErrorCheckNoSession(body, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func Update(c *config.Config, r *models.FmgCmdbRequest) (*models.FmgCmdbResponse, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &models.FmgCmdbResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during CREATE/UPDATE %s", err)
		return nil, err
	}

	// return empty if object doesn't exist
	if response.Result[0].Status.Message == "Object does not exist" {
		empty := &models.FmgCmdbResponse{}
		// shitmix to allow determine object doesn't exist needs some refactoring
		s := "notexist"
		f := false
		empty.Exists = &f
		empty.Session = &s
		empty.Result = append(empty.Result, models.Result{})
		return empty, nil
	}

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}
	t := true
	response.Exists = &t

	return response, err
}

func Read(c *config.Config, r *models.FmgCmdbRequest) (*models.FmgCmdbResponse, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &models.FmgCmdbResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during READ")
		return nil, err
	}

	// return empty if object doesn't exist
	if response.Result[0].Status.Message == "Object does not exist" {
		empty := &models.FmgCmdbResponse{}
		log.Printf("[WARN] Object not found, marking exist as false")
		f := false
		empty.Exists = &f
		empty.Result = append(empty.Result, models.Result{})
		return empty, nil
	}

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}

	t := true
	response.Exists = &t

	return response, err
}

func Read2(c *config.Config, r *models.FmgCmdbRequest) (*models.FmgCmdbResponse, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		// handle 404s (search for resource on delete at least)
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &models.FmgCmdbResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during READ")
		return nil, err
	}

	err = fortiErrorCheck(body, response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func Delete(c *config.Config, r *models.FmgCmdbRequest) (err error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := models.FmgCmdbResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("[ERROR] Error reading response body during DELETE")
		return err
	}

	err = fortiErrorCheck(body, &response)
	if err != nil {
		return err
	}

	return nil
}

func fortiErrorCheckNoSession(body []byte, res *models.FmgCmdbResponseNoSession) (err error) {
	if len(res.Result) < 1 || res.Result == nil {
		return fmt.Errorf("[ERROR] Result length < 1")
	}
	if res.Result[0].Status.Code != 0 {
		return fmt.Errorf("[ERROR] %s", res.Result[0].Status.Message)
	}
	return nil
}

func fortiErrorCheck(body []byte, res *models.FmgCmdbResponse) (err error) {
	if len(res.Result) < 1 || res.Result == nil {
		return fmt.Errorf("[ERROR] Result length < 1")
	}
	if res.Result[0].Status.Code != 0 {
		return fmt.Errorf("[ERROR] %s", res.Result[0].Status.Message)
	}
	return nil
}

func parseErrorCode(errorCode int64, errorString string) string {
	switch e := errorCode; e {
	case int64(-3):
		// parse errorString and return troubled datasource
		datasource := ""
		find := *regexp.MustCompile(`.*value parse error before '(.*)'.*`)
		match := find.FindAllStringSubmatch(errorString, -1)
		if len(match) == 1 {
			datasource = match[0][1]
		}
		return fmt.Sprintf("Referenced datasource %q does not exist\n", datasource)
	case int64(-5):
		return "Mkey already exists\n"
	case int64(-8):
		return "Invalid IP\n"
	case int64(-9):
		return "Invalid netmask\n"
	case int64(-15):
		return "Duplicated entry\n"
	case int64(-23):
		// potentially look at listing references?
		return "Unable to remove, target is referenced elsewhere\n"
	case int64(-651):
		return "Missing required attribute or attribute incorrect\n"
	default:
		return "No additional context available\n"
	}
}

func newRequest(c config.Config, r *models.FmgCmdbRequest) (*http.Request, error) {
	headers := buildHeaders(c, r)
	// set URL w/ url queries
	url := buildURL(c, r)

	r.Session = c.Auth.Session
	r.ID = 1
	r.Verbose = 1

	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	if body != nil {
		log.Printf("[DEBUG] BODY: %s", string(body))
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = *headers

	return req, nil
}

func buildHeaders(c config.Config, r *models.FmgCmdbRequest) *http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", c.UserAgent)

	return &headers
}

func buildURL(c config.Config, r *models.FmgCmdbRequest) *url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = c.Auth.Hostname
	// this shitmix is needed cause the fortios needs to query escape the mkey even though it's part of the path
	// path, _ := url.Parse(r.Path)
	u.Path = "jsonrpc"
	// u.RawPath = r.Path
	// q := marshalParams(&r.Params)
	// if r.Params.Vdom != "" {
	// 	q.Set("vdom", r.Params.Vdom)
	// } else {
	// 	q.Set("vdom", c.Auth.Vdom)
	// }
	// if r.NoVdom {
	// 	q.Del("vdom")
	// }

	// u.RawQuery = q.Encode()

	return &u
}
