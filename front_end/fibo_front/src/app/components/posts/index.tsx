"use client";
export interface PostResponse {
  status: number;
  message: string;
  data: RPost[];
}

export interface RPost {
  Id: number;
  UserId: number;
  Title: string;
  Content: string;
  Likes: number;
  IsPublished: boolean;
  created_at: string;
  Updated_at: string;
  Deleted_at: string;
}
import React from "react";
import { useQuery } from "@tanstack/react-query";
import { SinglePost } from "../singlePost";

export default function Posts() {
  const { isPending, error, data } = useQuery({
    queryKey: ["repoData"],
    queryFn: () =>
      fetch("http://localhost:3005/posts/published", {}).then(
        (res) => res.json() as Promise<PostResponse>,
      ),
  });
  if (isPending) return "Loading...";

  if (error) return "An error has occurred: " + error.message;

  if (!data.data) return "No data";

  return (
    <div className="flex flex-row flex-wrap">
      {data?.data.map((article: RPost) => (
        <div key={article.Id}>
          <SinglePost
            id={article.Id}
            likes={article.Likes}
            title={article.Title}
            content={article.Content}
          />
        </div>
      ))}
    </div>
  );
}
