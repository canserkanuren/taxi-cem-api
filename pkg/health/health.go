package health

import "fmt"

type Check struct {
	AppName string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func (c *Check) GoString() string {
	return fmt.Sprintf(`{
		Name: %s,
		Version: %s,
		Status: %s
	}`, c.AppName, c.Version, c.Status)
}
