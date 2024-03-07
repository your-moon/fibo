import { BACKEND_URL } from "@/app/provider";

export async function newPost(
  data: string,
  token: string,
  isPublish: boolean,
  title: string,
  categoryId: number,
): Promise<any> {
  const res = await fetch(`${BACKEND_URL}/posts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",

      Authorization: token,
    },
    body: JSON.stringify({
      title: title,
      is_published: isPublish,
      content: data,
      category_id: categoryId,
    }),
  });
  return await res.json();
}
