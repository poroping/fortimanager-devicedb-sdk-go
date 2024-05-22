package cmdb

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/fortimanager-devicedb-sdk-go/fmgrequest"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func createModelDevicePath() string {
	return "/dvm/cmd/add/device"
}

func createReadModelDevicePath() string {
	// if adom == "" {
	// adom = "root"
	// }
	return fmt.Sprintf("dvmdb/device/")
}

func createDeleteModelDevicePath() string {
	return "/dvm/cmd/del/device/"
}

func (c *Client) CreateModelDevice(payload *models.ModelDevice, params *models.CmdbRequestParams) (*models.FmgCmdbResponse, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.Data = payload

	exists, err := ModelDeviceExists(c, *payload.Device.Name, params)
	if err != nil {
		return nil, err
	}

	if exists != nil {
		if *exists {
			log.Printf("[WARN] Device already exists, attempting to update")
			// return c.ReadModelDevice(*payload.Device.Name, params)
			return c.UpdateModelDevice(*payload.Device.Name, payload, params)
		}
	}

	p.URL = createModelDevicePath()
	req.Params = append(req.Params, *p)
	req.Method = "exec"

	res, err := fmgrequest.Create(c.config, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ModelDeviceExists(c *Client, mkey string, params *models.CmdbRequestParams) (*bool, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createReadModelDevicePath() + mkey
	req.Params = append(req.Params, *p)
	req.Method = "get"

	res, err := fmgrequest.Read(c.config, req)
	if err != nil {
		return nil, err
	}

	if res.Exists == nil {
		log.Printf("[WARN] ModelDeviceExist seems to have failed and exists in null")
	}

	return res.Exists, nil

}

func (c *Client) ReadModelDevice(mkey string, params *models.CmdbRequestParams) (*models.Device, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createReadModelDevicePath() + mkey
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
		v := models.Device{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func (c *Client) UpdateModelDevice(mkey string, payload *models.ModelDevice, params *models.CmdbRequestParams) (*models.FmgCmdbResponse, error) {
	_, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	data := &models.Device{}

	// Loads of attributes are read only
	data.Desc = payload.Device.Desc
	data.PreferImgVer = payload.Device.PreferImgVer

	p.Data = data
	p.URL = createReadModelDevicePath() + mkey
	req.Params = append(req.Params, *p)
	req.Method = "set"

	res, err := fmgrequest.Update(c.config, req)
	if err != nil {
		return nil, err
	}

	if res.Exists != nil && params.AllowAppend != nil {
		if !*res.Exists && *params.AllowAppend {
			return c.CreateModelDevice(payload, params)
		}
	}

	return res, nil
}

func (c *Client) DeleteModelDevice(mkey string, params *models.CmdbRequestParams) error {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createDeleteModelDevicePath()
	p.Data = &models.DelDevice{
		Adom:   &params.Adom,
		Device: &mkey,
	}
	req.Params = append(req.Params, *p)
	req.Method = "exec"

	err := fmgrequest.Delete(c.config, req)
	return err
}
