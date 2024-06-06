import FadeBlock from "./components/FadeBlock";
import PatientCard from "./components/fhircards/PatientCard";
import { patient1 } from "./data";
import Home from "./articles/playground/1.Home.mdx";
import Fields from "./articles/physics/1.Fields.mdx";

export interface ArticleMeta {
  id: string;
  title: string;
  content: React.ReactNode;
  children?: ArticleMeta[];
}

const articles: Record<string, ArticleMeta[]> = {
  playground: [
    {
      id: "article-1",
      title: "Home",
      content: (
        <>
          <Home />
          <FadeBlock>{"I am text."}</FadeBlock>
        </>
      ),
      children: [
        {
          id: "article-1-1",
          title: "Sub Introduction 1",
          content: "This is the content for sub introduction 1.",
          children: [
            {
              id: "article-1-1-1",
              title: "Sub Sub Introduction 1",
              content: "This is the content for sub sub introduction 1.",
            },
          ],
        },
      ],
    },
    {
      id: "article-2",
      title: "Patient",
      content: <PatientCard patient={patient1} />,
      children: [],
    },
    {
      id: "article-3",
      title: "Advanced Topics",
      content: "This is the advanced topics article content.",
    },
  ],
  physics: [
    {
      id: "article-fields",
      title: "Fields",
      content: <Fields />,
    },
  ],
};

export default articles;
