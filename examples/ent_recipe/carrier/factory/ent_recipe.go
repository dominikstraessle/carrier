// Code generated by carrier, DO NOT EDIT.
package factory

import (
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent"

	"context"
)

type EntRecipeMutator struct {
	Author *ent.User

	AuthorID int

	Name string

	Servings int

	_creator *ent.RecipeCreate
}

func (m *EntRecipeMutator) EntCreator() *ent.RecipeCreate {
	return m._creator
}

type entRecipeMutation struct {
	authorType int
	authorFunc func(ctx context.Context, i *EntRecipeMutator, c int) error

	authorIDType int
	authorIDFunc func(ctx context.Context, i *EntRecipeMutator, c int) error

	nameType int
	nameFunc func(ctx context.Context, i *EntRecipeMutator, c int) error

	servingsType int
	servingsFunc func(ctx context.Context, i *EntRecipeMutator, c int) error

	_postStepsCountFunc func(ctx context.Context, set bool, obj *ent.Recipe, i int) error

	_postIngredientsCountFunc func(ctx context.Context, set bool, obj *ent.Recipe, i int) error

	beforeCreateFunc func(ctx context.Context, i *EntRecipeMutator) error
	afterCreateFunc  func(ctx context.Context, i *ent.Recipe) error
}
type EntRecipeMetaFactory struct {
	mutation entRecipeMutation

	veganTrait *entRecipeTrait

	ketoTrait *entRecipeTrait
}
type entRecipeTrait struct {
	mutation entRecipeMutation
	updates  []func(m *entRecipeMutation)
}

func EntRecipeTrait() *entRecipeTrait {
	return &entRecipeTrait{}
}
func (*entRecipeMutation) beforeCreateMutateFunc(fn func(ctx context.Context, i *EntRecipeMutator) error) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.beforeCreateFunc = fn
	}
}
func (*entRecipeMutation) afterCreateMutateFunc(fn func(ctx context.Context, i *ent.Recipe) error) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.afterCreateFunc = fn
	}
}

func (*entRecipeMutation) authorSequenceMutateFunc(fn func(ctx context.Context, i int) (*ent.User, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorType = TypeSequence
		m.authorFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetAuthor(value)

			i.Author = value
			return nil
		}
	}
}
func (*entRecipeMutation) authorLazyMutateFunc(fn func(ctx context.Context, i *EntRecipeMutator) (*ent.User, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorType = TypeLazy
		m.authorFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetAuthor(value)

			i.Author = value
			return nil
		}
	}
}
func (*entRecipeMutation) authorDefaultMutateFunc(v *ent.User) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorType = TypeDefault
		m.authorFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {

			i.EntCreator().SetAuthor(v)

			i.Author = v
			return nil
		}
	}
}
func (*entRecipeMutation) authorFactoryMutateFunc(fn func(ctx context.Context) (*ent.User, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorType = TypeFactory
		m.authorFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetAuthor(value)

			i.Author = value

			return nil
		}
	}
}

