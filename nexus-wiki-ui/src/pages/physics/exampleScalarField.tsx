import React from "react";
import { ScalarField } from "../../components/physics/ScalarField";

// Example usage
const exampleFieldFunction = (x: number, y: number) => {
  const centerX = 200;
  const centerY = 200;
  const distance = Math.sqrt((x - centerX) ** 2 + (y - centerY) ** 2);
  return Math.sin(distance / 20);
};

export const ExampleScalarField = () => {
  return (
    <ScalarField
      width={600}
      height={400}
      fieldFunction={exampleFieldFunction}
    />
  );
};
