// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"time"

	"github.com/Yiling-J/carrier/integration/ent"

	"context"
)

type EntCarMutator struct {
	Model string

	Owner *ent.User

	OwnerID int

	RegisteredAt time.Time
}

type entCarMutation struct {
	modelType int
	modelFunc func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error

	ownerType int
	ownerFunc func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error

	ownerIDType int
	ownerIDFunc func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error

	registeredAtType int
	registeredAtFunc func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error

	afterCreateFunc func(ctx context.Context, i *ent.Car) error
}
type EntCarMetaFactory struct {
	mutation entCarMutation
}
type entCarTrait struct {
	mutation entCarMutation
	updates  []func(m *entCarMutation)
}

func EntCarTrait() *entCarTrait {
	return &entCarTrait{}
}
func (*entCarMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.Car) error) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entCarMutation) modelSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.modelType = TypeSequence
		m.modelFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetModel(value)

			i.Model = value
			return nil
		}
	}
}
func (*entCarMutation) modelLazyMutateFunc(fn func(ctx context.Context, i *EntCarMutator) (string, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.modelType = TypeLazy
		m.modelFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetModel(value)

			i.Model = value
			return nil
		}
	}
}
func (*entCarMutation) modelDefaultMutateFunc(v string) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.modelType = TypeDefault
		m.modelFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {

			creator.SetModel(v)

			i.Model = v
			return nil
		}
	}
}
func (*entCarMutation) modelFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.modelType = TypeFactory
		m.modelFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetModel(value)

			i.Model = value

			return nil
		}
	}
}

