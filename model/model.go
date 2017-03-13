package model

// API请求
type ApiRequest struct {
	Retcode int `json:"retcode"`
	Res    interface{}   `json:"res"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}