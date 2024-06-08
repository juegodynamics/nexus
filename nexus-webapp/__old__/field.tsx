import React from "react";
import { ScalarField } from "../src/components/physics";

const generateTemperatureField = (rows: number, cols: number): number[][] => {
  // Generate a temperature gradient based on latitude
  return Array.from({ length: rows }, (_, rowIndex) =>
    Array.from({ length: cols }, () => {
      const latitude = (rowIndex / rows) * 180 - 90; // From -90 to 90 degrees
      const temperature = 30 - Math.abs(latitude); // Simplified temperature model
      return temperature;
    })
  );
};

const sampleTemperatureField = generateTemperatureField(40, 40);

export const SampleScalarField = () => (
  <ScalarField field={sampleTemperatureField} width={800} height={400} />
);
