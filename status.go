package kwgo

import (
    "strconv"
    "strings"
    "bytes"
    "encoding/json"
)


type Status struct {
    // Missing information to complete this struct
}

// List current import status of projects
func (c *KwClient) ImportStatus(
) ([]Status, error) {
    body, res, err := c.apiRequest("import_status", nil)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Status{}
        result := []Status{}
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

// Change the status, owner, and comment, or alternatively set the bug tracker id of issues
func (c *KwClient) UpdateStatus(
    project string, // Project name
    ids []uint64, // List of ids to change
    status *string, // (optional) New status to set
    comment *string, // (optional) New comment to set
    owner *string, // (optional) New owner to set
    bugTrackerId *uint64, // (optional) New bug tracker id to set
) (error) {
    postData := "&project=" + project
    idsText := []string{}
    for _, value := range ids {
        text := strconv.FormatUint(value, 10)
        idsText = append(idsText, text)
    }
    postData += "&ids=" + strings.Join(idsText, ",")
    if status != nil {
        postData += "&status=" + *status
    }
    if comment != nil {
        postData += "&comment=" + *comment
    }
    if owner != nil {
        postData += "&owner=" + *owner
    }
    if bugTrackerId != nil {
        postData += "&bug_tracker_id=" + strconv.FormatUint(*bugTrackerId, 10)
    }
    body, res, err := c.apiRequest("update_status", &postData)
    if err != nil {
        return err
    }
    if res.StatusCode == 200 {
        return nil
    }
    var kwErr kwError
    err = json.Unmarshal(body, &kwErr)
    if err != nil {
        return err
    }
    return &kwErr
}

