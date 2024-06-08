// src/components/Sidebar.tsx
import React from "react";
import {
  Drawer,
  AppBar,
  Toolbar,
  Box,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
} from "@mui/material";
import { useNavigate } from "react-router-dom";

import NestedList from "./NestedList";
import { pages } from "../../pages";

export interface SidebarProps {
  isOpen?: boolean;
  selectedArticle: string;
  setSelectedArticle: (id: string) => void;
  selectedProject: string;
  setSelectedProject: (project: string) => void;
}

export const Sidebar: React.FC<SidebarProps> = ({
  isOpen,
  selectedArticle,
  setSelectedArticle,
  selectedProject,
  setSelectedProject,
}) => {
  const navigate = useNavigate();

  return (
    <Drawer
      variant="permanent"
      // open={isOpen}
      sx={{
        width: 240,
        flexShrink: 0,
        zIndex: 0,
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
      {/* <AppBar>
        <Toolbar />
      </AppBar> */}
      <Box sx={{ pt: 10 }}>
        <FormControl sx={{ m: 1, width: "96%" }} size="small">
          <InputLabel id="project-label">Project</InputLabel>
          <Select
            labelId="project-label"
            value={selectedProject}
            onChange={(event) => {
              setSelectedProject(event.target.value);

              navigate(`/nexus/${event.target.value}`);
            }}
          >
            {Object.keys(pages).map((project) => (
              <MenuItem key={project} value={project}>
                {project}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <NestedList
          articles={pages[selectedProject]}
          setSelectedArticle={setSelectedArticle}
          selectedArticle={selectedArticle}
          selectedProject={selectedProject}
        />
      </Box>
    </Drawer>
  );
};
