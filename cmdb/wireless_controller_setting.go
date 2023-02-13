package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateWirelessControllerSetting(payload *models.WirelessControllerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	log.Printf("[INFO] Resource at path %q is complex type. Changing to UPDATE.", models.WirelessControllerSettingPath)
	return c.UpdateWirelessControllerSetting("", payload, params)
}

func (c *Client) ReadWirelessControllerSetting(mkey string, params *models.CmdbRequestParams) (*models.WirelessControllerSetting, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.WirelessControllerSettingPath
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
		v := models.WirelessControllerSetting{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateWirelessControllerSetting(mkey string, payload *models.WirelessControllerSetting, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.WirelessControllerSettingPath
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteWirelessControllerSetting(mkey string, params *models.CmdbRequestParams) error {
	payload := &models.WirelessControllerSetting{}
	_, err := c.UpdateWirelessControllerSetting("", payload, params)
	return err
}
