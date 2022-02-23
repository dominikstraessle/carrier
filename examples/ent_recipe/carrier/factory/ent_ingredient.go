// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent"

	"context"
)

type EntIngredientMutator struct {
	Name string

	_creator *ent.IngredientCreate
}

func (m *EntIngredientMutator) EntCreator() *ent.IngredientCreate {
	return m._creator
}

type entIngredientMutation struct {
	nameType int
	nameFunc func(ctx context.Context, i *EntIngredientMutator, c int) error

	beforeCreateFunc func(ctx context.Context, i *EntIngredientMutator) error
	afterCreateFunc  func(ctx context.Context, i *ent.Ingredient) error
}
type EntIngredientMetaFactory struct {
	mutation entIngredientMutation
}
type entIngredientTrait struct {
	mutation entIngredientMutation
	updates  []func(m *entIngredientMutation)
}

func EntIngredientTrait() *entIngredientTrait {
	return &entIngredientTrait{}
}
func (*entIngredientMutation) beforeCreateMutateFunc(fn func(ctx context.Context, i *EntIngredientMutator) error) func(m *entIngredientMutation) {
	return func(m *entIngredientMutation) {
		m.beforeCreateFunc = fn
	}
}
func (*entIngredientMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.Ingredient) error) func(m *entIngredientMutation) {
	return func(m *entIngredientMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entIngredientMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entIngredientMutation) {
	return func(m *entIngredientMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *EntIngredientMutator, c int) error {
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
func (*entIngredientMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *EntIngredientMutator) (string, error)) func(m *entIngredientMutation) {
	return func(m *entIngredientMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *EntIngredientMutator, c int) error {
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
func (*entIngredientMutation) nameDefaultMutateFunc(v string) func(m *entIngredientMutation) {
	return func(m *entIngredientMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *EntIngredientMutator, c int) error {

			i.EntCreator().SetName(v)

			i.Name = v
			return nil
		}
	}
}
func (*entIngredientMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entIngredientMutation) {
	return func(m *entIngredientMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *EntIngredientMutator, c int) error {
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
func (f *EntIngredientMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *EntIngredientMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (f *EntIngredientMetaFactory) SetNameLazy(fn func(ctx context.Context, i *EntIngredientMutator) (string, error)) *EntIngredientMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameDefault assign a default value to Name field
func (f *EntIngredientMetaFactory) SetNameDefault(v string) *EntIngredientMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (f *EntIngredientMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *EntIngredientMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameSequence register a function which accept a sequence counter and set return value to Name field
func (t *entIngredientTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *entIngredientTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (t *entIngredientTrait) SetNameLazy(fn func(ctx context.Context, i *EntIngredientMutator) (string, error)) *entIngredientTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}

// SetNameDefault assign a default value to Name field
func (t *entIngredientTrait) SetNameDefault(v string) *entIngredientTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (t *entIngredientTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *entIngredientTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

// SetAfterCreateFunc register a function to be called after struct create
func (f *EntIngredientMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Ingredient) error) *EntIngredientMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

// SetBeforeCreateFunc register a function to be called before struct create
func (f *EntIngredientMetaFactory) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntIngredientMutator) error) *EntIngredientMetaFactory {
	f.mutation.beforeCreateFunc = fn
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (t *entIngredientTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Ingredient) error) *entIngredientTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

// SetBeforeCreateFunc register a function to be called before struct create
func (t *entIngredientTrait) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntIngredientMutator) error) *entIngredientTrait {
	t.updates = append(t.updates, t.mutation.beforeCreateMutateFunc(fn))
	return t
}

// Build create a  EntIngredientFactory from EntIngredientMetaFactory
func (f *EntIngredientMetaFactory) Build() *EntIngredientFactory {
	return &EntIngredientFactory{meta: *f, counter: &Counter{}}
}

type EntIngredientFactory struct {
	meta    EntIngredientMetaFactory
	counter *Counter

	client *ent.Client
}

// SetName set the Name field
func (f *EntIngredientFactory) SetName(i string) *EntIngredientBuilder {
	builder := &EntIngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)

	builder.client = f.client

	return builder
}

// Create return a new *ent.Ingredient
func (f *EntIngredientFactory) Create(ctx context.Context) (*ent.Ingredient, error) {
	builder := &EntIngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.Create(ctx)
}

// CreateV return a new ent.Ingredient
func (f *EntIngredientFactory) CreateV(ctx context.Context) (ent.Ingredient, error) {
	builder := &EntIngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateV(ctx)
}

// CreateBatch return a []*ent.Ingredient slice
func (f *EntIngredientFactory) CreateBatch(ctx context.Context, n int) ([]*ent.Ingredient, error) {
	builder := &EntIngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatch(ctx, n)
}

// CreateBatchV return a []ent.Ingredient slice
func (f *EntIngredientFactory) CreateBatchV(ctx context.Context, n int) ([]ent.Ingredient, error) {
	builder := &EntIngredientBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatchV(ctx, n)
}

// Client set ent client to EntIngredientFactory
func (f *EntIngredientFactory) Client(c *ent.Client) *EntIngredientFactory {
	f.client = c
	return f
}

type EntIngredientBuilder struct {
	factory  *EntIngredientFactory
	mutation entIngredientMutation
	counter  *Counter

	nameOverride  string
	nameOverriden bool

	client *ent.Client
}

func (b *EntIngredientBuilder) Client(c *ent.Client) *EntIngredientBuilder {
	b.client = c
	return b
}

// SetName set the Name field
func (b *EntIngredientBuilder) SetName(i string) *EntIngredientBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

// CreateV return a new ent.Ingredient
func (b *EntIngredientBuilder) CreateV(ctx context.Context) (ent.Ingredient, error) {
	var d ent.Ingredient
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

// Create return a new *ent.Ingredient
func (b *EntIngredientBuilder) Create(ctx context.Context) (*ent.Ingredient, error) {

	var preSlice = []func(ctx context.Context, i *EntIngredientMutator, c int) error{}
	var lazySlice = []func(ctx context.Context, i *EntIngredientMutator, c int) error{}
	var postSlice = []func(ctx context.Context, i *ent.Ingredient, c int) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.Ingredient.Create()

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntIngredientMutator, c int) error {
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

	v := &EntIngredientMutator{}

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
func (b *EntIngredientBuilder) CreateBatch(ctx context.Context, n int) ([]*ent.Ingredient, error) {
	var results []*ent.Ingredient
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *EntIngredientBuilder) CreateBatchV(ctx context.Context, n int) ([]ent.Ingredient, error) {
	var results []ent.Ingredient
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}