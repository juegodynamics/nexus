// src/App.tsx
import React, { useState } from "react";
import { ThemeProvider, CssBaseline, Box, Typography } from "@mui/material";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { ScifiTheme } from "./theme";
import { pages } from "./pages";
import "./App.css";
import { Header, Sidebar } from "./components/layout";
import { Page } from "./components/ux";
import "./App.css";

const App: React.FC = () => {
  const [selectedArticle, setSelectedArticle] = useState<string>(
    pages[Object.keys(pages)[0]]?.[0].id
  );
  const [selectedProject, setSelectedProject] = useState<string>(
    Object.keys(pages)[0]
  );
  const [isSidebarOpen, setIsSidebarOpen] = useState<boolean>(true);

  return (
    <ThemeProvider theme={ScifiTheme}>
      <CssBaseline />
      <Router>
        <Box sx={{ display: "flex", flexDirection: "column" }}>
          <Header
            onClickMenuButton={() => setIsSidebarOpen((current) => !current)}
          />
          <Box sx={{ display: "flex" }}>
            <Sidebar
              isOpen={isSidebarOpen}
              selectedArticle={selectedArticle}
              setSelectedArticle={setSelectedArticle}
              selectedProject={selectedProject}
              setSelectedProject={setSelectedProject}
            />
            <Box component="main" sx={{ flexGrow: 1, p: 3, pt: 8 }}>
              <Routes>
                <Route
                  path="/nexus"
                  element={<Typography variant="h1">Welcome</Typography>}
                />
                <Route
                  path="/nexus/:projectId/:pageId"
                  element={
                    <div
                      style={{
                        paddingLeft: "10vw",
                        paddingRight: "10vw",
                        width: "80vw",
                        overflowY: "scroll",
                      }}
                    >
                      <Page />
                    </div>
                  }
                />
              </Routes>
            </Box>
          </Box>
        </Box>
      </Router>
    </ThemeProvider>
  );
};

export default App;
