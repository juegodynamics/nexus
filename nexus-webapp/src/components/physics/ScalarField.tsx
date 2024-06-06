// src/components/ScalarField.tsx
import React, { useRef, useEffect } from "react";
import { Box } from "@mui/material";
import { scaleSequential } from "d3-scale";
import { interpolateViridis } from "d3-scale-chromatic";

interface ScalarFieldProps {
  field: number[][];
  width: number;
  height: number;
}

export const ScalarField: React.FC<ScalarFieldProps> = ({
  field,
  width,
  height,
}) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    const numRows = field.length;
    const numCols = field[0].length;
    const colorScale = scaleSequential(interpolateViridis).domain([
      Math.min(...field.flat()),
      Math.max(...field.flat()),
    ]);

    // Set canvas size
    canvas.width = width;
    canvas.height = height;

    // Draw scalar field
    const cellWidth = width / numCols;
    const cellHeight = height / numRows;

    for (let i = 0; i < numRows; i++) {
      for (let j = 0; j < numCols; j++) {
        ctx.fillStyle = colorScale(field[i][j]);
        ctx.fillRect(j * cellWidth, i * cellHeight, cellWidth, cellHeight);
      }
    }

    // Draw axes
    ctx.strokeStyle = "#000";
    ctx.lineWidth = 1;

    // X-axis
    ctx.beginPath();
    ctx.moveTo(0, height);
    ctx.lineTo(width, height);
    ctx.stroke();

    // Y-axis
    ctx.beginPath();
    ctx.moveTo(0, 0);
    ctx.lineTo(0, height);
    ctx.stroke();

    // Draw labels
    ctx.font = "12px Arial";
    ctx.fillStyle = "#000";

    // X-axis labels
    for (let j = 0; j <= numCols; j++) {
      const x = j * cellWidth;
      ctx.fillText(j.toString(), x, height + 12);
    }

    // Y-axis labels
    for (let i = 0; i <= numRows; i++) {
      const y = i * cellHeight;
      ctx.fillText(i.toString(), -12, y);
    }
  }, [field, width, height]);

  return (
    <Box sx={{ position: "relative" }}>
      <canvas ref={canvasRef} style={{ border: "1px solid black" }} />
    </Box>
  );
};
