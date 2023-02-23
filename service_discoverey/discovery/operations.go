package discovery

import "github.com/hashicorp/consul/api"

func (c *client) Register(tags []string) error {
	reg := &api.AgentServiceRegistration{
		ID:      c.name,
		Name:    c.name,
		Tags:    tags,
		Port:    c.port,
		Address: c.address,
	}
	return c.client.Agent().ServiceRegister(reg)
}

func (c *client) Service(service string, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.client.Health().Service(service, tag, true, nil)
}