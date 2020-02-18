package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type Group struct {
    Name string `json:"name"`
    Readonly bool `json:"readonly"`
    Roles []aRole `json:"roles"`
    Groups []refGroup `json:"groups"`
    Users []string `json:"users"`
}

type refGroup struct {
    Name string `json:"name"`
    Readonly bool `json:"readonly"`

}

// List groups
func (c *KwClient) Groups(
    search *string, // (optional) Simple group name search pattern
    listUsers *bool, // (optional) Output user list for each group
    limit *uint64, // (optional) Maximum number of result to return
) ([]Group, error) {
    postData := ""
    if search != nil {
        postData += "&search=" + *search
    }
    if listUsers != nil {
        postData += "&list_users=" + strconv.FormatBool(*listUsers)
    }
    if limit != nil {
        postData += "&limit=" + strconv.FormatUint(*limit, 10)
    }
    body, res, err := c.apiRequest("groups", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Group{}
        result := []Group{}
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

// Create a group and optionally assign users to it
func (c *KwClient) CreateGroup(
    name string, // The name of the group to create
    users *string, // (optional) Comma separated list of users
) (error) {
    postData := "&name=" + name
    if users != nil {
        postData += "&users=" + *users
    }
    body, res, err := c.apiRequest("create_group", &postData)
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

// Delete a group
func (c *KwClient) DeleteGroup(
    name string, // The name of the group to delete
) (error) {
    postData := "&name=" + name
    body, res, err := c.apiRequest("delete_group", &postData)
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

// Update users in a group
func (c *KwClient) UpdateGroup(
    name string, // The name of the group to update
    users *string, // (optional) Comma separated list of users
    removeAll *bool, // (optional) If 'true', the group's user list will be cleared, Ignored if 'users' is not nil
) (error) {
    postData := "&name=" + name
    if users != nil {
        postData += "&users=" + *users
    } else if removeAll != nil {
        postData += "&remove_all=" + strconv.FormatBool(*removeAll)
    }
    body, res, err := c.apiRequest("update_group", &postData)
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

