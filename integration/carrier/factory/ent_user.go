// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/integration/ent"

	"context"
)

type EntUserMutator struct {
	Age int

	Email string

	Name string

	_creator *ent.UserCreate
}

func (m *EntUserMutator) EntCreator() *ent.UserCreate {
	return m._creator
}

type entUserMutation struct {
	ageType int
	ageFunc func(ctx context.Context, i *EntUserMutator, c int) error

	emailType int
	emailFunc func(ctx context.Context, i *EntUserMutator, c int) error

	nameType int
	nameFunc func(ctx context.Context, i *EntUserMutator, c int) error

	_postGroupsFunc func(ctx context.Context, set bool, obj *ent.User, i int) error

	beforeCreateFunc func(ctx context.Context, i *EntUserMutator) error
	afterCreateFunc  func(ctx context.Context, i *ent.User) error
}
type EntUserMetaFactory struct {
	mutation entUserMutation
}
type entUserTrait struct {
	mutation entUserMutation
	updates  []func(m *entUserMutation)
}

func EntUserTrait() *entUserTrait {
	return &entUserTrait{}
}
func (*entUserMutation) beforeCreateMutateFunc(fn func(ctx context.Context, i *EntUserMutator) error) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.beforeCreateFunc = fn
	}
}
func (*entUserMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.User) error) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entUserMutation) ageSequenceMutateFunc(fn func(ctx context.Context, i int) (int, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeSequence
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetAge(value)

			i.Age = value
			return nil
		}
	}
}
func (*entUserMutation) ageLazyMutateFunc(fn func(ctx context.Context, i *EntUserMutator) (int, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeLazy
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetAge(value)

			i.Age = value
			return nil
		}
	}
}
func (*entUserMutation) ageDefaultMutateFunc(v int) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeDefault
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int) error {

			i.EntCreator().SetAge(v)

			i.Age = v
			return nil
		}
	}
}
func (*entUserMutation) ageFactoryMutateFunc(fn func(ctx context.Context) (int, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeFactory
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetAge(value)

			i.Age = value

			return nil
		}
	}
}

// SetAgeSequence register a function which accept a sequence counter and set return value to Age field
func (f *EntUserMetaFactory) SetAgeSequence(fn func(ctx context.Context, i int) (int, error)) *EntUserMetaFactory {
	f.mutation.ageSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetAgeLazy register a function which accept the build struct and set return value to Age field
func (f *EntUserMetaFactory) SetAgeLazy(fn func(ctx context.Context, i *EntUserMutator) (int, error)) *EntUserMetaFactory {
	f.mutation.ageLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetAgeDefault assign a default value to Age field
func (f *EntUserMetaFactory) SetAgeDefault(v int) *EntUserMetaFactory {
	f.mutation.ageDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetAgeFactory register a factory function and assign return value to Age, you can also use related factory's Create/CreateV as input function here
func (f *EntUserMetaFactory) SetAgeFactory(fn func(ctx context.Context) (int, error)) *EntUserMetaFactory {
	f.mutation.ageFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetAgeSequence register a function which accept a sequence counter and set return value to Age field
func (t *entUserTrait) SetAgeSequence(fn func(ctx context.Context, i int) (int, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageSequenceMutateFunc(fn))
	return t
}

// SetAgeLazy register a function which accept the build struct and set return value to Age field
func (t *entUserTrait) SetAgeLazy(fn func(ctx context.Context, i *EntUserMutator) (int, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageLazyMutateFunc(fn))
	return t
}

// SetAgeDefault assign a default value to Age field
func (t *entUserTrait) SetAgeDefault(v int) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageDefaultMutateFunc(v))
	return t
}

// SetAgeFactory register a factory function and assign return value to Age, you can also use related factory's Create/CreateV as input function here
func (t *entUserTrait) SetAgeFactory(fn func(ctx context.Context) (int, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageFactoryMutateFunc(fn))
	return t
}

func (*entUserMutation) emailSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.emailType = TypeSequence
		m.emailFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetEmail(value)

			i.Email = value
			return nil
		}
	}
}
func (*entUserMutation) emailLazyMutateFunc(fn func(ctx context.Context, i *EntUserMutator) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.emailType = TypeLazy
		m.emailFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetEmail(value)

			i.Email = value
			return nil
		}
	}
}
func (*entUserMutation) emailDefaultMutateFunc(v string) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.emailType = TypeDefault
		m.emailFunc = func(ctx context.Context, i *EntUserMutator, c int) error {

			i.EntCreator().SetEmail(v)

			i.Email = v
			return nil
		}
	}
}
func (*entUserMutation) emailFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.emailType = TypeFactory
		m.emailFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetEmail(value)

			i.Email = value

			return nil
		}
	}
}

