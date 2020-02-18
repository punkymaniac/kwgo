package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
    "strconv"
)


type DefectType struct {
    Code string `json:"code"`
    Name string `json:"name"`
    Enabled bool `json:"enabled"`
    Severity uint64 `json:"severity"`
}

// Retrive the list of enabled defect types
func (c *KwClient) DefectTypes(
    project string, // Project name
    taxonomy *string, // (optional) Filter by taxonomy
) ([]DefectType, *http.Response, error) {
    postData := "&project=" + project
    if taxonomy != nil {
        postData += "&taxonomy=" + *taxonomy
    }
    body, res, err := c.apiRequest("defect_types", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := DefectType{}
        result := []DefectType{}
        for _, elem := range data {
            err := json.Unmarshal(elem, &target)
            if err != nil {
                return nil, nil, err
            }
            result = append(result, target)
        }
        return result, res, nil
    }
    var kwErr kwError
    err = json.Unmarshal(body, &kwErr)
    if err != nil {
        return nil, nil, err
    }
    return nil, res, &kwErr
}

// Enable or disable a defect
func (c *KwClient) UpdateDefectType(
    project string, // Project name
    code string, // Defect code
    enabled *bool, // (optional) true to enable, false to disable
    severity *uint64, // (optional) Specify new defect severity
) (*http.Response, error) {
    postData := "&project=" + project
    postData += "&code=" + code
    if enabled != nil {
        postData += "&enabled=" + strconv.FormatBool(*enabled)
    }
    if severity != nil {
        postData += "&severity=" + strconv.FormatUint(*severity, 10)
    }
    body, res, err := c.apiRequest("update_defect_type", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        return res, nil
    }
    var kwErr kwError
    err = json.Unmarshal(body, &kwErr)
    if err != nil {
        return nil, err
    }
    return res, &kwErr
}

