package kwgo

import (
    "net/http"
    "encoding/json"
)


// Import server configuration from another Klocwork server
func (c *KwClient) ImportServerConf(
    sourceUrl string, // Url to source Klocwork server
    sourceAdmin string, // Projects_root administrator account name
    sourcePassword string, // Projects_root administrator account password
) (*http.Response, error) {
    postData := "&sourceURL=" + sourceUrl
    postData += "&sourceAdmin=" + sourceAdmin
    postData += "&sourcePassword=" + sourcePassword
    body, res, err := c.apiRequest("import_server_configuration", &postData)
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

