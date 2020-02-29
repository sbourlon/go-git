package git

import (
	"math/rand"
	"strconv"
	"time"

	. "gopkg.in/check.v1"
)

type FileStatusSuite struct {
	BaseSuite
}

type FileStatusTestCase struct {
	In   FileStatus
	Want bool
}

var _ = Suite(&FileStatusSuite{})

func GenerateStatusCode(possible, impossible bool) StatusCode {
	possibleConflict := []StatusCode{
		UpdatedButUnmerged,
	}

	impossibleConflict := []StatusCode{
		Added,
		Deleted,
		Unmodified,
		Untracked,
		Modified,
		Renamed,
		Copied,
	}

	statuscode := make([]StatusCode, 0)

	if possible {
		statuscode = append(statuscode, possibleConflict...)
	}
	if impossible {
		statuscode = append(statuscode, impossibleConflict...)
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(statuscode))
	return statuscode[i]
}

func (s *FileStatusSuite) TestHasMergeConflict(c *C) {
	suite := make([]FileStatusTestCase, 0)

	var t FileStatusTestCase
	for i := 0; i < 8; i++ {
		// Staging UpdatedButUnmerged
		t = FileStatusTestCase{
			In: FileStatus{
				Staging:  UpdatedButUnmerged,
				Worktree: GenerateStatusCode(true, true),
			},
			Want: true,
		}
		suite = append(suite, t)

		// Worktree UpdatedButUnmerged
		t = FileStatusTestCase{
			In: FileStatus{
				Staging:  GenerateStatusCode(true, true),
				Worktree: UpdatedButUnmerged,
			},
			Want: true,
		}

		suite = append(suite, t)
	}

	// Staging and Worktree Deleted
	tBothDeleted := FileStatusTestCase{
		In: FileStatus{
			Staging:  Deleted,
			Worktree: Deleted,
		},
		Want: true,
	}

	// Staging and Worktree Added
	tBothAdded := FileStatusTestCase{
		In: FileStatus{
			Staging:  Added,
			Worktree: Added,
		},
		Want: true,
	}

	suite = append(suite, tBothDeleted, tBothAdded)

	for _, tc := range suite {
		in := tc.In
		want := tc.Want
		got := in.HasMergeConflict()
		c.Assert(got, Equals, want)
	}
}

type StatusSuite struct {
	BaseSuite
}

var _ = Suite(&FileStatusSuite{})

type StatusTestCase struct {
	In   Status
	Want bool
}

func (s *StatusSuite) TestHasMergeConflict(c *C) {
	suite := make([]StatusTestCase, 0)

	var tPassing StatusTestCase
	tPassing.In = make(Status)
	tPassing.Want = false

	passingSize := 8
	for i := 0; i < passingSize; i++ {
		filestatus := FileStatus{
			Staging:  GenerateStatusCode(false, true),
			Worktree: GenerateStatusCode(false, true),
		}
		tPassing.In[strconv.Itoa(i)] = &filestatus
	}

	suite = append(suite, tPassing)

	tFailing := tPassing
	tFailing.Want = true

	for i := passingSize; i < passingSize+8; i++ {
		filestatus := FileStatus{
			Staging:  GenerateStatusCode(true, false),
			Worktree: GenerateStatusCode(false, false),
		}
		tPassing.In[strconv.Itoa(i)] = &filestatus
	}

	suite = append(suite, tPassing)

	for _, tc := range suite {
		in := tc.In
		want := tc.Want
		got := in.HasMergeConflict()
		c.Assert(got, Equals, want)
	}
}
