"use client";

import React from "react";
import { newPost } from "./data";
import { Button, Checkbox, cn } from "@nextui-org/react";
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
      <div className="flex flex-col items-center justify-center my-4">
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
          {" "}
          Publish{" "}
        </Checkbox>
        <Button
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
