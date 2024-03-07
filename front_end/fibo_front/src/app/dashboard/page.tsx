"use client";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import Draft from "../components/myPosts";
import MyPublishedPosts from "../components/myPublishedPosts";
import { PostResponse } from "../components/posts";
import { useEffect } from "react";
import Cookie from "js-cookie";
import { BACKEND_URL } from "../provider";
import Loading from "../components/loader";
import { BigCard } from "../components/bigcard";

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
    queryKey: ["posts"],
    queryFn: () =>
      fetch(`${BACKEND_URL}/users/me/posts`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${token}`,
        },
      }).then((res) => res.json() as Promise<PostResponse>),
  });

  useEffect(() => {
    if (data) {
      console.log(data.data);
    }
  });

  if (isPending) return <Loading />;
  if (error) return <div>Error: {error.message}</div>;
  if (!token) {
    return (
      <div className="text-center mt-20">
        <h1 className="text-3xl font-bold">You are not logged in</h1>
      </div>
    );
  }

  if (!data.data)
    return (
      <div className="mx-64 mt-12">
        <div className="flex flex-row ">
          <BigCard
            title="Reputation"
            value={100}
            valueColor="text-violet-300"
          />
          <BigCard
            title="Ongoing Salary"
            value="1223$"
            valueColor="text-emerald-400"
          />
        </div>
        <h1 className="text-4xl font-bold my-6">My Published Posts</h1>
        <p>No data</p>
        <h1 className="text-4xl font-bold my-6">My Drafts</h1>
        <p>No data</p>
      </div>
    );

  return (
    <div className="mx-72 mt-12">
      <div className="flex flex-row ">
        <BigCard title="Reputation" value={100} valueColor="text-violet-300" />
        <BigCard
          title="Ongoing Salary"
          value="1223$"
          valueColor="text-emerald-400"
        />
      </div>
      <h1 className="text-4xl font-bold my-6">My Published Posts</h1>
      <MyPublishedPosts data={data.data} />
      <h1 className="text-4xl font-bold my-6">My Drafts</h1>
      <Draft data={data.data} />
    </div>
  );
}