// SetAuthorSequence register a function which accept a sequence counter and set return value to Author field
func (f *EntRecipeMetaFactory) SetAuthorSequence(fn func(ctx context.Context, i int) (*ent.User, error)) *EntRecipeMetaFactory {
	f.mutation.authorSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetAuthorLazy register a function which accept the build struct and set return value to Author field
func (f *EntRecipeMetaFactory) SetAuthorLazy(fn func(ctx context.Context, i *EntRecipeMutator) (*ent.User, error)) *EntRecipeMetaFactory {
	f.mutation.authorLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetAuthorDefault assign a default value to Author field
func (f *EntRecipeMetaFactory) SetAuthorDefault(v *ent.User) *EntRecipeMetaFactory {
	f.mutation.authorDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetAuthorFactory register a factory function and assign return value to Author, you can also use related factory's Create/CreateV as input function here
func (f *EntRecipeMetaFactory) SetAuthorFactory(fn func(ctx context.Context) (*ent.User, error)) *EntRecipeMetaFactory {
	f.mutation.authorFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetAuthorSequence register a function which accept a sequence counter and set return value to Author field
func (t *entRecipeTrait) SetAuthorSequence(fn func(ctx context.Context, i int) (*ent.User, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorSequenceMutateFunc(fn))
	return t
}

// SetAuthorLazy register a function which accept the build struct and set return value to Author field
func (t *entRecipeTrait) SetAuthorLazy(fn func(ctx context.Context, i *EntRecipeMutator) (*ent.User, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorLazyMutateFunc(fn))
	return t
}

// SetAuthorDefault assign a default value to Author field
func (t *entRecipeTrait) SetAuthorDefault(v *ent.User) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorDefaultMutateFunc(v))
	return t
}

// SetAuthorFactory register a factory function and assign return value to Author, you can also use related factory's Create/CreateV as input function here
func (t *entRecipeTrait) SetAuthorFactory(fn func(ctx context.Context) (*ent.User, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorFactoryMutateFunc(fn))
	return t
}

func (*entRecipeMutation) authorIDSequenceMutateFunc(fn func(ctx context.Context, i int) (int, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorIDType = TypeSequence
		m.authorIDFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetAuthorID(value)

			i.AuthorID = value
			return nil
		}
	}
}
func (*entRecipeMutation) authorIDLazyMutateFunc(fn func(ctx context.Context, i *EntRecipeMutator) (int, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorIDType = TypeLazy
		m.authorIDFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetAuthorID(value)

			i.AuthorID = value
			return nil
		}
	}
}
func (*entRecipeMutation) authorIDDefaultMutateFunc(v int) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorIDType = TypeDefault
		m.authorIDFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {

			i.EntCreator().SetAuthorID(v)

			i.AuthorID = v
			return nil
		}
	}
}
func (*entRecipeMutation) authorIDFactoryMutateFunc(fn func(ctx context.Context) (int, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.authorIDType = TypeFactory
		m.authorIDFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetAuthorID(value)

			i.AuthorID = value

			return nil
		}
	}
}

