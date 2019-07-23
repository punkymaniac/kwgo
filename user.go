package kwgo

import (
    "net/http"
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
) ([]User, *http.Response, error) {
    postData := ""
    if search != nil {
        postData += "&search=" + *search
    }
    if limit != nil {
        postData += "&limit=" + strconv.FormatUint(*limit, 10)
    }
    body, res, err := c.apiRequest("users", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := User{}
        result := []User{}
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

// Create a user
func (c *KwClient) CreateUser(
    name string, // The name of the user to create
    password *string, // (optional) assign a password to the new user
) (*http.Response, error) {
    postData := "&name=" + name
    if password != nil {
        postData += "&password=" + *password
    }
    body, res, err := c.apiRequest("create_user", &postData)
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

// Delete a user
func (c *KwClient) DeleteUser(
    name string, // The name of the user to delete
) (*http.Response, error) {
    postData := "&name=" + name
    body, res, err := c.apiRequest("delete_user", &postData)
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

