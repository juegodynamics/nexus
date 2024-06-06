import React from "react";
import { StoryFn, Meta } from "@storybook/react";
import { ThemeProvider, CssBaseline } from "@mui/material";
import Frame from "../Frame";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/Customizations",
  component: Frame,
} as Meta<typeof Frame>;

const Template: StoryFn<typeof Frame> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame {...args}>Customized Frame</Frame>
  </ThemeProvider>
);

export const CustomBorder = Template.bind({});
CustomBorder.args = {
  borderColor: "#ff00ff",
  borderWidth: 4,
};

export const Rounded = Template.bind({});
Rounded.args = {
  borderRadius: 20,
};

export const CustomShadowAndGlow = Template.bind({});
CustomShadowAndGlow.args = {
  shadowColor: "#ffcc00",
  glowColor: "#ff00ff",
};
