package cmdb

import (
	"encoding/json"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
	"github.com/poroping/fortimanager-devicedb-sdk-go/request"
)

func (c *Client) CreateVpnIpsecPhase1(payload *models.VpnIpsecPhase1, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	mkey := ""
	if payload.Name != nil && *params.AllowAppend {
		mkey = *payload.Name
		read, err := c.ReadVpnIpsecPhase1(mkey, params)
		if err != nil {
			return nil, err
		}
		if read != nil {
			log.Printf("[WARN] Resource at path %q with mkey %q detected upon CREATE with flag set to to overwrite. Changing to UPDATE.", models.VpnIpsecPhase1Path, mkey)
			return c.UpdateVpnIpsecPhase1(mkey, payload, params)
		}
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "POST"
	req.Payload = body
	req.Path = models.CmdbBasePath + models.VpnIpsecPhase1Path
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ReadVpnIpsecPhase1(mkey string, params *models.CmdbRequestParams) (*models.VpnIpsecPhase1, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VpnIpsecPhase1Path + mkey + "/"
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
		v := models.VpnIpsecPhase1{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateVpnIpsecPhase1(mkey string, payload *models.VpnIpsecPhase1, params *models.CmdbRequestParams) (*models.CmdbResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.CmdbRequest{}
	req.HTTPMethod = "PUT"
	req.Mkey = &mkey
	req.Payload = body
	req.Path = models.CmdbBasePath + models.VpnIpsecPhase1Path + mkey + "/"
	req.Params = *params

	res, err := request.CreateUpdate(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) DeleteVpnIpsecPhase1(mkey string, params *models.CmdbRequestParams) error {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "DELETE"
	req.Mkey = &mkey
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VpnIpsecPhase1Path + mkey + "/"
	req.Params = *params

	err := request.Delete(c.config, req)
	return err
}

func (c *Client) ListVpnIpsecPhase1(params *models.CmdbRequestParams) (*[]models.VpnIpsecPhase1, error) {
	req := &models.CmdbRequest{}
	req.HTTPMethod = "GET"
	req.Payload = nil
	req.Path = models.CmdbBasePath + models.VpnIpsecPhase1Path
	req.Params = *params

	res, err := request.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results

	if tmp, ok := res.Results.([]interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := []models.VpnIpsecPhase1{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}
