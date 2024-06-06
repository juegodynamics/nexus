import React from "react";
import { StoryFn, Meta } from "@storybook/react";
import { ThemeProvider, CssBaseline } from "@mui/material";
import Frame from "../Frame";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/InteractiveElements",
  component: Frame,
} as Meta<typeof Frame>;

const Template: StoryFn<typeof Frame> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame {...args} />
  </ThemeProvider>
);

export const WithButtons = Template.bind({});
WithButtons.args = {
  children: (
    <div>
      <button>Button 1</button>
      <button>Button 2</button>
    </div>
  ),
};
