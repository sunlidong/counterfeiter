package foo // import "github.com/maxbrunsfeld/counterfeiter/v6/fixtures/dup_packages/foo"

import (
	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures/dup_packages/a/foo"
	bfoo "github.com/maxbrunsfeld/counterfeiter/v6/fixtures/dup_packages/b/foo"
)

type S struct{}

//go:generate counterfeiter . MultiAB
type MultiAB interface {
	Mine() S
	foo.I
	bfoo.I
}
