"use client";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import Posts from "./components/posts";

const queryClient = new QueryClient();
export default function Home() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="mx-80 mt-20">
        <h1 className="text-4xl font-bold my-5">Publishing</h1>
        <Posts />
      </div>
    </QueryClientProvider>
  );
}
