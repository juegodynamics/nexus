import React from "react";
import { Typography, Box, ButtonBase, ButtonBaseProps } from "@mui/material";
import Frame from "../Frame";
import { Patient } from "fhir/r5";

interface PatientCardProps extends ButtonBaseProps {
  patient: Patient;
}

const PatientCard: React.FC<PatientCardProps> = ({
  patient,
  ...buttonBaseProps
}) => {
  const name = patient.name?.[0];
  const fullName = name
    ? `${name.given?.join(" ")} ${name.family}`
    : "Unknown Name";
  const address = patient.address?.[0];

  return (
    <ButtonBase
      {...buttonBaseProps}
      sx={{
        display: "block",
        textAlign: "left",
        width: "100%",
        cursor: "pointer",
        "&:hover": {
          ".frame": {
            boxShadow: "0 0 20px #00ffcc",
          },
        },
      }}
    >
      <Frame
        className="frame"
        borderColor="#00ffff"
        borderWidth={2}
        borderRadius={10}
        shadowColor="#00ffcc"
        glowColor="#00ffff"
        p={2}
      >
        <Typography variant="h5" component="h2" gutterBottom>
          {fullName}
        </Typography>
        <Typography variant="body1" gutterBottom>
          Gender: {patient.gender}
        </Typography>
        <Typography variant="body1" gutterBottom>
          Birth Date: {patient.birthDate}
        </Typography>
        {address && (
          <Box mt={2}>
            <Typography variant="body1">Address:</Typography>
            <Typography variant="body2">{address.line?.join(", ")}</Typography>
            <Typography variant="body2">
              {address.city}, {address.state} {address.postalCode}
            </Typography>
          </Box>
        )}
      </Frame>
    </ButtonBase>
  );
};

export default PatientCard;
