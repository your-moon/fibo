"use client";
import {
  Avatar,
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Link,
} from "@nextui-org/react";
import React, { useEffect, useState } from "react";
import PostContent from "../postcontent";
import { SingleCat } from "../category";
import { BACKEND_URL } from "@/app/provider";
import { useQuery } from "@tanstack/react-query";

export interface SingleCatResponse {
  status: string;
  message: string;
  data: SingleCat;
}

interface EditSinglePostProps {
  id: number;
  likes: number;
  title: string;
  content: string;
  categoryId: number;
}

export const EditSinglePost = ({
  id,
  likes,
  title,
  content,
  categoryId,
}: EditSinglePostProps) => {
  const [catName, setCatName] = useState<string>(""); // [1]
  const { isPending, error, data } = useQuery({
    queryKey: ["repoData"],
    queryFn: () =>
      fetch(`${BACKEND_URL}/categories/${categoryId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => res.json() as Promise<SingleCatResponse>),
  });

  useEffect(() => {
    if (data && data.data) {
      console.log(data.data);
      setCatName(data.data.Name); // [2]
    }
  }, [data]);

  return (
    <Card
      as={Link}
      href={`/posts/${id}/edit`}
      className="h-48 max-h-52 min-w-[260px] max-w-[340px] mx-1 mb-2"
    >
      <CardHeader className="justify-between">
        <div className="flex gap-5">
          <Avatar
            isBordered
            radius="full"
            size="md"
            src="/avatars/avatar-1.png"
          />
          <div className="flex flex-col gap-1 items-start justify-center">
            <h4 className="max-w-[130px] text-small font-semibold leading-none text-default-600 text-ellipsis overflow-hidden ...">
              {title}
            </h4>
            <h5 className="text-small tracking-tight text-default-400">
              @{catName}
            </h5>
          </div>
        </div>
      </CardHeader>
      <CardBody className="px-3 py-0 text-small text-default-400">
        <p>
          <PostContent content={content} />
        </p>
      </CardBody>
      <CardFooter className="gap-3">
        <div className="flex gap-1">
          <p className="font-semibold text-default-400 text-small">{likes}</p>
          <p className=" text-default-400 text-small">Likes</p>
        </div>
      </CardFooter>
    </Card>
  );
};
