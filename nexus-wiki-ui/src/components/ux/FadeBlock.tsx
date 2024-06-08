import React from "react";
import { Box, BoxProps } from "@mui/material";

interface FrameProps extends BoxProps {
  borderColor?: string;
}

export const FadeBlock: React.FC<FrameProps> = ({
  borderColor = "#00ffff",
  children,
  sx,
  ...rest
}) => {
  return (
    <Box
      sx={{
        borderLeft: `1px solid ${borderColor}`,
        background: `linear-gradient(to right, rgba(0,0,0,0.5) 0%, rgba(0,0,0,0) 100%)`,
        position: "relative",
        overflow: "scroll",
        p: 1,
        mt: 1,
        mb: 1,
        "&::before": {
          content: '""',
          position: "absolute",
          top: 0,
          left: 0,
          right: 0,
          bottom: 0,
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
