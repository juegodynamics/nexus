import React from "react";
import { Box, BoxProps } from "@mui/material";

interface FrameProps extends BoxProps {
  borderColor?: string;
  borderWidth?: number;
  borderRadius?: number;
  shadowColor?: string;
  glowColor?: string;
}

const Frame: React.FC<FrameProps> = ({
  borderColor = "#00ffff",
  borderWidth = 2,
  borderRadius = 10,
  shadowColor = "#00ffcc",
  glowColor = "#00ffff",
  children,
  sx,
  ...rest
}) => {
  return (
    <Box
      sx={{
        border: `${borderWidth}px solid ${borderColor}`,
        borderRadius: `${borderRadius}px`,
        boxShadow: `0 0 10px ${shadowColor}`,
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
          boxShadow: `0 0 20px ${glowColor}`,
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

export default Frame;
