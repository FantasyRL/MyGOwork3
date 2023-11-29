package TaskService

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0 doing 1 finished
}

type CheckTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
type SearchTaskService struct {
	Information string `json:"information" form:"information"`
	PageNum     int    `json:"page_num" form:"page_num"`
	PageSize    int    `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0 doing 1 finished
}

type DeleteTaskService struct {
}
