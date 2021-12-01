// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/integration/model"

	"context"
)

type userMutation struct {
	nameType int
	nameFunc func(ctx context.Context, i *model.User, c int) error

	emailType int
	emailFunc func(ctx context.Context, i *model.User, c int) error

	titleType int
	titleFunc func(ctx context.Context, i *model.User, c int) error

	groupType int
	groupFunc func(ctx context.Context, i *model.User, c int) error

	_postFooFunc func(ctx context.Context, set bool, obj *model.User, i string) error

	afterCreateFunc func(ctx context.Context, i *model.User) error
}
type UserMetaFactory struct {
	mutation userMutation

	defaultTrait *userTrait

	lazyTrait *userTrait

	sequenceTrait *userTrait

	factoryTrait *userTrait

	anonymousTrait *userTrait

	nilTrait *userTrait

	mixnameTrait *userTrait

	mixemailTrait *userTrait

	mixtitleTrait *userTrait
}
type userTrait struct {
	mutation userMutation
	updates  []func(m *userMutation)
}

func UserTrait() *userTrait {
	return &userTrait{}
}
func (*userMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *model.User) error) func(m *userMutation) {
	return func(m *userMutation) {
		m.afterCreateFunc = fn
	}
}

func (*userMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Name = value
			return nil
		}
	}
}
func (*userMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *model.User) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Name = value
			return nil
		}
	}
}
func (*userMutation) nameDefaultMutateFunc(v string) func(m *userMutation) {
	return func(m *userMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *model.User, c int) error {

			i.Name = v
			return nil
		}
	}
}
func (*userMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Name = value

			return nil
		}
	}
}

func (f *UserMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *UserMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetNameLazy(fn func(ctx context.Context, i *model.User) (string, error)) *UserMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetNameDefault(v string) *UserMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *UserMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *userTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}
func (t *userTrait) SetNameLazy(fn func(ctx context.Context, i *model.User) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}
func (t *userTrait) SetNameDefault(v string) *userTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}
func (t *userTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

func (*userMutation) emailSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.emailType = TypeSequence
		m.emailFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Email = value
			return nil
		}
	}
}
func (*userMutation) emailLazyMutateFunc(fn func(ctx context.Context, i *model.User) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.emailType = TypeLazy
		m.emailFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Email = value
			return nil
		}
	}
}
func (*userMutation) emailDefaultMutateFunc(v string) func(m *userMutation) {
	return func(m *userMutation) {
		m.emailType = TypeDefault
		m.emailFunc = func(ctx context.Context, i *model.User, c int) error {

			i.Email = v
			return nil
		}
	}
}
func (*userMutation) emailFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.emailType = TypeFactory
		m.emailFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Email = value

			return nil
		}
	}
}

func (f *UserMetaFactory) SetEmailSequence(fn func(ctx context.Context, i int) (string, error)) *UserMetaFactory {
	f.mutation.emailSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetEmailLazy(fn func(ctx context.Context, i *model.User) (string, error)) *UserMetaFactory {
	f.mutation.emailLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetEmailDefault(v string) *UserMetaFactory {
	f.mutation.emailDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetEmailFactory(fn func(ctx context.Context) (string, error)) *UserMetaFactory {
	f.mutation.emailFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *userTrait) SetEmailSequence(fn func(ctx context.Context, i int) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.emailSequenceMutateFunc(fn))
	return t
}
func (t *userTrait) SetEmailLazy(fn func(ctx context.Context, i *model.User) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.emailLazyMutateFunc(fn))
	return t
}
func (t *userTrait) SetEmailDefault(v string) *userTrait {
	t.updates = append(t.updates, t.mutation.emailDefaultMutateFunc(v))
	return t
}
func (t *userTrait) SetEmailFactory(fn func(ctx context.Context) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.emailFactoryMutateFunc(fn))
	return t
}

