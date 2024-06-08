import { InputHTMLAttributes, useCallback } from "react";
import { Handle, Position, NodeProps } from "reactflow";
import Frame from "../../../components/Frame";
import { Box } from "@mui/system";

const handleStyle = { left: 10 };

export const FrameNode = ({ data }: NodeProps<{ label: string }>) => {
  return (
    <>
      <Handle type="target" position={Position.Left} />
      {/* <Box style={{ width: "100px", height: "50px" }}> */}
      <Frame width={"100px"}>{data.label}</Frame>
      {/* </Box> */}
      <Handle type="source" position={Position.Right} />
    </>
  );
};
