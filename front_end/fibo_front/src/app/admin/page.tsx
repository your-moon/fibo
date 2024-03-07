"use client";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import { PostResponse } from "../components/posts";
import { useEffect } from "react";
import Cookie from "js-cookie";
import { BACKEND_URL } from "../provider";
import Loading from "../components/loader";
import { BigCard } from "../components/bigcard";
import Authors from "../components/admin/authors";
import AdminPosts from "../components/admin/posts";

export interface AllUserResponse {
  status: string;
  message: string;
  data: GetUser[];
}

export interface GetUser {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  reputation: number;
}

const queryClient = new QueryClient();
export default function Page() {
  return (
    <QueryClientProvider client={queryClient}>
      <Admin />
    </QueryClientProvider>
  );
}

function Admin() {
  const token = Cookie.get("token");
  const {
    isPending,
    error,
    data: users,
  } = useQuery({
    queryKey: ["users"],
    queryFn: () =>
      fetch(`${BACKEND_URL}/users/all`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${token}`,
        },
      }).then((res) => res.json() as Promise<AllUserResponse>),
  });

  const {
    isPending: isPostsPending,
    error: postsError,
    data: posts,
  } = useQuery({
    queryKey: ["posts"],
    queryFn: () =>
      fetch(`${BACKEND_URL}/posts`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${token}`,
        },
      }).then((res) => res.json() as Promise<PostResponse>),
  });

  useEffect(() => {
    if (users) {
      console.log(users.data);
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

  if (!users.data) return <div>No data</div>;

  return (
    <div className="mx-60 mt-12">
      <div className="flex flex-row ">
        <BigCard
          title="Total Authors"
          value={users.data.length}
          valueColor="text-violet-300"
        />
        <BigCard
          title="Total Posts"
          value={posts?.data.length}
          valueColor="text-emerald-400"
        />
        <BigCard
          title="Total Posts"
          value={posts?.data.length}
          valueColor="text-emerald-400"
        />
      </div>
      <h1 className="text-4xl font-bold my-6">Authors</h1>
      <Authors getUsers={users.data} />
      <h1 className="text-4xl font-bold my-6">All Posts</h1>
      <AdminPosts
        isPending={isPostsPending}
        error={postsError}
        data={posts?.data}
      />
    </div>
  );
}
