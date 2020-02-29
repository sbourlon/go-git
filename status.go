package git

import (
	"bytes"
	"fmt"
	"path/filepath"
)

// Status represents the current status of a Worktree.
// The key of the map is the path of the file.
type Status map[string]*FileStatus

// File returns the FileStatus for a given path, if the FileStatus doesn't
// exists a new FileStatus is added to the map using the path as key.
func (s Status) File(path string) *FileStatus {
	if _, ok := (s)[path]; !ok {
		s[path] = &FileStatus{Worktree: Untracked, Staging: Untracked}
	}

	return s[path]
}

// IsUntracked checks if file for given path is 'Untracked'
func (s Status) IsUntracked(path string) bool {
	stat, ok := (s)[filepath.ToSlash(path)]
	return ok && stat.Worktree == Untracked
}

// IsClean returns true if all the files are in Unmodified status.
func (s Status) IsClean() bool {
	for _, status := range s {
		if status.Worktree != Unmodified || status.Staging != Unmodified {
			return false
		}
	}

	return true
}

func (s Status) String() string {
	buf := bytes.NewBuffer(nil)
	for path, status := range s {
		if status.Staging == Unmodified && status.Worktree == Unmodified {
			continue
		}

		if status.Staging == Renamed {
			path = fmt.Sprintf("%s -> %s", path, status.Extra)
		}

		fmt.Fprintf(buf, "%c%c %s\n", status.Staging, status.Worktree, path)
	}

	return buf.String()
}

// HasMergeConflict returns true if the status has merge conflicts, else false
func (s Status) HasMergeConflict() bool {
	for _, filestatus := range s {
		if filestatus.HasMergeConflict() {
			return true
		}
	}
	return false
}

// FileStatus contains the status of a file in the worktree
type FileStatus struct {
	// Staging is the status of a file in the staging area
	Staging StatusCode
	// Worktree is the status of a file in the worktree
	Worktree StatusCode
	// Extra contains extra information, such as the previous name in a rename
	Extra string
}

// HasMergeConflict returns true if the status has merge conflicts, else false
func (f FileStatus) HasMergeConflict() bool {
	if f.Worktree == UpdatedButUnmerged ||
		f.Staging == UpdatedButUnmerged ||
		(f.Staging == Added && f.Worktree == Added) ||
		(f.Staging == Deleted && f.Worktree == Deleted) {
		return true
	}
	return false
}

// StatusCode status code of a file in the Worktree
type StatusCode byte

const (
	Unmodified         StatusCode = ' '
	Untracked          StatusCode = '?'
	Modified           StatusCode = 'M'
	Added              StatusCode = 'A'
	Deleted            StatusCode = 'D'
	Renamed            StatusCode = 'R'
	Copied             StatusCode = 'C'
	UpdatedButUnmerged StatusCode = 'U'
)