// SetAuthorIDSequence register a function which accept a sequence counter and set return value to AuthorID field
func (f *EntRecipeMetaFactory) SetAuthorIDSequence(fn func(ctx context.Context, i int) (int, error)) *EntRecipeMetaFactory {
	f.mutation.authorIDSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetAuthorIDLazy register a function which accept the build struct and set return value to AuthorID field
func (f *EntRecipeMetaFactory) SetAuthorIDLazy(fn func(ctx context.Context, i *EntRecipeMutator) (int, error)) *EntRecipeMetaFactory {
	f.mutation.authorIDLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetAuthorIDDefault assign a default value to AuthorID field
func (f *EntRecipeMetaFactory) SetAuthorIDDefault(v int) *EntRecipeMetaFactory {
	f.mutation.authorIDDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetAuthorIDFactory register a factory function and assign return value to AuthorID, you can also use related factory's Create/CreateV as input function here
func (f *EntRecipeMetaFactory) SetAuthorIDFactory(fn func(ctx context.Context) (int, error)) *EntRecipeMetaFactory {
	f.mutation.authorIDFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetAuthorIDSequence register a function which accept a sequence counter and set return value to AuthorID field
func (t *entRecipeTrait) SetAuthorIDSequence(fn func(ctx context.Context, i int) (int, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorIDSequenceMutateFunc(fn))
	return t
}

// SetAuthorIDLazy register a function which accept the build struct and set return value to AuthorID field
func (t *entRecipeTrait) SetAuthorIDLazy(fn func(ctx context.Context, i *EntRecipeMutator) (int, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorIDLazyMutateFunc(fn))
	return t
}

// SetAuthorIDDefault assign a default value to AuthorID field
func (t *entRecipeTrait) SetAuthorIDDefault(v int) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorIDDefaultMutateFunc(v))
	return t
}

// SetAuthorIDFactory register a factory function and assign return value to AuthorID, you can also use related factory's Create/CreateV as input function here
func (t *entRecipeTrait) SetAuthorIDFactory(fn func(ctx context.Context) (int, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.authorIDFactoryMutateFunc(fn))
	return t
}

func (*entRecipeMutation) nameSequenceMutateFunc(fn func(ctx context.Context, i int) (string, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.nameType = TypeSequence
		m.nameFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
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
func (*entRecipeMutation) nameLazyMutateFunc(fn func(ctx context.Context, i *EntRecipeMutator) (string, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.nameType = TypeLazy
		m.nameFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
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
func (*entRecipeMutation) nameDefaultMutateFunc(v string) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.nameType = TypeDefault
		m.nameFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {

			i.EntCreator().SetName(v)

			i.Name = v
			return nil
		}
	}
}
func (*entRecipeMutation) nameFactoryMutateFunc(fn func(ctx context.Context) (string, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.nameType = TypeFactory
		m.nameFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
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
func (f *EntRecipeMetaFactory) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *EntRecipeMetaFactory {
	f.mutation.nameSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (f *EntRecipeMetaFactory) SetNameLazy(fn func(ctx context.Context, i *EntRecipeMutator) (string, error)) *EntRecipeMetaFactory {
	f.mutation.nameLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameDefault assign a default value to Name field
func (f *EntRecipeMetaFactory) SetNameDefault(v string) *EntRecipeMetaFactory {
	f.mutation.nameDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (f *EntRecipeMetaFactory) SetNameFactory(fn func(ctx context.Context) (string, error)) *EntRecipeMetaFactory {
	f.mutation.nameFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetNameSequence register a function which accept a sequence counter and set return value to Name field
func (t *entRecipeTrait) SetNameSequence(fn func(ctx context.Context, i int) (string, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.nameSequenceMutateFunc(fn))
	return t
}

// SetNameLazy register a function which accept the build struct and set return value to Name field
func (t *entRecipeTrait) SetNameLazy(fn func(ctx context.Context, i *EntRecipeMutator) (string, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.nameLazyMutateFunc(fn))
	return t
}

// SetNameDefault assign a default value to Name field
func (t *entRecipeTrait) SetNameDefault(v string) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.nameDefaultMutateFunc(v))
	return t
}

// SetNameFactory register a factory function and assign return value to Name, you can also use related factory's Create/CreateV as input function here
func (t *entRecipeTrait) SetNameFactory(fn func(ctx context.Context) (string, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.nameFactoryMutateFunc(fn))
	return t
}

func (*entRecipeMutation) servingsSequenceMutateFunc(fn func(ctx context.Context, i int) (int, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.servingsType = TypeSequence
		m.servingsFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, c)
			if err != nil {
				return err
			}

			i.EntCreator().SetServings(value)

			i.Servings = value
			return nil
		}
	}
}
func (*entRecipeMutation) servingsLazyMutateFunc(fn func(ctx context.Context, i *EntRecipeMutator) (int, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.servingsType = TypeLazy
		m.servingsFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx, i)
			if err != nil {
				return err
			}

			i.EntCreator().SetServings(value)

			i.Servings = value
			return nil
		}
	}
}
func (*entRecipeMutation) servingsDefaultMutateFunc(v int) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.servingsType = TypeDefault
		m.servingsFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {

			i.EntCreator().SetServings(v)

			i.Servings = v
			return nil
		}
	}
}
func (*entRecipeMutation) servingsFactoryMutateFunc(fn func(ctx context.Context) (int, error)) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m.servingsType = TypeFactory
		m.servingsFunc = func(ctx context.Context, i *EntRecipeMutator, c int) error {
			if fn == nil {
				return nil
			}
			value, err := fn(ctx)
			if err != nil {
				return err
			}

			i.EntCreator().SetServings(value)

			i.Servings = value

			return nil
		}
	}
}

// SetServingsSequence register a function which accept a sequence counter and set return value to Servings field
func (f *EntRecipeMetaFactory) SetServingsSequence(fn func(ctx context.Context, i int) (int, error)) *EntRecipeMetaFactory {
	f.mutation.servingsSequenceMutateFunc(fn)(&f.mutation)
	return f
}

// SetServingsLazy register a function which accept the build struct and set return value to Servings field
func (f *EntRecipeMetaFactory) SetServingsLazy(fn func(ctx context.Context, i *EntRecipeMutator) (int, error)) *EntRecipeMetaFactory {
	f.mutation.servingsLazyMutateFunc(fn)(&f.mutation)
	return f
}

// SetServingsDefault assign a default value to Servings field
func (f *EntRecipeMetaFactory) SetServingsDefault(v int) *EntRecipeMetaFactory {
	f.mutation.servingsDefaultMutateFunc(v)(&f.mutation)
	return f
}

// SetServingsFactory register a factory function and assign return value to Servings, you can also use related factory's Create/CreateV as input function here
func (f *EntRecipeMetaFactory) SetServingsFactory(fn func(ctx context.Context) (int, error)) *EntRecipeMetaFactory {
	f.mutation.servingsFactoryMutateFunc(fn)(&f.mutation)
	return f
}

