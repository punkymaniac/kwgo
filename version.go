package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
)


type Version struct {
    MajorVersion string `jon:"majorVersion"`
    MinorVersion string `json:"minorVersion"`
}

// Retrive Klocwork server version
func (c *KwClient) Version(
) (*Version, *http.Response, error) {
    body, res, err := c.apiRequest("version", nil)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        result := Version{}
        err := json.Unmarshal(data[0], &result)
        if err != nil {
            return nil, nil, err
        }
        return &result, res, nil
    }
    var kwErr kwError
    err = json.Unmarshal(body, &kwErr)
    if err != nil {
        return nil, nil, err
    }
    return nil, res, &kwErr
}

