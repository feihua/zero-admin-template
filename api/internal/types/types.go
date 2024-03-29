// Code generated by goctl. DO NOT EDIT.
package types

type AddMenuReq struct {
	MenuName string `json:"menuName"`           // 菜单名称
	MenuType int64  `json:"menuType"`           // 菜单类型(1：目录   2：菜单   3：按钮)
	StatusId int64  `json:"statusId"`           // 状态(1:正常，0:禁用)
	Sort     int64  `json:"sort"`               // 排序
	ParentId int64  `json:"parentId,default=0"` // 父ID
	MenuUrl  string `json:"menuUrl,optional"`   // 路由路径
	ApiUrl   string `json:"apiUrl,optional"`    // 接口URL
	MenuIcon string `json:"menuIcon,optional"`  // 菜单图标
	Remark   string `json:"remark,optional"`    // 备注
}

type AddMenuResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type AddRoleReq struct {
	RoleName string `json:"roleName"` // 名称
	StatusId int64  `json:"statusId"` // 状态(1:正常，0:禁用)
	Sort     int64  `json:"sort"`     // 排序
	Remark   string `json:"remark"`   // 备注
}

type AddRoleResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type AddUserReq struct {
	Mobile   string `json:"mobile"`   // 手机
	UserName string `json:"userName"` // 姓名
	StatusId int64  `json:"statusId"` // 状态(1:正常，0:禁用)
	Sort     int64  `json:"sort"`     // 排序
	Remark   string `json:"remark"`   // 备注
}

type AddUserResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteLoginLogReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteLoginLogResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteMenuReq struct {
	Id int64 `json:"id"`
}

type DeleteMenuResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteOperateLogReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteOperateLogResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteRoleReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteRoleResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type DeleteUserReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteUserResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type LoginLogListData struct {
	Id        int64  `json:"id"`        // 编号
	UserName  string `json:"userName"`  // 用户名
	Status    string `json:"status"`    // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
	Ip        string `json:"ip"`        // IP地址
	LoginTime string `json:"loginTime"` // 登陆时间
}

type OperateLogListData struct {
	Id             int64  `json:"id"`             // 编号
	UserName       string `json:"userName"`       // 用户名
	Operation      string `json:"operation"`      // 用户操作
	Method         string `json:"method"`         // 请求方法
	RequestParams  string `json:"requestParams"`  // 请求参数
	ResponseParams string `json:"responseParams"` // 响应参数
	Time           int64  `json:"time"`           // 执行时长(毫秒)
	Ip             string `json:"ip"`             // IP地址
	OperationTime  string `json:"operationTime"`  // 操作时间
}

type QueryLoginLogListReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	UserName string `json:"userName,optional"`
	Ip       string `json:"ip,optional"` // IP地址
}

type QueryLoginLogListResp struct {
	Code    int64              `json:"code"`
	Msg     string             `json:"msg"`
	Data    []LoginLogListData `json:"data"`
	Success bool               `json:"success"`
	Total   int64              `json:"total"`
}

type QueryMenuListData struct {
	Id         int64  `json:"id"`         // 主键
	MenuName   string `json:"menuName"`   // 菜单名称
	MenuType   int64  `json:"menuType"`   // 菜单类型(1：目录   2：菜单   3：按钮)
	StatusId   int64  `json:"statusId"`   // 状态(1:正常，0:禁用)
	Sort       int64  `json:"sort"`       // 排序
	ParentId   int64  `json:"parentId"`   // 父ID
	MenuUrl    string `json:"menuUrl"`    // 路由路径
	ApiUrl     string `json:"apiUrl"`     // 接口URL
	MenuIcon   string `json:"menuIcon"`   // 菜单图标
	Remark     string `json:"remark"`     // 备注
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 修改时间
}

type QueryMenuListReq struct {
	MenuName string `json:"menuName,optional"`
}

type QueryMenuListResp struct {
	Code int64               `json:"code"`
	Msg  string              `json:"msg"`
	Data []QueryMenuListData `json:"data"`
}

type QueryOperateLogListReq struct {
	Current   int64  `json:"current,default=1"`
	PageSize  int64  `json:"pageSize,default=20"`
	UserName  string `json:"userName,optional"`  // 用户名
	Operation string `json:"operation,optional"` // 用户操作
	Method    string `json:"method,optional"`    // 请求方法
	Ip        string `json:"ip,optional"`        // IP地址
}

type QueryOperateLogListResp struct {
	Code    int64                `json:"code"`
	Msg     string               `json:"msg"`
	Data    []OperateLogListData `json:"data"`
	Success bool                 `json:"success"`
	Total   int64                `json:"total"`
}

type QueryRoleListReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	RoleName string `json:"roleName,optional"`
	StatusId int64  `json:"statusId,default=2"`
}

type QueryRoleListResp struct {
	Code    int64          `json:"code"`
	Msg     string         `json:"msg"`
	Data    []RoleListData `json:"data"`
	Success bool           `json:"success"`
	Total   int64          `json:"total"`
}

type QueryRoleMenuListData struct {
	RoleMenus    []int64        `json:"roleMenus"`
	MenuRoleList []RoleMenuList `json:"menuRoleList"`
}

type QueryRoleMenuListReq struct {
	RoleId int64 `json:"roleId"`
}

type QueryRoleMenuListResp struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data QueryRoleMenuListData `json:"data"`
}