// SetEmailSequence register a function which accept a sequence counter and set return value to Email field
func (f *EntUserMetaFactory) SetEmailSequence(fn func(ctx context.Context, i int) (string, error)) *EntUserMetaFactory {
	f.mutation.emailSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetEmailLazy register a function which accept the build struct and set return value to Email field
func (f *EntUserMetaFactory) SetEmailLazy(fn func(ctx context.Context, i *EntUserMutator) (string, error)) *EntUserMetaFactory {
	f.mutation.emailLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetEmailDefault assign a default value to Email field
func (f *EntUserMetaFactory) SetEmailDefault(v string) *EntUserMetaFactory {
	f.mutation.emailDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetEmailFactory register a factory function and assign return value to Email, you can also use related factory's Create/CreateV as input function here
func (f *EntUserMetaFactory) SetEmailFactory(fn func(ctx context.Context) (string, error)) *EntUserMetaFactory {
	f.mutation.emailFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetEmailSequence register a function which accept a sequence counter and set return value to Email field
func (t *entUserTrait) SetEmailSequence(fn func(ctx context.Context, i int) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.emailSequenceMutateFunc(fn))
	return t
}

// SetEmailLazy register a function which accept the build struct and set return value to Email field
func (t *entUserTrait) SetEmailLazy(fn func(ctx context.Context, i *EntUserMutator) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.emailLazyMutateFunc(fn))
	return t
}

// SetEmailDefault assign a default value to Email field
func (t *entUserTrait) SetEmailDefault(v string) *entUserTrait {
	t.updates = append(t.updates, t.mutation.emailDefaultMutateFunc(v))
	return t
}

// SetEmailFactory register a factory function and assign return value to Email, you can also use related factory's Create/CreateV as input function here
func (t *entUserTrait) SetEmailFactory(fn func(ctx context.Context) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.emailFactoryMutateFunc(fn))
	return t
}

func (*entUserMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetName(value)

			i.Name = value
			return nil
		}
	}
}
func (*entUserMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *EntUserMutator) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetName(value)

			i.Name = value
			return nil
		}
	}
}
func (*entUserMutation) nameDefaultMutateFunc(v string) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int) error {

			i.EntCreator().SetName(v)

			i.Name = v
			return nil
		}
	}
}
func (*entUserMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetName(value)

			i.Name = value

			return nil
		}
	}
}

// SetNameSequence register a function which accept a sequence counter and set return value to Name field
func (f *EntUserMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *EntUserMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (f *EntUserMetaFactory) SetNameLazy(fn func(ctx context.Context, i *EntUserMutator) (string, error)) *EntUserMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameDefault assign a default value to Name field
func (f *EntUserMetaFactory) SetNameDefault(v string) *EntUserMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (f *EntUserMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *EntUserMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameSequence register a function which accept a sequence counter and set return value to Name field
func (t *entUserTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (t *entUserTrait) SetNameLazy(fn func(ctx context.Context, i *EntUserMutator) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}

// SetNameDefault assign a default value to Name field
func (t *entUserTrait) SetNameDefault(v string) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (t *entUserTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

func (*entUserMutation) groupsPostMutateFunc(fn func(ctx context.Context, set bool, obj *ent.User, i int) error) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m._postGroupsFunc = fn
	}
}

// SetGroupsPostFunc register a post function which will be called in factory SetGroupsPost method
func (f *EntUserMetaFactory) SetGroupsPostFunc(fn func(ctx context.Context, set bool, obj *ent.User, i int) error) *EntUserMetaFactory {
	f.mutation.groupsPostMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entUserTrait) SetGroupsPostFunc(fn func(ctx context.Context, set bool, obj *ent.User, i int) error) *entUserTrait {
	t.updates = append(t.updates, t.mutation.groupsPostMutateFunc(fn))
	return t
}

// SetAfterCreateFunc register a function to be called after struct create
func (f *EntUserMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.User) error) *EntUserMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

// SetBeforeCreateFunc register a function to be called before struct create
func (f *EntUserMetaFactory) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntUserMutator) error) *EntUserMetaFactory {
	f.mutation.beforeCreateFunc = fn
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (t *entUserTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.User) error) *entUserTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

// SetBeforeCreateFunc register a function to be called before struct create
func (t *entUserTrait) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntUserMutator) error) *entUserTrait {
	t.updates = append(t.updates, t.mutation.beforeCreateMutateFunc(fn))
	return t
}

// Build create a  EntUserFactory from EntUserMetaFactory
func (f *EntUserMetaFactory) Build() *EntUserFactory {
	return &EntUserFactory{meta: *f, counter: &Counter{}}
}

type EntUserFactory struct {
	meta    EntUserMetaFactory
	counter *Counter

	client *ent.Client
}

// SetAge set the Age field
func (f *EntUserFactory) SetAge(i int) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetAge(i)

	builder.client = f.client

	return builder
}

// SetEmail set the Email field
func (f *EntUserFactory) SetEmail(i string) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetEmail(i)

	builder.client = f.client

	return builder
}

