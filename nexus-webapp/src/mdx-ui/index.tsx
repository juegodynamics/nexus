import React, { BlockquoteHTMLAttributes } from "react";
import FadeBlock from "../components/FadeBlock";
import { MDXProvider } from "@mdx-js/react";
import { Typography } from "@mui/material";

export const MDXBlockquote: React.FC<BlockquoteHTMLAttributes<{}>> = ({
  children,
}) => {
  console.log({ children });
  // Assuming children is an array with a single p element
  const text =
    (Array.isArray(children) && children?.[0]?.props?.children) || "";

  return (
    <FadeBlock>
      <Typography variant="body1"> {text}</Typography>
    </FadeBlock>
  );
};

export const mdxComponents: Parameters<typeof MDXProvider>[0]["components"] = {
  blockquote: MDXBlockquote,
};
