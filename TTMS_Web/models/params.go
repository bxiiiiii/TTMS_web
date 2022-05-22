package models

//存放前端发送请求参数的结构体

// ParamsLogin 登录接收的参数 form为接收的名称，required为必填字段
type ParamsLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ParamsRegister struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
}

type ParamsFrontPage struct {
	CarouselNum         int `json:"carouse_num" form:"carouse_num" binding:"required"`
	ShowingNum          int `json:"showing_num" form:"showing_num" binding:"required"`
	ComingNum           int `json:"coming_num" form:"coming_num" binding:"required"`
	ScoreRankingNum     int `json:"scoreRanking_num" form:"scoreRanking_num" binding:"required"`
	BoxOfficeRankingNum int `json:"boxofficeRanking_num" form:"boxofficeRanking_num" binding:"required"`
}

type ParamsMovie struct {
	Id int64 `json:"id" from:"id" binding:"required"`
}

type ParamsAdminmsg struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
	Identity    int    `json:"identity" from:"identity" binding:"required"`
}

type ParamsUpdateMsg struct {
	Id          int64  `json:"id" from:"id" binding:"required"`
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
	IsDelete    int    `json:"is_delete" from:"id_delete" binding:"required"`
	Identity    int    `json:"identity" from:"identity" binding:"required"`
	IsLogin     int    `json:"is_login" from:"is_login" binding:"required"`
}