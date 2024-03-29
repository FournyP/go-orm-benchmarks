// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/FournyP/go-orm-benchmarks/ent/comment"
	"github.com/FournyP/go-orm-benchmarks/ent/post"
	"github.com/FournyP/go-orm-benchmarks/ent/predicate"
	"github.com/FournyP/go-orm-benchmarks/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PostUpdate) SetTitle(s string) *PostUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetContent sets the "content" field.
func (pu *PostUpdate) SetContent(s string) *PostUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PostUpdate) SetUserID(u uint) *PostUpdate {
	pu.mutation.SetUserID(u)
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PostUpdate) SetUser(u *User) *PostUpdate {
	return pu.SetUserID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pu *PostUpdate) AddCommentIDs(ids ...uint) *PostUpdate {
	pu.mutation.AddCommentIDs(ids...)
	return pu
}

// AddComments adds the "comments" edges to the Comment entity.
func (pu *PostUpdate) AddComments(c ...*Comment) *PostUpdate {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PostUpdate) ClearUser() *PostUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearComments clears all "comments" edges to the Comment entity.
func (pu *PostUpdate) ClearComments() *PostUpdate {
	pu.mutation.ClearComments()
	return pu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (pu *PostUpdate) RemoveCommentIDs(ids ...uint) *PostUpdate {
	pu.mutation.RemoveCommentIDs(ids...)
	return pu
}

// RemoveComments removes "comments" edges to Comment entities.
func (pu *PostUpdate) RemoveComments(c ...*Comment) *PostUpdate {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PostUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if _, ok := pu.mutation.UserID(); pu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Post.user"`)
	}
	return nil
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !pu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetTitle sets the "title" field.
func (puo *PostUpdateOne) SetTitle(s string) *PostUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetContent sets the "content" field.
func (puo *PostUpdateOne) SetContent(s string) *PostUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetUserID sets the "user_id" field.
func (puo *PostUpdateOne) SetUserID(u uint) *PostUpdateOne {
	puo.mutation.SetUserID(u)
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PostUpdateOne) SetUser(u *User) *PostUpdateOne {
	return puo.SetUserID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (puo *PostUpdateOne) AddCommentIDs(ids ...uint) *PostUpdateOne {
	puo.mutation.AddCommentIDs(ids...)
	return puo
}

// AddComments adds the "comments" edges to the Comment entity.
func (puo *PostUpdateOne) AddComments(c ...*Comment) *PostUpdateOne {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCommentIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PostUpdateOne) ClearUser() *PostUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearComments clears all "comments" edges to the Comment entity.
func (puo *PostUpdateOne) ClearComments() *PostUpdateOne {
	puo.mutation.ClearComments()
	return puo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (puo *PostUpdateOne) RemoveCommentIDs(ids ...uint) *PostUpdateOne {
	puo.mutation.RemoveCommentIDs(ids...)
	return puo
}

// RemoveComments removes "comments" edges to Comment entities.
func (puo *PostUpdateOne) RemoveComments(c ...*Comment) *PostUpdateOne {
	ids := make([]uint, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCommentIDs(ids...)
}

// Where appends a list predicates to the PostUpdate builder.
func (puo *PostUpdateOne) Where(ps ...predicate.Post) *PostUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PostUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := post.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Post.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Post.content": %w`, err)}
		}
	}
	if _, ok := puo.mutation.UserID(); puo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Post.user"`)
	}
	return nil
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(post.Table, post.Columns, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUint))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !puo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
