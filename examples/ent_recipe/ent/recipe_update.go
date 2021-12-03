// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/category"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/predicate"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/recipe"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/recipeingredient"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/step"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/user"
)

// RecipeUpdate is the builder for updating Recipe entities.
type RecipeUpdate struct {
	config
	hooks    []Hook
	mutation *RecipeMutation
}

// Where appends a list predicates to the RecipeUpdate builder.
func (ru *RecipeUpdate) Where(ps ...predicate.Recipe) *RecipeUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RecipeUpdate) SetName(s string) *RecipeUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetServings sets the "servings" field.
func (ru *RecipeUpdate) SetServings(i int) *RecipeUpdate {
	ru.mutation.ResetServings()
	ru.mutation.SetServings(i)
	return ru
}

// AddServings adds i to the "servings" field.
func (ru *RecipeUpdate) AddServings(i int) *RecipeUpdate {
	ru.mutation.AddServings(i)
	return ru
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (ru *RecipeUpdate) SetAuthorID(id int) *RecipeUpdate {
	ru.mutation.SetAuthorID(id)
	return ru
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (ru *RecipeUpdate) SetNillableAuthorID(id *int) *RecipeUpdate {
	if id != nil {
		ru = ru.SetAuthorID(*id)
	}
	return ru
}

// SetAuthor sets the "author" edge to the User entity.
func (ru *RecipeUpdate) SetAuthor(u *User) *RecipeUpdate {
	return ru.SetAuthorID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Category entity by IDs.
func (ru *RecipeUpdate) AddTagIDs(ids ...int) *RecipeUpdate {
	ru.mutation.AddTagIDs(ids...)
	return ru
}

// AddTags adds the "tags" edges to the Category entity.
func (ru *RecipeUpdate) AddTags(c ...*Category) *RecipeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddTagIDs(ids...)
}

// AddStepIDs adds the "steps" edge to the Step entity by IDs.
func (ru *RecipeUpdate) AddStepIDs(ids ...int) *RecipeUpdate {
	ru.mutation.AddStepIDs(ids...)
	return ru
}

// AddSteps adds the "steps" edges to the Step entity.
func (ru *RecipeUpdate) AddSteps(s ...*Step) *RecipeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddStepIDs(ids...)
}

// AddRecipeIngredientIDs adds the "recipe_ingredients" edge to the RecipeIngredient entity by IDs.
func (ru *RecipeUpdate) AddRecipeIngredientIDs(ids ...int) *RecipeUpdate {
	ru.mutation.AddRecipeIngredientIDs(ids...)
	return ru
}

// AddRecipeIngredients adds the "recipe_ingredients" edges to the RecipeIngredient entity.
func (ru *RecipeUpdate) AddRecipeIngredients(r ...*RecipeIngredient) *RecipeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRecipeIngredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (ru *RecipeUpdate) Mutation() *RecipeMutation {
	return ru.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (ru *RecipeUpdate) ClearAuthor() *RecipeUpdate {
	ru.mutation.ClearAuthor()
	return ru
}

// ClearTags clears all "tags" edges to the Category entity.
func (ru *RecipeUpdate) ClearTags() *RecipeUpdate {
	ru.mutation.ClearTags()
	return ru
}

// RemoveTagIDs removes the "tags" edge to Category entities by IDs.
func (ru *RecipeUpdate) RemoveTagIDs(ids ...int) *RecipeUpdate {
	ru.mutation.RemoveTagIDs(ids...)
	return ru
}

// RemoveTags removes "tags" edges to Category entities.
func (ru *RecipeUpdate) RemoveTags(c ...*Category) *RecipeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveTagIDs(ids...)
}

// ClearSteps clears all "steps" edges to the Step entity.
func (ru *RecipeUpdate) ClearSteps() *RecipeUpdate {
	ru.mutation.ClearSteps()
	return ru
}

// RemoveStepIDs removes the "steps" edge to Step entities by IDs.
func (ru *RecipeUpdate) RemoveStepIDs(ids ...int) *RecipeUpdate {
	ru.mutation.RemoveStepIDs(ids...)
	return ru
}

// RemoveSteps removes "steps" edges to Step entities.
func (ru *RecipeUpdate) RemoveSteps(s ...*Step) *RecipeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveStepIDs(ids...)
}

