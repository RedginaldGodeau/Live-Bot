// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"backend/ent/gen/migrate"

	"backend/ent/gen/liveshow"
	"backend/ent/gen/upload"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// LiveShow is the client for interacting with the LiveShow builders.
	LiveShow *LiveShowClient
	// Upload is the client for interacting with the Upload builders.
	Upload *UploadClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.LiveShow = NewLiveShowClient(c.config)
	c.Upload = NewUploadClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("gen: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("gen: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		LiveShow: NewLiveShowClient(cfg),
		Upload:   NewUploadClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		LiveShow: NewLiveShowClient(cfg),
		Upload:   NewUploadClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		LiveShow.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.LiveShow.Use(hooks...)
	c.Upload.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.LiveShow.Intercept(interceptors...)
	c.Upload.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LiveShowMutation:
		return c.LiveShow.mutate(ctx, m)
	case *UploadMutation:
		return c.Upload.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("gen: unknown mutation type %T", m)
	}
}

// LiveShowClient is a client for the LiveShow schema.
type LiveShowClient struct {
	config
}

// NewLiveShowClient returns a client for the LiveShow from the given config.
func NewLiveShowClient(c config) *LiveShowClient {
	return &LiveShowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `liveshow.Hooks(f(g(h())))`.
func (c *LiveShowClient) Use(hooks ...Hook) {
	c.hooks.LiveShow = append(c.hooks.LiveShow, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `liveshow.Intercept(f(g(h())))`.
func (c *LiveShowClient) Intercept(interceptors ...Interceptor) {
	c.inters.LiveShow = append(c.inters.LiveShow, interceptors...)
}

// Create returns a builder for creating a LiveShow entity.
func (c *LiveShowClient) Create() *LiveShowCreate {
	mutation := newLiveShowMutation(c.config, OpCreate)
	return &LiveShowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LiveShow entities.
func (c *LiveShowClient) CreateBulk(builders ...*LiveShowCreate) *LiveShowCreateBulk {
	return &LiveShowCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LiveShowClient) MapCreateBulk(slice any, setFunc func(*LiveShowCreate, int)) *LiveShowCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LiveShowCreateBulk{err: fmt.Errorf("calling to LiveShowClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LiveShowCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LiveShowCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LiveShow.
func (c *LiveShowClient) Update() *LiveShowUpdate {
	mutation := newLiveShowMutation(c.config, OpUpdate)
	return &LiveShowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LiveShowClient) UpdateOne(ls *LiveShow) *LiveShowUpdateOne {
	mutation := newLiveShowMutation(c.config, OpUpdateOne, withLiveShow(ls))
	return &LiveShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LiveShowClient) UpdateOneID(id int) *LiveShowUpdateOne {
	mutation := newLiveShowMutation(c.config, OpUpdateOne, withLiveShowID(id))
	return &LiveShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LiveShow.
func (c *LiveShowClient) Delete() *LiveShowDelete {
	mutation := newLiveShowMutation(c.config, OpDelete)
	return &LiveShowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LiveShowClient) DeleteOne(ls *LiveShow) *LiveShowDeleteOne {
	return c.DeleteOneID(ls.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LiveShowClient) DeleteOneID(id int) *LiveShowDeleteOne {
	builder := c.Delete().Where(liveshow.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LiveShowDeleteOne{builder}
}

// Query returns a query builder for LiveShow.
func (c *LiveShowClient) Query() *LiveShowQuery {
	return &LiveShowQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLiveShow},
		inters: c.Interceptors(),
	}
}

// Get returns a LiveShow entity by its id.
func (c *LiveShowClient) Get(ctx context.Context, id int) (*LiveShow, error) {
	return c.Query().Where(liveshow.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LiveShowClient) GetX(ctx context.Context, id int) *LiveShow {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUpload queries the upload edge of a LiveShow.
func (c *LiveShowClient) QueryUpload(ls *LiveShow) *UploadQuery {
	query := (&UploadClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ls.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(liveshow.Table, liveshow.FieldID, id),
			sqlgraph.To(upload.Table, upload.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, liveshow.UploadTable, liveshow.UploadColumn),
		)
		fromV = sqlgraph.Neighbors(ls.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LiveShowClient) Hooks() []Hook {
	return c.hooks.LiveShow
}

// Interceptors returns the client interceptors.
func (c *LiveShowClient) Interceptors() []Interceptor {
	return c.inters.LiveShow
}

func (c *LiveShowClient) mutate(ctx context.Context, m *LiveShowMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LiveShowCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LiveShowUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LiveShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LiveShowDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("gen: unknown LiveShow mutation op: %q", m.Op())
	}
}

// UploadClient is a client for the Upload schema.
type UploadClient struct {
	config
}

// NewUploadClient returns a client for the Upload from the given config.
func NewUploadClient(c config) *UploadClient {
	return &UploadClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `upload.Hooks(f(g(h())))`.
func (c *UploadClient) Use(hooks ...Hook) {
	c.hooks.Upload = append(c.hooks.Upload, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `upload.Intercept(f(g(h())))`.
func (c *UploadClient) Intercept(interceptors ...Interceptor) {
	c.inters.Upload = append(c.inters.Upload, interceptors...)
}

// Create returns a builder for creating a Upload entity.
func (c *UploadClient) Create() *UploadCreate {
	mutation := newUploadMutation(c.config, OpCreate)
	return &UploadCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Upload entities.
func (c *UploadClient) CreateBulk(builders ...*UploadCreate) *UploadCreateBulk {
	return &UploadCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UploadClient) MapCreateBulk(slice any, setFunc func(*UploadCreate, int)) *UploadCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UploadCreateBulk{err: fmt.Errorf("calling to UploadClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UploadCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UploadCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Upload.
func (c *UploadClient) Update() *UploadUpdate {
	mutation := newUploadMutation(c.config, OpUpdate)
	return &UploadUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UploadClient) UpdateOne(u *Upload) *UploadUpdateOne {
	mutation := newUploadMutation(c.config, OpUpdateOne, withUpload(u))
	return &UploadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UploadClient) UpdateOneID(id int) *UploadUpdateOne {
	mutation := newUploadMutation(c.config, OpUpdateOne, withUploadID(id))
	return &UploadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Upload.
func (c *UploadClient) Delete() *UploadDelete {
	mutation := newUploadMutation(c.config, OpDelete)
	return &UploadDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UploadClient) DeleteOne(u *Upload) *UploadDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UploadClient) DeleteOneID(id int) *UploadDeleteOne {
	builder := c.Delete().Where(upload.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UploadDeleteOne{builder}
}

// Query returns a query builder for Upload.
func (c *UploadClient) Query() *UploadQuery {
	return &UploadQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUpload},
		inters: c.Interceptors(),
	}
}

// Get returns a Upload entity by its id.
func (c *UploadClient) Get(ctx context.Context, id int) (*Upload, error) {
	return c.Query().Where(upload.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UploadClient) GetX(ctx context.Context, id int) *Upload {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLiveShowUpload queries the live_show_upload edge of a Upload.
func (c *UploadClient) QueryLiveShowUpload(u *Upload) *LiveShowQuery {
	query := (&LiveShowClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(upload.Table, upload.FieldID, id),
			sqlgraph.To(liveshow.Table, liveshow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, upload.LiveShowUploadTable, upload.LiveShowUploadColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UploadClient) Hooks() []Hook {
	return c.hooks.Upload
}

// Interceptors returns the client interceptors.
func (c *UploadClient) Interceptors() []Interceptor {
	return c.inters.Upload
}

func (c *UploadClient) mutate(ctx context.Context, m *UploadMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UploadCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UploadUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UploadUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UploadDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("gen: unknown Upload mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		LiveShow, Upload []ent.Hook
	}
	inters struct {
		LiveShow, Upload []ent.Interceptor
	}
)
