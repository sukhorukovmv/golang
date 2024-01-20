package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"package/pkg"
)

func TestExecuteMergeDictsJob(t *testing.T) {
	a := assert.New(t)

	job, err := pkg.ExecuteMergeDictsJob(&pkg.MergeDictsJob{})
	a.Equal(pkg.ErrNotEnoughDicts, err)
	a.Equal(true, job.IsFinished)

	job, err = pkg.ExecuteMergeDictsJob(&pkg.MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
	}})
	a.Equal(pkg.ErrNotEnoughDicts, err)
	a.Equal(true, job.IsFinished)

	job, err = pkg.ExecuteMergeDictsJob(&pkg.MergeDictsJob{Dicts: []map[string]string{
		{"a": "b"},
		nil,
	}})
	a.Equal(pkg.ErrNilDict, err)
	a.Equal(true, job.IsFinished)

	job, err = pkg.ExecuteMergeDictsJob(&pkg.MergeDictsJob{Dicts: []map[string]string{
		{"a": "b", "b": "c"},
		{"d": "e", "f": "g"},
		{"a": "z", "f": "g"},
	}})
	a.NoError(err)
	a.Equal(true, job.IsFinished)
	a.Equal(map[string]string{
		"b": "c",
		"d": "e",
		"a": "z",
		"f": "g",
	}, job.Merged)
}
