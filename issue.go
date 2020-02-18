package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
    "strconv"
)


type IssueDetail struct {
    Id string `json:"id"`
    Status string `json:"status"`
    Severity string `json:"severity"`
    State string `json:"state"`
    Owner string `json:"owner"`
    Code string `json:"code"`
    Name string `json:"name"`
    Location string `json:"location"`
    Build string `json:"build`
    History []HistElem `json:"history"`
    Xsync string `json:"xsync"`
}

type HistElem struct {
    Date uint64 `json:"date"`
    UserId string `json:"userid"`
    Status string `json:"status"`
    Comment string `json:"comment"`
}

// Get details got the given issue id
func (c *KwClient) IssueDetails(
    project string, // Name of the project you want to search
    id uint64, // The id to search
    includeXsync *bool, // (optional) Boolean to return xSyncInfo
) (*IssueDetail, *http.Response, error) {
    postData := "&project=" + project
    postData += "&id=" + strconv.FormatUint(id, 10)
    if includeXsync != nil {
        postData += "&include_xsync=" + strconv.FormatBool(*includeXsync)
    }
    body, res, err := c.apiRequest("issue_details", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        result := IssueDetail{}
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

