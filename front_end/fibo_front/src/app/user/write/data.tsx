export async function newPost(
  data: string,
  token: string,
  isPublish: boolean,
): Promise<any> {
  const res = await fetch("http://localhost:3005/posts", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",

      Authorization: token,
    },
    body: JSON.stringify({
      title: "test",
      is_published: isPublish,
      content: data,
    }),
  });
  return await res.json();
}
