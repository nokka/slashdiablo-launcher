package d2

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total    float32
	Written  float32
	progress chan float32
}

// Write gets every write cycle reported on it.
func (wc *WriteCounter) Write(p []byte) (int, error) {
	// Bytes written this cycle.
	n := len(p)

	// Add the written bytes to the total.
	wc.Written += float32(n)

	// Calculate the percentage and send it on the channel.
	wc.progress <- wc.Written / wc.Total

	// Return the length of the written bytes this cycle.
	return n, nil
}
