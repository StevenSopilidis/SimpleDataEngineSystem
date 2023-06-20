package fs

type FsWriter interface {
	AppendSegment() bool
}
