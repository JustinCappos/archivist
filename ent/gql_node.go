// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/testifysec/archivist/ent/attestation"
	"github.com/testifysec/archivist/ent/attestationcollection"
	"github.com/testifysec/archivist/ent/dsse"
	"github.com/testifysec/archivist/ent/payloaddigest"
	"github.com/testifysec/archivist/ent/signature"
	"github.com/testifysec/archivist/ent/statement"
	"github.com/testifysec/archivist/ent/subject"
	"github.com/testifysec/archivist/ent/subjectdigest"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (a *Attestation) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "Attestation",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(a.Type); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "type",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "AttestationCollection",
		Name: "attestation_collection",
	}
	err = a.QueryAttestationCollection().
		Select(attestationcollection.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ac *AttestationCollection) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ac.ID,
		Type:   "AttestationCollection",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(ac.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Attestation",
		Name: "attestations",
	}
	err = ac.QueryAttestations().
		Select(attestation.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Statement",
		Name: "statement",
	}
	err = ac.QueryStatement().
		Select(statement.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (d *Dsse) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     d.ID,
		Type:   "Dsse",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(d.GitoidSha256); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "gitoid_sha256",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.PayloadType); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "payload_type",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Statement",
		Name: "statement",
	}
	err = d.QueryStatement().
		Select(statement.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Signature",
		Name: "signatures",
	}
	err = d.QuerySignatures().
		Select(signature.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "PayloadDigest",
		Name: "payload_digests",
	}
	err = d.QueryPayloadDigests().
		Select(payloaddigest.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (pd *PayloadDigest) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pd.ID,
		Type:   "PayloadDigest",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(pd.Algorithm); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "algorithm",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pd.Value); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "value",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Dsse",
		Name: "dsse",
	}
	err = pd.QueryDsse().
		Select(dsse.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Signature) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "Signature",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(s.KeyID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "key_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(s.Signature); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "signature",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Dsse",
		Name: "dsse",
	}
	err = s.QueryDsse().
		Select(dsse.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Statement) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "Statement",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(s.Predicate); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "predicate",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Subject",
		Name: "subjects",
	}
	err = s.QuerySubjects().
		Select(subject.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "AttestationCollection",
		Name: "attestation_collections",
	}
	err = s.QueryAttestationCollections().
		Select(attestationcollection.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Dsse",
		Name: "dsse",
	}
	err = s.QueryDsse().
		Select(dsse.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *Subject) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     s.ID,
		Type:   "Subject",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(s.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "SubjectDigest",
		Name: "subject_digests",
	}
	err = s.QuerySubjectDigests().
		Select(subjectdigest.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Statement",
		Name: "statement",
	}
	err = s.QueryStatement().
		Select(statement.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (sd *SubjectDigest) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     sd.ID,
		Type:   "SubjectDigest",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(sd.Algorithm); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "algorithm",
		Value: string(buf),
	}
	if buf, err = json.Marshal(sd.Value); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "value",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Subject",
		Name: "subject",
	}
	err = sd.QuerySubject().
		Select(subject.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case attestation.Table:
		query := c.Attestation.Query().
			Where(attestation.ID(id))
		query, err := query.CollectFields(ctx, "Attestation")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case attestationcollection.Table:
		query := c.AttestationCollection.Query().
			Where(attestationcollection.ID(id))
		query, err := query.CollectFields(ctx, "AttestationCollection")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case dsse.Table:
		query := c.Dsse.Query().
			Where(dsse.ID(id))
		query, err := query.CollectFields(ctx, "Dsse")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case payloaddigest.Table:
		query := c.PayloadDigest.Query().
			Where(payloaddigest.ID(id))
		query, err := query.CollectFields(ctx, "PayloadDigest")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case signature.Table:
		query := c.Signature.Query().
			Where(signature.ID(id))
		query, err := query.CollectFields(ctx, "Signature")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case statement.Table:
		query := c.Statement.Query().
			Where(statement.ID(id))
		query, err := query.CollectFields(ctx, "Statement")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subject.Table:
		query := c.Subject.Query().
			Where(subject.ID(id))
		query, err := query.CollectFields(ctx, "Subject")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subjectdigest.Table:
		query := c.SubjectDigest.Query().
			Where(subjectdigest.ID(id))
		query, err := query.CollectFields(ctx, "SubjectDigest")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case attestation.Table:
		query := c.Attestation.Query().
			Where(attestation.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Attestation")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case attestationcollection.Table:
		query := c.AttestationCollection.Query().
			Where(attestationcollection.IDIn(ids...))
		query, err := query.CollectFields(ctx, "AttestationCollection")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case dsse.Table:
		query := c.Dsse.Query().
			Where(dsse.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Dsse")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case payloaddigest.Table:
		query := c.PayloadDigest.Query().
			Where(payloaddigest.IDIn(ids...))
		query, err := query.CollectFields(ctx, "PayloadDigest")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case signature.Table:
		query := c.Signature.Query().
			Where(signature.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Signature")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case statement.Table:
		query := c.Statement.Query().
			Where(statement.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Statement")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subject.Table:
		query := c.Subject.Query().
			Where(subject.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Subject")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subjectdigest.Table:
		query := c.SubjectDigest.Query().
			Where(subjectdigest.IDIn(ids...))
		query, err := query.CollectFields(ctx, "SubjectDigest")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
