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
  userEmail: string;
  userName: string;
}

export const EditSinglePost = ({
  id,
  likes,
  title,
  content,
  categoryId,
  userName,
  userEmail,
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
      className=" flex-auto max-h-80 min-h-44 min-w-[220px] max-w-[280px] mx-1 mb-2"
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
            <h4 className="w-28 text-small font-semibold leading-none text-default-600 text-ellipsis overflow-hidden ...">
              {userName}
            </h4>
            <h5 className="text-small tracking-tight text-default-400">
              @{userEmail}
            </h5>
          </div>
        </div>
      </CardHeader>
      <CardBody className="px-4 py-0 text-3xl max-h-20 text-semibold text-default-600 text-ellipsis overflow-hidden ...">
        <p>{title}</p>
      </CardBody>
      <CardFooter className="gap-3">
        <div className="flex gap-10 w-full">
          <div className="flex flex-row w-1/2 left-0">
            <p className="font-semibold  text-small text-rose-300 mr-1">
              {likes}
            </p>
            <p className=" text-default-400 text-small">Likes</p>
          </div>
          <p className=" text-gray-950 text-small right-0 float-right ">
            @{catName}
          </p>
        </div>
      </CardFooter>
    </Card>
  );
};
