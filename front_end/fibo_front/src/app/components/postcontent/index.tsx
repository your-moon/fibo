import styles from "./editor.module.css";
import edjsHTML from "editorjs-html";

const edjsParser = edjsHTML();

const PostContent = (props: { content: string }) => {
  const json = JSON.parse(props.content);
  const html = edjsParser.parse(json);
  console.log(html);
  return (
    <div
      className={styles.postContent}
      dangerouslySetInnerHTML={{ __html: html }}
    />
  );
};

export default PostContent;
