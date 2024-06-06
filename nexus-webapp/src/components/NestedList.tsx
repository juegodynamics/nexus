import React, { useState } from "react";
import { List, ListItem, ListItemText, Collapse } from "@mui/material";
import { ExpandLess, ExpandMore } from "@mui/icons-material";
import { Link } from "react-router-dom";

interface Article {
  id: string;
  title: string;
  content: React.ReactNode;
  children?: Article[];
}

interface NestedListProps {
  articles: Article[];
  setSelectedArticle: (id: string) => void;
  selectedArticle: string;
  level?: number;
}

const NestedList: React.FC<NestedListProps> = ({
  articles,
  setSelectedArticle,
  selectedArticle,
  level = 0,
}) => {
  const [open, setOpen] = useState<{ [key: string]: boolean }>({});

  const handleClick = (id: string) => {
    setOpen((prevOpen) => ({ ...prevOpen, [id]: !prevOpen[id] }));
  };

  return (
    <List
      dense
      sx={{
        pl: 1,
        pt: 8,
      }}
    >
      {articles.map((article) => (
        <div key={article.id}>
          <ListItem
            button
            component={Link}
            to={`/${article.id}`}
            onClick={() => setSelectedArticle(article.id)}
            selected={selectedArticle === article.id}
            sx={{
              pl: 1 * level,
              mb: 0.5,
              borderLeft: `2px solid #00ffff`,
            }}
          >
            <ListItemText primary={article.title} sx={{ pl: 1 }} />
            {article.children && article.children.length > 0 ? (
              open[article.id] ? (
                <ExpandLess onClick={() => handleClick(article.id)} />
              ) : (
                <ExpandMore onClick={() => handleClick(article.id)} />
              )
            ) : null}
          </ListItem>
          {article.children && article.children.length > 0 && (
            <Collapse in={open[article.id]} timeout="auto" unmountOnExit>
              <NestedList
                articles={article.children}
                setSelectedArticle={setSelectedArticle}
                selectedArticle={selectedArticle}
                level={level + 1}
              />
            </Collapse>
          )}
        </div>
      ))}
    </List>
  );
};

export default NestedList;
