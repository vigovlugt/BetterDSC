// Code generated by entc, DO NOT EDIT.

package ent

import (
	"betterdsc/ent/serveremote"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServerEmoteCreate is the builder for creating a ServerEmote entity.
type ServerEmoteCreate struct {
	config
	mutation *ServerEmoteMutation
	hooks    []Hook
}

// SetServerID sets the "server_id" field.
func (sec *ServerEmoteCreate) SetServerID(s string) *ServerEmoteCreate {
	sec.mutation.SetServerID(s)
	return sec
}

// SetEmoteID sets the "emote_id" field.
func (sec *ServerEmoteCreate) SetEmoteID(s string) *ServerEmoteCreate {
	sec.mutation.SetEmoteID(s)
	return sec
}

// SetCode sets the "code" field.
func (sec *ServerEmoteCreate) SetCode(s string) *ServerEmoteCreate {
	sec.mutation.SetCode(s)
	return sec
}

// SetImageType sets the "image_type" field.
func (sec *ServerEmoteCreate) SetImageType(s string) *ServerEmoteCreate {
	sec.mutation.SetImageType(s)
	return sec
}

// Mutation returns the ServerEmoteMutation object of the builder.
func (sec *ServerEmoteCreate) Mutation() *ServerEmoteMutation {
	return sec.mutation
}

// Save creates the ServerEmote in the database.
func (sec *ServerEmoteCreate) Save(ctx context.Context) (*ServerEmote, error) {
	var (
		err  error
		node *ServerEmote
	)
	if len(sec.hooks) == 0 {
		if err = sec.check(); err != nil {
			return nil, err
		}
		node, err = sec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServerEmoteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sec.check(); err != nil {
				return nil, err
			}
			sec.mutation = mutation
			if node, err = sec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sec.hooks) - 1; i >= 0; i-- {
			if sec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sec *ServerEmoteCreate) SaveX(ctx context.Context) *ServerEmote {
	v, err := sec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sec *ServerEmoteCreate) Exec(ctx context.Context) error {
	_, err := sec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sec *ServerEmoteCreate) ExecX(ctx context.Context) {
	if err := sec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sec *ServerEmoteCreate) check() error {
	if _, ok := sec.mutation.ServerID(); !ok {
		return &ValidationError{Name: "server_id", err: errors.New(`ent: missing required field "ServerEmote.server_id"`)}
	}
	if _, ok := sec.mutation.EmoteID(); !ok {
		return &ValidationError{Name: "emote_id", err: errors.New(`ent: missing required field "ServerEmote.emote_id"`)}
	}
	if _, ok := sec.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "ServerEmote.code"`)}
	}
	if _, ok := sec.mutation.ImageType(); !ok {
		return &ValidationError{Name: "image_type", err: errors.New(`ent: missing required field "ServerEmote.image_type"`)}
	}
	return nil
}

func (sec *ServerEmoteCreate) sqlSave(ctx context.Context) (*ServerEmote, error) {
	_node, _spec := sec.createSpec()
	if err := sqlgraph.CreateNode(ctx, sec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sec *ServerEmoteCreate) createSpec() (*ServerEmote, *sqlgraph.CreateSpec) {
	var (
		_node = &ServerEmote{config: sec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: serveremote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serveremote.FieldID,
			},
		}
	)
	if value, ok := sec.mutation.ServerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serveremote.FieldServerID,
		})
		_node.ServerID = value
	}
	if value, ok := sec.mutation.EmoteID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serveremote.FieldEmoteID,
		})
		_node.EmoteID = value
	}
	if value, ok := sec.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serveremote.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := sec.mutation.ImageType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: serveremote.FieldImageType,
		})
		_node.ImageType = value
	}
	return _node, _spec
}

// ServerEmoteCreateBulk is the builder for creating many ServerEmote entities in bulk.
type ServerEmoteCreateBulk struct {
	config
	builders []*ServerEmoteCreate
}

// Save creates the ServerEmote entities in the database.
func (secb *ServerEmoteCreateBulk) Save(ctx context.Context) ([]*ServerEmote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(secb.builders))
	nodes := make([]*ServerEmote, len(secb.builders))
	mutators := make([]Mutator, len(secb.builders))
	for i := range secb.builders {
		func(i int, root context.Context) {
			builder := secb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServerEmoteMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, secb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, secb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, secb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (secb *ServerEmoteCreateBulk) SaveX(ctx context.Context) []*ServerEmote {
	v, err := secb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (secb *ServerEmoteCreateBulk) Exec(ctx context.Context) error {
	_, err := secb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (secb *ServerEmoteCreateBulk) ExecX(ctx context.Context) {
	if err := secb.Exec(ctx); err != nil {
		panic(err)
	}
}