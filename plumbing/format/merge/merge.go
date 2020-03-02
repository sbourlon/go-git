package merge

// Mode is the merge mode (e.g. "no-ff") given by the file MERGE_MODE in the
// .git folder
type Mode int

const (
	// Default when MERGE_MODE is empty
	Default Mode = iota
	// NoFF when MERGE_MODE is no-ff
	NoFF
)
