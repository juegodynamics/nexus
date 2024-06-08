package main

import (
	"fmt"

	datatypes "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/datatypes_go_proto"
	patient "github.com/google/fhir/go/proto/google/fhir/proto/r5/core/resources/patient_go_proto"
	"github.com/juegodynamics/nexus/nexus-services/lib/graphgorm"
)

func main() {
	fmt.Println(graphgorm.GenerateGraphQLSchema(patient.Patient{}))
	fmt.Println(graphgorm.GenerateGraphQLSchema(datatypes.HumanName{}))
}
