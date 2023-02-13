package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateLogMemoryGlobalSetting(payload *models.LogMemoryGlobalSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogMemoryGlobalSettingPath)
	return c.UpdateLogMemoryGlobalSetting("", payload, params)
}

func (c *Client) ReadLogMemoryGlobalSetting(mkey string, params *models.CmdbRequestParams) (*models.LogMemoryGlobalSetting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogMemoryGlobalSettingPath
	req.Params = *params

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results
	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp[0])
		if err != nil {
			return nil, err
		}
		v := models.LogMemoryGlobalSetting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateLogMemoryGlobalSetting(mkey string, payload *models.LogMemoryGlobalSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogMemoryGlobalSettingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogMemoryGlobalSetting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogMemoryGlobalSetting{}
	_, err := c.UpdateLogMemoryGlobalSetting("", payload, params)
	return err
}
