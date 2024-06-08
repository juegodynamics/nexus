// src/components/Page.tsx
import React from "react";
import { useParams } from "react-router-dom";
import { Typography } from "@mui/material";
import { findPageById } from "../../utils/findPageById";
import { pages } from "../../pages";

export interface PageParams {
  projectId: string;
  pageId: string;
}

export const Page: React.FC<PageParams> = ({ projectId, pageId }) => {
  const page = findPageById(pageId, pages[projectId]);

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
