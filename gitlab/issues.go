package gitlab

import (
	"encoding/json"
)

const (
	ProjectIssuesApiPath   = "/projects/:id/issues"
	ProjectIssueOneApiPath = "/projects/:id/issues/:issue_iid"
)

type Issue struct {
	Id                 int        `json:"id"`
	IId                int        `json:"iid"`
	ProjectId          int        `json:"project_id,omitempty"`
	Title              string     `json:"title,omitempty"`
	Description        string     `json:"description,omitempty"`
	Labels             []string   `json:"labels,omitempty"`
	Milestone          *Milestone `json:"milestone,omitempty"`
	Assignees          []*User    `json:"assignees"`
	Assignee           *User      `json:"assignee,omitempty"`
	Author             *User      `json:"author,omitempty"`
	State              string     `json:"state,omitempty"`
	CreatedAt          string     `json:"created_at,omitempty"`
	UpdatedAt          string     `json:"updated_at,omitempty"`
	UserNotesCount     int        `json:"user_notes_count,omitempty"`
	MergeRequestsCount int        `json:"merge_requests_count,omitempty" `
	Upvotes            int        `json:"upvotes,omitempty"`
	Downvotes          int        `json:"downvotes,omitempty"`
	DueDate            string     `json:"due_date,omitempty"`
	Confidential       bool       `json:"confidential,omitempty"`
	//DiscussionLocked     interface{}           `json:"discussion_locked,omitempty"`
	WebUrl               string                `json:"web_url,omitempty"`
	TimeStats            *TimeStats            `json:"time_stats,omitempty"`
	TaskCompletionStatus *TaskCompletionStatus `json:"task_completion_status,omitempty"`
	HasTasks             bool                  `json:"has_tasks,omitempty"`
	Subscribed           bool                  `json:"subscribed,omitempty"`
	//MovedToId  interface{} `json:"moved_to_id"`
}

type TaskCompletionStatus struct {
	Count          int `json:"count,omitempty"`
	CompletedCount int `json:"completed_count,omitempty"`
}

type IssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	AssigneeId  int    `json:"assignee_id,omitempty"`
	MilestoneId int    `json:"milestone_id,omitempty"`
	Labels      string `json:"labels,omitempty"`
}

func (g *Gitlab) GetIssue(projectId string, issueIId string) (issue *Issue, meta *ResponseMeta, err error) {
	params := map[string]string{
		":id":        projectId,
		":issue_iid": issueIId,
	}
	u := g.ResourceUrl(ProjectIssueOneApiPath, params)
	data, _, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	issue = new(Issue)
	err = json.Unmarshal(data, issue)
	if err != nil {
		panic(err)
	}
	return
}

func (g *Gitlab) AddIssue(projectId string, req *IssueRequest) (issue *Issue, meta *ResponseMeta, err error) {
	params := map[string]string{
		":id": projectId,
	}
	u := g.ResourceUrl(ProjectIssuesApiPath, params)

	encodedRequest, err := json.Marshal(req)
	if err != nil {
		return
	}

	data, _, err := g.buildAndExecRequest("POST", u.String(), encodedRequest)
	if err != nil {
		return
	}

	issue = new(Issue)
	err = json.Unmarshal(data, issue)
	if err != nil {
		panic(err)
	}

	return
}
