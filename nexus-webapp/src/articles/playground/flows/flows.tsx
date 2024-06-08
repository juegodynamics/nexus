import React, { useCallback, useEffect, useState } from "react";
import ReactFlow, {
  Node,
  Edge,
  useReactFlow,
  ReactFlowProvider,
} from "reactflow";
import ELK, { ElkNode } from "elkjs/lib/elk.bundled.js";
import Frame from "../../../components/Frame";
import MonacoEditor from "@monaco-editor/react";
import { Box, Button } from "@mui/material";
import { FrameNode } from "./nodes";
import "reactflow/dist/style.css";
import { debounce } from "../../../utils/debounce";

const elk = new ELK();

const Flowchart: React.FC<{ nodes: Node[]; edges: Edge[] }> = ({
  nodes,
  edges,
}) => {
  const { fitView } = useReactFlow();

  useEffect(() => {
    fitView();
  }, [nodes, edges, fitView]);

  const nodeTypes = React.useMemo(() => ({ frame: FrameNode }), []);

  return <ReactFlow nodeTypes={nodeTypes} nodes={nodes} edges={edges} />;
};

export const FlowPlayground: React.FC = () => {
  const [code, setCode] = useState<string>("");
  const [nodes, setNodes] = useState<Node[]>([]);
  const [edges, setEdges] = useState<Edge[]>([]);
  const [error, setError] = useState<string | null>(null);

  const parseInput = (input: string): { nodes: Node[]; edges: Edge[] } => {
    const lines = input.split("\n");
    const nodes: Node[] = [];
    const edges: Edge[] = [];

    lines.forEach((line) => {
      const [type, from, ...rest] = line.split(" ");
      const to = rest.join(" ");

      if (type === "node") {
        nodes.push({
          id: from,
          type: "frame",
          data: { label: to || from },
          position: { x: 0, y: 0 },
        });
      } else if (type === "edge") {
        edges.push({ id: `${from}-${to}`, source: from, target: to });
      }
    });

    return { nodes, edges };
  };

  const generateFlowchart = useCallback(
    debounce(async (input: string) => {
      try {
        const { nodes: parsedNodes, edges: parsedEdges } = parseInput(input);

        const elkNodes: ElkNode[] = parsedNodes.map((node) => ({
          id: node.id,
          width: 100,
          height: 50,
        }));

        const elkEdges = parsedEdges.map((edge) => ({
          id: edge.id,
          sources: [edge.source],
          targets: [edge.target],
        }));

        const layout = await elk.layout({
          id: "root",
          children: elkNodes,
          edges: elkEdges,
        });

        const positionedNodes = parsedNodes.map((node) => {
          const elkNode = layout.children?.find((n) => n.id === node.id);
          return elkNode
            ? { ...node, position: { x: elkNode.x || 0, y: elkNode.y || 0 } }
            : node;
        });

        setNodes(positionedNodes);
        setEdges(parsedEdges);
        setError(null); // Clear any previous errors
      } catch (e) {
        setError("Error generating layout. Please check your input format.");
        console.error("Error generating layout:", e);
      }
    }, 500),
    []
  );
  useEffect(() => {
    generateFlowchart(code);
  }, [code, generateFlowchart]);

  return (
    <Box style={{ display: "flex", height: "100vh" }}>
      <Box style={{ paddingRight: "20px" }}>
        <MonacoEditor
          height="80vh"
          width="30vw"
          defaultLanguage="plaintext"
          value={code}
          onChange={(value) => setCode(value || "")}
          options={{
            minimap: { enabled: false },
            automaticLayout: true,
          }}
        />
        {error && (
          <div style={{ color: "red", marginTop: "10px" }}>{error}</div>
        )}
      </Box>
      <Frame style={{ width: "40vw", height: "80vh" }}>
        <ReactFlowProvider>
          <Flowchart nodes={nodes} edges={edges} />
        </ReactFlowProvider>
      </Frame>
    </Box>
  );
};
