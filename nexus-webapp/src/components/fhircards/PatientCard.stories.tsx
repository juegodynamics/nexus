import React from "react";
import { Meta, StoryFn } from "@storybook/react";
import { Patient } from "fhir/r5";
import PatientCard from "./PatientCard";
import { CssBaseline, ThemeProvider } from "@mui/material";
import theme from "../../theme/theme";

export default {
  title: "Components/PatientCard",
  component: PatientCard,
} as Meta<typeof PatientCard>;

const examplePatient: Patient = {
  resourceType: "Patient",
  id: "12345",
  name: [
    {
      family: "Doe",
      given: ["John", "A."],
    },
  ],
  gender: "male",
  birthDate: "1980-05-20",
  address: [
    {
      line: ["123 Main St"],
      city: "Anytown",
      state: "CA",
      postalCode: "90210",
    },
  ],
};

const Template: StoryFn<typeof PatientCard> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <PatientCard {...args} />
  </ThemeProvider>
);

export const Default = Template.bind({});
Default.args = {
  patient: examplePatient,
  onClick: () => alert("Patient Card Clicked!"),
};
