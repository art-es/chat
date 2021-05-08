package ioc

import (
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	c := dig.New()

	providers := []Provider{
		repositoriesProvider,
		sessionProvider,
		websocketProvider,
		restProvider,
	}
	for _, p := range providers {
		if err := p.Provide(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

type Provider []interface{}

func (p Provider) Provide(c *dig.Container) error {
	for _, construct := range p {
		if err := c.Provide(construct); err != nil {
			return err
		}
	}
	return nil
}
