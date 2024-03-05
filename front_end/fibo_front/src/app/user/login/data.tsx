import { Login } from "./page";

export interface LoginResponseSuccess {
  status: number;
  message: string;
  data?: {
    id: number;
    email: string;
    firstName: string;
    lastName: string;
    token: string;
  };
}

export interface LoginResponseError {
  status: number;
  message: string;
  data?: null;
}

export type LoginResponse = LoginResponseSuccess | LoginResponseError;

export async function fiboLogin(login: Login): Promise<LoginResponse> {
  const res = await fetch("http://localhost:3005/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: login.email,
      password: login.password,
    }),
  });
  if (res.ok) {
    const data = (await res.json()) as LoginResponseSuccess;
    return data;
  }
  const data = (await res.json()) as LoginResponseError;

  return data;
}
