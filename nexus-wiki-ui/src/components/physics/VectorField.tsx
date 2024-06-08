import React, { useRef, useEffect, useState } from "react";
import { Box, Paper, Typography } from "@mui/material";
import "katex/dist/katex.min.css";
import katex from "katex";

export interface VectorFieldProps {
  width: number;
  height: number;
  fieldFunction: (x: number, y: number) => [number, number];
}

const getColorFromMagnitude = (magnitude: number): string => {
  const maxMagnitude = 20; // Adjust this value based on your field function's maximum magnitude
  const normalizedMagnitude = Math.min(1, magnitude / maxMagnitude);
  const r = Math.min(255, Math.max(0, normalizedMagnitude * 255));
  const g = Math.min(255, Math.max(0, (1 - normalizedMagnitude) * 255));
  return `rgb(${r}, ${g}, 0)`; // Gradient from red to green
};

export const VectorField: React.FC<VectorFieldProps> = ({
  width,
  height,
  fieldFunction,
}) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [mousePos, setMousePos] = useState<{ x: number; y: number } | null>(
    null
  );
  const [vectorValue, setVectorValue] = useState<[number, number] | null>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const context = canvas.getContext("2d");
    if (!context) return;

    context.clearRect(0, 0, width, height);

    for (let y = 0; y < height; y += 20) {
      for (let x = 0; x < width; x += 20) {
        const [vx, vy] = fieldFunction(x, y);
        const magnitude = Math.sqrt(vx * vx + vy * vy);
        drawVector(context, x, y, vx, vy, getColorFromMagnitude(magnitude));
      }
    }
  }, [width, height, fieldFunction]);

  const drawVector = (
    context: CanvasRenderingContext2D,
    x: number,
    y: number,
    vx: number,
    vy: number,
    color: string
  ) => {
    const arrowLength = 10; // Length of the arrowhead
    const angle = Math.atan2(vy, vx);

    // Draw the main line
    context.beginPath();
    context.moveTo(x, y);
    context.lineTo(x + vx, y + vy);
    context.strokeStyle = color;
    context.stroke();

    // Draw the arrowhead
    context.beginPath();
    context.moveTo(x + vx, y + vy);
    context.lineTo(
      x + vx - arrowLength * Math.cos(angle - Math.PI / 6),
      y + vy - arrowLength * Math.sin(angle - Math.PI / 6)
    );
    context.lineTo(
      x + vx - arrowLength * Math.cos(angle + Math.PI / 6),
      y + vy - arrowLength * Math.sin(angle + Math.PI / 6)
    );
    context.lineTo(x + vx, y + vy);
    context.fillStyle = color;
    context.fill();
  };

  const handleMouseMove = (event: React.MouseEvent<HTMLCanvasElement>) => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const rect = canvas.getBoundingClientRect();
    const x = Math.floor(event.clientX - rect.left);
    const y = Math.floor(event.clientY - rect.top);

    if (x >= 0 && x < width && y >= 0 && y < height) {
      setMousePos({ x, y });
      setVectorValue(fieldFunction(x, y));
    } else {
      setMousePos(null);
      setVectorValue(null);
    }
  };

  const renderLatex = (x: number, y: number, [vx, vy]: [number, number]) => {
    return katex.renderToString(
      `\\vec{\\Phi}(${x}, ${y}) = \\begin{pmatrix} ${vx.toFixed(2)} \\\\ ${vy.toFixed(2)} \\end{pmatrix}`
    );
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
        {mousePos && vectorValue !== null ? (
          <Typography
            variant="h6"
            dangerouslySetInnerHTML={{
              __html: renderLatex(mousePos.x, mousePos.y, vectorValue),
            }}
          />
        ) : (
          <Typography variant="h6">Move the mouse over the field</Typography>
        )}
      </Box>
    </Box>
  );
};

// Example usage
// const ExampleFieldFunction = (x: number, y: number) => {
//   const centerX = 200;
//   const centerY = 200;
//   const angle = Math.atan2(y - centerY, x - centerX);
//   const magnitude = 10;
//   return [magnitude * Math.cos(angle), magnitude * Math.sin(angle)];
// };
