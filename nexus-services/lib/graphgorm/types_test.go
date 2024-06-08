package graphgorm_test

import (
	"reflect"
	"testing"

	"github.com/juegodynamics/nexus/nexus-services/lib/graphgorm"
)

type Identifier struct {
	Use    string `graphql:"use"`
	System string `graphql:"system"`
	Value  string `graphql:"value"`
}

type HumanName struct {
	Use    string   `graphql:"use"`
	Text   string   `graphql:"text"`
	Family string   `graphql:"family"`
	Given  []string `graphql:"given"`
}

type Address struct {
	Use        string   `graphql:"use"`
	Type       string   `graphql:"type"`
	Text       string   `graphql:"text"`
	Line       []string `graphql:"line"`
	City       string   `graphql:"city"`
	District   string   `graphql:"district"`
	State      string   `graphql:"state"`
	PostalCode string   `graphql:"postalCode"`
	Country    string   `graphql:"country"`
}

type Patient struct {
	ID         string       `graphql:"id"`
	Identifier []Identifier `graphql:"identifier"`
	Name       []HumanName  `graphql:"name"`
	Gender     string       `graphql:"gender"`
	BirthDate  string       `graphql:"birthDate"`
	Address    []Address    `graphql:"address"`
}

func TestGenerateGraphQLSchema(t *testing.T) {
	schema := graphgorm.GenerateGraphQLSchema(Patient{})
	expected := `type Patient {
  id: id
  identifier: [Identifier]
  name: [HumanName]
  gender: gender
  birthDate: birthDate
  address: [Address]
}
`
	if schema != expected {
		t.Errorf("expected %s but got %s", expected, schema)
	}
}

func TestGenerateGraphQLQuery(t *testing.T) {
	query := graphgorm.GenerateGraphQLQuery(Patient{})
	expected := `query getPatient($id: ID!) {
  getPatient(id: $id) {
    id
    identifier {
      use
      system
      value
    }
    name {
      use
      text
      family
      given
    }
    gender
    birthDate
    address {
      use
      type
      text
      line
      city
      district
      state
      postalCode
      country
    }
  }
}`
	if query != expected {
		t.Errorf("expected %s but got %s", expected, query)
	}
}

func TestGenerateGraphQLMutation(t *testing.T) {
	mutation := graphgorm.GenerateGraphQLMutation(Patient{})
	expected := `mutation addPatient($input: AddPatientInput!) {
  addPatient(input: $input) {
    patient {
      id
      identifier {
        use
        system
        value
      }
      name {
        use
        text
        family
        given
      }
      gender
      birthDate
      address {
        use
        type
        text
        line
        city
        district
        state
        postalCode
        country
      }
    }
  }
}`
	if mutation != expected {
		t.Errorf("expected %s but got %s", expected, mutation)
	}
}

func TestGenerateGraphQLInput(t *testing.T) {
	input := graphgorm.GenerateGraphQLInput(Patient{})
	expected := `input AddPatientInput {
  id: id
  identifier: [Identifier]
  name: [HumanName]
  gender: gender
  birthDate: birthDate
  address: [Address]
}
`
	if input != expected {
		t.Errorf("expected %s but got %s", expected, input)
	}
}

func TestGenerateGraphQLEnum(t *testing.T) {
	type EnumStruct struct {
		Status string `graphql:"status"`
	}
	enum := graphgorm.GenerateGraphQLEnum(reflect.TypeOf(EnumStruct{}))
	expected := `enum EnumStruct {
  Status
}
`
	if enum != expected {
		t.Errorf("expected %s but got %s", expected, enum)
	}
}

func TestGenerateGraphQLSchemaWithNestedStruct(t *testing.T) {
	schema := graphgorm.GenerateGraphQLSchema(Patient{})
	expected := `type Patient {
  id: id
  identifier: [Identifier]
  name: [HumanName]
  gender: gender
  birthDate: birthDate
  address: [Address]
}
`
	if schema != expected {
		t.Errorf("expected %s but got %s", expected, schema)
	}
}

func TestGenerateGraphQLQueryWithNestedStruct(t *testing.T) {
	query := graphgorm.GenerateGraphQLQuery(Patient{})
	expected := `query getPatient($id: ID!) {
  getPatient(id: $id) {
    id
    identifier {
      use
      system
      value
    }
    name {
      use
      text
      family
      given
    }
    gender
    birthDate
    address {
      use
      type
      text
      line
      city
      district
      state
      postalCode
      country
    }
  }
}`
	if query != expected {
		t.Errorf("expected %s but got %s", expected, query)
	}
}

func TestGenerateGraphQLMutationWithNestedStruct(t *testing.T) {
	mutation := graphgorm.GenerateGraphQLMutation(Patient{})
	expected := `mutation addPatient($input: AddPatientInput!) {
  addPatient(input: $input) {
    patient {
      id
      identifier {
        use
        system
        value
      }
      name {
        use
        text
        family
        given
      }
      gender
      birthDate
      address {
        use
        type
        text
        line
        city
        district
        state
        postalCode
        country
      }
    }
  }
}`
	if mutation != expected {
		t.Errorf("expected %s but got %s", expected, mutation)
	}
}

func TestGenerateGraphQLInputWithNestedStruct(t *testing.T) {
	input := graphgorm.GenerateGraphQLInput(Patient{})
	expected := `input AddPatientInput {
  id: id
  identifier: [Identifier]
  name: [HumanName]
  gender: gender
  birthDate: birthDate
  address: [Address]
}
`
	if input != expected {
		t.Errorf("expected %s but got %s", expected, input)
	}
}
