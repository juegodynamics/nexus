import { PageMeta } from "types/PageMeta";
import _1_Introduction from "./health/1.introduction";
import _1_Fields from "./physics/1.Fields.mdx";

export const pages: Record<string, PageMeta[]> = {
  health: [{ id: "health-1", title: "Introduction", content: _1_Introduction }],
  physics: [{ id: "physics-1", title: "Fields", content: <_1_Fields /> }],
};