func (*userMutation) titleSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.titleType = TypeSequence
		m.titleFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Title = value
			return nil
		}
	}
}
func (*userMutation) titleLazyMutateFunc(fn func(ctx context.Context, i *model.User) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.titleType = TypeLazy
		m.titleFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Title = value
			return nil
		}
	}
}
func (*userMutation) titleDefaultMutateFunc(v string) func(m *userMutation) {
	return func(m *userMutation) {
		m.titleType = TypeDefault
		m.titleFunc = func(ctx context.Context, i *model.User, c int) error {

			i.Title = v
			return nil
		}
	}
}
func (*userMutation) titleFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.titleType = TypeFactory
		m.titleFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Title = value

			return nil
		}
	}
}

func (f *UserMetaFactory) SetTitleSequence(fn func(ctx context.Context, i int) (string, error)) *UserMetaFactory {
	f.mutation.titleSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetTitleLazy(fn func(ctx context.Context, i *model.User) (string, error)) *UserMetaFactory {
	f.mutation.titleLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetTitleDefault(v string) *UserMetaFactory {
	f.mutation.titleDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetTitleFactory(fn func(ctx context.Context) (string, error)) *UserMetaFactory {
	f.mutation.titleFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *userTrait) SetTitleSequence(fn func(ctx context.Context, i int) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.titleSequenceMutateFunc(fn))
	return t
}
func (t *userTrait) SetTitleLazy(fn func(ctx context.Context, i *model.User) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.titleLazyMutateFunc(fn))
	return t
}
func (t *userTrait) SetTitleDefault(v string) *userTrait {
	t.updates = append(t.updates, t.mutation.titleDefaultMutateFunc(v))
	return t
}
func (t *userTrait) SetTitleFactory(fn func(ctx context.Context) (string, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.titleFactoryMutateFunc(fn))
	return t
}

func (*userMutation) groupSequenceMutateFunc(fn func(ctx context.Context, i int) (*model.Group, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.groupType = TypeSequence
		m.groupFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.Group = value
			return nil
		}
	}
}
func (*userMutation) groupLazyMutateFunc(fn func(ctx context.Context, i *model.User) (*model.Group, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.groupType = TypeLazy
		m.groupFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.Group = value
			return nil
		}
	}
}
func (*userMutation) groupDefaultMutateFunc(v *model.Group) func(m *userMutation) {
	return func(m *userMutation) {
		m.groupType = TypeDefault
		m.groupFunc = func(ctx context.Context, i *model.User, c int) error {

			i.Group = v
			return nil
		}
	}
}
func (*userMutation) groupFactoryMutateFunc(fn func(ctx context.Context) (*model.Group, error)) func(m *userMutation) {
	return func(m *userMutation) {
		m.groupType = TypeFactory
		m.groupFunc = func(ctx context.Context, i *model.User, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.Group = value

			return nil
		}
	}
}

