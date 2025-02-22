package interactors

type MakeProductScreenshot interface {
	Call(id string) ([]byte, error)
}