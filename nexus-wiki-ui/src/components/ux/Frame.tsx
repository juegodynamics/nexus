import { Box, BoxProps, useTheme } from "@mui/material";
import React from "react";

export interface FrameProps extends BoxProps {
  borderColor?: string;
  borderWidth?: number;
  borderRadius?: number;
  shadowColor?: string;
  glowColor?: string;
}

export const Frame: React.FC<FrameProps> = ({
  borderColor,
  borderWidth = 2,
  borderRadius = 10,
  shadowColor,
  glowColor,
  children,
  sx,
  ...rest
}) => {
  const theme = useTheme();

  return (
    <Box
      sx={{
        border: `${borderWidth}px solid ${borderColor || theme.palette.primary.main}`,
        borderRadius: `${borderRadius}px`,
        boxShadow: `0 0 10px ${shadowColor || theme.palette.background.paper}`,
        position: "relative",
        overflow: "hidden",
        p: 1,
        "&::before": {
          content: '""',
          position: "absolute",
          top: 0,
          left: 0,
          right: 0,
          bottom: 0,
          borderRadius: `${borderRadius}px`,
          boxShadow: `0 0 20px ${glowColor || theme.palette.primary.main}`,
          pointerEvents: "none",
          zIndex: 1,
        },
        ...sx,
      }}
      {...rest}
    >
      {children}
    </Box>
  );
};
