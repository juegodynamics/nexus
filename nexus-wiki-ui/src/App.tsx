// src/App.tsx
import React, { useState } from "react";
import { ThemeProvider, CssBaseline, Box, Typography } from "@mui/material";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useSearchParams,
  useLocation,
  useNavigate,
} from "react-router-dom";
import { ScifiTheme } from "./theme";
import { pages } from "./pages";
import "./App.css";
import { DRAWER_WIDTH, Header, Sidebar } from "./components/layout";
import { Page } from "./components/ux";
import "./App.css";

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

const Layout: React.FC = () => {
  const query = useQuery();
  const queryProjectId = query.get("project");
  const queryPageId = query.get("page");

  const pageId = queryPageId || pages[Object.keys(pages)[0]]?.[0].id;
  const projectId = queryProjectId || Object.keys(pages)[0];

  const navigate = useNavigate();

  // const [selectedArticle, setSelectedArticle] = useState<string>(
  //   pageId || pages[Object.keys(pages)[0]]?.[0].id
  // );
  // const [selectedProject, setSelectedProject] = useState<string>(
  //   projectId || Object.keys(pages)[0]
  // );
  const [isSidebarOpen, setIsSidebarOpen] = useState<boolean>(true);

  return (
    <Box sx={{ display: "flex", flexDirection: "column" }}>
      <Header
        onClickMenuButton={() => setIsSidebarOpen((current) => !current)}
      />
      <Box sx={{ display: "flex" }}>
        <Sidebar
          isOpen={isSidebarOpen}
          selectedArticle={pageId}
          setSelectedArticle={(id: string) => {
            navigate(`/nexus?project=${projectId}&page=${id}`);
          }}
          selectedProject={projectId}
          setSelectedProject={(id: string) => {
            navigate(`/nexus?project=${id}&page=${pages[id][0].id}`);
          }}
        />
        <Box
          component="main"
          sx={{
            flexGrow: 1,
            pr: 3,
            pb: 3,
            pt: 8,
            paddingLeft: isSidebarOpen ? `${DRAWER_WIDTH + 10}px` : 1,
            transition: "padding-left 200ms ease-in-out",
          }}
        >
          <div
            style={{
              paddingLeft: "10vw",
              paddingRight: "10vw",
              width: "80vw",
              overflowY: "scroll",
            }}
          >
            <Page pageId={pageId} projectId={projectId} />
          </div>
        </Box>
      </Box>
    </Box>
  );
};

const App = () => (
  <ThemeProvider theme={ScifiTheme}>
    <CssBaseline />
    <Router>
      <Routes>
        <Route path="/nexus" element={<Layout />} />
      </Routes>
    </Router>
  </ThemeProvider>
);

export default App;
