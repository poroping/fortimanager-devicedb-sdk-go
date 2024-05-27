package cmdb

import (
	"encoding/json"
	"fmt"

	"github.com/poroping/fortimanager-devicedb-sdk-go/fmgrequest"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func createInstallScriptPath() string {
	return fmt.Sprintf("/dvmdb/script/execute")
}

func (c *Client) CreateInstallScript(payload *models.InstallScriptPayload, params *models.CmdbRequestParams) (*models.InstallScriptResponse, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.Data = payload
	p.URL = createInstallScriptPath()
	req.Params = append(req.Params, *p)
	req.Method = "exec"

	res, err := fmgrequest.Read2(c.config, req)
	if err != nil {
		return nil, err
	}

	// marshal/unmarshal results
	if tmp, ok := res.Result[0].Data.(interface{}); ok {
		jsontmp, err := json.Marshal(tmp)
		if err != nil {
			return nil, err
		}
		v := models.InstallScriptResponse{}
		json.Unmarshal(jsontmp, &v)
		if v.Task != nil {
			taskid := *v.Task
			_, err = c.ReadTaskUntilCompletion(taskid, params)
			if err != nil {
				return nil, err
			}
		}
		return &v, nil
	}

	return nil, err
}
