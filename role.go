package kwgo

import (
    "net/http"
    "bytes"
    "encoding/json"
    "strconv"
)


type Role struct {
    Name string `json:"name"`
    Readonly bool `json:"readonly"`
    Permissions []RolePerm `json:"permissions"`
    StatusPermissions []RoleStatusPerm `json:"statusPermissions"`
}

type RolePerm struct {
    Name string `json:"name"`
    Enabled bool `json:"enabled"`
}

type RoleStatusPerm struct {
    From string `json:"from"`
    To string `json:"to"`
}

type RoleAssignment struct {
    Name string `json:"name"`
    Readonly bool `json:"readonly"`
    Assignments []aRole2 `json:"assignments"`
}

type aRole struct {
    Name string `json:"name"`
    ProjectId string `json:"projectId"`
}

type aRole2 struct {
    aRole
    Group bool `json:"group"`
}

type Permission struct {
    CreateProject bool `json:"create_project"`
    ManageRoles bool `json:"manage_roles"`
    ManageUsers bool `json:"manage_users"`
    AccessSourceFiles bool `json:"access_source_files"`
    AssignRole bool `json:"assign_role"`
    ChangeProjectSettings bool `json:"change_project_settings"`
    CreateBuild bool `json:"create_build"`
    DeleteBuild bool `json:"delete_build"`
    DeleteProject bool `json:"delete_project"`
    ManageModules bool `json:"manage_modules"`
    UseLocalConfiguration bool `json:"use_local_configuration"`
    ChangeIssueStatus bool `json:"change_issue_status"`
    WebApiAccess bool `json:"webapi_access"`
    ExecuteKwxsync bool `json:"execute_kwxsync"`
}

type Transition struct {
    Items []RoleStatusPerm `json:"item"`
}

// Add a status transition rules
func (a *Transition) New(
    from string,
    to string,
) {
    new := RoleStatusPerm{From: from, To: to}
    a.Items = append(a.Items, new)
}

// Clear all status transition rules
func (a *Transition) Clear(
) {
    a.Items = []RoleStatusPerm{}
}

func (a *Transition) postString(
) (string) {
    s := ""
    for _, item := range a.Items {
        s += item.From + "," + item.To + ";"
    }
    return s
}

// Deny all permissions
func (p *Permission) DenyAll(
) {
    p.CreateProject = false
    p.ManageRoles = false
    p.ManageUsers = false
    p.AccessSourceFiles = false
    p.AssignRole = false
    p.ChangeProjectSettings = false
    p.CreateBuild = false
    p.DeleteBuild = false
    p.DeleteProject = false
    p.ManageModules = false
    p.UseLocalConfiguration = false
    p.ChangeIssueStatus = false
    p.WebApiAccess = false
    p.ExecuteKwxsync = false
}

// Allow all permissions
func (p *Permission) AllowAll(
) {
    p.CreateProject = true
    p.ManageRoles = true
    p.ManageUsers = true
    p.AccessSourceFiles = true
    p.AssignRole = true
    p.ChangeProjectSettings = true
    p.CreateBuild = true
    p.DeleteBuild = true
    p.DeleteProject = true
    p.ManageModules = true
    p.UseLocalConfiguration = true
    p.ChangeIssueStatus = true
    p.WebApiAccess = true
    p.ExecuteKwxsync = true
}

func (p *Permission) postString(
) (string) {
    s := ""
    s += "&create_project=" + strconv.FormatBool(p.CreateProject)
    s += "&manage_roles=" + strconv.FormatBool(p.ManageRoles)
    s += "&manage_users=" + strconv.FormatBool(p.ManageUsers)
    s += "&access_source_files=" + strconv.FormatBool(p.AccessSourceFiles)
    s += "&assign_role=" + strconv.FormatBool(p.AssignRole)
    s += "&change_project_settings=" + strconv.FormatBool(p.ChangeProjectSettings)
    s += "&create_build=" + strconv.FormatBool(p.CreateBuild)
    s += "&delete_build=" + strconv.FormatBool(p.DeleteBuild)
    s += "&delete_project=" + strconv.FormatBool(p.DeleteProject)
    s += "&manage_modules=" + strconv.FormatBool(p.ManageModules)
    s += "&use_local_configuration=" + strconv.FormatBool(p.UseLocalConfiguration)
    s += "&change_issue_status=" + strconv.FormatBool(p.ChangeIssueStatus)
    s += "&webapi_access=" + strconv.FormatBool(p.WebApiAccess)
    s += "&execute_kwxsync=" + strconv.FormatBool(p.ExecuteKwxsync)
    return s
}


