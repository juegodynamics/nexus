package graphgorm

import (
	"context"
	"testing"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDgraphClient is a mock implementation of the DgraphClient interface
type MockDgraphClient struct {
	mock.Mock
}

func (m *MockDgraphClient) Alter(ctx context.Context, op *api.Operation) error {
	args := m.Called(ctx, op)
	return args.Error(0)
}

func (m *MockDgraphClient) NewTxn() DgraphTxn {
	args := m.Called()
	return args.Get(0).(DgraphTxn)
}

// MockDgraphTxn is a mock implementation of the DgraphTxn interface
type MockDgraphTxn struct {
	mock.Mock
}

func (m *MockDgraphTxn) Mutate(ctx context.Context, mu *api.Mutation) (*api.Response, error) {
	args := m.Called(ctx, mu)
	return args.Get(0).(*api.Response), args.Error(0)
}

func (m *MockDgraphTxn) Query(ctx context.Context, q string) (*api.Response, error) {
	args := m.Called(ctx, q)
	return args.Get(0).(*api.Response), args.Error(0)
}

func (m *MockDgraphTxn) Discard(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestNewDgraphORM(t *testing.T) {
	orm, err := NewDgraphORM("localhost:9080")
	assert.NoError(t, err)
	assert.NotNil(t, orm)
}

func TestDgraphORM_CreateSchema(t *testing.T) {
	mockClient := new(MockDgraphClient)
	orm := &DgraphORM{client: mockClient}
	schema := "type Person { name: string }"

	mockClient.On("Alter", mock.Anything, &api.Operation{Schema: schema}).Return(nil)

	err := orm.CreateSchema(schema)
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDgraphORM_AddData(t *testing.T) {
	mockClient := new(MockDgraphClient)
	mockTxn := new(MockDgraphTxn)
	orm := &DgraphORM{client: mockClient}
	data := map[string]string{"name": "John"}

	mockClient.On("NewTxn").Return(mockTxn)
	expectedResponse := &api.Response{Json: []byte(`{"uid": "0x1"}`)}
	mockTxn.On("Mutate", mock.Anything, mock.Anything).Return(expectedResponse, nil).Run(func(args mock.Arguments) {
		mutation := args.Get(1).(*api.Mutation)
		t.Logf("Received mutation: %+v", mutation)
		if string(mutation.SetJson) != `{"name":"John"}` {
			t.Fatalf("Expected mutation data to be %s but got %s", `{"name":"John"}`, string(mutation.SetJson))
		}
	})
	mockTxn.On("Discard", mock.Anything).Return(nil)

	resp, err := orm.AddData(data)
	t.Logf("\nReceived response: %s\n", resp)
	assert.NoError(t, err)
	assert.Equal(t, `{"uid": "0x1"}`, resp)
	mockClient.AssertExpectations(t)
	mockTxn.AssertExpectations(t)
}

func TestDgraphORM_QueryData(t *testing.T) {
	mockClient := new(MockDgraphClient)
	mockTxn := new(MockDgraphTxn)
	orm := &DgraphORM{client: mockClient}
	query := "{ me(func: has(name)) { name } }"
	expectedResult := map[string]interface{}{"data": "test"}

	mockClient.On("NewTxn").Return(mockTxn)
	mockTxn.On("Query", mock.Anything, query).Return(&api.Response{Json: mustMarshal(expectedResult)}, nil)

	result, err := orm.QueryData(query)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	mockClient.AssertExpectations(t)
	mockTxn.AssertExpectations(t)
}

func TestDgraphORM_CreateGraphQLSchema(t *testing.T) {
	mockClient := new(MockDgraphClient)
	orm := &DgraphORM{client: mockClient}
	schema := "type Person { name: string }"

	mockClient.On("Alter", mock.Anything, &api.Operation{Schema: schema}).Return(nil)

	err := orm.CreateGraphQLSchema(struct {
		Name string `json:"name"`
	}{})
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDgraphORM_CreateGraphQLInput(t *testing.T) {
	mockClient := new(MockDgraphClient)
	orm := &DgraphORM{client: mockClient}
	input := "input PersonInput { name: string }"

	mockClient.On("Alter", mock.Anything, &api.Operation{Schema: input}).Return(nil)

	err := orm.CreateGraphQLInput(struct {
		Name string `json:"name"`
	}{})
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDgraphORM_CreateGraphQLEnum(t *testing.T) {
	mockClient := new(MockDgraphClient)
	orm := &DgraphORM{client: mockClient}
	enum := "enum Status { ACTIVE INACTIVE }"

	mockClient.On("Alter", mock.Anything, &api.Operation{Schema: enum}).Return(nil)

	err := orm.CreateGraphQLEnum(struct {
		Status string `json:"status"`
	}{})
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}
