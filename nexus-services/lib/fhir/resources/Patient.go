package resources

// Patient: Demographics and other administrative information about an
// individual or animal receiving care or other health-related services.
type Patient struct {
	// Demographics and other administrative information about an individual or
	// animal receiving care or other health-related services.
	Patient interface{} `json:"Patient,omitempty"`

	// The logical id of the resource, as used in the URL for the resource. Once
	// assigned, this value never changes. Within the context of the FHIR RESTful
	// interactions, the resource has an id except for cases like the create and
	// conditional update. Otherwise, the use of the resouce id depends on the given
	// use case.
	PatientId string `json:"PatientId,omitempty"`

	// The metadata about the resource. This is content that is maintained by the
	// infrastructure. Changes to the content might not always be associated with
	// version changes to the resource.
	PatientMeta *Meta `json:"PatientMeta,omitempty"`
	// A reference to a set of rules that were followed when the resource was
	// constructed, and which must be understood when processing the content. Often,
	// this is a reference to an implementation guide that defines the special rules
	// along with other profiles etc. Asserting this rule set restricts the content
	// to be only understood by a limited set of trading partners. This inherently
	// limits the usefulness of the data in the long term. However, the existing
	// health eco-system is highly fractured, and not yet ready to define, collect,
	// and exchange data in a generally computable sense. Wherever possible,
	// implementers and/or specification writers should avoid using this element.
	// Often, when used, the URL is a reference to an implementation guide that
	// defines these special rules as part of its narrative along with other
	// profiles, value sets, etc.
	PatientImplicitRules *uri `json:"PatientImplicitRules,omitempty"`
	// The base language in which the resource is written. Language is provided to
	// support indexing and accessibility (typically, services such as text to
	// speech use the language tag). The html language tag in the narrative applies
	// to the narrative. The language tag on the resource may be used to specify the
	// language of other presentations generated from the data in the resource. Not
	// all the content has to be in the base language. The Resource.language should
	// not be assumed to apply to the narrative automatically. If a language is
	// specified, it should it also be specified on the div element in the html (see
	// rules in HTML5 for information about the relationship between xml:lang and
	// the html lang attribute).
	PatientLanguage *code `json:"PatientLanguage,omitempty"`
	// A human-readable narrative that contains a summary of the resource and can be
	// used to represent the content of the resource to a human. The narrative need
	// not encode all the structured data, but is required to contain sufficient
	// detail to make it "clinically safe" for a human to just read the narrative.
	// Resource definitions may define what content should be represented in the
	// narrative to ensure clinical safety. Contained resources do not have a
	// narrative. Resources that are not contained SHOULD have a narrative. In some
	// cases, a resource may only have text with little or no additional discrete
	// data (as long as all minOccurs=1 elements are satisfied). This may be
	// necessary for data from legacy systems where information is captured as a
	// "text blob" or where text is additionally entered raw or narrated and encoded
	// information is added later.
	PatientText *Narrative `json:"PatientText,omitempty"`
	// These resources do not have an independent existence apart from the resource
	// that contains them - they cannot be identified independently, nor can they
	// have their own independent transaction scope. This is allowed to be a
	// Parameters resource if and only if it is referenced by a resource that
	// provides context/meaning. This should never be done when the content can be
	// identified properly, as once identification is lost, it is extremely
	// difficult (and context dependent) to restore it again. Contained resources
	// may have profiles and tags in their meta elements, but SHALL NOT have
	// security labels.
	PatientContained []*Resource `json:"PatientContained,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the resource. To make the use of extensions safe and managable,
	// there is a strict set of governance applied to the definition and use of
	// extensions. Though any implementer can define an extension, there is a set of
	// requirements that SHALL be met as part of the definition of the extension.
	// There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientExtension []*Extension `json:"PatientExtension,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the resource and that modifies the understanding of the element
	// that contains it and/or the understanding of the containing element's
	// descendants. Usually modifier elements provide negation or qualification. To
	// make the use of extensions safe and managable, there is a strict set of
	// governance applied to the definition and use of extensions. Though any
	// implementer is allowed to define an extension, there is a set of requirements
	// that SHALL be met as part of the definition of the extension. Applications
	// processing a resource are required to check for modifier extensions. Modifier
	// extensions SHALL NOT change the meaning of any elements on Resource or
	// DomainResource (including cannot change the meaning of modifierExtension
	// itself). There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientModifierExtension []*Extension `json:"PatientModifierExtension,omitempty"`
	// An identifier for this patient.
	PatientIdentifier []*Identifier `json:"PatientIdentifier,omitempty"`
	// Whether this patient record is in active use. Many systems use this property
	// to mark as non-current patients, such as those that have not been seen for a
	// period of time based on an organization's business rules. It is often used to
	// filter patient lists to exclude inactive patients Deceased patients may also
	// be marked as inactive for the same reasons, but may be active for some time
	// after death. If a record is inactive, and linked to an active record, then
	// future patient/record updates should occur on the other patient.
	PatientActive *boolean `json:"PatientActive,omitempty"`
	// A name associated with the individual. A patient may have multiple names with
	// different uses or applicable periods. For animals, the name is a "HumanName"
	// in the sense that is assigned and used by humans and has the same patterns.
	// Animal names may be communicated as given names, and optionally may include a
	// family name.
	PatientName []*HumanName `json:"PatientName,omitempty"`
	// A contact detail (e.g. a telephone number or an email address) by which the
	// individual may be contacted. A Patient may have multiple ways to be contacted
	// with different uses or applicable periods. May need to have options for
	// contacting the person urgently and also to help with identification. The
	// address might not go directly to the individual, but may reach another party
	// that is able to proxy for the patient (i.e. home phone, or pet owner's
	// phone).
	PatientTelecom []*ContactPoint `json:"PatientTelecom,omitempty"`
	// Administrative Gender - the gender that the patient is considered to have for
	// administration and record keeping purposes. The gender might not match the
	// biological sex as determined by genetics or the individual's preferred
	// identification. Note that for both humans and particularly animals, there are
	// other legitimate possibilities than male and female, though the vast majority
	// of systems and contexts only support male and female. Systems providing
	// decision support or enforcing business rules should ideally do this on the
	// basis of Observations dealing with the specific sex or gender aspect of
	// interest (anatomical, chromosomal, social, etc.) However, because these
	// observations are infrequently recorded, defaulting to the administrative
	// gender is common practice. Where such defaulting occurs, rule enforcement
	// should allow for the variation between administrative and biological,
	// chromosomal and other gender aspects. For example, an alert about a
	// hysterectomy on a male should be handled as a warning or overridable error,
	// not a "hard" error. See the Patient Gender and Sex section for additional
	// information about communicating patient gender and sex.
	PatientGender *code `json:"PatientGender,omitempty"`
	// The date of birth for the individual. Partial dates are allowed if the
	// specific date of birth is unknown. There is a standard extension
	// "patient-birthTime" available that should be used where Time is required
	// (such as in maternity/infant care systems).
	PatientBirthDate *date `json:"PatientBirthDate,omitempty"`
	// Indicates if the individual is deceased or not. If there's no value in the
	// instance, it means there is no statement on whether or not the individual is
	// deceased. Most systems will interpret the absence of a value as a sign of the
	// person being alive.
	PatientDeceased [X]*boolean `json:"PatientDeceased[X],omitempty"`
	// An address for the individual. Patient may have multiple addresses with
	// different uses or applicable periods.
	PatientAddress []*Address `json:"PatientAddress,omitempty"`
	// This field contains a patient's most recent marital (civil) status.
	PatientMaritalStatus *CodeableConcept `json:"PatientMaritalStatus,omitempty"`
	// Indicates whether the patient is part of a multiple (boolean) or indicates
	// the actual birth order (integer). Where the valueInteger is provided, the
	// number is the birth number in the sequence. E.g. The middle birth in triplets
	// would be valueInteger=2 and the third born would have valueInteger=3 If a
	// boolean value was provided for this triplets example, then all 3 patient
	// records would have valueBoolean=true (the ordering is not indicated).
	PatientMultipleBirth [X]*boolean `json:"PatientMultipleBirth[X],omitempty"`
	// Image of the patient. Guidelines: * Use id photos, not clinical photos. *
	// Limit dimensions to thumbnail. * Keep byte count low to ease resource
	// updates.
	PatientPhoto []*Attachment `json:"PatientPhoto,omitempty"`
	// A contact party (e.g. guardian, partner, friend) for the patient. Contact
	// covers all kinds of contact parties: family members, business contacts,
	// guardians, caregivers. Not applicable to register pedigree and family ties
	// beyond use of having contact.
	PatientContact []*BackboneElement `json:"PatientContact,omitempty"`
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	PatientContactId string `json:"PatientContactId,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the element. To make the use of extensions safe and managable,
	// there is a strict set of governance applied to the definition and use of
	// extensions. Though any implementer can define an extension, there is a set of
	// requirements that SHALL be met as part of the definition of the extension.
	// There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientContactExtension []*Extension `json:"PatientContactExtension,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the element and that modifies the understanding of the element
	// in which it is contained and/or the understanding of the containing element's
	// descendants. Usually modifier elements provide negation or qualification. To
	// make the use of extensions safe and managable, there is a strict set of
	// governance applied to the definition and use of extensions. Though any
	// implementer can define an extension, there is a set of requirements that
	// SHALL be met as part of the definition of the extension. Applications
	// processing a resource are required to check for modifier extensions. Modifier
	// extensions SHALL NOT change the meaning of any elements on Resource or
	// DomainResource (including cannot change the meaning of modifierExtension
	// itself). There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientContactModifierExtension []*Extension `json:"PatientContactModifierExtension,omitempty"`
	// The nature of the relationship between the patient and the contact person.
	PatientContactRelationship []*CodeableConcept `json:"PatientContactRelationship,omitempty"`
	// A name associated with the contact person.
	PatientContactName *HumanName `json:"PatientContactName,omitempty"`
	// A contact detail for the person, e.g. a telephone number or an email address.
	// Contact may have multiple ways to be contacted with different uses or
	// applicable periods. May need to have options for contacting the person
	// urgently, and also to help with identification.
	PatientContactTelecom []*ContactPoint `json:"PatientContactTelecom,omitempty"`
	// Address for the contact person.
	PatientContactAddress *Address `json:"PatientContactAddress,omitempty"`
	// Administrative Gender - the gender that the contact person is considered to
	// have for administration and record keeping purposes.
	PatientContactGender *code `json:"PatientContactGender,omitempty"`
	// Organization on behalf of which the contact is acting or for which the
	// contact is working.
	PatientContactOrganization *Reference `json:"PatientContactOrganization,omitempty"`
	// The period during which this contact person or organization is valid to be
	// contacted relating to this patient.
	PatientContactPeriod *Period `json:"PatientContactPeriod,omitempty"`
	// A language which may be used to communicate with the patient about his or her
	// health. If no language is specified, this *implies* that the default local
	// language is spoken. If you need to convey proficiency for multiple modes,
	// then you need multiple Patient.Communication associations. For animals,
	// language is not a relevant field, and should be absent from the instance. If
	// the Patient does not speak the default local language, then the Interpreter
	// Required Standard can be used to explicitly declare that an interpreter is
	// required.
	PatientCommunication []*BackboneElement `json:"PatientCommunication,omitempty"`
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	PatientCommunicationId string `json:"PatientCommunicationId,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the element. To make the use of extensions safe and managable,
	// there is a strict set of governance applied to the definition and use of
	// extensions. Though any implementer can define an extension, there is a set of
	// requirements that SHALL be met as part of the definition of the extension.
	// There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientCommunicationExtension []*Extension `json:"PatientCommunicationExtension,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the element and that modifies the understanding of the element
	// in which it is contained and/or the understanding of the containing element's
	// descendants. Usually modifier elements provide negation or qualification. To
	// make the use of extensions safe and managable, there is a strict set of
	// governance applied to the definition and use of extensions. Though any
	// implementer can define an extension, there is a set of requirements that
	// SHALL be met as part of the definition of the extension. Applications
	// processing a resource are required to check for modifier extensions. Modifier
	// extensions SHALL NOT change the meaning of any elements on Resource or
	// DomainResource (including cannot change the meaning of modifierExtension
	// itself). There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientCommunicationModifierExtension []*Extension `json:"PatientCommunicationModifierExtension,omitempty"`
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally
	// followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper
	// case; e.g. "en" for English, or "en-US" for American English versus "en-AU"
	// for Australian English. The structure aa-BB with this exact casing is one the
	// most widely used notations for locale. However not all systems actually code
	// this but instead have it as free text. Hence CodeableConcept instead of code
	// as the data type.
	PatientCommunicationLanguage *CodeableConcept `json:"PatientCommunicationLanguage,omitempty"`
	// Indicates whether or not the patient prefers this language (over other
	// languages he masters up a certain level). This language is specifically
	// identified for communicating healthcare information.
	PatientCommunicationPreferred *boolean `json:"PatientCommunicationPreferred,omitempty"`
	// Patient's nominated care provider. This may be the primary care provider (in
	// a GP context), or it may be a patient nominated care manager in a
	// community/disability setting, or even organization that will provide people
	// to perform the care provider roles. It is not to be used to record Care
	// Teams, these should be in a CareTeam resource that may be linked to the
	// CarePlan or EpisodeOfCare resources. Multiple GPs may be recorded against the
	// patient for various reasons, such as a student that has his home GP listed
	// along with the GP at university during the school semesters, or a
	// "fly-in/fly-out" worker that has the onsite GP also included with his home GP
	// to remain aware of medical issues. Jurisdictions may decide that they can
	// profile this down to 1 if desired, or 1 per type.
	PatientGeneralPractitioner []*Reference `json:"PatientGeneralPractitioner,omitempty"`
	// Organization that is the custodian of the patient record. There is only one
	// managing organization for a specific patient record. Other organizations will
	// have their own Patient record, and may use the Link property to join the
	// records together (or a Person resource which can include confidence ratings
	// for the association).
	PatientManagingOrganization *Reference `json:"PatientManagingOrganization,omitempty"`
	// Link to a Patient or RelatedPerson resource that concerns the same actual
	// individual. There is no assumption that linked patient records have mutual
	// links.
	PatientLink []*BackboneElement `json:"PatientLink,omitempty"`
	// Unique id for the element within a resource (for internal references). This
	// may be any string value that does not contain spaces.
	PatientLinkId string `json:"PatientLinkId,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the element. To make the use of extensions safe and managable,
	// there is a strict set of governance applied to the definition and use of
	// extensions. Though any implementer can define an extension, there is a set of
	// requirements that SHALL be met as part of the definition of the extension.
	// There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientLinkExtension []*Extension `json:"PatientLinkExtension,omitempty"`
	// May be used to represent additional information that is not part of the basic
	// definition of the element and that modifies the understanding of the element
	// in which it is contained and/or the understanding of the containing element's
	// descendants. Usually modifier elements provide negation or qualification. To
	// make the use of extensions safe and managable, there is a strict set of
	// governance applied to the definition and use of extensions. Though any
	// implementer can define an extension, there is a set of requirements that
	// SHALL be met as part of the definition of the extension. Applications
	// processing a resource are required to check for modifier extensions. Modifier
	// extensions SHALL NOT change the meaning of any elements on Resource or
	// DomainResource (including cannot change the meaning of modifierExtension
	// itself). There can be no stigma associated with the use of extensions by any
	// application, project, or standard - regardless of the institution or
	// jurisdiction that uses or defines the extensions. The use of extensions is
	// what allows the FHIR specification to retain a core level of simplicity for
	// everyone.
	PatientLinkModifierExtension []*Extension `json:"PatientLinkModifierExtension,omitempty"`
	// Link to a Patient or RelatedPerson resource that concerns the same actual
	// individual. Referencing a RelatedPerson here removes the need to use a Person
	// record to associate a Patient and RelatedPerson as the same individual.
	PatientLinkOther *Reference `json:"PatientLinkOther,omitempty"`
	// The type of link between this patient resource and another patient resource.
	PatientLinkType *code `json:"PatientLinkType,omitempty"`
}
