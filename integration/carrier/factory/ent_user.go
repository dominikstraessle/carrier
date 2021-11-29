// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/integration/ent"

	"context"
)

type EntUserMutator struct {
	Age int

	Name string
}

type entUserMutation struct {
	ageType int
	ageFunc func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error

	nameType int
	nameFunc func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error

	afterCreateFunc func(ctx context.Context, i *ent.User) error
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
func (*entUserMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.User) error) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entUserMutation) ageSequenceMutateFunc(fn func(ctx context.Context, i int) (int, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeSequence
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetAge(value)

			return nil
		}
	}
}
func (*entUserMutation) ageLazyMutateFunc(fn func(ctx context.Context, i *EntUserMutator) (int, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeLazy
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetAge(value)

			return nil
		}
	}
}
func (*entUserMutation) ageDefaultMutateFunc(v int) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeDefault
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {

			creator.SetAge(v)

			return nil
		}
	}
}
func (*entUserMutation) ageFactoryMutateFunc(fn func(ctx context.Context) (int, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.ageType = TypeFactory
		m.ageFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetAge(value)

			return nil
		}
	}
}

func (f *EntUserMetaFactory) SetAgeSequence(fn func(ctx context.Context, i int) (int, error)) *EntUserMetaFactory {
	f.mutation.ageSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntUserMetaFactory) SetAgeLazy(fn func(ctx context.Context, i *EntUserMutator) (int, error)) *EntUserMetaFactory {
	f.mutation.ageLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntUserMetaFactory) SetAgeDefault(v int) *EntUserMetaFactory {
	f.mutation.ageDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *EntUserMetaFactory) SetAgeFactory(fn func(ctx context.Context) (int, error)) *EntUserMetaFactory {
	f.mutation.ageFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entUserTrait) SetAgeSequence(fn func(ctx context.Context, i int) (int, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageSequenceMutateFunc(fn))
	return t
}
func (t *entUserTrait) SetAgeLazy(fn func(ctx context.Context, i *EntUserMutator) (int, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageLazyMutateFunc(fn))
	return t
}
func (t *entUserTrait) SetAgeDefault(v int) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageDefaultMutateFunc(v))
	return t
}
func (t *entUserTrait) SetAgeFactory(fn func(ctx context.Context) (int, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.ageFactoryMutateFunc(fn))
	return t
}

func (*entUserMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetName(value)

			return nil
		}
	}
}
func (*entUserMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *EntUserMutator) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetName(value)

			return nil
		}
	}
}
func (*entUserMutation) nameDefaultMutateFunc(v string) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {

			creator.SetName(v)

			return nil
		}
	}
}
func (*entUserMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entUserMutation) {
	return func(m *entUserMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetName(value)

			return nil
		}
	}
}

func (f *EntUserMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *EntUserMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntUserMetaFactory) SetNameLazy(fn func(ctx context.Context, i *EntUserMutator) (string, error)) *EntUserMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntUserMetaFactory) SetNameDefault(v string) *EntUserMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *EntUserMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *EntUserMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entUserTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}
func (t *entUserTrait) SetNameLazy(fn func(ctx context.Context, i *EntUserMutator) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}
func (t *entUserTrait) SetNameDefault(v string) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}
func (t *entUserTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *entUserTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

func (f *EntUserMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.User) error) *EntUserMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

func (f *EntUserMetaFactory) Build() *EntUserFactory {
	return &EntUserFactory{meta: *f, counter: &Counter{}}
}

type EntUserFactory struct {
	meta    EntUserMetaFactory
	counter *Counter

	client *ent.Client
}

func (f *EntUserFactory) SetAge(i int) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetAge(i)
	return builder
}

func (f *EntUserFactory) SetName(i string) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)
	return builder
}

func (f *EntUserFactory) Create(ctx context.Context) (*ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	return builder.Create(ctx)
}
func (f *EntUserFactory) CreateV(ctx context.Context) (ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	return builder.CreateV(ctx)
}
func (f *EntUserFactory) CreateBatch(ctx context.Context, n int) ([]*ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	return builder.CreateBatch(ctx, n)
}
func (f *EntUserFactory) CreateBatchV(ctx context.Context, n int) ([]ent.User, error) {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	return builder.CreateBatchV(ctx, n)
}

func (f *EntUserFactory) Client(c *ent.Client) *EntUserBuilder {
	builder := &EntUserBuilder{mutation: f.meta.mutation, counter: f.counter, client: c, factory: f}
	return builder
}

type EntUserBuilder struct {
	factory  *EntUserFactory
	mutation entUserMutation
	counter  *Counter

	ageOverride  int
	ageOverriden bool

	nameOverride  string
	nameOverriden bool

	client *ent.Client
}

func (b *EntUserBuilder) Client(c *ent.Client) *EntUserBuilder {
	b.client = c
	return b
}

func (b *EntUserBuilder) SetAge(i int) *EntUserBuilder {
	b.ageOverride = i
	b.ageOverriden = true
	return b
}

func (b *EntUserBuilder) SetName(i string) *EntUserBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

func (b *EntUserBuilder) CreateV(ctx context.Context) (ent.User, error) {
	var d ent.User
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

func (b *EntUserBuilder) Create(ctx context.Context) (*ent.User, error) {

	var preSlice = []func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error{}
	var lazySlice = []func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error{}
	var postSlice = []func(ctx context.Context, i *ent.User, c int, creator *ent.UserCreate) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.User.Create()

	if b.ageOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			value := b.ageOverride

			creator.SetAge(value)

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

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntUserMutator, c int, creator *ent.UserCreate) error {
			value := b.nameOverride

			creator.SetName(value)

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

	v := &EntUserMutator{}
	for _, f := range preSlice {

		err := f(ctx, v, index, entBuilder)

		if err != nil {
			return nil, err
		}
	}
	for _, f := range lazySlice {

		err := f(ctx, v, index, entBuilder)

		if err != nil {
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

		err := f(ctx, new, index, entBuilder)

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
