package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type Churn struct {
    baseReport
}

// Generate file churns report
func (c *KwClient) Fchurns(
    project string, // Name of the project you want to create a report for
    view *string, // (optional) The view you want to set for the report
    viewCreator *string, // (optional) View creator name
    latestsBuilds *uint64, // (optional) The number of builds you want to show in the report
    component *string, // (optional) Root component
) (*Churn, error) {
    postData := "&project=" + project
    if view != nil {
        postData += "&view=" + *view
    }
    if viewCreator != nil {
        postData += "&viewCreator=" + *viewCreator
    }
    if latestsBuilds != nil {
        postData += "&latestsBuilds=" + strconv.FormatUint(*latestsBuilds, 10)
    }
    if component != nil {
        postData += "&component=" + *component
    }
    body, res, err := c.apiRequest("fchurns", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        result := Churn{}
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

