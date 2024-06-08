export interface PageMeta {
  id: string;
  title: string;
  content: React.ReactNode;
  children?: PageMeta[];
}
