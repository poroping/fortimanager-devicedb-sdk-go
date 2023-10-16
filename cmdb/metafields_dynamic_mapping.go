package cmdb

import (
	"encoding/json"
	"fmt"

	"github.com/poroping/fortimanager-devicedb-sdk-go/fmgrequest"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func createMetafieldDynamicMappingPath(adom, variable string) string {
	if adom == "" {
		adom = "root"
	}
	return fmt.Sprintf("pm/config/adom/%s/obj/fmg/variable/%s/dynamic_mapping/", adom, variable)
}

func (c *Client) CreateMetafieldDynamicMapping(variable string, payload *models.DynamicMapping, params *models.CmdbRequestParams) (*models.FmgCmdbResponse, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.Data = payload
	p.URL = createMetafieldDynamicMappingPath(params.Adom, variable)
	req.Params = append(req.Params, *p)
	req.Method = "add"

	res, err := fmgrequest.Create(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ReadMetafieldDynamicMapping(variable, mkey string, params *models.CmdbRequestParams) (*models.DynamicMapping, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createMetafieldDynamicMappingPath(params.Adom, variable) + mkey
	req.Params = append(req.Params, *p)
	req.Method = "get"

	res, err := fmgrequest.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results
	if tmp, ok := res.Result[0].Data.(interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := models.DynamicMapping{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateMetafieldDynamicMapping(variable, mkey string, payload *models.DynamicMapping, params *models.CmdbRequestParams) (*models.FmgCmdbResponse, error) {
	_, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.Data = payload
	p.URL = createMetafieldDynamicMappingPath(params.Adom, variable) + mkey
	req.Params = append(req.Params, *p)
	req.Method = "update"

	res, err := fmgrequest.Update(c.config, req)
	if err != nil {
		return nil, err
	}

	if res.Session != nil && params.AllowAppend != nil {
		if *res.Session == "notexist" && *params.AllowAppend {
			return c.CreateMetafieldDynamicMapping(variable, payload, params)
		}
	}

	return res, nil
}

func (c *Client) DeleteMetafieldDynamicMapping(variable, mkey string, params *models.CmdbRequestParams) error {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createMetafieldDynamicMappingPath(params.Adom, variable) + mkey
	req.Params = append(req.Params, *p)
	req.Method = "delete"

	err := fmgrequest.Delete(c.config, req)
	return err
}
