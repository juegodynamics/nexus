import React from "react";
import { StoryFn, Meta } from "@storybook/react";
import Frame from "../Frame";
import { CssBaseline, ThemeProvider } from "@mui/material";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/AnimatedFrame",
  component: Frame,
} as Meta<typeof Frame>;

const Template: StoryFn<typeof Frame> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame {...args} />
  </ThemeProvider>
);

export const AnimatedFrame = Template.bind({});
AnimatedFrame.args = {
  sx: {
    "@keyframes pulse": {
      "0%": {
        boxShadow: "0 0 10px #00ffff",
      },
      "50%": {
        boxShadow: "0 0 20px #00ffff",
      },
      "100%": {
        boxShadow: "0 0 10px #00ffff",
      },
    },
    animation: "pulse 2s infinite",
  },
  children: "Animated Frame",
};
