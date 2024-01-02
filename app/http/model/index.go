package model

type ApiRecordModel struct {
	GroupId int    `json:"group_id"  form:"group_id"`
	Name    string `json:"name" form:"name"`
	Url     string `json:"url" form:"url"`
	Method  string `json:"method" form:"method"`
	Ip      string `json:"ip" form:"ip"`
}

type AuthTokenModel struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
}
