// Package procedures shows how Go handles functions and methods.
package procedures

type container struct{}

func (c *container) methodOne() error {
	return nil
}
