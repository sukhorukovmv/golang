package pkg

import (
	"errors"
)

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

var (
	ErrNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	ErrNilDict        = errors.New("nil dictionary")
)

func (j *MergeDictsJob) setFinish() {
	j.IsFinished = true
}

// BEGIN (write your solution here)
func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	defer job.setFinish()
	if len(job.Dicts) < 2 {
		return job, ErrNotEnoughDicts
	}

	job.Merged = make(map[string]string)
	for _, dict := range job.Dicts {
		if dict == nil {
			return job, ErrNilDict
		}

		for k, v := range dict {
			job.Merged[k] = v
		}
	}
	return job, nil
}

// END
