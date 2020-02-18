
// Package kwgo implements access to klocwork api.
package kwgo

import (
    "fmt"
    "net/http"
    "time"
    "io/ioutil"
    "bytes"
)


// Klocwork client to interact with server
type KwClient struct {
    client *http.Client
    serverUrl string
    user string
    ltoken string
}

// Body data returned by the api on error
type kwError struct {
    Status uint `json:"status"`
    Message string `json:"message"`
}

// Implement error interface
func (e *kwError) Error() string {
    return fmt.Sprintf("Status %d: %s", e.Status, e.Message)
}

// Return a KwClient
func NewKwClient(
    url string, // Url of the klocwork server
    user string, // user used to login on the api
    ltoken string, // ltoken of the user
) (KwClient) {
    newClient := KwClient{
                            client: &http.Client{ Timeout: 10 * time.Second },
                            serverUrl: url + "/review/api",
                            user: user,
                            ltoken: ltoken,
                         }
    return newClient
}

// Made request to the klocwork server
func (c *KwClient) apiRequest(
    action string, // Action parameter
    data *string, // Data parameter
) ([]byte, *http.Response, error) {
    var postData string
    if data != nil {
        postData = "user=" + c.user + "&ltoken=" + c.ltoken + "&action=" + action + *data
    } else {
        postData = "user=" + c.user + "&ltoken=" + c.ltoken + "&action=" + action
    }
    req, err := http.NewRequest("POST", c.serverUrl, bytes.NewBuffer([]byte(postData)))
    if err != nil {
        return nil, nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    res, err := c.client.Do(req)
    if err != nil {
        return nil, nil, err
    }
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, nil, err
    }
    return body, res, nil
}

