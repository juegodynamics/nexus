// src/components/HyperbolicSpace.tsx
import React, { useMemo } from "react";
import { Canvas } from "@react-three/fiber";
// import { OrbitControls } from "@react-three/drei";
import * as THREE from "three";

export const HyperbolicSpace: React.FC = () => {
  const hyperbolicGeometry = useMemo(() => {
    const geometry = new THREE.BufferGeometry();
    const vertices = generateHyperbolicVertices();
    const indices = generateHyperbolicIndices(vertices.length);

    geometry.setAttribute(
      "position",
      new THREE.Float32BufferAttribute(vertices, 3)
    );
    geometry.setIndex(indices);

    return geometry;
  }, []);

  return (
    <Canvas camera={{ position: [0, 0, 5], fov: 75 }}>
      <color attach="background" args={["#000"]} />
      <ambientLight intensity={0.5} />
      <pointLight position={[10, 10, 10]} />

      <mesh geometry={hyperbolicGeometry}>
        <meshStandardMaterial color="#ff6347" wireframe />
      </mesh>

      {/* <OrbitControls /> */}
    </Canvas>
  );
};

const generateHyperbolicVertices = (): Float32Array => {
  const vertices: number[] = [];
  for (let i = 0; i < 100; i++) {
    const angle = (i / 100) * Math.PI * 2;
    const radius = Math.sinh(i / 10);
    vertices.push(radius * Math.cos(angle), radius * Math.sin(angle), 0);
  }
  return new Float32Array(vertices);
};

const generateHyperbolicIndices = (vertexCount: number): number[] => {
  const indices: number[] = [];
  for (let i = 0; i < vertexCount - 1; i++) {
    indices.push(i, i + 1, (i + 2) % vertexCount);
  }
  return indices;
};
