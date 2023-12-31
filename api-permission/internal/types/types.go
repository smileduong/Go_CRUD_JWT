// Code generated by goctl. DO NOT EDIT.
package types

type CreateUserRequest struct {
	HeadImg  string  `json:"headimg"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	// RoleIds  []int64 `json:"role_ids"`
}

type CreateUserResponse struct {
	Id        int64  `json:"id"`
	HeadImg   string `json:"headimg"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

type UpdateUserRequest struct {
	Id       int64   `json:"id"`
	HeadImg  string  `json:"headimg"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	RoleIds  []int64 `json:"role_ids"`
}

type UpdateUserResponse struct {
	Id       int64  `json:"id"`
	HeadImg  string `json:"headimg"`
	Username string `json:"username"`
}

type ViewUserRequest struct {
	Id int64 `json:"id"`
}

type ViewUserResponse struct {
	Id       int64  `json:"id"`
	HeadImg  string `json:"headimg"`
	Username string `json:"username"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ListUserItem struct {
	Id       int64  `json:"id"`
	HeadImg  string `json:"headimg"`
	Username string `json:"username"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ListUserResponse struct {
	Items  []ListUserItem `json:"items"`
	Paging ListPageItem   `json:"paging"`
}

type LoginUsernameReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

type LoginUsernameResp struct {
	Token string `json:"token"`
}

type ListPageItem struct {
	Limit     int64 `json:"limit"`
	Page      int64 `json:"page"`
	TotalPage int64 `json:"total_page"`
	TotalRows int64 `json:"total_rows"`
}

type ListCommonRequest struct {
	Limit  int64                        `json:"limit"`
	Page   int64                        `json:"page"`
	Filter map[string]map[string]string `json:"filter,optional"`
}

type CommonResponse struct {
	Success bool `json:"success"`
}

type AddRuleRequest struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type AddRuleResponse struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type UpdateRuleRequest struct {
	Id   int64  `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type UpdateRuleResponse struct {
	Id       int64  `json:"id"`
	Uri      string `json:"uri"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ViewRuleRequest struct {
	Id int64 `json:"id"`
}

type ViewRuleResponse struct {
	Id       int64  `json:"id"`
	Uri      string `json:"uri"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ListRuleItem struct {
	Id       int64  `json:"id"`
	Uri      string `json:"uri"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ListRuleResponse struct {
	Items  []ListRuleItem `json:"items"`
	Paging ListPageItem   `json:"paging"`
}

type CreateRoleRequest struct {
	Name    string  `json:"name"`
	Pid     int64   `json:"pid"`
	MaxRole int     `json:"max_role"`
	Rules   []int64 `json:"rules"`
	Slug    string  `json:"slug"`
}

type CreateRoleResponse struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Pid     int64  `json:"pid"`
	MaxRole int    `json:"max_role"`
	Slug    string `json:"slug"`
}

type UpdateRoleRequest struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Pid     int64   `json:"pid"`
	MaxRole int     `json:"max_role"`
	Rules   []int64 `json:"rules"`
	Slug    string  `json:"slug"`
}

type UpdateRoleResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Pid      int64  `json:"pid"`
	MaxRole  int    `json:"max_role"`
	Slug     string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ViewRoleRequest struct {
	Id int64 `json:"id"`
}

type ViewRoleResponse struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Pid      int64  `json:"pid"`
	MaxRole  int    `json:"max_role"`
	Slug     string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ListRoleItem struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Pid      int64  `json:"pid"`
	MaxRole  int    `json:"max_role"`
	Slug     string `json:"slug"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type ListRoleResponse struct {
	Items  []ListRoleItem `json:"items"`
	Paging ListPageItem   `json:"paging"`
}

type AddRulesRequest struct {
	RoleId int64   `json:"role_id"`
	Rules  []int64 `json:"rules"`
}