// SetName set the Name field
func (f *EntUserFactory) SetName(i string) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)

	builder.client = f.client

	return builder
}

// SetGroupsPost call the post function with int input
func (f *EntUserFactory) SetGroupsPost(i int) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetGroupsPost(i)

	builder.client = f.client

	return builder
}

// Create return a new *ent.User
func (f *EntUserFactory) Create(ctx context.Context) (*ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.Create(ctx)
}

// CreateV return a new ent.User
func (f *EntUserFactory) CreateV(ctx context.Context) (ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateV(ctx)
}

// CreateBatch return a []*ent.User slice
func (f *EntUserFactory) CreateBatch(ctx context.Context, n int) ([]*ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatch(ctx, n)
}

// CreateBatchV return a []ent.User slice
func (f *EntUserFactory) CreateBatchV(ctx context.Context, n int) ([]ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatchV(ctx, n)
}

// Client set ent client to EntUserFactory
func (f *EntUserFactory) Client(c *ent.Client) *EntUserFactory {
	f.client = c
	return f
}

type EntUserBuilder struct {
	factory  *EntUserFactory
	mutation entUserMutation
	counter  *Counter

	ageOverride  int
	ageOverriden bool

	emailOverride  string
	emailOverriden bool

	nameOverride  string
	nameOverriden bool

	_postGroups    int
	_postGroupsSet bool

	client *ent.Client
}

func (b *EntUserBuilder) Client(c *ent.Client) *EntUserBuilder {
	b.client = c
	return b
}

// SetAge set the Age field
func (b *EntUserBuilder) SetAge(i int) *EntUserBuilder {
	b.ageOverride = i
	b.ageOverriden = true
	return b
}

// SetEmail set the Email field
func (b *EntUserBuilder) SetEmail(i string) *EntUserBuilder {
	b.emailOverride = i
	b.emailOverriden = true
	return b
}

// SetName set the Name field
func (b *EntUserBuilder) SetName(i string) *EntUserBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

// SetGroupsPost call the post function with int input
func (b *EntUserBuilder) SetGroupsPost(i int) *EntUserBuilder {
	b._postGroups = i
	b._postGroupsSet = true
	return b
}

// CreateV return a new ent.User
func (b *EntUserBuilder) CreateV(ctx context.Context) (ent.User, error) {
	var d ent.User
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

// Create return a new *ent.User
func (b *EntUserBuilder) Create(ctx context.Context) (*ent.User, error) {

	var preSlice = []func(ctx context.Context, i *EntUserMutator, c int) error{}
	var lazySlice = []func(ctx context.Context, i *EntUserMutator, c int) error{}
	var postSlice = []func(ctx context.Context, i *ent.User, c int) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.User.Create()

	if b.ageOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntUserMutator, c int) error {
			value := b.ageOverride

			i.EntCreator().SetAge(value)

			i.Age = value
			return nil
		})
	} else {
		switch b.mutation.ageType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.ageFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.ageFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.ageFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.ageFunc)
		}
	}

	if b.emailOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntUserMutator, c int) error {
			value := b.emailOverride

			i.EntCreator().SetEmail(value)

			i.Email = value
			return nil
		})
	} else {
		switch b.mutation.emailType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.emailFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.emailFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.emailFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.emailFunc)
		}
	}

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntUserMutator, c int) error {
			value := b.nameOverride

			i.EntCreator().SetName(value)

			i.Name = value
			return nil
		})
	} else {
		switch b.mutation.nameType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.nameFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.nameFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.nameFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.nameFunc)
		}
	}

	if b.mutation._postGroupsFunc != nil {
		postSlice = append(postSlice, func(ctx context.Context, i *ent.User, c int) error {
			err := b.mutation._postGroupsFunc(ctx, b._postGroupsSet, i, b._postGroups)
			return err
		})
	}

	v := &EntUserMutator{}

	v._creator = entBuilder

	for _, f := range preSlice {

		err := f(ctx, v, index)

		if err != nil {
			return nil, err
		}
	}
	for _, f := range lazySlice {

		err := f(ctx, v, index)

		if err != nil {
			return nil, err
		}
	}
	if b.mutation.beforeCreateFunc != nil {
		if err := b.mutation.beforeCreateFunc(ctx, v); err != nil {
			return nil, err
		}
	}

	new, err := entBuilder.Save(ctx)
	if err != nil {
		return nil, err
	}

	if b.mutation.afterCreateFunc != nil {
		err := b.mutation.afterCreateFunc(ctx, new)
		if err != nil {
			return nil, err
		}
	}
	for _, f := range postSlice {
		err := f(ctx, new, index)
		if err != nil {
			return nil, err
		}
	}

	return new, nil
}
func (b *EntUserBuilder) CreateBatch(ctx context.Context, n int) ([]*ent.User, error) {
	var results []*ent.User
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *EntUserBuilder) CreateBatchV(ctx context.Context, n int) ([]ent.User, error) {
	var results []ent.User
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
