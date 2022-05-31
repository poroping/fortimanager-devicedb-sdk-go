package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateLogFortiguardSetting(payload *models.LogFortiguardSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogFortiguardSettingPath)
	return c.UpdateLogFortiguardSetting("", payload, params)
}

func (c *Client) ReadLogFortiguardSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortiguardSetting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogFortiguardSettingPath
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
		v := models.LogFortiguardSetting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateLogFortiguardSetting(mkey string, payload *models.LogFortiguardSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogFortiguardSettingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogFortiguardSetting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogFortiguardSetting{}
	_, err := c.UpdateLogFortiguardSetting("", payload, params)
	return err
}
