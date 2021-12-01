// Code generated by carrier, DO NOT EDIT.
package carrier

import (
	"github.com/Yiling-J/carrier/integration/carrier/factory"

	"github.com/Yiling-J/carrier/integration/ent"
)

type Factory struct {
	groupCategoryFactory *factory.GroupCategoryFactory

	groupFactory *factory.GroupFactory

	userFactory *factory.UserFactory
}

func GroupCategoryMetaFactory() *factory.GroupCategoryMetaFactory {
	return &factory.GroupCategoryMetaFactory{}
}
func (f *Factory) SetGroupCategoryFactory(c *factory.GroupCategoryFactory) *Factory {
	f.groupCategoryFactory = c
	return f
}

func (f *Factory) GroupCategoryFactory() *factory.GroupCategoryFactory {
	return f.groupCategoryFactory
}

func GroupMetaFactory() *factory.GroupMetaFactory {
	return &factory.GroupMetaFactory{}
}
func (f *Factory) SetGroupFactory(c *factory.GroupFactory) *Factory {
	f.groupFactory = c
	return f
}

func (f *Factory) GroupFactory() *factory.GroupFactory {
	return f.groupFactory
}

func UserMetaFactory() *factory.UserMetaFactory {
	return &factory.UserMetaFactory{}
}
func (f *Factory) SetUserFactory(c *factory.UserFactory) *Factory {
	f.userFactory = c
	return f
}

func (f *Factory) UserFactory() *factory.UserFactory {
	return f.userFactory
}

type EntFactory struct {
	userFactory *factory.EntUserFactory

	carFactory *factory.EntCarFactory

	groupFactory *factory.EntGroupFactory

	client *ent.Client
}

func (f *EntFactory) SetClient(c *ent.Client) {
	f.client = c

	_ = f.userFactory.Client(f.client)

	_ = f.carFactory.Client(f.client)

	_ = f.groupFactory.Client(f.client)

}

func (f *EntFactory) Client() *ent.Client {
	return f.client
}

func EntUserMetaFactory() *factory.EntUserMetaFactory {
	return &factory.EntUserMetaFactory{}
}
func (f *EntFactory) SetUserFactory(c *factory.EntUserFactory) *EntFactory {
	f.userFactory = c
	return f
}
func (f *EntFactory) UserFactory() *factory.EntUserFactory {
	return f.userFactory
}

func EntCarMetaFactory() *factory.EntCarMetaFactory {
	return &factory.EntCarMetaFactory{}
}
func (f *EntFactory) SetCarFactory(c *factory.EntCarFactory) *EntFactory {
	f.carFactory = c
	return f
}
func (f *EntFactory) CarFactory() *factory.EntCarFactory {
	return f.carFactory
}

func EntGroupMetaFactory() *factory.EntGroupMetaFactory {
	return &factory.EntGroupMetaFactory{}
}
func (f *EntFactory) SetGroupFactory(c *factory.EntGroupFactory) *EntFactory {
	f.groupFactory = c
	return f
}
func (f *EntFactory) GroupFactory() *factory.EntGroupFactory {
	return f.groupFactory
}
