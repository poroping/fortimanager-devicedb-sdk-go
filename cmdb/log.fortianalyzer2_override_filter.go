package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateLogFortianalyzer2OverrideFilter(payload *models.LogFortianalyzer2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.LogFortianalyzer2OverrideFilterPath)
	return c.UpdateLogFortianalyzer2OverrideFilter("", payload, params)
}

func (c *Client) ReadLogFortianalyzer2OverrideFilter(mkey string, params *models.CmdbRequestParams) (*models.LogFortianalyzer2OverrideFilter, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.LogFortianalyzer2OverrideFilterPath
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
		v := models.LogFortianalyzer2OverrideFilter{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateLogFortianalyzer2OverrideFilter(mkey string, payload *models.LogFortianalyzer2OverrideFilter, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.LogFortianalyzer2OverrideFilterPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteLogFortianalyzer2OverrideFilter(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.LogFortianalyzer2OverrideFilter{}
	_, err := c.UpdateLogFortianalyzer2OverrideFilter("", payload, params)
	return err
}