// SetServingsSequence register a function which accept a sequence counter and set return value to Servings field
func (t *entRecipeTrait) SetServingsSequence(fn func(ctx context.Context, i int) (int, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.servingsSequenceMutateFunc(fn))
	return t
}

// SetServingsLazy register a function which accept the build struct and set return value to Servings field
func (t *entRecipeTrait) SetServingsLazy(fn func(ctx context.Context, i *EntRecipeMutator) (int, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.servingsLazyMutateFunc(fn))
	return t
}

// SetServingsDefault assign a default value to Servings field
func (t *entRecipeTrait) SetServingsDefault(v int) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.servingsDefaultMutateFunc(v))
	return t
}

// SetServingsFactory register a factory function and assign return value to Servings, you can also use related factory's Create/CreateV as input function here
func (t *entRecipeTrait) SetServingsFactory(fn func(ctx context.Context) (int, error)) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.servingsFactoryMutateFunc(fn))
	return t
}

func (*entRecipeMutation) stepsCountPostMutateFunc(fn func(ctx context.Context, set bool, obj *ent.Recipe, i int) error) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m._postStepsCountFunc = fn
	}
}

// SetStepsCountPostFunc register a post function which will be called in factory SetStepsCountPost method
func (f *EntRecipeMetaFactory) SetStepsCountPostFunc(fn func(ctx context.Context, set bool, obj *ent.Recipe, i int) error) *EntRecipeMetaFactory {
	f.mutation.stepsCountPostMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entRecipeTrait) SetStepsCountPostFunc(fn func(ctx context.Context, set bool, obj *ent.Recipe, i int) error) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.stepsCountPostMutateFunc(fn))
	return t
}

func (*entRecipeMutation) ingredientsCountPostMutateFunc(fn func(ctx context.Context, set bool, obj *ent.Recipe, i int) error) func(m *entRecipeMutation) {
	return func(m *entRecipeMutation) {
		m._postIngredientsCountFunc = fn
	}
}

// SetIngredientsCountPostFunc register a post function which will be called in factory SetIngredientsCountPost method
func (f *EntRecipeMetaFactory) SetIngredientsCountPostFunc(fn func(ctx context.Context, set bool, obj *ent.Recipe, i int) error) *EntRecipeMetaFactory {
	f.mutation.ingredientsCountPostMutateFunc(fn)(&f.mutation)
	return f
}
func (t *entRecipeTrait) SetIngredientsCountPostFunc(fn func(ctx context.Context, set bool, obj *ent.Recipe, i int) error) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.ingredientsCountPostMutateFunc(fn))
	return t
}

// SetVeganTrait accept a entRecipeTrait, will override builder using Trait's methods if enable
func (f *EntRecipeMetaFactory) SetVeganTrait(t *entRecipeTrait) *EntRecipeMetaFactory {
	f.veganTrait = t
	return f
}

// SetKetoTrait accept a entRecipeTrait, will override builder using Trait's methods if enable
func (f *EntRecipeMetaFactory) SetKetoTrait(t *entRecipeTrait) *EntRecipeMetaFactory {
	f.ketoTrait = t
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (f *EntRecipeMetaFactory) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Recipe) error) *EntRecipeMetaFactory {
	f.mutation.afterCreateFunc = fn
	return f
}

// SetBeforeCreateFunc register a function to be called before struct create
func (f *EntRecipeMetaFactory) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntRecipeMutator) error) *EntRecipeMetaFactory {
	f.mutation.beforeCreateFunc = fn
	return f
}

// SetAfterCreateFunc register a function to be called after struct create
func (t *entRecipeTrait) SetAfterCreateFunc(fn func(ctx context.Context, i *ent.Recipe) error) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.afterCreateMutateFunc(fn))
	return t
}

// SetBeforeCreateFunc register a function to be called before struct create
func (t *entRecipeTrait) SetBeforeCreateFunc(fn func(ctx context.Context, i *EntRecipeMutator) error) *entRecipeTrait {
	t.updates = append(t.updates, t.mutation.beforeCreateMutateFunc(fn))
	return t
}

// Build create a  EntRecipeFactory from EntRecipeMetaFactory
func (f *EntRecipeMetaFactory) Build() *EntRecipeFactory {
	return &EntRecipeFactory{meta: *f, counter: &Counter{}}
}

type EntRecipeFactory struct {
	meta    EntRecipeMetaFactory
	counter *Counter

	client *ent.Client
}

