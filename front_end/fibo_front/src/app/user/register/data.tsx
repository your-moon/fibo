import { Register } from "./page";

export interface RegisterResponseSuccess {
  status: number;
  message: string;
  data?: number;
}
export interface RegisterResponseError {
  status: number;
  message: string;
  data?: null;
}

export type RegisterResponse = RegisterResponseSuccess | RegisterResponseError;

export async function fiboRegister(login: Register): Promise<RegisterResponse> {
  const res = await fetch("http://localhost:3005/users", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      firstName: login.firstName,
      lastName: login.lastName,
      email: login.email,
      password: login.password,
    }),
  });
  if (res.ok) {
    const data = (await res.json()) as RegisterResponseSuccess;
    return data;
  }
  const data = (await res.json()) as RegisterResponseError;

  return data;
}
