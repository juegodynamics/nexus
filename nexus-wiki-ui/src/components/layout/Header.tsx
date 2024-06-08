// src/components/Header.tsx
import React from "react";
import { AppBar, Toolbar, IconButton, Typography, Button } from "@mui/material";
import Icon from "@mdi/react";
import { mdiMenu } from "@mdi/js";

interface HeaderProps {
  onClickMenuButton?: React.MouseEventHandler<HTMLButtonElement>;
}

export const Header: React.FC<HeaderProps> = ({ onClickMenuButton }) => (
  <AppBar position="fixed" sx={{ zIndex: 1 }}>
    <Toolbar>
      <IconButton
        size="large"
        edge="start"
        color="inherit"
        aria-label="menu"
        onClick={(e) => onClickMenuButton?.(e)}
        sx={{ mr: 2 }}
      >
        <Icon path={mdiMenu} size={1} />
      </IconButton>
      <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
        News
      </Typography>
      <Button color="inherit">Login</Button>
    </Toolbar>
  </AppBar>
);
