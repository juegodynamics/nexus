import FadeBlock from "./components/FadeBlock";
import PatientCard from "./components/fhircards/PatientCard";
import { patient1 } from "./data";

const articles = [
  {
    id: "article-1",
    title: "Introduction",
    content: (
      <>
        {"This is the introduction article content."}
        <FadeBlock>{"This is more article content."}</FadeBlock>
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
];

export default articles;
