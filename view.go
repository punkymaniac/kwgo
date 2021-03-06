package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type View struct {
    Id uint64 `json:"id"`
    Name string `json:"name"`
    Query string `json:"query"`
    Creator string `json:"creator"`
    IsPublic bool `json:"is_public"`
    Tags []string `json:"tags"`
}

// Retrive list of views
func (c *KwClient) Views(
    project string, // Project name
) ([]View, error) {
    postData := "&project=" + project
    body, res, err := c.apiRequest("views", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := View{}
        result := []View{}
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

// Create a view for a project
func (c *KwClient) CreateView(
    project string, // Project name
    name string, // View name
    query string, // Search query for the view
    tags *string, // (optional) List of comma separated tags
    isPublic *bool, // (optional) Whether the views is visible to all users with access to this project
) (error) {
    postData := "&project=" + project
    postData += "&name=" + name
    postData += "&query=" + query
    if tags != nil {
        postData += "&tags=" + *tags
    }
    if isPublic != nil {
        postData += "&is_public=" + strconv.FormatBool(*isPublic)
    }
    body, res, err := c.apiRequest("create_view", &postData)
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

// Delete a view
func (c *KwClient) DeleteView(
    project string, // Project name
    name string, // View name
) (error) {
    postData := "&project=" + project
    postData += "&name=" + name
    body, res, err := c.apiRequest("delete_view", &postData)
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

// Update a view
func (c *KwClient) UpdateView(
    project string, // Project name
    name string, // Current view name
    newName *string, // (optional) New view name
    query *string, // (optional) Search query for the view
    tags *string, // (optional) List of comma separated tags
    isPublic *bool, // (optional) Whether the views is visible to all users with access to this project
) (error) {
    postData := "&project=" + project
    postData += "&name=" + name
    if newName != nil {
       postData += "&new_name=" + *newName
    }
    if query != nil {
       postData += "&query=" + *query
    }
    if tags != nil {
        postData += "&tags=" + *tags
    }
    if isPublic != nil {
        postData += "&is_public=" + strconv.FormatBool(*isPublic)
    }
    body, res, err := c.apiRequest("update_view", &postData)
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

