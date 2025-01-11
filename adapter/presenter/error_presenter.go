
package presenter

import "fmt"

type ErrorPresenter struct{}

func NewErrorPresenter() *ErrorPresenter {
    return &ErrorPresenter{}
}

func (p *ErrorPresenter) ShowError(err error) {
    fmt.Printf("❌ Error: %v\n", err)
}
