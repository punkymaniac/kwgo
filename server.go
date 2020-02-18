package kwgo

import (
    "encoding/json"
)


// Import server configuration from another Klocwork server
func (c *KwClient) ImportServerConf(
    sourceUrl string, // Url to source Klocwork server
    sourceAdmin string, // Projects_root administrator account name
    sourcePassword string, // Projects_root administrator account password
) (error) {
    postData := "&sourceURL=" + sourceUrl
    postData += "&sourceAdmin=" + sourceAdmin
    postData += "&sourcePassword=" + sourcePassword
    body, res, err := c.apiRequest("import_server_configuration", &postData)
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

