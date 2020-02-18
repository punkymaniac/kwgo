package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
)


type Taxonomie struct {
    Name string `json:"name"`
    IsCustom bool `json:"is_custom"`
}

// Retrive the list of taxonomy terms for a project
func (c *KwClient) Taxonomies(
	project string, // Project name
) ([]Taxonomie, *http.Response, error) {
    postData := "&project=" + project
    body, res, err := c.apiRequest("taxonomies", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Taxonomie{}
        result := []Taxonomie{}
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

