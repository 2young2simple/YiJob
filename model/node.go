package model


type Node struct{
	IsHealth bool `json:"is_health"`
	IP string `json:"ip"`
	Port string `json:"port"`
	Name string `json:"name"`

	TaskList []string `json:"task_list"`
}
