import styles from "./editor.module.css";
import edjsHTML from "editorjs-html";

const edjsParser = edjsHTML();

const PostContent = (props: { content: string }) => {
  const json = JSON.parse(props.content);
  console.log(json);
  let htmlArray = edjsParser.parse(json);
  console.log(htmlArray);
  return (
    <div>
      {htmlArray.map((html, index) => (
        <div key={index} dangerouslySetInnerHTML={{ __html: html }} />
      ))}
    </div>
  );
};

export default PostContent;
