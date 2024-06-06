import React from "react";
import { StoryFn, Meta } from "@storybook/react";
import { ThemeProvider, CssBaseline } from "@mui/material";
import Frame from "../Frame";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/CustomContent",
  component: Frame,
} as Meta<typeof Frame>;

const Template: StoryFn<typeof Frame> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame {...args} />
  </ThemeProvider>
);

export const CustomContent = Template.bind({});
CustomContent.args = {
  children: (
    <div>
      <h2>Custom Content</h2>
      <p>This is an example of a frame with custom content inside.</p>
    </div>
  ),
};
