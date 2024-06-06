import fhir from "fhir/r5";

export const patient1: fhir.Patient = {
  resourceType: "Patient",
  id: "1",
  name: [{ use: "official", family: "Rodriguez", given: ["Jacovie"] }],
  gender: "male",
  birthDate: "1993-08-06",
  address: [
    {
      line: ["1003 GREENE AVE APT 6E"],
      city: "BROOKLYN",
      state: "NY",
      postalCode: "11221",
      use: "home",
    },
  ],
  telecom: [
    { system: "phone", value: "773-746-3306", use: "mobile" },
    { system: "email", value: "jacovie@alum.mit.edu" },
  ],
};

export const allergyIntolerance2: fhir.AllergyIntolerance = {
  resourceType: "AllergyIntolerance",
  id: "2",
  clinicalStatus: {
    coding: [
      {
        system:
          "http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical",
        code: "active",
      },
    ],
  },
  verificationStatus: {
    coding: [
      {
        system:
          "http://terminology.hl7.org/CodeSystem/allergyintolerance-verification",
        code: "confirmed",
      },
    ],
  },
  patient: { reference: "urn:uuid:1" },
  reaction: [],
};
