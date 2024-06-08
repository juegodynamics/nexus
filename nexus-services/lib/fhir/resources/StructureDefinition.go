package resources

type StructureDefinition struct {
	ID          string                       `json:"id"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Snapshot    *StructureDefinitionSnapshot `json:"snapshot"`
}

type StructureDefinitionSnapshot struct {
	Element []*StructureDefinitionSnapshotElement `json:"element"`
}

type StructureDefinitionSnapshotElement struct {
	ID         string                                    `json:"id"`
	Path       string                                    `json:"path"`
	Min        uint                                      `json:"min"`
	Max        string                                    `json:"max"`
	Type       []*StructureDefinitionSnapshotElementType `json:"type"`
	Definition string                                    `json:"definition"`
	Comment    string                                    `json:"comment"`
}

type StructureDefinitionSnapshotElementType struct {
	Extension []*Extension `json:"extension,omitempty"`
	Code      string       `json:"code"`
}

type Extension struct {
	URL            string    `json:"url"`
	ValueBoolean   bool      `json:"valueBoolean,omitempty"`
	ValueInteger   int       `json:"valueInteger,omitempty"`
	ValueString    string    `json:"valueString,omitempty"`
	ValueDecimal   float64   `json:"valueDecimal,omitempty"`
	ValueDate      string    `json:"valueDate,omitempty"`
	ValueDateTime  string    `json:"valueDateTime,omitempty"`
	ValueTime      string    `json:"valueTime,omitempty"`
	ValueCode      string    `json:"valueCode,omitempty"`
	ValueCoding    Coding    `json:"valueCoding,omitempty"`
	ValueQuantity  Quantity  `json:"valueQuantity,omitempty"`
	ValueReference Reference `json:"valueReference,omitempty"`
}

type Coding struct {
	System  string `json:"system,omitempty"`
	Version string `json:"version,omitempty"`
	Code    string `json:"code,omitempty"`
	Display string `json:"display,omitempty"`
}

type Quantity struct {
	Value  float64 `json:"value,omitempty"`
	Unit   string  `json:"unit,omitempty"`
	System string  `json:"system,omitempty"`
	Code   string  `json:"code,omitempty"`
}

type Reference struct {
	Reference  string     `json:"reference,omitempty"`
	Type       string     `json:"type,omitempty"`
	Identifier Identifier `json:"identifier,omitempty"`
	Display    string     `json:"display,omitempty"`
}

type Identifier struct {
	Use    string          `json:"use,omitempty"`
	Type   CodeableConcept `json:"type,omitempty"`
	System string          `json:"system,omitempty"`
	Value  string          `json:"value,omitempty"`
}

type CodeableConcept struct {
	Coding []Coding `json:"coding,omitempty"`
	Text   string   `json:"text,omitempty"`
}
