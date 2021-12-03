// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Yiling-J/carrier/examples/ent_recipe/ent/ingredient"
)

// Ingredient is the model entity for the Ingredient schema.
type Ingredient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IngredientQuery when eager-loading is set.
	Edges IngredientEdges `json:"edges"`
}

// IngredientEdges holds the relations/edges for other nodes in the graph.
type IngredientEdges struct {
	// RecipeIngredients holds the value of the recipe_ingredients edge.
	RecipeIngredients []*RecipeIngredient `json:"recipe_ingredients,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecipeIngredientsOrErr returns the RecipeIngredients value or an error if the edge
// was not loaded in eager-loading.
func (e IngredientEdges) RecipeIngredientsOrErr() ([]*RecipeIngredient, error) {
	if e.loadedTypes[0] {
		return e.RecipeIngredients, nil
	}
	return nil, &NotLoadedError{edge: "recipe_ingredients"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ingredient) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ingredient.FieldID:
			values[i] = new(sql.NullInt64)
		case ingredient.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ingredient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ingredient fields.
func (i *Ingredient) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case ingredient.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case ingredient.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		}
	}
	return nil
}

// QueryRecipeIngredients queries the "recipe_ingredients" edge of the Ingredient entity.
func (i *Ingredient) QueryRecipeIngredients() *RecipeIngredientQuery {
	return (&IngredientClient{config: i.config}).QueryRecipeIngredients(i)
}

// Update returns a builder for updating this Ingredient.
// Note that you need to call Ingredient.Unwrap() before calling this method if this Ingredient
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Ingredient) Update() *IngredientUpdateOne {
	return (&IngredientClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Ingredient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Ingredient) Unwrap() *Ingredient {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ingredient is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Ingredient) String() string {
	var builder strings.Builder
	builder.WriteString("Ingredient(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", name=")
	builder.WriteString(i.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Ingredients is a parsable slice of Ingredient.
type Ingredients []*Ingredient

func (i Ingredients) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
