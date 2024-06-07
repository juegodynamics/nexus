import FadeBlock from "./components/FadeBlock";
import PatientCard from "./components/fhircards/PatientCard";
import { patient1 } from "./data";
import Home from "./articles/playground/1.Home.mdx";
import HyperbolicSpace from "./articles/physics/spacetime/1.HyperbolicSpace.mdx";
import { FlowPlayground } from "./articles/playground/flows/flows";

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
          title: "Flow Playground",
          content: <FlowPlayground />,
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
      id: "article-hyperbolic",
      title: "Hyperbolic Spaces",
      content: <HyperbolicSpace />,
    },
    // {`
    //   id: "article-fields",
    //   title: "Fields",
    //   content: <Fields />,
    // },`
  ],
};

export default articles;
