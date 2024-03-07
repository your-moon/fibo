"use client";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import Posts from "./components/posts";

const queryClient = new QueryClient();
export default function Home() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="w-full flex flex-col mt-20 items-center ">
        <h1 className="flex text-6xl font-bold my-8 w-8/12">Publishing</h1>
        <Posts />
      </div>
    </QueryClientProvider>
  );
}
