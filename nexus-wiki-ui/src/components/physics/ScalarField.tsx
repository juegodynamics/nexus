import React, { useState, useRef, useEffect } from "react";
import { Box, Paper, Typography } from "@mui/material";
import "katex/dist/katex.min.css";
import katex from "katex";

export interface ScalarFieldProps {
  width: number;
  height: number;
  fieldFunction: (x: number, y: number) => number;
}

export const ScalarField: React.FC<ScalarFieldProps> = ({
  width,
  height,
  fieldFunction,
}) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [mousePos, setMousePos] = useState<{ x: number; y: number } | null>(
    null
  );
  const [fieldValue, setFieldValue] = useState<number | null>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const context = canvas.getContext("2d");
    if (!context) return;

    const imageData = context.createImageData(width, height);
    for (let y = 0; y < height; y++) {
      for (let x = 0; x < width; x++) {
        const value = fieldFunction(x, y);
        const color = getColorFromValue(value);
        const index = (y * width + x) * 4;
        imageData.data[index] = color[0];
        imageData.data[index + 1] = color[1];
        imageData.data[index + 2] = color[2];
        imageData.data[index + 3] = 255; // Alpha channel
      }
    }
    context.putImageData(imageData, 0, 0);
  }, [width, height, fieldFunction]);

  const getColorFromValue = (value: number): [number, number, number] => {
    const r = Math.min(255, Math.max(0, value * 255));
    const g = Math.min(255, Math.max(0, (1 - value) * 255));
    const b = 128; // Fixed blue value for simplicity
    return [r, g, b];
  };

  const handleMouseMove = (event: React.MouseEvent<HTMLCanvasElement>) => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const rect = canvas.getBoundingClientRect();
    const x = Math.floor(event.clientX - rect.left);
    const y = Math.floor(event.clientY - rect.top);

    if (x >= 0 && x < width && y >= 0 && y < height) {
      setMousePos({ x, y });
      setFieldValue(fieldFunction(x, y));
    } else {
      setMousePos(null);
      setFieldValue(null);
    }
  };

  const renderLatex = (x: number, y: number, value: number) => {
    return katex.renderToString(`\\Phi(${x}, ${y}) = ${value.toFixed(2)}`);
  };

  return (
    <Box display="flex" justifyContent="center" alignItems="center" p={2}>
      <Paper elevation={3}>
        <canvas
          ref={canvasRef}
          width={width}
          height={height}
          onMouseMove={handleMouseMove}
          style={{ cursor: "crosshair" }}
        />
      </Paper>
      <Box ml={2} width={"400px"}>
        {mousePos && fieldValue !== null ? (
          <Typography
            variant="h6"
            dangerouslySetInnerHTML={{
              __html: renderLatex(mousePos.x, mousePos.y, fieldValue),
            }}
          />
        ) : (
          <Typography variant="h6">Move the mouse over the field</Typography>
        )}
      </Box>
    </Box>
  );
};
