"use client";
import { Button, Input } from "@nextui-org/react";
import React, { useEffect } from "react";
import { fiboRegister, RegisterResponseSuccess } from "./data";
import { useRouter } from "next/navigation";
export interface Register {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

export default function Login() {
  const router = useRouter();
  const [firstName, setFirstName] = React.useState("");
  const [lastName, setLastName] = React.useState("");
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [error, setError] = React.useState("");
  const [isInvalid, setIsInvalid] = React.useState(false);
  const onSubm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const register: Register = {
      firstName: firstName,
      lastName: lastName,
      email: email,
      password,
    };
    console.log(register);
    let res = await fiboRegister(register);
    console.log(res);
    if (res.status === 200) {
      const data = res as RegisterResponseSuccess;
      router.push("/admin");
    }
    if (res.status === 409 || res.status === 400) {
      setError(res.message);
      setIsInvalid(true);
    }
  };
  return (
    <div className="h-screen flex flex-row items-center justify-center">
      <div className="flex flex-col items-center">
        <form onSubmit={onSubm} className="w-52">
          <Input
            variant="bordered"
            placeholder="First Name"
            value={firstName}
            onValueChange={setFirstName}
            className="mb-2"
          />
          <Input
            variant="bordered"
            placeholder="Last Name"
            value={lastName}
            onValueChange={setLastName}
            className="mb-2"
          />
          <Input
            variant="bordered"
            placeholder="Email"
            value={email}
            onValueChange={setEmail}
            isInvalid={isInvalid}
            errorMessage={error}
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
            Register
          </Button>
        </form>
      </div>
    </div>
  );
}