// SetAuthor set the Author field
func (f *EntRecipeFactory) SetAuthor(i *ent.User) *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetAuthor(i)

	builder.client = f.client

	return builder
}

// SetAuthorID set the AuthorID field
func (f *EntRecipeFactory) SetAuthorID(i int) *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetAuthorID(i)

	builder.client = f.client

	return builder
}

// SetName set the Name field
func (f *EntRecipeFactory) SetName(i string) *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetName(i)

	builder.client = f.client

	return builder
}

// SetServings set the Servings field
func (f *EntRecipeFactory) SetServings(i int) *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetServings(i)

	builder.client = f.client

	return builder
}

// SetStepsCountPost call the post function with int input
func (f *EntRecipeFactory) SetStepsCountPost(i int) *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetStepsCountPost(i)

	builder.client = f.client

	return builder
}

// SetIngredientsCountPost call the post function with int input
func (f *EntRecipeFactory) SetIngredientsCountPost(i int) *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}
	builder.SetIngredientsCountPost(i)

	builder.client = f.client

	return builder
}

// WithVeganTrait() enable the Vegan trait
func (f *EntRecipeFactory) WithVeganTrait() *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	builder.client = f.client

	if f.meta.veganTrait == nil {
		return builder
	}
	for _, u := range f.meta.veganTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

// WithKetoTrait() enable the Keto trait
func (f *EntRecipeFactory) WithKetoTrait() *EntRecipeBuilder {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter}
	builder.factory = f

	builder.client = f.client

	if f.meta.ketoTrait == nil {
		return builder
	}
	for _, u := range f.meta.ketoTrait.updates {
		u(&builder.mutation)
	}
	return builder
}

// Create return a new *ent.Recipe
func (f *EntRecipeFactory) Create(ctx context.Context) (*ent.Recipe, error) {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.Create(ctx)
}

// CreateV return a new ent.Recipe
func (f *EntRecipeFactory) CreateV(ctx context.Context) (ent.Recipe, error) {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateV(ctx)
}

// CreateBatch return a []*ent.Recipe slice
func (f *EntRecipeFactory) CreateBatch(ctx context.Context, n int) ([]*ent.Recipe, error) {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatch(ctx, n)
}

// CreateBatchV return a []ent.Recipe slice
func (f *EntRecipeFactory) CreateBatchV(ctx context.Context, n int) ([]ent.Recipe, error) {
	builder := &EntRecipeBuilder{mutation: f.meta.mutation, counter: f.counter, factory: f}

	builder.client = f.client

	return builder.CreateBatchV(ctx, n)
}

// Client set ent client to EntRecipeFactory
func (f *EntRecipeFactory) Client(c *ent.Client) *EntRecipeFactory {
	f.client = c
	return f
}

type EntRecipeBuilder struct {
	factory  *EntRecipeFactory
	mutation entRecipeMutation
	counter  *Counter

	authorOverride  *ent.User
	authorOverriden bool

	authorIDOverride  int
	authorIDOverriden bool

	nameOverride  string
	nameOverriden bool

	servingsOverride  int
	servingsOverriden bool

	_postStepsCount    int
	_postStepsCountSet bool

	_postIngredientsCount    int
	_postIngredientsCountSet bool

	client *ent.Client
}

func (b *EntRecipeBuilder) Client(c *ent.Client) *EntRecipeBuilder {
	b.client = c
	return b
}

// SetAuthor set the Author field
func (b *EntRecipeBuilder) SetAuthor(i *ent.User) *EntRecipeBuilder {
	b.authorOverride = i
	b.authorOverriden = true
	return b
}

// SetAuthorID set the AuthorID field
func (b *EntRecipeBuilder) SetAuthorID(i int) *EntRecipeBuilder {
	b.authorIDOverride = i
	b.authorIDOverriden = true
	return b
}

// SetName set the Name field
func (b *EntRecipeBuilder) SetName(i string) *EntRecipeBuilder {
	b.nameOverride = i
	b.nameOverriden = true
	return b
}

// SetServings set the Servings field
func (b *EntRecipeBuilder) SetServings(i int) *EntRecipeBuilder {
	b.servingsOverride = i
	b.servingsOverriden = true
	return b
}

