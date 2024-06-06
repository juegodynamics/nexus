import React, { useState } from "react";
import {
  ThemeProvider,
  CssBaseline,
  Drawer,
  Box,
  Typography,
  AppBar,
  Toolbar,
  IconButton,
  Button,
} from "@mui/material";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useParams,
} from "react-router-dom";
import arwesTheme from "./theme/theme";
import articles from "./articles";
import NestedList from "./components/NestedList";
import "./App.css";

const findArticleById = (id: string, articles: any[]): any | null => {
  for (const article of articles) {
    if (article.id === id) {
      return article;
    }
    if (article.children && article.children.length > 0) {
      const found = findArticleById(id, article.children);
      if (found) {
        return found;
      }
    }
  }
  return null;
};

const Article: React.FC = () => {
  const { articleId } = useParams<{ articleId: string }>();
  const article = findArticleById(articleId || "", articles);

  if (!article) {
    return <Typography variant="h1">Article Not Found</Typography>;
  }

  return (
    <>
      <Typography variant="h1">{article.title}</Typography>
      <Typography variant="body1">{article.content}</Typography>
    </>
  );
};

const App: React.FC = () => {
  const [selectedArticle, setSelectedArticle] = useState<string>(
    articles[0].id
  );

  return (
    <ThemeProvider theme={arwesTheme}>
      <CssBaseline />
      <Router>
        <Box sx={{ display: "flex", flexDirection: "column" }}>
          <AppBar position="static">
            <Toolbar>
              <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="menu"
                sx={{ mr: 2 }}
              >
                {/* <MenuIcon /> */}
              </IconButton>
              <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                News
              </Typography>
              <Button color="inherit">Login</Button>
            </Toolbar>
          </AppBar>

          <Box sx={{ display: "flex" }}>
            <Drawer
              variant="permanent"
              sx={{
                width: 240,
                flexShrink: 0,
                [`& .MuiDrawer-paper`]: {
                  width: 240,
                  boxSizing: "border-box",
                  boxShadow: "inset -10px 0 10px -10px rgba(37, 50, 56, 0.7)",
                  background: "rgba(0,0,0,0)",
                  backdropFilter: "blur(2px)",
                  borderRight: "0px",
                },
              }}
            >
              <AppBar>
                <Toolbar />
              </AppBar>
              <NestedList
                articles={articles}
                setSelectedArticle={setSelectedArticle}
                selectedArticle={selectedArticle}
              />
            </Drawer>
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
              <Routes>
                <Route
                  path="/"
                  element={<Typography variant="h1">Welcome</Typography>}
                />
                <Route path="/:articleId" element={<Article />} />
              </Routes>
            </Box>
          </Box>
        </Box>
      </Router>
    </ThemeProvider>
  );
};

export default App;
