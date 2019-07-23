package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
)


type Status struct {
    // Missing information to complete this struct
}

// List current import status of projects
func (c *KwClient) ImportStatus(
) ([]Status, *http.Response, error) {
    body, res, err := c.apiRequest("import_status", nil)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Status{}
        result := []Status{}
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

