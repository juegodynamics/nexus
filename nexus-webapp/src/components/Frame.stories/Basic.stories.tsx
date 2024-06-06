import React from "react";
import { StoryFn, Meta } from "@storybook/react";
import { ThemeProvider, CssBaseline } from "@mui/material";
import Frame from "../Frame";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/Basic",
  component: Frame,
} as Meta<typeof Frame>;

const Template: StoryFn<typeof Frame> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame {...args}>Basic Frame</Frame>
  </ThemeProvider>
);

export const Basic = Template.bind({});
Basic.args = {};
