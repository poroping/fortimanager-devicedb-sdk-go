package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateLogFortianalyzerCloudOverrideSetting(payload *models.LogFortianalyzerCloudOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogFortianalyzerCloudOverrideSettingPath)
	return c.UpdateLogFortianalyzerCloudOverrideSetting("", payload, params)
}

func (c *Client) ReadLogFortianalyzerCloudOverrideSetting(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzerCloudOverrideSetting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogFortianalyzerCloudOverrideSettingPath
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
		v := models.LogFortianalyzerCloudOverrideSetting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateLogFortianalyzerCloudOverrideSetting(mkey string, payload *models.LogFortianalyzerCloudOverrideSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogFortianalyzerCloudOverrideSettingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogFortianalyzerCloudOverrideSetting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogFortianalyzerCloudOverrideSetting{}
	_, err := c.UpdateLogFortianalyzerCloudOverrideSetting("", payload, params)
	return err
}
