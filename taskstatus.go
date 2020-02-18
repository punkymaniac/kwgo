package kwgo

import (
    "bytes"
    "encoding/json"
)


type TaskStatus struct {
    // Missing information to complete this struct
}

// List statuses of all tasks running on Klocwork server
func (c *KwClient) TaskStatus(
) ([]TaskStatus, error) {
    body, res, err := c.apiRequest("task_status", nil)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := TaskStatus{}
        result := []TaskStatus{}
        for _, elem := range data {
            err := json.Unmarshal(elem, &target)
            if err != nil {
                return nil, err
            }
            result = append(result, target)
        }
        return result, nil
    }
    var kwErr kwError
    err = json.Unmarshal(body, &kwErr)
    if err != nil {
        return nil, err
    }
    return nil, &kwErr
}

