package cmdb

import (
	"encoding/json"
	"fmt"

	"github.com/poroping/fortimanager-devicedb-sdk-go/fmgrequest"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func createMetafieldPath(adom string) string {
	if adom == "" {
		adom = "root"
	}
	return fmt.Sprintf("pm/config/adom/%s/obj/fmg/variable/", adom)
}

func (c *Client) CreateMetafields(payload *models.Metafields, params *models.CmdbRequestParams) (*models.FmgCmdbResponse, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.Data = payload
	p.URL = createMetafieldPath(params.Vdom)
	req.Params = append(req.Params, *p)
	req.Method = "add"

	res, err := fmgrequest.Create(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) ReadMetafields(mkey string, params *models.CmdbRequestParams) (*models.Metafields, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createMetafieldPath(params.Vdom) + mkey
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
		v := models.Metafields{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateMetafields(mkey string, payload *models.Metafields, params *models.CmdbRequestParams) (*models.FmgCmdbResponse, error) {
	_, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.Data = payload
	p.URL = createMetafieldPath(params.Vdom) + mkey
	req.Params = append(req.Params, *p)
	req.Method = "update"

	res, err := fmgrequest.Update(c.config, req)
	if err != nil {
		return nil, err
	}

	if *res.Session == "notexist" && params.AllowAppend != nil {
		if *params.AllowAppend {
			return c.CreateMetafields(payload, params)
		}
	}

	return res, nil
}

func (c *Client) DeleteMetafields(mkey string, params *models.CmdbRequestParams) error {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createMetafieldPath(params.Vdom) + mkey
	req.Params = append(req.Params, *p)
	req.Method = "delete"

	err := fmgrequest.Delete(c.config, req)
	return err
}
