package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type Metric struct {
    FilePath string `json:"filepath"`
    Entity string `json:"entity"`
    EntityId uint64 `json:"entity_id"`
    Tag string `json:"tag"`
    MetricValue float64 `json:"metricValue"`
}

type MetricStat struct {
    Tag string `json:"tag"`
    Sum float64 `json:"sum"`
    Min float64 `json:"min"`
    Max float64 `json:"max"`
    Entries uint64 `json:"entires"`
}

// Retrive the list of metrics
func (c *KwClient) Metrics(
    project string, // Project name
    query *string, // (optional) Search query, such as narrowing by file
    view *string, // (optional) View name
    limit *uint64, // (optional) Search result limit
) ([]Metric, error) {
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
    body, res, err := c.apiRequest("metrics", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Metric{}
        result := []Metric{}
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

// Retrive the statistic of metrics
func (c *KwClient) MetricStat(
    project string, // Project name
    query *string, // (optional) Search query, such as narrowing by file
    view *string, // (optional) View name
    limit *uint64, // (optional) Search result limit
) ([]MetricStat, error) {
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
    postData += "&aggregate=true"
    body, res, err := c.apiRequest("metrics", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := MetricStat{}
        result := []MetricStat{}
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

