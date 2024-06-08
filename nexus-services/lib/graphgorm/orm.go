package graphgorm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// DgraphClient is an interface for Dgraph client methods
type DgraphClient interface {
	Alter(ctx context.Context, op *api.Operation) error
	NewTxn() DgraphTxn
}

type dgraphClient struct {
	*dgo.Dgraph
}

func (dc *dgraphClient) NewTxn() DgraphTxn {
	return dc.NewTxn()
}

func newDgraphClient(conn *grpc.ClientConn) DgraphClient {
	return &dgraphClient{dgo.NewDgraphClient(api.NewDgraphClient(conn))}
}

// DgraphTxn is an interface for Dgraph transaction methods
type DgraphTxn interface {
	Mutate(ctx context.Context, mu *api.Mutation) (*api.Response, error)
	Query(ctx context.Context, q string) (*api.Response, error)
	Discard(ctx context.Context) error
}

// DgraphORM provides methods to interact with Dgraph
type DgraphORM struct {
	client DgraphClient
}

// NewDgraphORM creates a new instance of DgraphORM
func NewDgraphORM(dgraphURL string) (*DgraphORM, error) {
	conn, err := grpc.Dial(dgraphURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Dgraph: %v", err)
	}

	return &DgraphORM{client: newDgraphClient(conn)}, nil
}

// CreateSchema creates the GraphQL schema in Dgraph
func (orm *DgraphORM) CreateSchema(schema string) error {
	op := &api.Operation{Schema: schema}
	return orm.client.Alter(context.Background(), op)
}

// AddData adds data to Dgraph
func (orm *DgraphORM) AddData(data interface{}) (*api.Response, error) {
	mutation := &api.Mutation{
		SetJson:   mustMarshal(data),
		CommitNow: true,
	}
	txn := orm.client.NewTxn()
	defer txn.Discard(context.Background())

	resp, err := txn.Mutate(context.Background(), mutation)
	if err != nil {
		return &api.Response{}, fmt.Errorf("unable to add data to Dgraph: %v", err)
	}
	return resp, nil
}

// QueryData queries data from Dgraph
func (orm *DgraphORM) QueryData(query string) (map[string]interface{}, error) {
	resp, err := orm.client.NewTxn().Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to query data from Dgraph: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Json, &result); err != nil {
		return nil, fmt.Errorf("unable to unmarshal query response: %v", err)
	}
	return result, nil
}

// mustMarshal marshals data into JSON and panics on error
func mustMarshal(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("unable to marshal data: %v", err)
	}
	return data
}

// CreateGraphQLSchema creates a GraphQL schema for a given struct
func (orm *DgraphORM) CreateGraphQLSchema(v interface{}) error {
	schema := GenerateGraphQLSchema(v)
	return orm.CreateSchema(schema)
}

// CreateGraphQLInput creates a GraphQL input type for a given struct
func (orm *DgraphORM) CreateGraphQLInput(v interface{}) error {
	input := GenerateGraphQLInput(v)
	return orm.CreateSchema(input)
}

// CreateGraphQLEnum creates a GraphQL enum for a given struct
func (orm *DgraphORM) CreateGraphQLEnum(v interface{}) error {
	enum := GenerateGraphQLEnum(v)
	return orm.CreateSchema(enum)
}
