package commonvos

type Sex uint8

const (
	Man Sex = iota
	Woman
)

type InsideUserInfo struct {
	Id              int64    `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Sex             Sex      `json:"sex,omitempty"`
	Age             int      `json:"age,omitempty"`
	Birthday        string   `json:"birthday,omitempty"`
	IdCode          string   `json:"idCode,omitempty"`
	Phone           string   `json:"phone,omitempty"`
	Email           string   `json:"email,omitempty"`
	CompanyName     string   `json:"companyName,omitempty"`
	JobName         string   `json:"jobName,omitempty"`
	PostsName       string   `json:"postsName,omitempty"`
	DeptName        string   `json:"deptName,omitempty"`
	DeptMain        bool     `json:"deptMain,omitempty"`
	UserName        string   `json:"userName,omitempty"`
	RegistryTime    string   `json:"registryTime,omitempty"`
	DataInTime      string   `json:"dataInTime,omitempty"`
	EndUpdateTime   string   `json:"endUpdateTime,omitempty"`
	EndLoginTime    string   `json:"endLoginTime,omitempty"`
	Icon            string   `json:"icon,omitempty"`
	BackgroundImage string   `json:"backgroundImage,omitempty"`
	BackgroundColor string   `json:"backgroundColor,omitempty"`
	SysFlagList     []string `json:"sysFlagList,omitempty"`
	QueueHost       string   `json:"queueHost,omitempty"`
	QueuePort       int      `json:"queuePort,omitempty"`
	ServerPort      string   `json:"serverPort,omitempty"`
	Token           string   `json:"token,omitempty"`
}