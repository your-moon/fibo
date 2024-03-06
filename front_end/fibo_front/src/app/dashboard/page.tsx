"use client";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import Draft from "../components/myPosts";
import MyPublishedPosts from "../components/myPublishedPosts";
import { Reputation } from "../components/reputation";
import { PostResponse } from "../components/posts";
import { useEffect } from "react";
import Cookie from "js-cookie";

const queryClient = new QueryClient();
export default function Page() {
  return (
    <QueryClientProvider client={queryClient}>
      <Dashboard />
    </QueryClientProvider>
  );
}

function Dashboard() {
  const token = Cookie.get("token");
  const { isPending, error, data } = useQuery({
    queryKey: ["user"],
    queryFn: () =>
      fetch("http://localhost:3005/users/me/posts", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${token}`,
        },
      }).then((res) => res.json() as Promise<PostResponse>),
  });

  useEffect(() => {
    if (data) {
      console.log(data);
    }
  });

  if (isPending) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;
  if (!token) {
    return (
      <div className="text-center mt-20">
        <h1 className="text-3xl font-bold">You are not logged in</h1>
      </div>
    );
  }

  if (!data.data) return <div>No data</div>;

  return (
    <div className="mx-80 mt-12">
      <Reputation />
      <h1 className="text-3xl font-bold my-3">My Published Posts</h1>
      <MyPublishedPosts data={data.data} />
      <h1 className="text-3xl font-bold my-3">My Drafts</h1>
      <Draft data={data.data} />
    </div>
  );
}
