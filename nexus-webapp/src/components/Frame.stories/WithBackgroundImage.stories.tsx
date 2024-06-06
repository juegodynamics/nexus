import React from "react";
import { StoryFn, Meta } from "@storybook/react";
import { ThemeProvider, CssBaseline } from "@mui/material";
import Frame from "../Frame";
import theme from "../../theme/theme";

export default {
  title: "Components/Frame/WithBackgroundImage",
  component: Frame,
} as Meta<typeof Frame>;

const Template: StoryFn<typeof Frame> = (args) => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <Frame {...args} />
  </ThemeProvider>
);

export const WithBackgroundImage = Template.bind({});
WithBackgroundImage.args = {
  sx: {
    backgroundImage: 'url("https://via.placeholder.com/150")',
    backgroundSize: "cover",
  },
  children: "Frame with Background Image",
};
