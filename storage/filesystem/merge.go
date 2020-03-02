package filesystem

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/format/merge"
	"gopkg.in/src-d/go-git.v4/storage/filesystem/dotgit"
)

// MergeStorage gives the content of the file
// MERGE_HEAD, MERGE_MODE, MERGE_MSG in the
// .git folder during a merge
type MergeStorage struct {
	dir  *dotgit.DotGit
	head plumbing.Hash
	mode merge.Mode
	msg  string
}

// MergeHead returns the content of MERGE_HEAD
func (m *MergeStorage) MergeHead() plumbing.Hash {
	head, _ := m.dir.MergeHead()
	return head
}

// MergeMode returns the content of MERGE_MODE
func (m *MergeStorage) MergeMode() merge.Mode {
	mode, _ := m.dir.MergeMode()
	return mode
}

// MergeMsg returns the content of MERGE_MSG
func (m *MergeStorage) MergeMsg() string {
	msg, _ := m.dir.MergeMsg()
	return msg
}
