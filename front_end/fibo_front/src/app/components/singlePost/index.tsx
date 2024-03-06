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
import React from "react";
import { HeartIcon } from "./heartIcon";
import PostContent from "../postcontent";

interface SinglePostProps {
  id: number;
  likes: number;
  title: string;
  content: string;
}

export const SinglePost = ({ id, likes, title, content }: SinglePostProps) => {
  const [isLiked, setIsLiked] = React.useState(false);

  return (
    <Card
      as={Link}
      href={`/posts/${id}`}
      className="h-40 max-h-80 min-w-[260px] max-w-[340px] mx-1 mb-2"
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
            <h4 className="text-small font-semibold leading-none text-default-600">
              {title}
            </h4>
            <h5 className="text-small tracking-tight text-default-400">
              @zoeylang
            </h5>
          </div>
        </div>
        <Button
          isIconOnly
          color="danger"
          variant="flat"
          onPress={() => setIsLiked(!isLiked)}
          onClick={(e) => {
            e.stopPropagation();
            e.preventDefault();
          }}
        >
          <HeartIcon filled={isLiked} />
        </Button>
      </CardHeader>
      <CardBody className="px-4 py-0 text-small text-default-400">
        <PostContent content={content} />
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
