// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/integration/model"

	"context"
)

type foodMutation struct {
	fooType int
	fooFunc func(ctx context.Context, i *model.Food, c int) error

	categoryType int
	categoryFunc func(ctx context.Context, i *model.Food, c int) error

	afterCreateFunc func(ctx context.Context, i *model.Food) error
}
type FoodMetaFactory struct {
	mutation foodMutation
}
type foodTrait struct {
	mutation foodMutation
	updates  []func(m *foodMutation)
}

func FoodTrait() *foodTrait {
	return &foodTrait{}
}
func (*foodMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *model.Food) error) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.afterCreateFunc = fn
	}
}

func (*foodMutation) fooSequenceMutateFunc(fn func(ctx context.Context, i int) (model.Foo, error)) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.fooType = TypeSequence
		m.fooFunc = func(ctx context.Context, i *model.Food, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Foo = value
			return nil
		}
	}
}
func (*foodMutation) fooLazyMutateFunc(fn func(ctx context.Context, i *model.Food) (model.Foo, error)) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.fooType = TypeLazy
		m.fooFunc = func(ctx context.Context, i *model.Food, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Foo = value
			return nil
		}
	}
}
func (*foodMutation) fooDefaultMutateFunc(v model.Foo) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.fooType = TypeDefault
		m.fooFunc = func(ctx context.Context, i *model.Food, c int) error {

			i.Foo = v
			return nil
		}
	}
}
func (*foodMutation) fooFactoryMutateFunc(fn func(ctx context.Context) (model.Foo, error)) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.fooType = TypeFactory
		m.fooFunc = func(ctx context.Context, i *model.Food, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Foo = value

			return nil
		}
	}
}

func (f *FoodMetaFactory) SetFooSequence(fn func(ctx context.Context, i int) (model.Foo, error)) *FoodMetaFactory {
	f.mutation.fooSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *FoodMetaFactory) SetFooLazy(fn func(ctx context.Context, i *model.Food) (model.Foo, error)) *FoodMetaFactory {
	f.mutation.fooLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *FoodMetaFactory) SetFooDefault(v model.Foo) *FoodMetaFactory {
	f.mutation.fooDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *FoodMetaFactory) SetFooFactory(fn func(ctx context.Context) (model.Foo, error)) *FoodMetaFactory {
	f.mutation.fooFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *foodTrait) SetFooSequence(fn func(ctx context.Context, i int) (model.Foo, error)) *foodTrait {
	t.updates = append(t.updates, t.mutation.fooSequenceMutateFunc(fn))
	return t
}
func (t *foodTrait) SetFooLazy(fn func(ctx context.Context, i *model.Food) (model.Foo, error)) *foodTrait {
	t.updates = append(t.updates, t.mutation.fooLazyMutateFunc(fn))
	return t
}
func (t *foodTrait) SetFooDefault(v model.Foo) *foodTrait {
	t.updates = append(t.updates, t.mutation.fooDefaultMutateFunc(v))
	return t
}
func (t *foodTrait) SetFooFactory(fn func(ctx context.Context) (model.Foo, error)) *foodTrait {
	t.updates = append(t.updates, t.mutation.fooFactoryMutateFunc(fn))
	return t
}

func (*foodMutation) categorySequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.categoryType = TypeSequence
		m.categoryFunc = func(ctx context.Context, i *model.Food, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Category = value
			return nil
		}
	}
}
func (*foodMutation) categoryLazyMutateFunc(fn func(ctx context.Context, i *model.Food) (string, error)) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.categoryType = TypeLazy
		m.categoryFunc = func(ctx context.Context, i *model.Food, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Category = value
			return nil
		}
	}
}
func (*foodMutation) categoryDefaultMutateFunc(v string) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.categoryType = TypeDefault
		m.categoryFunc = func(ctx context.Context, i *model.Food, c int) error {

			i.Category = v
			return nil
		}
	}
}
func (*foodMutation) categoryFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *foodMutation) {
	return func(m *foodMutation) {
		m.categoryType = TypeFactory
		m.categoryFunc = func(ctx context.Context, i *model.Food, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Category = value

			return nil
		}
	}
}

