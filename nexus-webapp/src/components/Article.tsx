// src/components/Article.tsx
import React from "react";
import { useParams } from "react-router-dom";
import { Typography } from "@mui/material";
import { findArticleById } from "../utils/findArticleById";
import articles from "../articles";

const Article: React.FC = () => {
  const { projectId, articleId } = useParams<{
    projectId: string;
    articleId: string;
  }>();
  const article = findArticleById(
    articleId || "",
    articles[projectId || Object.keys(articles)[0]]
  );

  if (!article) {
    return <Typography variant="h1">Article Not Found</Typography>;
  }

  return (
    <>
      <Typography variant="h1">{article.title}</Typography>
      {article.content}
    </>
  );
};

export default Article;
