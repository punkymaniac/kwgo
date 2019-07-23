package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
    "strconv"
    "time"
    "strings"
)


type Project struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Creator string `json:"creator"`
    Description string `json:"description"`
    Tags []string `json:"tags"`
}

type ProjectConf struct {
    Build string `json:"build"`
    CreationDate kwTimeFormat `json:"creationDate"`
    Version string `json:"version"`
    NumberOfFiles string `json:"numberOfFiles"`
    CFilesAnalyzed string `json:"cFilesAnalyzed"`
    SystemFilesAnalyzed string `json:"systemFilesAnalyzed"`
    LinesOfCode string `json:"linesOfCode"`
    LinesOfComments string `json:"linesOfComments"`
    NumberOfEntities string `json:"numberOfEntities"`
    NumberOfFunctions string `json:"numberOfFunctions"`
    NumberOfClasses string `json:"numberOfClasses"`
    Taxonomies string `json:"taxonomies"`
}

type kwTimeFormat struct {
    time.Time
}

func (sd *kwTimeFormat) UnmarshalJSON(
    input []byte,
) (error) {
    strInput := string(input)
    strInput = strings.Trim(strInput, `"`)
    newTime, err := time.Parse("Mon Jan 2 15:04:05 CEST 2006", strInput)
    if err != nil {
        return err
    }
    sd.Time = newTime
    return nil
}


// Retrive list of projects
func (c *KwClient) Projects(
) ([]Project, *http.Response, error) {
    body, res, err := c.apiRequest("projects", nil)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Project{}
        result := []Project{}
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

// Delete a project
func (c *KwClient) DeleteProject(
    name string, // Project name
) (*http.Response, error) {
    postData := "&name=" + name
    body, res, err := c.apiRequest("delete_project", &postData)
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

// Import project from another Klocwork server
func (c *KwClient) ImportProject(
    project string, // Project name
    sourceUrl string, // Url to source Klocwork server
    sourceAdmin string, // Projects_root administrator account name
    sourcePassword string, // Projects_root administrator account password
) (*http.Response, error) {
    postData := "&project=" + project
    postData += "&sourceURL=" + sourceUrl
    postData += "&sourceAdmin=" + sourceAdmin
    postData += "&sourcePassword=" + sourcePassword
    body, res, err := c.apiRequest("import_project", &postData)
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

// Import project from another Klocwork server
func (c *KwClient) UpdateProject(
    name string, // Project name
    newName *string, // (optional) New project name
    description *string, // (optional) New project description
    tags *string, // (optional) List of comma separated tags
    autoDeleteBuilds *bool, // (optional) Whether the builds in the project should automatically be deleted
    autoDeleteThreshold *bool, // (optional) The number of builds to keep in the project if auto_delete_builds it true (default: 20)
) (*http.Response, error) {
    postData := "&name=" + name
    if newName != nil {
        postData += "&new_name=" + *newName
    }
    if description != nil {
        postData += "&description=" + *description
    }
    if tags != nil {
        postData += "&tags=" + *tags
    }
    if autoDeleteBuilds != nil {
        postData += "&auto_delete_builds=" + strconv.FormatBool(*autoDeleteBuilds)
    }
    if autoDeleteThreshold != nil {
        postData += "&auto_delete_threshold=" + strconv.FormatBool(*autoDeleteThreshold)
    }
    body, res, err := c.apiRequest("update_project", &postData)
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

// Generate project configuration report
func (c *KwClient) ProjectConfiguration(
    project string, // Name of the project you want to create a report for
    build *string, // (optional) Name fo the build you want to focus your report on
) (*ProjectConf, *http.Response, error) {
    postData := "&project=" + project
    if build != nil {
        postData += "&build=" + *build
    }
    body, res, err := c.apiRequest("project_configuration", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        result := ProjectConf{}
        err := json.Unmarshal(data[0], &result)
        if err != nil {
            return nil, nil, err
        }
        return &result, res, nil
    }
    err = json.Unmarshal(body, &c.KwErr)
    if err != nil {
        return nil, nil, err
    }
    return nil, res, nil
}