func (f *EntCarMetaFactory) SetModelSequence(fn func(ctx context.Context, i int) (string, error)) *EntCarMetaFactory {
	f.mutation.modelSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetModelLazy(fn func(ctx context.Context, i *EntCarMutator) (string, error)) *EntCarMetaFactory {
	f.mutation.modelLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetModelDefault(v string) *EntCarMetaFactory {
	f.mutation.modelDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetModelFactory(fn func(ctx context.Context) (string, error)) *EntCarMetaFactory {
	f.mutation.modelFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entCarTrait) SetModelSequence(fn func(ctx context.Context, i int) (string, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.modelSequenceMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetModelLazy(fn func(ctx context.Context, i *EntCarMutator) (string, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.modelLazyMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetModelDefault(v string) *entCarTrait {
	t.updates = append(t.updates, t.mutation.modelDefaultMutateFunc(v))
	return t
}
func (t *entCarTrait) SetModelFactory(fn func(ctx context.Context) (string, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.modelFactoryMutateFunc(fn))
	return t
}

func (*entCarMutation) ownerSequenceMutateFunc(fn func(ctx context.Context, i int) (*ent.User, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerType = TypeSequence
		m.ownerFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetOwner(value)

			i.Owner = value
			return nil
		}
	}
}
func (*entCarMutation) ownerLazyMutateFunc(fn func(ctx context.Context, i *EntCarMutator) (*ent.User, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerType = TypeLazy
		m.ownerFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetOwner(value)

			i.Owner = value
			return nil
		}
	}
}
func (*entCarMutation) ownerDefaultMutateFunc(v *ent.User) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerType = TypeDefault
		m.ownerFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {

			creator.SetOwner(v)

			i.Owner = v
			return nil
		}
	}
}
func (*entCarMutation) ownerFactoryMutateFunc(fn func(ctx context.Context) (*ent.User, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerType = TypeFactory
		m.ownerFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetOwner(value)

			i.Owner = value

			return nil
		}
	}
}

func (f *EntCarMetaFactory) SetOwnerSequence(fn func(ctx context.Context, i int) (*ent.User, error)) *EntCarMetaFactory {
	f.mutation.ownerSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetOwnerLazy(fn func(ctx context.Context, i *EntCarMutator) (*ent.User, error)) *EntCarMetaFactory {
	f.mutation.ownerLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetOwnerDefault(v *ent.User) *EntCarMetaFactory {
	f.mutation.ownerDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetOwnerFactory(fn func(ctx context.Context) (*ent.User, error)) *EntCarMetaFactory {
	f.mutation.ownerFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entCarTrait) SetOwnerSequence(fn func(ctx context.Context, i int) (*ent.User, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerSequenceMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetOwnerLazy(fn func(ctx context.Context, i *EntCarMutator) (*ent.User, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerLazyMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetOwnerDefault(v *ent.User) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerDefaultMutateFunc(v))
	return t
}
func (t *entCarTrait) SetOwnerFactory(fn func(ctx context.Context) (*ent.User, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerFactoryMutateFunc(fn))
	return t
}

func (*entCarMutation) ownerIDSequenceMutateFunc(fn func(ctx context.Context, i int) (int, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerIDType = TypeSequence
		m.ownerIDFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetOwnerID(value)

			i.OwnerID = value
			return nil
		}
	}
}
func (*entCarMutation) ownerIDLazyMutateFunc(fn func(ctx context.Context, i *EntCarMutator) (int, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerIDType = TypeLazy
		m.ownerIDFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetOwnerID(value)

			i.OwnerID = value
			return nil
		}
	}
}
func (*entCarMutation) ownerIDDefaultMutateFunc(v int) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerIDType = TypeDefault
		m.ownerIDFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {

			creator.SetOwnerID(v)

			i.OwnerID = v
			return nil
		}
	}
}
func (*entCarMutation) ownerIDFactoryMutateFunc(fn func(ctx context.Context) (int, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.ownerIDType = TypeFactory
		m.ownerIDFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetOwnerID(value)

			i.OwnerID = value

			return nil
		}
	}
}

func (f *EntCarMetaFactory) SetOwnerIDSequence(fn func(ctx context.Context, i int) (int, error)) *EntCarMetaFactory {
	f.mutation.ownerIDSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetOwnerIDLazy(fn func(ctx context.Context, i *EntCarMutator) (int, error)) *EntCarMetaFactory {
	f.mutation.ownerIDLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetOwnerIDDefault(v int) *EntCarMetaFactory {
	f.mutation.ownerIDDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetOwnerIDFactory(fn func(ctx context.Context) (int, error)) *EntCarMetaFactory {
	f.mutation.ownerIDFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entCarTrait) SetOwnerIDSequence(fn func(ctx context.Context, i int) (int, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerIDSequenceMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetOwnerIDLazy(fn func(ctx context.Context, i *EntCarMutator) (int, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerIDLazyMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetOwnerIDDefault(v int) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerIDDefaultMutateFunc(v))
	return t
}
func (t *entCarTrait) SetOwnerIDFactory(fn func(ctx context.Context) (int, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.ownerIDFactoryMutateFunc(fn))
	return t
}

func (*entCarMutation) registeredAtSequenceMutateFunc(fn func(ctx context.Context, i int) (time.Time, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.registeredAtType = TypeSequence
		m.registeredAtFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			creator.SetRegisteredAt(value)

			i.RegisteredAt = value
			return nil
		}
	}
}
func (*entCarMutation) registeredAtLazyMutateFunc(fn func(ctx context.Context, i *EntCarMutator) (time.Time, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.registeredAtType = TypeLazy
		m.registeredAtFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			creator.SetRegisteredAt(value)

			i.RegisteredAt = value
			return nil
		}
	}
}
func (*entCarMutation) registeredAtDefaultMutateFunc(v time.Time) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.registeredAtType = TypeDefault
		m.registeredAtFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {

			creator.SetRegisteredAt(v)

			i.RegisteredAt = v
			return nil
		}
	}
}
func (*entCarMutation) registeredAtFactoryMutateFunc(fn func(ctx context.Context) (time.Time, error)) func(m *entCarMutation) {
	return func(m *entCarMutation) {
		m.registeredAtType = TypeFactory
		m.registeredAtFunc = func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			creator.SetRegisteredAt(value)

			i.RegisteredAt = value

			return nil
		}
	}
}

func (f *EntCarMetaFactory) SetRegisteredAtSequence(fn func(ctx context.Context, i int) (time.Time, error)) *EntCarMetaFactory {
	f.mutation.registeredAtSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetRegisteredAtLazy(fn func(ctx context.Context, i *EntCarMutator) (time.Time, error)) *EntCarMetaFactory {
	f.mutation.registeredAtLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetRegisteredAtDefault(v time.Time) *EntCarMetaFactory {
	f.mutation.registeredAtDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *EntCarMetaFactory) SetRegisteredAtFactory(fn func(ctx context.Context) (time.Time, error)) *EntCarMetaFactory {
	f.mutation.registeredAtFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entCarTrait) SetRegisteredAtSequence(fn func(ctx context.Context, i int) (time.Time, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.registeredAtSequenceMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetRegisteredAtLazy(fn func(ctx context.Context, i *EntCarMutator) (time.Time, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.registeredAtLazyMutateFunc(fn))
	return t
}
func (t *entCarTrait) SetRegisteredAtDefault(v time.Time) *entCarTrait {
	t.updates = append(t.updates, t.mutation.registeredAtDefaultMutateFunc(v))
	return t
}
func (t *entCarTrait) SetRegisteredAtFactory(fn func(ctx context.Context) (time.Time, error)) *entCarTrait {
	t.updates = append(t.updates, t.mutation.registeredAtFactoryMutateFunc(fn))
	return t
}

func (f *EntCarMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Car) error) *EntCarMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}
func (t *entCarTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Car) error) *entCarTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

func (f *EntCarMetaFactory) Build() *EntCarFactory {
	return &EntCarFactory{meta: *f, counter: &Counter{}}
}

type EntCarFactory struct {
	meta    EntCarMetaFactory
	counter *Counter

	client *ent.Client
}

func (f *EntCarFactory) SetModel(i string) *EntCarBuilder {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetModel(i)

	builder.client = f.client

	return builder
}

func (f *EntCarFactory) SetOwner(i *ent.User) *EntCarBuilder {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetOwner(i)

	builder.client = f.client

	return builder
}

func (f *EntCarFactory) SetOwnerID(i int) *EntCarBuilder {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetOwnerID(i)

	builder.client = f.client

	return builder
}

func (f *EntCarFactory) SetRegisteredAt(i time.Time) *EntCarBuilder {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetRegisteredAt(i)

	builder.client = f.client

	return builder
}

func (f *EntCarFactory) Create(ctx context.Context) (*ent.Car, error) {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.Create(ctx)
}
func (f *EntCarFactory) CreateV(ctx context.Context) (ent.Car, error) {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateV(ctx)
}
func (f *EntCarFactory) CreateBatch(ctx context.Context, n int) ([]*ent.Car, error) {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatch(ctx, n)
}
func (f *EntCarFactory) CreateBatchV(ctx context.Context, n int) ([]ent.Car, error) {
	builder := &EntCarBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatchV(ctx, n)
}

func (f *EntCarFactory) Client(c *ent.Client) *EntCarFactory {
	f.client = c
	return f
}

type EntCarBuilder struct {
	factory  *EntCarFactory
	mutation entCarMutation
	counter  *Counter

	modelOverride  string
	modelOverriden bool

	ownerOverride  *ent.User
	ownerOverriden bool

	ownerIDOverride  int
	ownerIDOverriden bool

	registeredAtOverride  time.Time
	registeredAtOverriden bool

	client *ent.Client
}

func (b *EntCarBuilder) Client(c *ent.Client) *EntCarBuilder {
	b.client = c
	return b
}

func (b *EntCarBuilder) SetModel(i string) *EntCarBuilder {
	b.modelOverride = i
	b.modelOverriden = true
	return b
}

func (b *EntCarBuilder) SetOwner(i *ent.User) *EntCarBuilder {
	b.ownerOverride = i
	b.ownerOverriden = true
	return b
}

func (b *EntCarBuilder) SetOwnerID(i int) *EntCarBuilder {
	b.ownerIDOverride = i
	b.ownerIDOverriden = true
	return b
}

func (b *EntCarBuilder) SetRegisteredAt(i time.Time) *EntCarBuilder {
	b.registeredAtOverride = i
	b.registeredAtOverriden = true
	return b
}

func (b *EntCarBuilder) CreateV(ctx context.Context) (ent.Car, error) {
	var d ent.Car
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

func (b *EntCarBuilder) Create(ctx context.Context) (*ent.Car, error) {

	var preSlice = []func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error{}
	var lazySlice = []func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error{}
	var postSlice = []func(ctx context.Context, i *ent.Car, c int, creator *ent.CarCreate) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.Car.Create()

	if b.modelOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			value := b.modelOverride

			creator.SetModel(value)

			i.Model = value
			return nil
		})
	} else {
		switch b.mutation.modelType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.modelFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.modelFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.modelFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.modelFunc)
		}
	}

	if b.ownerOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			value := b.ownerOverride

			creator.SetOwner(value)

			i.Owner = value
			return nil
		})
	} else {
		switch b.mutation.ownerType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.ownerFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.ownerFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.ownerFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.ownerFunc)
		}
	}

	if b.ownerIDOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			value := b.ownerIDOverride

			creator.SetOwnerID(value)

			i.OwnerID = value
			return nil
		})
	} else {
		switch b.mutation.ownerIDType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.ownerIDFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.ownerIDFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.ownerIDFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.ownerIDFunc)
		}
	}

	if b.registeredAtOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntCarMutator, c int, creator *ent.CarCreate) error {
			value := b.registeredAtOverride

			creator.SetRegisteredAt(value)

			i.RegisteredAt = value
			return nil
		})
	} else {
		switch b.mutation.registeredAtType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.registeredAtFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.registeredAtFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.registeredAtFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.registeredAtFunc)
		}
	}

	v := &EntCarMutator{}
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
func (b *EntCarBuilder) CreateBatch(ctx context.Context, n int) ([]*ent.Car, error) {
	var results []*ent.Car
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *EntCarBuilder) CreateBatchV(ctx context.Context, n int) ([]ent.Car, error) {
	var results []ent.Car
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
