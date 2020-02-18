package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
    "strconv"
)


type Module struct {
    Name string `json:"name"`
    AllowAll bool `json:"allowAll"`
    Paths []string `json:"paths"`
    Tags []string `json:"tags"`
}

// Retrive list of modules for a project
func (c *KwClient) Modules(
    project string, // Project name
) ([]Module, *http.Response, error) {
    postData := "&project=" + project
    body, res, err := c.apiRequest("modules", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Module{}
        result := []Module{}
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

// Create a module for a project
func (c *KwClient) CreateModule(
    project string, // Project name
    name string, // Module name
    allowAll *bool, // (optional) Module access
    allowUsers *string, // (optional) Grant access to users
    allowGroups *string, // (optional) Grant access to groups
    denyUsers *string, // (optional) Deny access to users
    denyGroups *string, // (optional) Deny access to groups
    paths string, // List of comma separated path regexps
    tags *string, // (optional) List of comma separated tags
) (*http.Response, error) {
    postData := "&project=" + project
    postData += "&name=" + name
    if allowAll != nil {
        postData += "&allowAll=" + strconv.FormatBool(*allowAll)
    }
    if allowUsers != nil {
        postData += "&allow_users=" + *allowUsers
    }
    if allowGroups != nil {
        postData += "&allow_groups=" + *allowGroups
    }
    if denyUsers != nil {
        postData += "&deny_users=" + *denyUsers
    }
    if denyGroups != nil {
        postData += "&deny_groups=" + *denyGroups
    }
    postData += "&paths=" + paths
    if tags != nil {
        postData += "&tags=" + *tags
    }
    body, res, err := c.apiRequest("create_module", &postData)
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

// Delete a module
func (c *KwClient) DeleteModule(
    project string, // Project name
    name string, // View name
) (*http.Response, error) {
    postData := "&project=" + project
    postData += "&name=" + name
    body, res, err := c.apiRequest("delete_module", &postData)
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

// Update a module
func (c *KwClient) UpdateModule(
    project string, // Project name
    name string, // Module name
    newName *string, // (optional) New module name
    allowAll *bool, // (optional) Module access
    allowUsers *string, // (optional) Grant access to users
    allowGroups *string, // (optional) Grant access to groups
    denyUsers *string, // (optional) Deny access to users
    denyGroups *string, // (optional) Deny access to groups
    paths *string, // (optional) List of comma separated path regexps
    tags *string, // (optional) List of comma separated tags
) (*http.Response, error) {
    postData := "&project=" + project
    postData += "&name=" + name
    if newName != nil {
       postData += "&new_name=" + *newName
    }
    if allowAll != nil {
        postData += "&allowAll=" + strconv.FormatBool(*allowAll)
    }
    if allowUsers != nil {
        postData += "&allow_users=" + *allowUsers
    }
    if allowGroups != nil {
        postData += "&allow_groups=" + *allowGroups
    }
    if denyUsers != nil {
        postData += "&deny_users=" + *denyUsers
    }
    if denyGroups != nil {
        postData += "&deny_groups=" + *denyGroups
    }
    if paths != nil {
       postData += "&paths=" + *paths
    }
    if tags != nil {
        postData += "&tags=" + *tags
    }
    body, res, err := c.apiRequest("update_module", &postData)
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

