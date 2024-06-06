import React from "react";
import { Meta } from "@storybook/react";
import { ThemeProvider, CssBaseline } from "@mui/material";
import Frame from "../Frame";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/NestedFrames",
  component: Frame,
} as Meta<typeof Frame>;

export const NestedFrames = () => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame
      borderColor="#ff0000"
      borderWidth={5}
      borderRadius={15}
      shadowColor="#ff0000"
      glowColor="#ff0000"
    >
      <Frame
        borderColor="#00ff00"
        borderWidth={4}
        borderRadius={10}
        shadowColor="#00ff00"
        glowColor="#00ff00"
      >
        <Frame
          borderColor="#0000ff"
          borderWidth={3}
          borderRadius={5}
          shadowColor="#0000ff"
          glowColor="#0000ff"
        >
          Nested Frame
        </Frame>
      </Frame>
    </Frame>
  </ThemeProvider>
);
