package foo // import "github.com/maxbrunsfeld/counterfeiter/v6/fixtures/dup_packages/b/foo"

type S struct{}

//go:generate counterfeiter . I
type I interface {
	FromB() S
}
