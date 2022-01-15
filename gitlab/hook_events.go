package gitlab

type Events struct {
	ObjectKind string         `json:"object_kind"`
	EventType  string         `json:"event_type"`
	Project    *EventsProject `json:"project"`
	IssuesText *IssuesText    `json:"object_attributes"`
	Repository *hRepository   `json:"repository"`
	Assignees  []*Assignee    `json:"assignees"`
	Changes    *Changes       `json:"changes"`
	Issue      *IssuesText    `json:"issue"`
	User       *User          `json:"user"`
}

type Changes struct {
	UpdatedAt *UpdatedAt `json:"updated_at"`
	Assignees *Assignees `json:"assignees"`
	Labels    *Labels    `json:"labels"`
}

/*
Changes 对象 内部实体
标签跟新时 对象
*/
type Labels struct {
	Previous []*Label `json:"previous"`
	Current  []*Label `json:"current"`
}

/*
Changes 对象 内部实体
截止日期跟新时 对象
*/
type UpdatedAt struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
}

/*
Changes 对象 内部实体
指派人跟新时对象
*/
type Assignees struct {
	Previous []*Assignee `json:"previous"`
	Current  []*Assignee `json:"current"`
}

type Assignee struct {
	User
	Mobile string `json:"mobile"`
}

type Label struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type IssuesText struct {
	AuthorId         int      `json:"author_id"`
	Confidential     bool     `json:"confidential"`
	CreatedAt        string   `json:"created_at"`
	Description      string   `json:"description"`
	DueDate          string   `json:"due_date"`
	Id               int      `json:"id"`
	Iid              int      `json:"iid"`
	ProjectId        int      `json:"project_id"`
	RelativePosition int      `json:"relative_position"`
	StateId          int      `json:"state_id"`
	TimeEstimate     int      `json:"time_estimate"`
	Title            string   `json:"title"`
	UpdatedAt        string   `json:"updated_at"`
	UpdatedById      int      `json:"updated_by_id"`
	Url              string   `json:"url"`
	TotalTimeSpent   int      `json:"total_time_spent"`
	AssigneeIds      []int    `json:"assignee_ids"`
	AssigneeId       int      `json:"assignee_id"`
	Labels           []*Label `json:"labels"`
	State            string   `json:"state"`
	Action           string   `json:"action"`
	Note             string   `json:"note"`
}
