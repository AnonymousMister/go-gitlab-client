package gitlab

import (
	"github.com/AnonymousMister/go-gitlab-client/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitlab_ProjectBranches(t *testing.T) {
	ts := test.CreateMockServer(t, []string{
		"branches/project_1_branches",
	})
	defer ts.Close()
	gitlab := NewGitlab(ts.URL, "", "")

	c, meta, err := gitlab.ProjectBranches("1", nil)

	assert.NoError(t, err)

	assert.Equal(t, 10, len(c.Items))

	assert.IsType(t, new(ResponseMeta), meta)
	assert.Equal(t, 1, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
}
