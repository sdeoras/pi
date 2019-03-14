package pi

// Getter gets digits of pi
type Getter interface {
	Get(start, numberOfDigits int) ([]byte, int, error)
}
