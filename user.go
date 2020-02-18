package kwgo

import (
    "bytes"
    "encoding/json"
    "strconv"
)


type User struct {
    Name string `json:"name"`
    Readonly bool `json:"readonly"`
    Roles []aRole `json:"roles"`
    Groups []refGroup `json:"groups"`
}

// List users
func (c *KwClient) Users(
    search *string, // (optional) Simple username search pattern
    limit *uint64, // (optional) Maximum number of result to return
) ([]User, error) {
    postData := ""
    if search != nil {
        postData += "&search=" + *search
    }
    if limit != nil {
        postData += "&limit=" + strconv.FormatUint(*limit, 10)
    }
    body, res, err := c.apiRequest("users", &postData)
    if err != nil {
        return nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := User{}
        result := []User{}
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

// Create a user
func (c *KwClient) CreateUser(
    name string, // The name of the user to create
    password *string, // (optional) assign a password to the new user
) (error) {
    postData := "&name=" + name
    if password != nil {
        postData += "&password=" + *password
    }
    body, res, err := c.apiRequest("create_user", &postData)
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

// Delete a user
func (c *KwClient) DeleteUser(
    name string, // The name of the user to delete
) (error) {
    postData := "&name=" + name
    body, res, err := c.apiRequest("delete_user", &postData)
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

