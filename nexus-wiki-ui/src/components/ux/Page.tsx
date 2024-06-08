// src/components/Page.tsx
import React from "react";
import { useParams } from "react-router-dom";
import { Typography } from "@mui/material";
import { findPageById } from "../../utils/findPageById";
import { pages } from "../../pages";

export const Page: React.FC = () => {
  const { projectId, pageId } = useParams<{
    projectId: string;
    pageId: string;
  }>();
  const page = findPageById(
    pageId || "",
    pages[projectId || Object.keys(pages)[0]]
  );

  if (!page) {
    return <Typography variant="h1">Page Not Found</Typography>;
  }

  return (
    <>
      <Typography variant="h1">{page.title}</Typography>
      {page.content}
    </>
  );
};