func (f *UserMetaFactory) SetGroupSequence(fn func(ctx context.Context, i int) (*model.Group, error)) *UserMetaFactory {
	f.mutation.groupSequenceMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetGroupLazy(fn func(ctx context.Context, i *model.User) (*model.Group, error)) *UserMetaFactory {
	f.mutation.groupLazyMutateFunc(fn)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetGroupDefault(v *model.Group) *UserMetaFactory {
	f.mutation.groupDefaultMutateFunc(v)(&f.mutation)
	return f
}
func (f *UserMetaFactory) SetGroupFactory(fn func(ctx context.Context) (*model.Group, error)) *UserMetaFactory {
	f.mutation.groupFactoryMutateFunc(fn)(&f.mutation)
	return f
}
func (t *userTrait) SetGroupSequence(fn func(ctx context.Context, i int) (*model.Group, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.groupSequenceMutateFunc(fn))
	return t
}
func (t *userTrait) SetGroupLazy(fn func(ctx context.Context, i *model.User) (*model.Group, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.groupLazyMutateFunc(fn))
	return t
}
func (t *userTrait) SetGroupDefault(v *model.Group) *userTrait {
	t.updates = append(t.updates, t.mutation.groupDefaultMutateFunc(v))
	return t
}
func (t *userTrait) SetGroupFactory(fn func(ctx context.Context) (*model.Group, error)) *userTrait {
	t.updates = append(t.updates, t.mutation.groupFactoryMutateFunc(fn))
	return t
}

func (*userMutation) fooPostMutateFunc(fn func(ctx context.Context, set bool, obj *model.User, i string) error) func(m *userMutation) {
	return func(m *userMutation) {
		m._postFooFunc = fn
	}
}
func (f *UserMetaFactory) SetFooPostFunc(fn func(ctx context.Context, set bool, obj *model.User, i string) error) *UserMetaFactory {
	f.mutation.fooPostMutateFunc(fn)(&f.mutation)
	return f
}
func (t *userTrait) SetFooPostFunc(fn func(ctx context.Context, set bool, obj *model.User, i string) error) *userTrait {
	t.updates = append(t.updates, t.mutation.fooPostMutateFunc(fn))
	return t
}

func (f *UserMetaFactory) SetDefaultTrait(t *userTrait) *UserMetaFactory {
	f.defaultTrait = t
	return f
}

func (f *UserMetaFactory) SetLazyTrait(t *userTrait) *UserMetaFactory {
	f.lazyTrait = t
	return f
}

func (f *UserMetaFactory) SetSequenceTrait(t *userTrait) *UserMetaFactory {
	f.sequenceTrait = t
	return f
}

func (f *UserMetaFactory) SetFactoryTrait(t *userTrait) *UserMetaFactory {
	f.factoryTrait = t
	return f
}

func (f *UserMetaFactory) SetAnonymousTrait(t *userTrait) *UserMetaFactory {
	f.anonymousTrait = t
	return f
}

func (f *UserMetaFactory) SetNilTrait(t *userTrait) *UserMetaFactory {
	f.nilTrait = t
	return f
}

func (f *UserMetaFactory) SetMixnameTrait(t *userTrait) *UserMetaFactory {
	f.mixnameTrait = t
	return f
}

func (f *UserMetaFactory) SetMixemailTrait(t *userTrait) *UserMetaFactory {
	f.mixemailTrait = t
	return f
}

func (f *UserMetaFactory) SetMixtitleTrait(t *userTrait) *UserMetaFactory {
	f.mixtitleTrait = t
	return f
}

func (f *UserMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *model.User) error) *UserMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}
func (t *userTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *model.User) error) *userTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

func (f *UserMetaFactory) Build() *UserFactory {
	return &UserFactory{meta: *f, counter: &Counter{}}
}

type UserFactory struct {
	meta    UserMetaFactory
	counter *Counter
}

func (f *UserFactory) SetName(i string) *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)

	return builder
}

func (f *UserFactory) SetEmail(i string) *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetEmail(i)

	return builder
}

func (f *UserFactory) SetTitle(i string) *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetTitle(i)

	return builder
}

func (f *UserFactory) SetGroup(i *model.Group) *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetGroup(i)

	return builder
}

func (f *UserFactory) SetFooPost(i string) *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetFooPost(i)

	return builder
}

