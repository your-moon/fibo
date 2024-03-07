"use client";

import Loading from "@/app/components/loader";
import PostContent from "@/app/components/postcontent";
import { RPost } from "@/app/components/posts";
import { BACKEND_URL } from "@/app/provider";
import { Button } from "@nextui-org/react";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";

interface SinglePostResponse {
  status: number;
  message: string;
  data: RPost;
}

export default function Page({ params }: { params: { id: number } }) {
  return (
    <QueryClientProvider client={new QueryClient()}>
      <SinglePostPage params={params} />
    </QueryClientProvider>
  );
}

function SinglePostPage({ params }: { params: { id: number } }) {
  const { isPending, error, data } = useQuery({
    queryKey: ["post", params.id],
    queryFn: () =>
      fetch(`${BACKEND_URL}/posts/${params.id}`).then(
        (res) => res.json() as Promise<SinglePostResponse>,
      ),
  });

  if (isPending) return <Loading />;

  if (error) return <p>Error: {error.message}</p>;

  return (
    <div className="mt-20 flex flex-col items-center mx-70">
      <h1 className="mb-12 min-w-96  ">{data.data.Title}</h1>
      <div className="min-w-[300px] max-w-screen-lg">
        <PostContent content={data.data.Content} />
      </div>
      <div className="flex flex-row-reverse my-20">
        <Button disabled={true}>{data.data.Likes} Likes</Button>
      </div>
    </div>
  );
}
