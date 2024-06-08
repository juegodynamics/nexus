import fhir from "fhir/r5";

export const patient1: fhir.Patient = {
  id: "eEKEB3r9spcEhxqyBeMr8cRtuNFaastLBKU7cpoC.Iqk3",
  birthDate: "1993-08-06",
  resourceType: "Patient",
  active: true,
  telecom: [
    {
      system: "phone",
      value: "773-746-3306",
      use: "mobile",
    },
    {
      system: "email",
      value: "jacovie@alum.mit.edu",
      rank: 1,
    },
  ],
  contact: [
    {
      name: {
        use: "usual",
        text: "o'sullivan,john",
      },
      telecom: [
        {
          system: "phone",
          value: "786-376-2956",
          use: "mobile",
        },
      ],
      relationship: [
        {
          text: "Friend",
          coding: [
            {
              system: "http://terminology.hl7.org/CodeSystem/v2-0131",
              display: "Emergency Contact",
              code: "C",
            },
            {
              system: "urn:oid:1.2.840.114350.1.13.535.2.7.4.827665.1000",
              display: "Friend",
              code: "5",
            },
          ],
        },
      ],
    },
    {
      name: {
        use: "usual",
        text: "Martin,Tony",
      },
      telecom: [
        {
          system: "phone",
          value: "646-251-7869",
          use: "mobile",
        },
      ],
      relationship: [
        {
          text: "Other",
          coding: [
            {
              system: "http://terminology.hl7.org/CodeSystem/v2-0131",
              display: "Unknown",
              code: "U",
            },
            {
              system: "urn:oid:1.2.840.114350.1.13.535.2.7.4.827665.1000",
              display: "Other",
              code: "10",
            },
          ],
        },
      ],
    },
  ],
  deceasedBoolean: false,
  identifier: [
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.65.2.7.3.688884.100",
      value: "WMCLNDF5KVPW7NN",
      type: {
        text: "CEID",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.5.737384.0",
      value: "12420075",
      type: {
        text: "WCMC",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.2.698084",
      value: "Z12512843",
      type: {
        text: "EXTERNAL",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.5.737384.999334707",
      value: "A184X000005IQ5BAAE",
      type: {
        text: "FDTC",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.5.737384.999334707",
      value: "0014X00001QR2DCAAD",
      type: {
        text: "FDTC",
      },
    },
    {
      use: "usual",
      system:
        "http://open.epic.com/FHIR/StructureDefinition/patient-dstu2-fhir-id",
      value: "T.hO1YWbRr8.mhh9gKDt0JUtPKIqe6pRk-z-SZBTudb4B",
      type: {
        text: "FHIR",
      },
    },
    {
      use: "usual",
      system: "http://open.epic.com/FHIR/StructureDefinition/patient-fhir-id",
      value: "eEKEB3r9spcEhxqyBeMr8cRtuNFaastLBKU7cpoC.Iqk3",
      type: {
        text: "FHIR STU3",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.2.698084",
      value: " Z12512843",
      type: {
        text: "INTERNAL",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.3.878082.110",
      value: "JACOVIE",
      type: {
        text: "MYCHARTLOGIN",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.5.8.2.7",
      value: "1400956250",
      type: {
        text: "UI",
      },
    },
    {
      use: "usual",
      system: "urn:oid:1.2.840.114350.1.13.535.2.7.2.878082",
      value: "2740171",
      type: {
        text: "WPRINTERNAL",
      },
    },
    {
      use: "usual",
      system: "https://open.epic.com/FHIR/StructureDefinition/PayerMemberId",
      value: "10431563100",
    },
    {
      use: "usual",
      system: "https://open.epic.com/FHIR/StructureDefinition/PayerMemberId",
      value: "CFT866W16791",
    },
  ],
  address: [
    {
      use: "home",
      city: "BROOKLYN",
      district: "KINGS",
      postalCode: "11221-6223",
      country: "USA",
      line: ["1003 GREENE AVE APT 6E"],
      state: "NY",
    },
    {
      use: "old",
      city: "BROOKLYN",
      district: "KINGS",
      postalCode: "11221",
      country: "USA",
      line: ["1003 GREENE AVE, APT 6E"],
      state: "NY",
    },
  ],
  maritalStatus: {
    text: "SINGLE",
  },
  communication: [
    {
      preferred: true,
      language: {
        text: "English",
        coding: [
          {
            system: "urn:ietf:bcp:47",
            display: "English",
            code: "en",
          },
        ],
      },
    },
  ],
  extension: [
    {
      valueCodeableConcept: {
        coding: [
          {
            system:
              "urn:oid:1.2.840.114350.1.13.535.2.7.10.698084.130.657370.258999",
            display: "male",
            code: "male",
          },
        ],
      },
      url: "http://open.epic.com/FHIR/StructureDefinition/extension/legal-sex",
    },
    {
      valueCodeableConcept: {
        coding: [
          {
            system: "http://hl7.org/fhir/gender-identity",
            display: "male",
            code: "male",
          },
        ],
      },
      url: "http://hl7.org/fhir/StructureDefinition/patient-genderIdentity",
    },
    {
      url: "http://hl7.org/fhir/us/core/StructureDefinition/us-core-birthsex",
      valueCode: "M",
    },
    {
      url: "http://hl7.org/fhir/us/core/StructureDefinition/us-core-race",
      extension: [
        {
          valueCoding: {
            system: "urn:oid:2.16.840.1.113883.6.238",
            display: "Other Race",
            code: "2131-1",
          },
          url: "ombCategory",
        },
        {
          valueString: "OTHER COMBINATIONS NOT DESCRIBED",
          url: "text",
        },
      ],
    },
    {
      url: "http://hl7.org/fhir/us/core/StructureDefinition/us-core-ethnicity",
      extension: [
        {
          valueCoding: {
            system: "urn:oid:2.16.840.1.113883.6.238",
            display: "Hispanic or Latino",
            code: "2135-2",
          },
          url: "ombCategory",
        },
        {
          valueString: "Hispanic or Latino",
          url: "text",
        },
      ],
    },
  ],
  managingOrganization: {
    display: "Nyp/Columbia University",
    reference: "Organization/esC9DyQ2M8L8I3hEfySc-.Q3",
  },
  name: [
    {
      use: "official",
      text: "Jacovie Rodriguez",
      family: "Rodriguez",
      given: ["Jacovie"],
    },
    {
      use: "usual",
      text: "Jacovie Rodriguez",
      family: "Rodriguez",
      given: ["Jacovie"],
    },
  ],
  gender: "male",
};
