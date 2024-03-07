import { EditSinglePost } from "../editSinglePost";
import { PostResponse, RPost } from "../posts";

interface MyPublishedPostsProps {
  data: RPost[];
}

export default function MyPublishedPosts({ data }: MyPublishedPostsProps) {
  return (
    <div className="flex flex-row flex-wrap gap-2 ">
      {data
        .filter((posts) => posts.IsPublished)
        .map((post: RPost) => (
          <EditSinglePost
            key={post.Id}
            id={post.Id}
            likes={post.Likes}
            title={post.Title}
            content={post.Content}
            categoryId={post.CategoryId}
            userEmail={post.UserEmail}
            userName={post.UserName}
          />
        ))}
    </div>
  );
}