// SetStepsCountPost call the post function with int input
func (b *EntRecipeBuilder) SetStepsCountPost(i int) *EntRecipeBuilder {
	b._postStepsCount = i
	b._postStepsCountSet = true
	return b
}

// SetIngredientsCountPost call the post function with int input
func (b *EntRecipeBuilder) SetIngredientsCountPost(i int) *EntRecipeBuilder {
	b._postIngredientsCount = i
	b._postIngredientsCountSet = true
	return b
}

// WithVeganTrait() enable the Vegan trait
func (b *EntRecipeBuilder) WithVeganTrait() *EntRecipeBuilder {
	if b.factory.meta.veganTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.veganTrait.updates {
		u(&b.mutation)
	}
	return b
}

// WithKetoTrait() enable the Keto trait
func (b *EntRecipeBuilder) WithKetoTrait() *EntRecipeBuilder {
	if b.factory.meta.ketoTrait == nil {
		return b
	}
	for _, u := range b.factory.meta.ketoTrait.updates {
		u(&b.mutation)
	}
	return b
}

// CreateV return a new ent.Recipe
func (b *EntRecipeBuilder) CreateV(ctx context.Context) (ent.Recipe, error) {
	var d ent.Recipe
	p, err := b.Create(ctx)
	if err == nil {
		d = *p
	}
	return d, err
}

// Create return a new *ent.Recipe
func (b *EntRecipeBuilder) Create(ctx context.Context) (*ent.Recipe, error) {

	var preSlice = []func(ctx context.Context, i *EntRecipeMutator, c int) error{}
	var lazySlice = []func(ctx context.Context, i *EntRecipeMutator, c int) error{}
	var postSlice = []func(ctx context.Context, i *ent.Recipe, c int) error{}

	index := b.counter.Get()
	_ = index

	client := b.client
	entBuilder := client.Recipe.Create()

	if b.authorOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntRecipeMutator, c int) error {
			value := b.authorOverride

			i.EntCreator().SetAuthor(value)

			i.Author = value
			return nil
		})
	} else {
		switch b.mutation.authorType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.authorFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.authorFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.authorFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.authorFunc)
		}
	}

	if b.authorIDOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntRecipeMutator, c int) error {
			value := b.authorIDOverride

			i.EntCreator().SetAuthorID(value)

			i.AuthorID = value
			return nil
		})
	} else {
		switch b.mutation.authorIDType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.authorIDFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.authorIDFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.authorIDFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.authorIDFunc)
		}
	}

	if b.nameOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntRecipeMutator, c int) error {
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

	if b.servingsOverriden {
		preSlice = append(preSlice, func(ctx context.Context, i *EntRecipeMutator, c int) error {
			value := b.servingsOverride

			i.EntCreator().SetServings(value)

			i.Servings = value
			return nil
		})
	} else {
		switch b.mutation.servingsType {
		case TypeDefault:
			preSlice = append(preSlice, b.mutation.servingsFunc)
		case TypeLazy:
			lazySlice = append(lazySlice, b.mutation.servingsFunc)
		case TypeSequence:
			preSlice = append(preSlice, b.mutation.servingsFunc)
		case TypeFactory:
			preSlice = append(preSlice, b.mutation.servingsFunc)
		}
	}

	if b.mutation._postStepsCountFunc != nil {
		postSlice = append(postSlice, func(ctx context.Context, i *ent.Recipe, c int) error {
			err := b.mutation._postStepsCountFunc(ctx, b._postStepsCountSet, i, b._postStepsCount)
			return err
		})
	}

	if b.mutation._postIngredientsCountFunc != nil {
		postSlice = append(postSlice, func(ctx context.Context, i *ent.Recipe, c int) error {
			err := b.mutation._postIngredientsCountFunc(ctx, b._postIngredientsCountSet, i, b._postIngredientsCount)
			return err
		})
	}

	v := &EntRecipeMutator{}

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
func (b *EntRecipeBuilder) CreateBatch(ctx context.Context, n int) ([]*ent.Recipe, error) {
	var results []*ent.Recipe
	for i := 0; i < n; i++ {
		d, err := b.Create(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
func (b *EntRecipeBuilder) CreateBatchV(ctx context.Context, n int) ([]ent.Recipe, error) {
	var results []ent.Recipe
	for i := 0; i < n; i++ {
		d, err := b.CreateV(ctx)
		if err != nil {
			return results, err
		}
		results = append(results, d)
	}
	return results, nil
}
