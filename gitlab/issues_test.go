package gitlab

import (
	"github.com/AnonymousMister/go-gitlab-client/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIssue(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"notes/project_issue",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	issue, meta, err := gitlab.GetIssue("4", "194")
	assert.NoError(t, err)
	assert.Equal(t, 4, issue.ProjectId)
	assert.Equal(t, 194, issue.IId)
	assert.IsType(t, new(ResponseMeta), meta)

}

func TestGetMyIssues(t *testing.T) {

	gitlab := NewGitlab("http://192.168.1.8:18080", "", "XZAHghaEr5gmRU1GscPG")

	issues, meta, err := gitlab.GetMyIssues(&QIssuesRequest{
		Labels: "优化",
	})

	assert.Equal(t, 4, len(issues))
	assert.NoError(t, err)
	assert.IsType(t, new(ResponseMeta), meta)

}

func TestUpIssue(t *testing.T) {

	gitlab := NewGitlab("http://192.168.1.8:18080", "", "XZAHghaEr5gmRU1GscPG")

	issue, meta, err := gitlab.UpIssue("4", "216", &UpIssueRequest{
		StateEvent: "close",
	})

	assert.Equal(t, 4, issue.ProjectId)
	assert.NoError(t, err)
	assert.IsType(t, new(ResponseMeta), meta)

}
