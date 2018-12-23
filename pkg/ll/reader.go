package ll

type Reader interface {

	Peek(n int) ([]byte, error)

	Discard(n int) (int, error)
}
