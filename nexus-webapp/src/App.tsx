// src/App.tsx
import React, { useState } from "react";
import { ThemeProvider, CssBaseline, Box, Typography } from "@mui/material";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import arwesTheme from "./theme/theme";
import articles from "./articles";
import "./App.css";
import { MDXProvider } from "@mdx-js/react";
import { mdxComponents } from "./mdx-ui";
import { Header, Sidebar } from "./components/layout";
import Article from "./components/Article";

const App: React.FC = () => {
  const [selectedArticle, setSelectedArticle] = useState<string>(
    articles[Object.keys(articles)[0]]?.[0].id
  );
  const [selectedProject, setSelectedProject] = useState<string>(
    Object.keys(articles)[0]
  );

  return (
    <MDXProvider components={mdxComponents}>
      <ThemeProvider theme={arwesTheme}>
        <CssBaseline />
        <Router>
          <Box sx={{ display: "flex", flexDirection: "column" }}>
            <Header />
            <Box sx={{ display: "flex" }}>
              <Sidebar
                selectedArticle={selectedArticle}
                setSelectedArticle={setSelectedArticle}
                selectedProject={selectedProject}
                setSelectedProject={setSelectedProject}
              />
              <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                <Routes>
                  <Route
                    path="/"
                    element={<Typography variant="h1">Welcome</Typography>}
                  />
                  <Route
                    path="/:projectId/:articleId"
                    element={
                      <div
                        style={{
                          paddingLeft: "2vw",
                          paddingRight: "5vw",
                          width: "80vw",
                        }}
                      >
                        <Article />
                      </div>
                    }
                  />
                </Routes>
              </Box>
            </Box>
          </Box>
        </Router>
      </ThemeProvider>
    </MDXProvider>
  );
};

export default App;
