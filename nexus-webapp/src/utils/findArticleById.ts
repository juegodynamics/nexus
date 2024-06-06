// src/utils/findArticleById.ts
export const findArticleById = (id: string, articles: any[]): any | null => {
  for (const article of articles) {
    if (article.id === id) {
      return article;
    }
    if (article.children && article.children.length > 0) {
      const found = findArticleById(id, article.children);
      if (found) {
        return found;
      }
    }
  }
  return null;
};
