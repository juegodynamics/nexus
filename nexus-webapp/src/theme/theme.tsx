import { createTheme } from "@mui/material/styles";
import { teal, cyan } from "@mui/material/colors";

const arwesTheme = createTheme({
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
      primary: "#00ffff",
      secondary: "#00ffcc",
    },
  },
  typography: {
    fontFamily: '"Orbitron", "Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      color: "#00ffff",
      textShadow: "0 0 10px #00ffff",
    },
    h2: {
      color: "#00ffcc",
      textShadow: "0 0 8px #00ffcc",
    },
    body1: {
      color: "#00ffff",
    },
    body2: {
      color: "#00ffcc",
    },
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: "5px",
          border: "1px solid #00ffff",
          backgroundColor: "#00334d",
          color: "#00ffff",
          textShadow: "0 0 5px #00ffff",
          "&:hover": {
            backgroundColor: "#004d66",
            boxShadow: "0 0 10px #00ffff",
          },
        },
      },
    },
  },
});

export default arwesTheme;
