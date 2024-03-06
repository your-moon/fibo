"use client";

import { NextUIProvider } from "@nextui-org/react";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import getenv from "getenv";
export const BACKEND_URL = getenv.string(
  "BACKEND_URL",
  "http://localhost:3005",
);

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <NextUIProvider className="h-full">
      <NextThemesProvider
        attribute="class"
        enableSystem={false}
        defaultTheme="light"
      >
        {children}
      </NextThemesProvider>
    </NextUIProvider>
  );
}
