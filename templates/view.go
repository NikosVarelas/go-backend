package templates

import "fmt"

func Asset(name string) string {
	return fmt.Sprintf("/static/css/%s", name)
}