"use client";

import dynamic from "next/dynamic";
import React, { useState } from "react";
import { Button, Checkbox, Input, cn } from "@nextui-org/react";
import Cookie from "js-cookie";
import { useRouter } from "next/navigation";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { newPost } from "./data";
import CategoryComp from "@/app/components/category";

const DynamicEditor = dynamic(() => import("@/app/components/editor"), {
  ssr: false,
});

export default function Page() {
  return (
    <QueryClientProvider client={new QueryClient()}>
      <Write />
    </QueryClientProvider>
  );
}

// Initial Data
const INITIAL_DATA = {
  time: new Date().getTime(),
  blocks: [
    {
      type: "header",
      data: {
        text: "This is my awesome editor!",
        level: 1,
      },
    },
  ],
};

function Write() {
  const router = useRouter();
  const [data, setData] = React.useState<any>();
  const [title, setTitle] = React.useState<string>("");
  const [isPublished, setIsPublished] = React.useState(false);
  const [isSaving, setIsSaving] = React.useState(false);
  const [categoryId, setCategoryId] = useState<string>("");

  const bOnClick = async () => {
    const token = Cookie.get("token");

    if (!token) {
      router.push("/user/login");
      return;
    }

    let res = await newPost(
      JSON.stringify(data),
      token,
      isPublished,
      title,
      Number.parseInt(categoryId),
    );
    if (res.status === 200) {
      setIsSaving(true);
      router.push("/");
    }
    console.log(res);
  };

  return (
    <div className="flex flex-col items-center justify-center">
      <Input
        className="max-w-[800px] mx-56 my-20"
        size="lg"
        placeholder="Title"
        value={title}
        onValueChange={setTitle}
        variant="underlined"
      />
      <div className="w-full" id="editorjs"></div>
      <DynamicEditor
        data={data}
        onChange={(data) => {
          setData(data);
        }}
        editorBlock="editorjs"
      />
      <div id="editorjs"></div>
      <div className="flex flex-col w-full items-center justify-center">
        <div className="min-w-[250px] flex flex-row items-center justify-center my-4">
          <Checkbox
            classNames={{
              base: cn(
                "inline-flex max-w-[250px] min-w-[200px] my-2 bg-content1",
                "hover:bg-content2 items-center justify-start",
                "cursor-pointer rounded-lg gap-2 p-4 border-2 border-transparent bg-stone-100",
                "data-[selected=true]:border-primary",
              ),
              label: "w-full",
            }}
            isSelected={isPublished}
            onValueChange={setIsPublished}
          >
            Publish
          </Checkbox>
          <CategoryComp categoryId={categoryId} setCategoryId={setCategoryId} />
        </div>
        <Button
          size="lg"
          className="max-w-[250px] min-w-[200px] mt-2"
          color={isSaving ? "success" : "default"}
          onClick={bOnClick}
        >
          Save
        </Button>
      </div>
    </div>
  );
}