// ClearRecipeIngredients clears all "recipe_ingredients" edges to the RecipeIngredient entity.
func (ru *RecipeUpdate) ClearRecipeIngredients() *RecipeUpdate {
	ru.mutation.ClearRecipeIngredients()
	return ru
}

// RemoveRecipeIngredientIDs removes the "recipe_ingredients" edge to RecipeIngredient entities by IDs.
func (ru *RecipeUpdate) RemoveRecipeIngredientIDs(ids ...int) *RecipeUpdate {
	ru.mutation.RemoveRecipeIngredientIDs(ids...)
	return ru
}

// RemoveRecipeIngredients removes "recipe_ingredients" edges to RecipeIngredient entities.
func (ru *RecipeUpdate) RemoveRecipeIngredients(r ...*RecipeIngredient) *RecipeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRecipeIngredientIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecipeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecipeUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecipeUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecipeUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RecipeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipe.Table,
			Columns: recipe.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipe.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldName,
		})
	}
	if value, ok := ru.mutation.Servings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldServings,
		})
	}
	if value, ok := ru.mutation.AddedServings(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldServings,
		})
	}
	if ru.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recipe.AuthorTable,
			Columns: []string{recipe.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recipe.AuthorTable,
			Columns: []string{recipe.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedTagsIDs(); len(nodes) > 0 && !ru.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedStepsIDs(); len(nodes) > 0 && !ru.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RecipeIngredientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRecipeIngredientsIDs(); len(nodes) > 0 && !ru.mutation.RecipeIngredientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RecipeIngredientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipe.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RecipeUpdateOne is the builder for updating a single Recipe entity.
type RecipeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecipeMutation
}

// SetName sets the "name" field.
func (ruo *RecipeUpdateOne) SetName(s string) *RecipeUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetServings sets the "servings" field.
func (ruo *RecipeUpdateOne) SetServings(i int) *RecipeUpdateOne {
	ruo.mutation.ResetServings()
	ruo.mutation.SetServings(i)
	return ruo
}

// AddServings adds i to the "servings" field.
func (ruo *RecipeUpdateOne) AddServings(i int) *RecipeUpdateOne {
	ruo.mutation.AddServings(i)
	return ruo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (ruo *RecipeUpdateOne) SetAuthorID(id int) *RecipeUpdateOne {
	ruo.mutation.SetAuthorID(id)
	return ruo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (ruo *RecipeUpdateOne) SetNillableAuthorID(id *int) *RecipeUpdateOne {
	if id != nil {
		ruo = ruo.SetAuthorID(*id)
	}
	return ruo
}

// SetAuthor sets the "author" edge to the User entity.
func (ruo *RecipeUpdateOne) SetAuthor(u *User) *RecipeUpdateOne {
	return ruo.SetAuthorID(u.ID)
}

// AddTagIDs adds the "tags" edge to the Category entity by IDs.
func (ruo *RecipeUpdateOne) AddTagIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.AddTagIDs(ids...)
	return ruo
}

// AddTags adds the "tags" edges to the Category entity.
func (ruo *RecipeUpdateOne) AddTags(c ...*Category) *RecipeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddTagIDs(ids...)
}

// AddStepIDs adds the "steps" edge to the Step entity by IDs.
func (ruo *RecipeUpdateOne) AddStepIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.AddStepIDs(ids...)
	return ruo
}

// AddSteps adds the "steps" edges to the Step entity.
func (ruo *RecipeUpdateOne) AddSteps(s ...*Step) *RecipeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddStepIDs(ids...)
}

// AddRecipeIngredientIDs adds the "recipe_ingredients" edge to the RecipeIngredient entity by IDs.
func (ruo *RecipeUpdateOne) AddRecipeIngredientIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.AddRecipeIngredientIDs(ids...)
	return ruo
}