func (f *UserFactory) WithDefaultTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.defaultTrait == nil {
		return builder
	}
	for _, u := range f.meta.defaultTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithLazyTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.lazyTrait == nil {
		return builder
	}
	for _, u := range f.meta.lazyTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithSequenceTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.sequenceTrait == nil {
		return builder
	}
	for _, u := range f.meta.sequenceTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithFactoryTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.factoryTrait == nil {
		return builder
	}
	for _, u := range f.meta.factoryTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithAnonymousTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.anonymousTrait == nil {
		return builder
	}
	for _, u := range f.meta.anonymousTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithNilTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.nilTrait == nil {
		return builder
	}
	for _, u := range f.meta.nilTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithMixnameTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.mixnameTrait == nil {
		return builder
	}
	for _, u := range f.meta.mixnameTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithMixemailTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.mixemailTrait == nil {
		return builder
	}
	for _, u := range f.meta.mixemailTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) WithMixtitleTrait() *UserBuilder {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	if f.meta.mixtitleTrait == nil {
		return builder
	}
	for _, u := range f.meta.mixtitleTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

func (f *UserFactory) Create(ctx context.Context) (*model.User, error) {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.Create(ctx)
}
func (f *UserFactory) CreateV(ctx context.Context) (model.User, error) {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateV(ctx)
}
func (f *UserFactory) CreateBatch(ctx context.Context, n int) ([]*model.User, error) {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatch(ctx, n)
}
func (f *UserFactory) CreateBatchV(ctx context.Context, n int) ([]model.User, error) {
	builder := &UserBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	return builder.CreateBatchV(ctx, n)
}

type UserBuilder struct {
	factory  *UserFactory
	mutation userMutation
	counter  *Counter

	nameOverride  string
	nameOverriden bool

	emailOverride  string
	emailOverriden bool

	titleOverride  string
	titleOverriden bool

	groupOverride  *model.Group
	groupOverriden bool

	_postFoo    string
	_postFooSet bool
}

func (b *UserBuilder) SetName(i string) *UserBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

func (b *UserBuilder) SetEmail(i string) *UserBuilder {
	b.emailOverride = i
	b.emailOverriden = true
	return b
}

func (b *UserBuilder) SetTitle(i string) *UserBuilder {
	b.titleOverride = i
	b.titleOverriden = true
	return b
}

func (b *UserBuilder) SetGroup(i *model.Group) *UserBuilder {
	b.groupOverride = i
	b.groupOverriden = true
	return b
}

func (b *UserBuilder) SetFooPost(i string) *UserBuilder {
	b._postFoo = i
	b._postFooSet = true
	return b
}

func (b *UserBuilder) WithDefaultTrait() *UserBuilder {
	if b.factory.meta.defaultTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.defaultTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithLazyTrait() *UserBuilder {
	if b.factory.meta.lazyTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.lazyTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithSequenceTrait() *UserBuilder {
	if b.factory.meta.sequenceTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.sequenceTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithFactoryTrait() *UserBuilder {
	if b.factory.meta.factoryTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.factoryTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithAnonymousTrait() *UserBuilder {
	if b.factory.meta.anonymousTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.anonymousTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithNilTrait() *UserBuilder {
	if b.factory.meta.nilTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.nilTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithMixnameTrait() *UserBuilder {
	if b.factory.meta.mixnameTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.mixnameTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithMixemailTrait() *UserBuilder {
	if b.factory.meta.mixemailTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.mixemailTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) WithMixtitleTrait() *UserBuilder {
	if b.factory.meta.mixtitleTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.mixtitleTrait.updates {
		u(&b.mutation)
	}
	return b
}

func (b *UserBuilder) CreateV(ctx context.Context) (model.User, error) {
	var d model.User
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

func (b *UserBuilder) Create(ctx context.Context) (*model.User, error) {

	var preSlice = []func(ctx context.Context, i *model.User, c int) error{}
	var lazySlice = []func(ctx context.Context, i *model.User, c int) error{}
	var postSlice = []func(ctx context.Context, i *model.User, c int) error{}

	index := b.counter.Get()
	_ = index

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.User, c int) error {
			value := b.nameOverride

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

	if b.emailOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.User, c int) error {
			value := b.emailOverride

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

	if b.titleOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.User, c int) error {
			value := b.titleOverride

			i.Title = value
			return nil
		})
	} else {
		switch b.mutation.titleType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.titleFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.titleFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.titleFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.titleFunc)
		}
	}

	if b.groupOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *model.User, c int) error {
			value := b.groupOverride

			i.Group = value
			return nil
		})
	} else {
		switch b.mutation.groupType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.groupFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.groupFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.groupFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.groupFunc)
		}
	}

	if b.mutation._postFooFunc != nil {
		postSlice = append(postSlice, func(ctx context.Context, i *model.User, c int) error {
			err := b.mutation._postFooFunc(ctx, b._postFooSet, i, b._postFoo)
			return err
		})
	}

	v := &model.User{}
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
func (b *UserBuilder) CreateBatch(ctx context.Context, n int) ([]*model.User, error) {
	var results []*model.User
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *UserBuilder) CreateBatchV(ctx context.Context, n int) ([]model.User, error) {
	var results []model.User
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
