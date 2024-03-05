"use client";

import React from "react";
import { newPost } from "./data";
import { Button, Checkbox } from "@nextui-org/react";
import Editor from "@/app/components/editor";
import Cookie from "js-cookie";
import { useRouter } from "next/navigation";
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

export default function Write() {
  const router = useRouter();
  const [data, setData] = React.useState<any>();
  const [isPublished, setIsPublished] = React.useState(false);
  const [isSaving, setIsSaving] = React.useState(false);
  const token = Cookie.get("token");

  const bOnClick = async () => {
    let res = await newPost(JSON.stringify(data), token, isPublished);
    if (res.status === 200) {
      setIsSaving(true);
      router.push("/");
    }
    console.log(res);
  };
  return (
    <div className="h-screen">
      <div id="editorjs"></div>
      <Editor data={data} onChange={setData} editorBlock="editorjs-container" />
      <div id="editorjs-container"></div>
      <div className="flex flex-col items-center justify-center">
        <Checkbox
          className="mb-2"
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
