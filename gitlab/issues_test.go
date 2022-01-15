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
