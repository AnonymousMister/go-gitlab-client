package gitlab

import (
	"encoding/json"
	"strconv"
)

const (
	ProjectIssuesApiPath   = "/projects/:id/issues"
	ProjectIssueOneApiPath = "/projects/:id/issues/:issue_iid"
	IssuesApiPath          = "/issues"
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

type QIssuesRequest struct {
	PaginationOptions
	State  string `json:"state,omitempty" url:"state,omitempty"`
	Labels string `json:"labels,omitempty" url:"labels,omitempty"`
}

type UpIssueRequest struct {
	StateEvent string `json:"state_event,omitempty"`
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
func (g *Gitlab) GetMyIssues(qIssues *QIssuesRequest) (issues []*Issue, meta *ResponseMeta, err error) {

	u := g.ResourceUrlQ(IssuesApiPath, nil, qIssues)
	data, meta, err := g.buildAndExecRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	issues = []*Issue{}
	err = json.Unmarshal(data, &issues)
	if err != nil {
		panic(err)
	}
	return
}

func (g *Gitlab) UpIssue(projectId int, issueIId int, upIssues *UpIssueRequest) (issue *Issue, meta *ResponseMeta, err error) {
	params := map[string]string{
		":id":        strconv.Itoa(projectId),
		":issue_iid": strconv.Itoa(issueIId),
	}
	u := g.ResourceUrl(ProjectIssueOneApiPath, params)
	upIssuesBody, err := json.Marshal(upIssues)
	if err != nil {
		return
	}
	data, meta, err := g.buildAndExecRequest("PUT", u.String(), upIssuesBody)
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
