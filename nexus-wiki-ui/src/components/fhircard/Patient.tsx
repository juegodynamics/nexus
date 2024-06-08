import React from "react";
import {
  Card,
  CardContent,
  Typography,
  Grid,
  List,
  ListItem,
  ListItemText,
  Divider,
  Box,
  Avatar,
  Tooltip,
  useTheme,
} from "@mui/material";
import fhir from "fhir/r5";
import { LabelValue } from "./LabelValue";

const PatientCard: React.FC<{ patient: fhir.Patient }> = ({ patient }) => {
  const theme = useTheme();

  return (
    <Card
      sx={{
        transition: "transform 0.2s ease-in-out",
        maxWidth: "20vw",
        "&:hover": {
          transform: "scale(1.02)",
        },
        "& .MuiCardContent-root": {
          padding: "16px",
        },
      }}
    >
      <CardContent>
        <Box display="flex" alignItems="center" mb={2} flexDirection={"column"}>
          <Avatar alt={patient.name?.[0]?.text || "Unknown"}>
            {patient.name?.[0]?.text?.charAt(0) || "?"}
          </Avatar>
          <Typography variant="h4" sx={{ ml: 2 }}>
            {patient.name?.[0]?.text || "Unknown"}
          </Typography>
        </Box>
        <Grid container spacing={2}>
          <Grid item xs={3}>
            <LabelValue label="Gender" value={patient.gender || "Unknown"} />
            <LabelValue
              label="Birth Date"
              value={patient.birthDate || "Unknown"}
            />
            <LabelValue
              label="Active"
              value={
                patient.active !== undefined
                  ? patient.active
                    ? "Yes"
                    : "No"
                  : "Unknown"
              }
            />
            <LabelValue
              label="Deceased"
              value={
                patient.deceasedBoolean !== undefined
                  ? patient.deceasedBoolean
                    ? "Yes"
                    : "No"
                  : "Unknown"
              }
            />
            <Typography variant="h6" gutterBottom>
              Contact Information
            </Typography>
            <List dense>
              {patient.telecom?.length ? (
                patient.telecom.map((contact, index) => (
                  <Tooltip
                    key={index}
                    title={`Use: ${contact.use || "Unknown"}`}
                    arrow
                  >
                    <ListItem>
                      <ListItemText
                        primary={`${contact.system || "Unknown"}: ${contact.value || "Unknown"}`}
                      />
                    </ListItem>
                  </Tooltip>
                ))
              ) : (
                <ListItem>
                  <ListItemText primary="No contact information available" />
                </ListItem>
              )}
            </List>
          </Grid>
          {/* <Grid item xs={12} md={6}>
            <Typography variant="h6" gutterBottom>
              Addresses
            </Typography>
            <List dense>
              {patient.address?.length ? (
                patient.address.map((address, index) => (
                  <Tooltip
                    key={index}
                    title={`Use: ${address.use || "Unknown"}`}
                    arrow
                  >
                    <ListItem>
                      <ListItemText
                        primary={`${address.line?.join(", ") || "Unknown"}, ${address.city || "Unknown"}, ${address.state || "Unknown"} ${address.postalCode || "Unknown"}, ${address.country || "Unknown"}`}
                      />
                    </ListItem>
                  </Tooltip>
                ))
              ) : (
                <ListItem>
                  <ListItemText primary="No address information available" />
                </ListItem>
              )}
            </List>
            <Divider />
            <Typography variant="h6" gutterBottom>
              Identifiers
            </Typography>
            <List dense>
              {patient.identifier?.length ? (
                patient.identifier.map((id, index) => (
                  <Tooltip
                    key={index}
                    title={`Use: ${id.use || "Unknown"}${id.type ? `, Type: ${id.type.text || "Unknown"}` : ""}`}
                    arrow
                  >
                    <ListItem>
                      <ListItemText
                        primary={`${id.system || "Unknown"}: ${id.value || "Unknown"}`}
                      />
                    </ListItem>
                  </Tooltip>
                ))
              ) : (
                <ListItem>
                  <ListItemText primary="No identifier information available" />
                </ListItem>
              )}
            </List>
          </Grid> */}
        </Grid>
        <Divider />
        <Typography variant="h6" gutterBottom>
          Contact Persons
        </Typography>
        <List dense>
          {patient.contact?.length ? (
            patient.contact.map((contact, index) => (
              <Tooltip key={index} title="Click for more info" arrow>
                <ListItem>
                  <ListItemText
                    primary={contact.name?.text || "Unknown"}
                    secondary={
                      contact.telecom
                        ?.map(
                          (tele, idx) =>
                            `${tele.system || "Unknown"}: ${tele.value || "Unknown"}${tele.use ? ` (Use: ${tele.use})` : ""}`
                        )
                        .join(", ") || "No contact telecom information"
                    }
                  />
                </ListItem>
              </Tooltip>
            ))
          ) : (
            <ListItem>
              <ListItemText primary="No contact information available" />
            </ListItem>
          )}
        </List>
        <Divider />
        <Typography variant="h6" gutterBottom>
          Managing Organization
        </Typography>
        <Typography variant="body1" gutterBottom>
          {patient.managingOrganization?.display || "Unknown"}
        </Typography>
      </CardContent>
    </Card>
  );
};

export default PatientCard;
