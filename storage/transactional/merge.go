package transactional

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/format/merge"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

type MergeStorage struct {
	storer.MergeStorer
	temporal storer.MergeStorer
}

// NewMergeStorage returns a new MergeStorer based on a base storer and a
// temporal storer.
func NewMergeStorage(s, temporal storer.MergeStorer) *MergeStorage {
	return &MergeStorage{MergeStorer: s, temporal: temporal}
}

// MergeHead returns the hash of the merge
func (m MergeStorage) MergeHead() plumbing.Hash {
	return m.MergeStorer.MergeHead()
}

// MergeMode returns the mode of the merge, e.g. no-ff
func (m MergeStorage) MergeMode() merge.Mode {
	return m.MergeStorer.MergeMode()
}

// MergeMsg returns the message of the merge
func (m MergeStorage) MergeMsg() string {
	return m.MergeStorer.MergeMsg()
}
