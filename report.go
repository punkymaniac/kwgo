package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type baseReport struct {
    Rows []reportElem `json:"rows"`
    Columns []reportElem `json:"columns"`
    Data [][]int64 `json:"data"`
}

type ReportData struct {
    baseReport
    Warnings []string `json:"warnings"`
}

type reportElem struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
}

// Generate build summary report
func (c *KwClient) Report(
    project string, // Name of the project you want to create a report for
    build *string, // (optional) Name of the build you want to focus your report on
    filterQuery *string, // (optional) The filter query you want to set got the report
    view *string, // (optional) The view you want to set for the report
    x *string, // (optional) The value you want to set along the x-axis
    xDrilldown *string, // (optional) Row drill-down item id
    y *string, // (optional) The variable you want to set along the y-axis
    yDrilldown *string, // (optional) Column drill-down item id
    groupIssues *bool, // (optional) Show grouped issues
) (*ReportData, error) {
    postData := "&project=" + project
    if build != nil {
        postData += "&build=" + *build
    }
    if filterQuery != nil {
        postData += "&filterQuery=" + *filterQuery
    }
    if view != nil {
        postData += "&view=" + *view
    }
    if x != nil {
        postData += "&x=" + *x
    }
    if xDrilldown != nil {
        postData += "&xDrilldown=" + *xDrilldown
    }
    if y != nil {
        postData += "&y=" + *y
    }
    if yDrilldown != nil {
        postData += "&yDrilldown=" + *yDrilldown
    }
    if groupIssues != nil {
        postData += "&group_issues=" + strconv.FormatBool(*groupIssues)
    }
    body, res, err := c.apiRequest("report", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        result := ReportData{}
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

