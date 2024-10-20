"use client";
// app/page.tsx
import React, { useEffect, useState } from "react";
import Post, { PostProps } from "./components/Post";
import config from "@/config";
import Sidebar from "./components/Sidebar";

const HomePage = () => {
  const [posts, setPosts] = useState<PostProps[]>([]);
  useEffect(() => {
    const fetchPosts = async () => {
      const response = await fetch(`${config.backendUrl}/api/v1/posts`, {
        credentials: "include",
      });
      const data = await response.json();
      // Assuming the data is in the expected format
      const postsRes = data.data;
      setPosts(postsRes);
    };

    fetchPosts();
  }, []);

  return (
    <div className="flex space-x-4">
      {" "}
      {/* Use flex to create a horizontal layout */}
      <div className="flex-1">
        {" "}
        {/* This will take up the remaining space */}
        <div className="">
          {posts.map((post) => (
            <Post key={post.id} {...post} />
          ))}
        </div>
      </div>
      <div className="w-1/2">
        <Sidebar />
      </div>
    </div>
  );
};

export default HomePage;