type QueryUserListReq struct {
	Current  int64  `json:"current,default=1"`
	PageSize int64  `json:"pageSize,default=20"`
	Mobile   string `json:"mobile,optional"`    // 手机
	UserName string `json:"userName,optional"`  // 姓名
	StatusId int64  `json:"statusId,default=2"` // 状态(1:正常，0:禁用)
}

type QueryUserListResp struct {
	Code    int64          `json:"code"`
	Msg     string         `json:"msg"`
	Data    []UserListData `json:"data"`
	Success bool           `json:"success"`
	Total   int64          `json:"total"`
}

type QueryUserMenuListData struct {
	SysMenu []UserMenuList `json:"sysMenu"`
	BtnMenu []string       `json:"btnMenu"`
	Avatar  string         `json:"avatar"`
	Name    string         `json:"name"`
}

type QueryUserMenuListResp struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data QueryUserMenuListData `json:"data"`
}

type QueryUserRoleListData struct {
	AllRoles       []UserRoleList `json:"allRoles"`
	AllUserRoleIds []int64        `json:"allUserRoleIds"`
}

type QueryUserRoleListReq struct {
	UserId int64 `json:"userId"`
}

type QueryUserRoleListResp struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data QueryUserRoleListData `json:"data"`
}

type RoleListData struct {
	Id         int64  `json:"id"`         // 主键
	RoleName   string `json:"roleName"`   // 名称
	StatusId   int64  `json:"statusId"`   // 状态(1:正常，0:禁用)
	Sort       int64  `json:"sort"`       // 排序
	Remark     string `json:"remark"`     // 备注
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 修改时间
}

type RoleMenuList struct {
	Id       int64  `json:"id"`
	ParentID int64  `json:"parentId"`
	Title    string `json:"title"`
	Key      string `json:"key"`
}

type StatisticsLoginLogData struct {
	DayLoginCount   int32 `json:"dayLoginCount"`   //查询当天登录人数（根据IP）
	WeekLoginCount  int32 `json:"weekLoginCount"`  //统计当前周登录人数（根据IP）
	MonthLoginCount int32 `json:"monthLoginCount"` //统计当前月登录人数（根据IP）
}

type StatisticsLoginLogReq struct {
}

type StatisticsLoginLogResp struct {
	Code    int64                  `json:"code"`
	Message string                 `json:"message"`
	Data    StatisticsLoginLogData `json:"data"`
}

type UpdateMenuReq struct {
	Id       int64  `json:"id"`                 // 主键
	MenuName string `json:"menuName"`           // 菜单名称
	MenuType int64  `json:"menuType"`           // 菜单类型(1：目录   2：菜单   3：按钮)
	StatusId int64  `json:"statusId"`           // 状态(1:正常，0:禁用)
	Sort     int64  `json:"sort"`               // 排序
	ParentId int64  `json:"parentId,default=0"` // 父ID
	MenuUrl  string `json:"menuUrl,optional"`   // 路由路径
	ApiUrl   string `json:"apiUrl,optional"`    // 接口URL
	MenuIcon string `json:"menuIcon,optional"`  // 菜单图标
	Remark   string `json:"remark,optional"`    // 备注
}

type UpdateMenuResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type UpdateRoleMenuListReq struct {
	MenuIds []int64 `json:"menuIds"`
	RoleId  int64   `json:"roleId"`
}

type UpdateRoleMenuListResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type UpdateRoleReq struct {
	Id       int64  `json:"id"`       // 主键
	RoleName string `json:"roleName"` // 名称
	StatusId int64  `json:"statusId"` // 状态(1:正常，0:禁用)
	Sort     int64  `json:"sort"`     // 排序
	Remark   string `json:"remark"`   // 备注
}

type UpdateRoleResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type UpdateUserReq struct {
	Id       int64  `json:"id"`       // 主键
	Mobile   string `json:"mobile"`   // 手机
	UserName string `json:"userName"` // 姓名
	StatusId int64  `json:"statusId"` // 状态(1:正常，0:禁用)
	Sort     int64  `json:"sort"`     // 排序
	Remark   string `json:"remark"`   // 备注
}

type UpdateUserResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type UpdateUserRoleListReq struct {
	UserId  int64   `json:"userId"`
	RoleIds []int64 `json:"roleIds"`
}

type UpdateUserRoleListResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type UserListData struct {
	Id         int64  `json:"id"`         // 主键
	Mobile     string `json:"mobile"`     // 手机
	UserName   string `json:"userName"`   // 姓名
	StatusId   int64  `json:"statusId"`   // 状态(1:正常，0:禁用)
	Sort       int64  `json:"sort"`       // 排序
	Remark     string `json:"remark"`     // 备注
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 修改时间
}

type UserLoginData struct {
	UserId int64  `json:"userId"`
	Mobile string `json:"mobile"`
	Token  string `json:"token"`
}

type UserLoginReq struct {
	Mobile   string `json:"account"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	Code int64         `json:"code"`
	Msg  string        `json:"msg"`
	Data UserLoginData `json:"data"`
}

type UserMenuList struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parentId"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	ApiUrl   string `json:"apiUrl"`
	MenuType int64  `json:"menuType"`
	Icon     string `json:"icon"`
}

type UserRoleList struct {
	Id         int64  `json:"id"`         // 主键
	RoleName   string `json:"roleName"`   // 名称
	StatusId   int64  `json:"statusId"`   // 状态(1:正常，0:禁用)
	Sort       int64  `json:"sort"`       // 排序
	Remark     string `json:"remark"`     // 备注
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 修改时间
}
