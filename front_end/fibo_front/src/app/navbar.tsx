"use client";
import React, { useEffect } from "react";
import Cookie from "js-cookie";
import {
  Button,
  Navbar as NavbarComponent,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
  Link,
  NavbarMenuToggle,
} from "@nextui-org/react";

export default function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = React.useState(false);
  const [active, setActive] = React.useState<number>(0);
  const [isLogged, setIsLogged] = React.useState<boolean>(false);

  useEffect(() => {
    if (Cookie.get("token")) {
      setIsLogged(true);
    }
  }, []);

  if (isLogged) {
    return (
      <NavbarComponent>
        <NavbarContent>
          <NavbarMenuToggle
            aria-label={isMenuOpen ? "Close menu" : "Open menu"}
            className="sm:hidden"
          />
          <NavbarBrand>
            <Link href="/">
              <p className="font-bold text-inherit">FIBO</p>
            </Link>
          </NavbarBrand>
        </NavbarContent>
        <NavbarContent className="hidden sm:flex gap-4" justify="center">
          <NavbarItem>
            <Link color="foreground" href="/">
              Articles
            </Link>
          </NavbarItem>
          <NavbarItem>
            <Link color="foreground" href="/dashboard" aria-current="page">
              BackOffice
            </Link>
          </NavbarItem>
          <NavbarItem>
            <Link color="foreground" href="/user/write">
              Write
            </Link>
          </NavbarItem>
          <NavbarItem>
            <Link color="foreground" href="/admin">
              Admin
            </Link>
          </NavbarItem>
        </NavbarContent>
        <NavbarContent justify="end">
          <NavbarItem>
            <Button
              as={Link}
              color="danger"
              href="#"
              variant="flat"
              onClick={() => {
                Cookie.remove("token");
                setIsLogged(false);
              }}
            >
              Log Out
            </Button>
          </NavbarItem>
        </NavbarContent>
      </NavbarComponent>
    );
  }

  return (
    <NavbarComponent>
      <NavbarBrand>
        <Link href="/">
          <p className="font-bold text-inherit">FIBO</p>
        </Link>
      </NavbarBrand>
      <NavbarContent className="hidden sm:flex gap-4" justify="center">
        <NavbarItem>
          <Link color="foreground" href="/">
            Articles
          </Link>
        </NavbarItem>
        <NavbarItem isActive>
          <Link href="/dashboard" aria-current="page">
            BackOffice
          </Link>
        </NavbarItem>
        <NavbarItem>
          <Link color="foreground" href="/user/write">
            Write
          </Link>
        </NavbarItem>
      </NavbarContent>
      <NavbarContent justify="end">
        <NavbarItem className="hidden lg:flex">
          <Link href="/user/login">Login</Link>
        </NavbarItem>
        <NavbarItem>
          <Button
            as={Link}
            color="primary"
            href="/user/register"
            variant="flat"
          >
            Sign Up
          </Button>
        </NavbarItem>
      </NavbarContent>
    </NavbarComponent>
  );
}
