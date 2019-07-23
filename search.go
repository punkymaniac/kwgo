package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
    "strconv"
)


type Issue struct {
    Id uint64 `json:"id"`
    Status string `json:"status"`
    Severity string `json:"severity"`
    State string `json:"state"`
    Owner string `json:"owner"`
    SeverityCode uint64 `json:"severityCode"`
    Code string `json:"code"`
    Title string `json:"title"`
    Message string `json:"message"`
    File string `json:"file"`
    Method string `json:"method"`
    TaxonomyName string `json:"taxonomyName"`
    DateOriginated uint64 `json:"dateOriginated"`
    Url string `json:"url"`
    IssueIds []uint64 `json:"issueIds"`
}

// Retrive the list of detected issues
func (c *KwClient) Search(
    project string, // Project name
    query *string, // (optional) Search query, such as narrowing by file
    view *string, // (optional) View name
    limit *uint64, // (optional) Search result limit
    summary *string, // (optional) Include summary record to output stream
) ([]Issue, *http.Response, error) {
    postData := "&project=" + project
    if query != nil {
        postData += "&query=" + *query
    }
    if view != nil {
        postData += "&view=" + *view
    }
    if limit != nil {
        postData += "&limit=" + strconv.FormatUint(*limit, 10)
    }
    if summary != nil {
        postData += "&summary=" + *summary
    }
    body, res, err := c.apiRequest("search", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Issue{}
        result := []Issue{}
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

