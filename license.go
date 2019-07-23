package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
)


type License struct {
    Feature string `json:"feature"`
    Version string `json:"version"`
    Count uint64 `json:"count"`
    Permanent bool `json:"permanent"`
    DaysToExpiry int64 `json:"days_to_expiry"`
    Servers []string `json:"servers"`
}

// Retrive the number of licenses issued for a feature
func (c *KwClient) License(
) ([]License, *http.Response, error) {
    body, res, err := c.apiRequest("license_count", nil)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := License{}
        result := []License{}
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

