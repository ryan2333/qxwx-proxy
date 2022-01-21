package httpdata

type TextData struct {
	commonData
	Text map[string]interface{} `json:"text"`
}

type MediaData struct {
	commonData
	Image map[string]interface{} `json:"image"`
}

type FileData struct {
	commonData
	File map[string]interface{} `json:"file"`
}

type TextCardData struct {
	commonData
	TextCard textCard `json:"textcard"`
}

type textCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	BtnTxt      string `json:"btntxt"`
}

type MarkDownData struct {
	commonData
	MarkDown map[string]interface{} `json:"markdown"`
}

type TaskCard struct {
	commonData
	TaskCard taskCard `json:"taskCard"`
}

type taskCard struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Task_Id     string   `json:"task_id"`
	Button      []button `json:"btn"`
}

type button struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	Replace_Name string `json:"replace_name"`
	Color        string `json:"color"`
	Is_Bold      bool   `json:"is_bold"`
}

type ReceiveData struct {
	commonData
	TemplateId string `json:"templateId"`
}