func (f *FoodMetaFactory) SetCategorySequence(fn func(ctx context.Context, i int) (string, error)) *FoodMetaFactory {
	f.mutation.categorySequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *FoodMetaFactory) SetCategoryLazy(fn func(ctx context.Context, i *model.Food) (string, error)) *FoodMetaFactory {
	f.mutation.categoryLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *FoodMetaFactory) SetCategoryDefault(v string) *FoodMetaFactory {
	f.mutation.categoryDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *FoodMetaFactory) SetCategoryFactory(fn func(ctx context.Context) (string, error)) *FoodMetaFactory {
	f.mutation.categoryFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *foodTrait) SetCategorySequence(fn func(ctx context.Context, i int) (string, error)) *foodTrait {
	t.updates = append(t.updates, t.mutation.categorySequenceMutateFunc(fn))
	return t
}
func (t *foodTrait) SetCategoryLazy(fn func(ctx context.Context, i *model.Food) (string, error)) *foodTrait {
	t.updates = append(t.updates, t.mutation.categoryLazyMutateFunc(fn))
	return t
}
func (t *foodTrait) SetCategoryDefault(v string) *foodTrait {
	t.updates = append(t.updates, t.mutation.categoryDefaultMutateFunc(v))
	return t
}
func (t *foodTrait) SetCategoryFactory(fn func(ctx context.Context) (string, error)) *foodTrait {
	t.updates = append(t.updates, t.mutation.categoryFactoryMutateFunc(fn))
	return t
}

func (f *FoodMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *model.Food) error) *FoodMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}
func (t *foodTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *model.Food) error) *foodTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

func (f *FoodMetaFactory) Build() *FoodFactory {
	return &FoodFactory{meta: *f, counter: &Counter{}}
}

type FoodFactory struct {
	meta    FoodMetaFactory
	counter *Counter
}

func (f *FoodFactory) SetFoo(i model.Foo) *FoodBuilder {
	builder := &FoodBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetFoo(i)

	return builder
}

func (f *FoodFactory) SetCategory(i string) *FoodBuilder {
	builder := &FoodBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetCategory(i)

	return builder
}

func (f *FoodFactory) Create(ctx context.Context) (*model.Food, error) {
	builder := &FoodBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.Create(ctx)
}
func (f *FoodFactory) CreateV(ctx context.Context) (model.Food, error) {
	builder := &FoodBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateV(ctx)
}
func (f *FoodFactory) CreateBatch(ctx context.Context, n int) ([]*model.Food, error) {
	builder := &FoodBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatch(ctx, n)
}
func (f *FoodFactory) CreateBatchV(ctx context.Context, n int) ([]model.Food, error) {
	builder := &FoodBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatchV(ctx, n)
}

type FoodBuilder struct {
	factory  *FoodFactory
	mutation foodMutation
	counter  *Counter

	fooOverride  model.Foo
	fooOverriden bool

	categoryOverride  string
	categoryOverriden bool
}

func (b *FoodBuilder) SetFoo(i model.Foo) *FoodBuilder {
	b.fooOverride = i
	b.fooOverriden = true
	return b
}

func (b *FoodBuilder) SetCategory(i string) *FoodBuilder {
	b.categoryOverride = i
	b.categoryOverriden = true
	return b
}

func (b *FoodBuilder) CreateV(ctx context.Context) (model.Food, error) {
	var d model.Food
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

func (b *FoodBuilder) Create(ctx context.Context) (*model.Food, error) {

	var preSlice = []func(ctx context.Context, i *model.Food, c int) error{}
	var lazySlice = []func(ctx context.Context, i *model.Food, c int) error{}
	var postSlice = []func(ctx context.Context, i *model.Food, c int) error{}

	index := b.counter.Get()
	_ = index

	if b.fooOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.Food, c int) error {
			value := b.fooOverride

			i.Foo = value
			return nil
		})
	} else {
		switch b.mutation.fooType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.fooFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.fooFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.fooFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.fooFunc)
		}
	}

	if b.categoryOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.Food, c int) error {
			value := b.categoryOverride

			i.Category = value
			return nil
		})
	} else {
		switch b.mutation.categoryType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.categoryFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.categoryFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.categoryFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.categoryFunc)
		}
	}

	v := &model.Food{}
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

	new := v

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
func (b *FoodBuilder) CreateBatch(ctx context.Context, n int) ([]*model.Food, error) {
	var results []*model.Food
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *FoodBuilder) CreateBatchV(ctx context.Context, n int) ([]model.Food, error) {
	var results []model.Food
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}