// Code generated by ent, DO NOT EDIT.

package gen

import (
	"backend/ent/gen/liveshow"
	"backend/ent/gen/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LiveShowDelete is the builder for deleting a LiveShow entity.
type LiveShowDelete struct {
	config
	hooks    []Hook
	mutation *LiveShowMutation
}

// Where appends a list predicates to the LiveShowDelete builder.
func (lsd *LiveShowDelete) Where(ps ...predicate.LiveShow) *LiveShowDelete {
	lsd.mutation.Where(ps...)
	return lsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lsd *LiveShowDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lsd.sqlExec, lsd.mutation, lsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lsd *LiveShowDelete) ExecX(ctx context.Context) int {
	n, err := lsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lsd *LiveShowDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(liveshow.Table, sqlgraph.NewFieldSpec(liveshow.FieldID, field.TypeInt))
	if ps := lsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lsd.mutation.done = true
	return affected, err
}

// LiveShowDeleteOne is the builder for deleting a single LiveShow entity.
type LiveShowDeleteOne struct {
	lsd *LiveShowDelete
}

// Where appends a list predicates to the LiveShowDelete builder.
func (lsdo *LiveShowDeleteOne) Where(ps ...predicate.LiveShow) *LiveShowDeleteOne {
	lsdo.lsd.mutation.Where(ps...)
	return lsdo
}

// Exec executes the deletion query.
func (lsdo *LiveShowDeleteOne) Exec(ctx context.Context) error {
	n, err := lsdo.lsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{liveshow.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lsdo *LiveShowDeleteOne) ExecX(ctx context.Context) {
	if err := lsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
