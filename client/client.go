package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/poroping/fortimanager-devicedb-sdk-go/auth"
	"github.com/poroping/fortimanager-devicedb-sdk-go/cmdb"
	"github.com/poroping/fortimanager-devicedb-sdk-go/config"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

type FortiSDKClient struct {
	Config config.Config
	Cmdb   cmdb.Endpoints
}

func NewClient(auth *auth.Auth) (*FortiSDKClient, error) {
	c, err := NewClientBase(auth)
	if err != nil {
		return nil, err
	}

	c.Cmdb = cmdb.New(&c.Config)

	return c, nil
}

func NewClientBase(auth *auth.Auth) (*FortiSDKClient, error) {
	insecure := auth.Insecure

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	dc := &http.Client{}
	dc.Transport = tr
	dc.Timeout = 30 * time.Second

	c := &FortiSDKClient{}
	c.Config.Auth = auth
	c.Config.HTTPCon = dc
	c.Config.FwTarget = auth.Hostname
	c.Config.UserAgent = "FortiManager DeviceDB Go-SDK"
	c.Config.Retries = 5
	c.Login()

	// if !auth.AutoVersion {
	// 	log.Print("[INFO] Skipping FortiOS version check.")
	// } else {
	// 	log.Print("[INFO] Attempting to determine FortiOS version.")
	// 	req := &models.CmdbRequest{}
	// 	req.HTTPMethod = "GET"
	// 	req.NoVdom = true
	// 	req.Path = models.MonitorBasePath + "system/status/"
	// 	res, err := request.Read(&c.Config, req)
	// 	if err != nil || res == nil {
	// 		log.Print("[WARN] Error attempting to determine FortiOS version, skipping.")
	// 	}
	// 	if res.Version != nil {
	// 		log.Printf("[INFO] FortiOS version detected as %s", *res.Version)
	// 		c.Config.Fv = *res.Version
	// 	}
	// }

	return c, nil
}

func (c *FortiSDKClient) Login() {
	// make this from model and return errors instead.
	s := fmt.Sprintf(`{
		"id": 1,
		"method": "exec",
		"params": [
			{
				"data": {
					"passwd": "%s",
					"user": "%s"
				},
				"url": "/sys/login/user"
			}
		]
	}`, c.Config.Auth.Password, c.Config.Auth.Username)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/jsonrpc", c.Config.Auth.Hostname), strings.NewReader(s))
	if err != nil {
		log.Fatalf("[FATAL] Error creating login request")
	}
	res, err := c.Config.HTTPCon.Do(req)
	if err != nil {
		log.Fatalf("[FATAL] Error performing login request, %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("[FATAL] Error getting body from login response")
	}
	m := &models.FmgCmdbResponse{}
	err = json.Unmarshal(body, m)
	if err != nil {
		log.Fatalf("[FATAL] Error unmarshalling login response")
	}
	if m.Session != nil {
		c.Config.Auth.Session = *m.Session
	} else {
		log.Print("[ERROR] Session not found in login response.")
	}
}

func (c *FortiSDKClient) Logout() {
	if c.Config.Auth.Session == "" {
		return
	}
	s := fmt.Sprintf(`{
		"id": 1,
		"method": "exec",
		"params": [
			{
				"url": "/sys/logout"
			}
		],
		"session": %s
	}`, c.Config.Auth.Session)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/jsonrpc", c.Config.Auth.Hostname), strings.NewReader(s))
	if err != nil {
		log.Print("[INFO] Error creating logout request")
	}
	res, err := c.Config.HTTPCon.Do(req)
	if err != nil {
		log.Print("[WARN] Error performing logout request, %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print("[WARN] Error getting body from logout response")
	}
	m := &models.FmgCmdbResponse{}
	err = json.Unmarshal(body, m)
	if err != nil {
		log.Print("[WARN] Error unmarshalling logout response")
	}
	if m.Result[0].Status.Message == "OK" {
		log.Print("[INFO] Successfully logged out")
	} else {
		log.Print("[ERROR] Error during logout. Check admin session is not hanging")
	}
}
