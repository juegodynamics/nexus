import React from "react";
import { VectorField } from "../../components/physics/VectorField";

const exampleFieldFunction = (x: number, y: number): [number, number] => {
  const centerX = 200;
  const centerY = 200;
  const angle = Math.atan2(y - centerY, x - centerX);
  const magnitude = 10;
  return [magnitude * Math.cos(angle), magnitude * Math.sin(angle)];
};

export const ExampleVectorField = () => {
  return (
    <VectorField
      width={600}
      height={400}
      fieldFunction={exampleFieldFunction}
    />
  );
};
