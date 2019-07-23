package kwgo

import (
    "net/http"
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
) ([]Build, *http.Response, error) {
    postData := "&project=" + project
    body, res, err := c.apiRequest("builds", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Build{}
        result := []Build{}
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

// Delete a build
func (c *KwClient) DeleteBuild(
    project string, // Project name
    name string, // Build name
) (*http.Response, error) {
    postData := "&project=" + project
    postData += "&name=" + name
    body, res, err := c.apiRequest("delete_build", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        return res, nil
    }
    err = json.Unmarshal(body, &c.KwErr)
    if err != nil {
        return nil, err
    }
    return res, nil
}

// Update a build
func (c *KwClient) UpdateBuild(
    project string, // Project name
    name string, // Build name
    newName *string, // (optional) New build name
    keepit *bool, // (optional) Whether this build will be deleted by the auto-delete build feature
) (*http.Response, error) {
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
        return nil, err
    }
    if res.StatusCode == 200 {
        return res, nil
    }
    err = json.Unmarshal(body, &c.KwErr)
    if err != nil {
        return nil, err
    }
    return res, err
}