// List roles
func (c *KwClient) Roles(
    search *string, // (optional) Simple role filter
) ([]Role, *http.Response, error) {
    postData := ""
    if search != nil {
        postData += "&search=" + *search
    }
    body, res, err := c.apiRequest("roles", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := Role{}
        result := []Role{}
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

// List roles assignments
func (c *KwClient) RoleAssignments(
    search *string, // (optional) Simple role filter
) ([]RoleAssignment, *http.Response, error) {
    postData := ""
    if search != nil {
        postData += "&search=" + *search
    }
    body, res, err := c.apiRequest("role_assignments", &postData)
    if err != nil {
        return nil, nil, err
    }
    if res.StatusCode == 200 {
        data := bytes.Split(body, []byte{'\n'})
        data = data[:len(data) - 1]
        target := RoleAssignment{}
        result := []RoleAssignment{}
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

// Create a role
func (c *KwClient) CreateRole(
    name string, // The name of the role to create
    permission *Permission, // (optional) Grant or revoke permission
    allowedStatusTransitions *Transition, // (optional) Set the allowed status transition (ex: Any,Analyze;Analyse,Fix)
) (*http.Response, error) {
    postData := "&name=" + name
    if permission != nil {
        postData += permission.postString()
    }
    if allowedStatusTransitions != nil {
        postData += "&allowed_status_transitions=" + allowedStatusTransitions.postString()
    }
    body, res, err := c.apiRequest("create_role", &postData)
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

// Delete a role
func (c *KwClient) DeleteRole(
    name string, // The name of the role to delete
) (*http.Response, error) {
    postData := "&name=" + name
    body, res, err := c.apiRequest("delete_role", &postData)
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

// Add or remove a user to/from a role
func (c *KwClient) UpdateRoleAssignment(
    name string, // The name of the role to update
    project *string, // (optional) The project id the role should be updated for, if any
    account string, // The name of the account to add or remove
    group *bool, // (optional) Set to 'true' if the account is a group
    remove *bool, // (optional) Set to 'true' to remove the account from the role
) (*http.Response, error) {
    postData := "&name=" + name
    if project != nil {
        postData += "&project=" + *project
    }
    postData += "&account=" + account
    if group != nil {
        postData += "&group=" + strconv.FormatBool(*group)
    }
    if remove != nil {
        postData += "&remove=" + strconv.FormatBool(*remove)
    }
    body, res, err := c.apiRequest("update_role_assignment", &postData)
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

// Update role permissions
func (c *KwClient) UpdateRolePermission(
    name string, // The name of the role to update
    permission *Permission, // (optional) Grant or revoke permission
    allowedStatusTransitions *Transition, // (optional) Set the allowed status transition (ex: Any,Analyze;Analyse,Fix)
) (*http.Response, error) {
    postData := "&name=" + name
    if permission != nil {
        postData += permission.postString()
    }
    if allowedStatusTransitions != nil {
        postData += "&allowed_status_transitions=" + allowedStatusTransitions.postString()
    }
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

//
// Granular role permisssions update
//

// Update create project permissions
func (c *KwClient) RoleCreateProject(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&create_project=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update manage roles permissions
func (c *KwClient) RoleManageRoles(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&manage_roles=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update manage users permissions
func (c *KwClient) RoleManageUsers(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&manage_users=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update access source files permissions
func (c *KwClient) RoleAccessSourceFiles(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&access_source_files=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update assign role permissions
func (c *KwClient) RoleAssignRole(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&assign_role=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update project settings permissions
func (c *KwClient) RoleChangeProjectSettings(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&change_project_settings=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update create build permissions
func (c *KwClient) RoleCreateBuild(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&create_build=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update delete build permissions
func (c *KwClient) RoleDeleteBuild(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&delete_build=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update delete project permissions
func (c *KwClient) RoleDeleteProject(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&delete_project=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update manage modules permissions
func (c *KwClient) RoleManageModules(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&manage_modules=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update use local configuration permissions
func (c *KwClient) RoleUseLocalConfiguration(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&use_local_configuration=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update change issue status permissions
func (c *KwClient) RoleChangeIssueStatus(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&change_issue_status=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update web api access permissions
func (c *KwClient) RoleWebApiAccess(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&webapi_access=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

// Update perform cross-project synchronization permissions
func (c *KwClient) RoleExecKwxsync(
    name string, // The name of the role to update
    value bool, // Grant or revoke permission
) (*http.Response, error) {
    postData := "&name=" + name
    postData += "&execute_kwxsync=" + strconv.FormatBool(value)
    body, res, err := c.apiRequest("update_role_permissions", &postData)
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

