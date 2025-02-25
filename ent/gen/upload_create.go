// Code generated by ent, DO NOT EDIT.

package gen

import (
	"backend/ent/gen/liveshow"
	"backend/ent/gen/upload"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UploadCreate is the builder for creating a Upload entity.
type UploadCreate struct {
	config
	mutation *UploadMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (uc *UploadCreate) SetCreateTime(t time.Time) *UploadCreate {
	uc.mutation.SetCreateTime(t)
	return uc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uc *UploadCreate) SetNillableCreateTime(t *time.Time) *UploadCreate {
	if t != nil {
		uc.SetCreateTime(*t)
	}
	return uc
}

// SetFilePath sets the "file_path" field.
func (uc *UploadCreate) SetFilePath(s string) *UploadCreate {
	uc.mutation.SetFilePath(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UploadCreate) SetName(s string) *UploadCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetFileType sets the "file_type" field.
func (uc *UploadCreate) SetFileType(ut upload.FileType) *UploadCreate {
	uc.mutation.SetFileType(ut)
	return uc
}

// AddLiveShowUploadIDs adds the "live_show_upload" edge to the LiveShow entity by IDs.
func (uc *UploadCreate) AddLiveShowUploadIDs(ids ...int) *UploadCreate {
	uc.mutation.AddLiveShowUploadIDs(ids...)
	return uc
}

// AddLiveShowUpload adds the "live_show_upload" edges to the LiveShow entity.
func (uc *UploadCreate) AddLiveShowUpload(l ...*LiveShow) *UploadCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uc.AddLiveShowUploadIDs(ids...)
}

// Mutation returns the UploadMutation object of the builder.
func (uc *UploadCreate) Mutation() *UploadMutation {
	return uc.mutation
}

// Save creates the Upload in the database.
func (uc *UploadCreate) Save(ctx context.Context) (*Upload, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UploadCreate) SaveX(ctx context.Context) *Upload {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UploadCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UploadCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UploadCreate) defaults() {
	if _, ok := uc.mutation.CreateTime(); !ok {
		v := upload.DefaultCreateTime()
		uc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UploadCreate) check() error {
	if _, ok := uc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`gen: missing required field "Upload.create_time"`)}
	}
	if _, ok := uc.mutation.FilePath(); !ok {
		return &ValidationError{Name: "file_path", err: errors.New(`gen: missing required field "Upload.file_path"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`gen: missing required field "Upload.name"`)}
	}
	if _, ok := uc.mutation.FileType(); !ok {
		return &ValidationError{Name: "file_type", err: errors.New(`gen: missing required field "Upload.file_type"`)}
	}
	if v, ok := uc.mutation.FileType(); ok {
		if err := upload.FileTypeValidator(v); err != nil {
			return &ValidationError{Name: "file_type", err: fmt.Errorf(`gen: validator failed for field "Upload.file_type": %w`, err)}
		}
	}
	return nil
}

func (uc *UploadCreate) sqlSave(ctx context.Context) (*Upload, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UploadCreate) createSpec() (*Upload, *sqlgraph.CreateSpec) {
	var (
		_node = &Upload{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(upload.Table, sqlgraph.NewFieldSpec(upload.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.CreateTime(); ok {
		_spec.SetField(upload.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := uc.mutation.FilePath(); ok {
		_spec.SetField(upload.FieldFilePath, field.TypeString, value)
		_node.FilePath = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(upload.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.FileType(); ok {
		_spec.SetField(upload.FieldFileType, field.TypeEnum, value)
		_node.FileType = value
	}
	if nodes := uc.mutation.LiveShowUploadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   upload.LiveShowUploadTable,
			Columns: []string{upload.LiveShowUploadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(liveshow.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UploadCreateBulk is the builder for creating many Upload entities in bulk.
type UploadCreateBulk struct {
	config
	err      error
	builders []*UploadCreate
}

// Save creates the Upload entities in the database.
func (ucb *UploadCreateBulk) Save(ctx context.Context) ([]*Upload, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Upload, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UploadMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UploadCreateBulk) SaveX(ctx context.Context) []*Upload {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UploadCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UploadCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
