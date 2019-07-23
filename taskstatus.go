package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
)


type TaskStatus struct {
    // Missing information to complete this struct
}

// List statuses of all tasks running on Klocwork server
func (c *KwClient) TaskStatus(
) ([]TaskStatus, *http.Response, error) {
    body, res, err := c.apiRequest("task_status", nil)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := TaskStatus{}
        result := []TaskStatus{}
        for _, elem := range data {
            err := json.Unmarshal(elem, &target)
            if err != nil {
                return nil, nil, err
            }
            result = append(result, target)
        }
        return result, res, nil
    }
    err = json.Unmarshal(body, &c.KwErr)
    if err != nil {
        return nil, nil, err
    }
    return nil, res, nil
}

