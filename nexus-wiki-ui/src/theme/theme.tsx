import { createTheme } from "@mui/material/styles";
import { teal, cyan } from "@mui/material/colors";
import { Colors } from "./colors";

export const ScifiTheme = createTheme({
  palette: {
    primary: {
      main: cyan[500],
    },
    secondary: {
      main: teal[500],
    },
    background: {
      default: "#253238",
      paper: "#192226",
    },
    text: {
      primary: Colors.Aqua,
      secondary: Colors.LemonGlacier,
    },
  },
  typography: {
    fontFamily: '"Orbitron", "Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      color: Colors.Aqua,
      textShadow: `0 0 10px ${Colors.Aqua}`,
    },
    h2: {
      color: Colors.Aqua,
      textShadow: `0 0 8px ${Colors.Aqua}`,
    },
    body1: {
      color: Colors.Aqua,
    },
    body2: {
      color: Colors.Aqua,
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: "5px",
          border: `1px solid ${Colors.Aqua}`,
          backgroundColor: "#00334d",
          color: Colors.Aqua,
          textShadow: `0 0 5px ${Colors.Aqua}`,
          "&:hover": {
            backgroundColor: "#004d66",
            boxShadow: "0 0 10px #00ffff",
          },
        },
      },
    },
  },
});
