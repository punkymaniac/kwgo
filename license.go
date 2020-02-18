package kwgo

import (
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
) ([]License, error) {
    body, res, err := c.apiRequest("license_count", nil)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := License{}
        result := []License{}
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

