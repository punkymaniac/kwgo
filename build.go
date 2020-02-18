package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type Build struct {
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Date uint64 `json:"date"`
    Keepit bool `json:"keepit"`
}

// Retrive the list of builds for a project
func (c *KwClient) Builds(
    project string, // Project name
) ([]Build, error) {
    postData := "&project=" + project
    body, res, err := c.apiRequest("builds", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Build{}
        result := []Build{}
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

// Delete a build
func (c *KwClient) DeleteBuild(
    project string, // Project name
    name string, // Build name
) (error) {
    postData := "&project=" + project
    postData += "&name=" + name
    body, res, err := c.apiRequest("delete_build", &postData)
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

// Update a build
func (c *KwClient) UpdateBuild(
    project string, // Project name
    name string, // Build name
    newName *string, // (optional) New build name
    keepit *bool, // (optional) Whether this build will be deleted by the auto-delete build feature
) (error) {
    postData := "&project=" + project
    postData += "&name=" + name
    if newName != nil {
        postData += "&new_name=" + *newName
    }
    if keepit != nil {
        postData += "&keepit=" + strconv.FormatBool(*keepit)
    }
    body, res, err := c.apiRequest("update_build", &postData)
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