// AddRecipeIngredients adds the "recipe_ingredients" edges to the RecipeIngredient entity.
func (ruo *RecipeUpdateOne) AddRecipeIngredients(r ...*RecipeIngredient) *RecipeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRecipeIngredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (ruo *RecipeUpdateOne) Mutation() *RecipeMutation {
	return ruo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (ruo *RecipeUpdateOne) ClearAuthor() *RecipeUpdateOne {
	ruo.mutation.ClearAuthor()
	return ruo
}

// ClearTags clears all "tags" edges to the Category entity.
func (ruo *RecipeUpdateOne) ClearTags() *RecipeUpdateOne {
	ruo.mutation.ClearTags()
	return ruo
}

// RemoveTagIDs removes the "tags" edge to Category entities by IDs.
func (ruo *RecipeUpdateOne) RemoveTagIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.RemoveTagIDs(ids...)
	return ruo
}

// RemoveTags removes "tags" edges to Category entities.
func (ruo *RecipeUpdateOne) RemoveTags(c ...*Category) *RecipeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveTagIDs(ids...)
}

// ClearSteps clears all "steps" edges to the Step entity.
func (ruo *RecipeUpdateOne) ClearSteps() *RecipeUpdateOne {
	ruo.mutation.ClearSteps()
	return ruo
}

// RemoveStepIDs removes the "steps" edge to Step entities by IDs.
func (ruo *RecipeUpdateOne) RemoveStepIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.RemoveStepIDs(ids...)
	return ruo
}

// RemoveSteps removes "steps" edges to Step entities.
func (ruo *RecipeUpdateOne) RemoveSteps(s ...*Step) *RecipeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveStepIDs(ids...)
}

// ClearRecipeIngredients clears all "recipe_ingredients" edges to the RecipeIngredient entity.
func (ruo *RecipeUpdateOne) ClearRecipeIngredients() *RecipeUpdateOne {
	ruo.mutation.ClearRecipeIngredients()
	return ruo
}

// RemoveRecipeIngredientIDs removes the "recipe_ingredients" edge to RecipeIngredient entities by IDs.
func (ruo *RecipeUpdateOne) RemoveRecipeIngredientIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.RemoveRecipeIngredientIDs(ids...)
	return ruo
}

// RemoveRecipeIngredients removes "recipe_ingredients" edges to RecipeIngredient entities.
func (ruo *RecipeUpdateOne) RemoveRecipeIngredients(r ...*RecipeIngredient) *RecipeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRecipeIngredientIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecipeUpdateOne) Select(field string, fields ...string) *RecipeUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Recipe entity.
func (ruo *RecipeUpdateOne) Save(ctx context.Context) (*Recipe, error) {
	var (
		err  error
		node *Recipe
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecipeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecipeUpdateOne) SaveX(ctx context.Context) *Recipe {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecipeUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecipeUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RecipeUpdateOne) sqlSave(ctx context.Context) (_node *Recipe, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recipe.Table,
			Columns: recipe.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recipe.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Recipe.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recipe.FieldID)
		for _, f := range fields {
			if !recipe.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recipe.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recipe.FieldName,
		})
	}
	if value, ok := ruo.mutation.Servings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldServings,
		})
	}
	if value, ok := ruo.mutation.AddedServings(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recipe.FieldServings,
		})
	}
	if ruo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recipe.AuthorTable,
			Columns: []string{recipe.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recipe.AuthorTable,
			Columns: []string{recipe.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !ruo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   recipe.TagsTable,
			Columns: recipe.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedStepsIDs(); len(nodes) > 0 && !ruo.mutation.StepsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.StepsTable,
			Columns: []string{recipe.StepsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RecipeIngredientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRecipeIngredientsIDs(); len(nodes) > 0 && !ruo.mutation.RecipeIngredientsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RecipeIngredientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeIngredientsTable,
			Columns: []string{recipe.RecipeIngredientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recipeingredient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Recipe{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipe.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
