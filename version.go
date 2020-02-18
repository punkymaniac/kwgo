package kwgo

import (
    "bytes"
    "encoding/json"
)


type Version struct {
    MajorVersion string `jon:"majorVersion"`
    MinorVersion string `json:"minorVersion"`
}

// Retrive Klocwork server version
func (c *KwClient) Version(
) (*Version, error) {
    body, res, err := c.apiRequest("version", nil)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        result := Version{}
        err := json.Unmarshal(data[0], &result)
        if err != nil {
            return nil, err
        }
        return &result, nil
    }
    var kwErr kwError
    err = json.Unmarshal(body, &kwErr)
    if err != nil {
        return nil, err
    }
    return nil, &kwErr
}

