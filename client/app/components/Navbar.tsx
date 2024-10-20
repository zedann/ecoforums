"use client"; // Ensure this is a client-side component

import Link from "next/link";
import { useState } from "react";
import { useAuth } from "../context/AuthContext";
import { useRouter } from "next/navigation";

export default function Navbar() {
  const [isOpen, setIsOpen] = useState(false); // State for mobile menu visibility
  const { isLoggedIn, logout } = useAuth(); // State to track login status

  const router = useRouter();

  const handleLogout = async (e: React.FormEvent) => {
    e.preventDefault();

    // Make a request to the backend
    const res = await fetch("http://127.0.0.1:8080/api/v1/users/logout", {
      method: "GET",
    });

    if (res.ok) {
      logout();
      router.push("/login");
    }
  };

  return (
    <nav className="bg-blue-600 p-4">
      <div className="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
        <div className="relative flex items-center justify-between h-16">
          <div className="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
            <div className="flex-shrink-0">
              <Link href="/" className="text-white text-2xl font-bold">
                Investors Forums
              </Link>
            </div>
            <div className="hidden sm:block sm:ml-6">
              <div className="flex space-x-4">
                <Link
                  href="/"
                  className="text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                >
                  Home
                </Link>
                <Link
                  href="/about"
                  className="text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                >
                  About
                </Link>
              </div>
            </div>
          </div>
          {/* Right section for buttons */}
          <div className="flex items-center space-x-4">
            <button
              onClick={() => setIsOpen(!isOpen)} // Toggle mobile menu visibility
              className="sm:hidden text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >
              {/* Hamburger icon for mobile menu */}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M4 6h16M4 12h16m-7 6h7"
                />
              </svg>
            </button>

            {isLoggedIn ? (
              <button
                onClick={handleLogout}
                className="text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
              >
                Logout
              </button>
            ) : (
              <>
                <Link
                  href="/login"
                  className="text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                >
                  Login
                </Link>
                <Link
                  href="/signup"
                  className="text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                >
                  Signup
                </Link>
              </>
            )}
            <Link
              href="/settings"
              className="text-gray-200 hover:bg-blue-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >
              Settings
            </Link>
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      <div
        className={`${isOpen ? "block" : "hidden"} sm:hidden`} // Show or hide based on isOpen state
        id="mobile-menu"
      >
        <div className="px-2 pt-2 pb-3 space-y-1">
          <Link
            href="/"
            className="text-gray-200 hover:bg-blue-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
          >
            Home
          </Link>
          <Link
            href="/about"
            className="text-gray-200 hover:bg-blue-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
          >
            About
          </Link>
          {isLoggedIn ? (
            <button
              onClick={handleLogout}
              className="text-gray-200 hover:bg-blue-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
            >
              Logout
            </button>
          ) : (
            <>
              <Link
                href="/login"
                className="text-gray-200 hover:bg-blue-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
              >
                Login
              </Link>
              <Link
                href="/signup"
                className="text-gray-200 hover:bg-blue-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
              >
                Signup
              </Link>
            </>
          )}
        </div>
      </div>
    </nav>
  );
}
