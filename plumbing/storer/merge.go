package storer

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/format/merge"
)

// MergeStorer access the merge information
// In .git folder, that's the files
// MERGE_HEAD, MERGE_MODE, MERGE_MSG
type MergeStorer interface {
	MergeHead() plumbing.Hash
	MergeMode() merge.Mode
	MergeMsg() string
}
