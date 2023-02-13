package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateLogTacacsaccounting2Setting(payload *models.LogTacacsaccounting2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogTacacsaccounting2SettingPath)
	return c.UpdateLogTacacsaccounting2Setting("", payload, params)
}

func (c *Client) ReadLogTacacsaccounting2Setting(mkey string, params *models.CmdbRequestParams) (*models.LogTacacsaccounting2Setting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogTacacsaccounting2SettingPath
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
		v := models.LogTacacsaccounting2Setting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateLogTacacsaccounting2Setting(mkey string, payload *models.LogTacacsaccounting2Setting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogTacacsaccounting2SettingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogTacacsaccounting2Setting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogTacacsaccounting2Setting{}
	_, err := c.UpdateLogTacacsaccounting2Setting("", payload, params)
	return err
}
