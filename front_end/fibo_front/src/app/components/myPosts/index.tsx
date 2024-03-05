import { EditSinglePost } from "../editSinglePost";
import { RPost } from "../posts";

interface DraftProps {
  data: RPost[];
}

export default function Draft({ data }: DraftProps) {
  return (
    <div className="flex flex-row flex-wrap">
      {data
        .filter((post) => !post.IsPublished)
        .map((post: RPost) => (
          <EditSinglePost
            id={post.Id}
            likes={post.Likes}
            title={post.Title}
            content={post.Content}
          />
        ))}
    </div>
  );
}
