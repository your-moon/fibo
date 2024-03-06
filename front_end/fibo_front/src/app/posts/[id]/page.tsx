"use client";

import Loading from "@/app/components/loader";
import PostContent from "@/app/components/postcontent";
import { RPost } from "@/app/components/posts";
import { BACKEND_URL } from "@/app/provider";
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
    <div className="mx-96 mt-20">
      <PostContent content={data.data.Content} />
    </div>
  );
}