package fixtures

//go:generate counterfeiter . Something
type Something interface {
	DoThings(string, uint64) (int, error)
	DoNothing()
	DoASlice([]byte)
	DoAnArray([4]byte)
}
