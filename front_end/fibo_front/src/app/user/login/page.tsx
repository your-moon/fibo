"use client";
import { Button, Input } from "@nextui-org/react";
import React from "react";
import { fiboLogin, LoginResponseSuccess } from "./data";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie"; // import js-cookie
export interface Login {
  email: string;
  password: string;
}

export default function Login() {
  const router = useRouter();
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const onSubm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const login: Login = {
      email: email,
      password,
    };
    let res = (await fiboLogin(login)) as LoginResponseSuccess;
    if (res.status === 200) {
      if (!res.data?.token) {
        return;
      }
      Cookies.set("token", res?.data.token);
      router.push("/dashboard");
    }
  };
  return (
    <div className="h-screen flex flex-row justify-center">
      <div className="flex flex-col items-center mt-40">
        <h1 className="text-6xl font-extrabold mb-20 text-indigo-300">Login</h1>
        <form onSubmit={onSubm} className="w-52">
          <Input
            variant="bordered"
            placeholder="Username"
            value={email}
            onValueChange={setEmail}
            className="mb-2"
          />
          <Input
            variant="bordered"
            placeholder="Password"
            value={password}
            onValueChange={setPassword}
            type="password"
            className="mb-2"
          />
          <Button type="submit" className="w-full">
            Login
          </Button>
        </form>
      </div>
    </div>
  );
}
