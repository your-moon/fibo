"use client";

import Editor from "@/app/components/editor";
import { RPost } from "@/app/components/posts";
import { Button, Checkbox } from "@nextui-org/react";
import {
  QueryClient,
  QueryClientProvider,
  useQuery,
} from "@tanstack/react-query";
import { useEffect, useState } from "react";
import Cookie from "js-cookie";
import { useRouter } from "next/navigation";

interface SinglePostResponse {
  status: number;
  message: string;
  data: RPost;
}

export default function Page({ params }: { params: { id: number } }) {
  return (
    <QueryClientProvider client={new QueryClient()}>
      <EditSinglePost params={params} />
    </QueryClientProvider>
  );
}

function EditSinglePost({ params }: { params: { id: number } }) {
  const router = useRouter();
  const [content, setContent] = useState<string>("");
  const [isPublished, setIsPublished] = useState<boolean>(false);
  const [isSaving, setIsSaving] = useState<boolean>(false);
  const { isPending, error, data } = useQuery({
    queryKey: ["post", params.id],
    queryFn: () =>
      fetch(`http://localhost:3005/posts/${params.id}`).then((res) => {
        return res.json() as Promise<SinglePostResponse>;
      }),
  });
  useEffect(() => {
    if (data && data.data) {
      setContent(data.data.Content);
      setIsPublished(data.data.IsPublished);
    }
    console.log(data);
  }, [data]);

  const bOnClick = async () => {
    const token = Cookie.get("token");
    if (!token) {
      alert("You need to be logged in to do that");
      return;
    }
    if (!data) return;

    const res = await fetch(`http://localhost:3005/posts/${params.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `${token}`,
      },
      body: JSON.stringify({
        Title: data.data.Title,
        Content:
          typeof content === "string" ? content : JSON.stringify(content),
        Is_published: isPublished,
      }),
    });
    if (res.status === 200) {
      setIsSaving(true);
      setTimeout(() => {
        setIsSaving(false);
      }, 1000);
      router.push(`/posts/${params.id}`);
    }
  };

  if (isPending) return <p>Loading...</p>;

  if (error) return <p>Error: {error.message}</p>;

  if (!data.data) return <p>No data</p>;

  return (
    <div className="mx-96 mt-20">
      <div id="editorjs"></div>
      <Editor
        data={JSON.parse(data.data.Content)}
        onChange={setContent}
        editorBlock="editorjs"
      />
      <div className="flex flex-col items-center justify-center mt-4">
        <Checkbox
          className="mb-4"
          isSelected={isPublished}
          onValueChange={setIsPublished}
        >
          Publish
        </Checkbox>
        <Button
          className="max-w-[250px] min-w-[200px]"
          color={isSaving ? "success" : "default"}
          onClick={bOnClick}
        >
          Save
        </Button>
      </div>
    </div>
  );
}
