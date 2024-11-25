package cache

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

// NewByteView creates a new ByteView  struct
func NewByteView(b []byte) ByteView {
	return ByteView{b: b}
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return CloneBytes(v.b)
}

// ByteSource returns the byte slice of source.
func (v ByteView) ByteSource() []byte {
	return v.b
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

func CloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
