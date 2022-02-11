// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/integration/ent"

	"context"
)

type EntGroupMutator struct {
	Name string
}

type entGroupMutation struct {
	nameType int
	nameFunc func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error

	beforeCreateFunc func(ctx context.Context, creator *ent.GroupCreate) error
	afterCreateFunc  func(ctx context.Context, i *ent.Group) error
}
type EntGroupMetaFactory struct {
	mutation entGroupMutation

	nouserTrait *entGroupTrait
}
type entGroupTrait struct {
	mutation entGroupMutation
	updates  []func(m *entGroupMutation)
}

func EntGroupTrait() *entGroupTrait {
	return &entGroupTrait{}
}
func (*entGroupMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.Group) error) func(m *entGroupMutation) {
	return func(m *entGroupMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entGroupMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entGroupMutation) {
	return func(m *entGroupMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetName(value)

			i.Name = value
			return nil
		}
	}
}
func (*entGroupMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *EntGroupMutator) (string, error)) func(m *entGroupMutation) {
	return func(m *entGroupMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetName(value)

			i.Name = value
			return nil
		}
	}
}
func (*entGroupMutation) nameDefaultMutateFunc(v string) func(m *entGroupMutation) {
	return func(m *entGroupMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error {

			creator.SetName(v)

			i.Name = v
			return nil
		}
	}
}
func (*entGroupMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entGroupMutation) {
	return func(m *entGroupMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetName(value)

			i.Name = value

			return nil
		}
	}
}

// SetNameSequence register a function which accept a sequence counter and set return value to Name field
func (f *EntGroupMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *EntGroupMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (f *EntGroupMetaFactory) SetNameLazy(fn func(ctx context.Context, i *EntGroupMutator) (string, error)) *EntGroupMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameDefault assign a default value to Name field
func (f *EntGroupMetaFactory) SetNameDefault(v string) *EntGroupMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (f *EntGroupMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *EntGroupMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameSequence register a function which accept a sequence counter and set return value to Name field
func (t *entGroupTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *entGroupTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (t *entGroupTrait) SetNameLazy(fn func(ctx context.Context, i *EntGroupMutator) (string, error)) *entGroupTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}

// SetNameDefault assign a default value to Name field
func (t *entGroupTrait) SetNameDefault(v string) *entGroupTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (t *entGroupTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *entGroupTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

// SetNouserTrait accept a entGroupTrait, will override builder using Trait's methods if enable
func (f *EntGroupMetaFactory) SetNouserTrait(t *entGroupTrait) *EntGroupMetaFactory {
	f.nouserTrait = t
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (f *EntGroupMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Group) error) *EntGroupMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

// SetBeforeCreateFunc register a function to be called after struct create
func (f *EntGroupMetaFactory) SetBeforeCreateFunc(fn func(ctx context.Context, creator *ent.GroupCreate) error) *EntGroupMetaFactory {
	f.mutation.beforeCreateFunc = fn
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (t *entGroupTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Group) error) *entGroupTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

// Build create a  EntGroupFactory from EntGroupMetaFactory
func (f *EntGroupMetaFactory) Build() *EntGroupFactory {
	return &EntGroupFactory{meta: *f, counter: &Counter{}}
}

type EntGroupFactory struct {
	meta    EntGroupMetaFactory
	counter *Counter

	client *ent.Client
}

// SetName set the Name field
func (f *EntGroupFactory) SetName(i string) *EntGroupBuilder {
	builder := &EntGroupBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)

	builder.client = f.client

	return builder
}

// WithNouserTrait() enable the Nouser trait
func (f *EntGroupFactory) WithNouserTrait() *EntGroupBuilder {
	builder := &EntGroupBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	builder.client = f.client

	if f.meta.nouserTrait == nil {
		return builder
	}
	for _, u := range f.meta.nouserTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

// Create return a new *ent.Group
func (f *EntGroupFactory) Create(ctx context.Context) (*ent.Group, error) {
	builder := &EntGroupBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.Create(ctx)
}

// CreateV return a new ent.Group
func (f *EntGroupFactory) CreateV(ctx context.Context) (ent.Group, error) {
	builder := &EntGroupBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateV(ctx)
}

// CreateBatch return a []*ent.Group slice
func (f *EntGroupFactory) CreateBatch(ctx context.Context, n int) ([]*ent.Group, error) {
	builder := &EntGroupBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatch(ctx, n)
}

// CreateBatchV return a []ent.Group slice
func (f *EntGroupFactory) CreateBatchV(ctx context.Context, n int) ([]ent.Group, error) {
	builder := &EntGroupBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatchV(ctx, n)
}

// Client set ent client to EntGroupFactory
func (f *EntGroupFactory) Client(c *ent.Client) *EntGroupFactory {
	f.client = c
	return f
}

type EntGroupBuilder struct {
	factory  *EntGroupFactory
	mutation entGroupMutation
	counter  *Counter

	nameOverride  string
	nameOverriden bool

	client *ent.Client
}

func (b *EntGroupBuilder) Client(c *ent.Client) *EntGroupBuilder {
	b.client = c
	return b
}

// SetName set the Name field
func (b *EntGroupBuilder) SetName(i string) *EntGroupBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

// WithNouserTrait() enable the Nouser trait
func (b *EntGroupBuilder) WithNouserTrait() *EntGroupBuilder {
	if b.factory.meta.nouserTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.nouserTrait.updates {
		u(&b.mutation)
	}
	return b
}

// CreateV return a new ent.Group
func (b *EntGroupBuilder) CreateV(ctx context.Context) (ent.Group, error) {
	var d ent.Group
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

// Create return a new *ent.Group
func (b *EntGroupBuilder) Create(ctx context.Context) (*ent.Group, error) {

	var preSlice = []func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error{}
	var lazySlice = []func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error{}
	var postSlice = []func(ctx context.Context, i *ent.Group, c int, creator *ent.GroupCreate) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.Group.Create()

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntGroupMutator, c int, creator *ent.GroupCreate) error {
			value := b.nameOverride

			creator.SetName(value)

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

	v := &EntGroupMutator{}
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

	if b.mutation.beforeCreateFunc != nil {
		if err := b.mutation.beforeCreateFunc(ctx, entBuilder); err != nil {
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
func (b *EntGroupBuilder) CreateBatch(ctx context.Context, n int) ([]*ent.Group, error) {
	var results []*ent.Group
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *EntGroupBuilder) CreateBatchV(ctx context.Context, n int) ([]ent.Group, error) {
	var results []ent.Group
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
