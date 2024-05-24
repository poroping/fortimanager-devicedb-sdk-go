package cmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/poroping/fortimanager-devicedb-sdk-go/fmgrequest"
	"github.com/poroping/fortimanager-devicedb-sdk-go/models"
)

func createTaskPath(mkey int64) string {
	return fmt.Sprintf("/task/task/%d", mkey)
}

func (c *Client) ReadTask(mkey int64, params *models.CmdbRequestParams) (*models.TaskResultData, error) {
	req := &models.FmgCmdbRequest{}
	p := &models.FmgCmdbRequestPayload{}
	p.URL = createTaskPath(mkey)
	req.Params = append(req.Params, *p)
	req.Method = "get"

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
		v := models.TaskResultData{}
		json.Unmarshal(jsontmp, &v)
		return &v, nil
	}
	return nil, err
}

func isTaskCompleted(obj *models.TaskResultData) bool {
	if obj.Percent != nil {
		if *obj.Percent == 100 {
			log.Printf("[DEBUG] Task completed")
			return true
		}
	}

	return false
}

func (c *Client) ReadTaskUntilCompletion(mkey int64, params *models.CmdbRequestParams) (*models.TaskResultData, error) {
	complete := false
	retry := 0

	res, err := c.ReadTask(mkey, params)
	if err != nil {
		return nil, err
	}
	complete = isTaskCompleted(res)

	if retry == 120 {
		return nil, errors.New("task seems to be hanging, check Task Log on FortiManager")
	}

	if !complete {
		time.Sleep(5 * time.Second)
		log.Printf("[DEBUG] Task not complete, reading again")
		retry += 1
		c.ReadTaskUntilCompletion(mkey, params)
	}

	if res.NumErr != nil {
		if *res.NumErr != 0 {
			errs := []string{}
			for _, line := range res.Line {
				errs = append(errs, *line.Detail)
			}
			err = errors.New(strings.Join(errs, "\n"))
			return nil, err
		}
	}

	return res, err
}
