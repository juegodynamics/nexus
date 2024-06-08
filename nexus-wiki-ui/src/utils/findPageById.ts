// src/utils/findArticleById.ts
import { PageMeta } from "../types";

export const findPageById = (id: string, pages: PageMeta[]): any | null => {
  for (const page of pages) {
    if (page.id === id) {
      return page;
    }
    if (page.children && page.children.length > 0) {
      const found = findPageById(id, page.children);
      if (found) {
        return found;
      }
    }
  }
  return null;
};
