package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/poroping/fortimanager-devicedb-sdk-go/config"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func Fmg2FgtResp(r1 *models.FmgCmdbResponse) (*models.CmdbResponse, error) {
	r2 := &models.CmdbResponse{}
	res := make([]interface{}, 0)
	res = append(res, r1.Result[0].Data)
	r2.Results = res
	var mkey interface{}
	if data, ok := r1.Result[0].Data.(map[string]interface{}); ok {
		if len(data) == 1 {
			for _, v := range data {
				mkey = v
			}
		}
	}
	r2.Mkey = mkey
	return r2, nil
}

func Fgt2FmgReq(c *config.Config, r1 *models.CmdbRequest) (*models.FmgCmdbRequest, error) {
	r2 := &models.FmgCmdbRequest{}
	if r1.HTTPMethod == "POST" {
		r2.Method = "add"
	} else if r1.HTTPMethod == "PUT" {
		r2.Method = "update"
	} else if r1.HTTPMethod == "GET" {
		r2.Method = "get"
	} else if r1.HTTPMethod == "DELETE" {
		r2.Method = "delete"
	}
	if r1.Params.Vdom == "" {
		r1.Params.Vdom = "root"
	}
	r1.Params.Vdom = "vdom/" + r1.Params.Vdom
	if r1.Params.Vdom == "vdom/global" {
		r1.NoVdom = true
	}
	if r1.NoVdom {
		r1.Params.Vdom = "global"
	}
	targetDevice := c.FwTarget
	if r1.Params.Scope != "" {
		targetDevice = r1.Params.Scope
	}
	path := strings.TrimPrefix(r1.Path, "/pm/config/device/")
	r2.ID = 1
	r2.Verbose = 1
	params := models.FmgCmdbRequestPayload{}
	if r1.Payload != nil {
		var data interface{}
		log.Println(string(r1.Payload))
		err := json.Unmarshal(r1.Payload, &data)
		if err != nil {
			log.Printf("[ERROR] unable to unmarshal.")
			return nil, err
		}
		params.Data = data
	}
	params.URL = fmt.Sprintf("/pm/config/device/%s/%s/%s", targetDevice, r1.Params.Vdom, path)
	r2.Params = append(r2.Params, params)
	r2.Session = c.Auth.Session
	return r2, nil
}
