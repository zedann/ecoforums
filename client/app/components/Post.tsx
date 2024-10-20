// components/Post.tsx
import config from "@/config";
import React from "react";

export interface PostProps {
  id: number;
  title: string;
  content: string;
  image: string;
  ups_number: string;
  downs_number: string;
  username: string;
  created_at: string;
}

const Post: React.FC<PostProps> = ({
  title,
  content,
  image,
  ups_number,
  downs_number,
  username,
  created_at,
}) => {
  return (
    <div className="max-w-xl bg-white rounded-lg shadow-md overflow-hidden my-4">
      <img
        src={`${config.backendUrl}/images/${image}`}
        alt={title}
        className="w-full h-48 object-cover mb-4" // Added margin bottom for spacing
      />
      <div className="p-4">
        <h2 className="text-xl text-black font-bold">{title}</h2>
        <p className="text-gray-700 my-2">{content}</p>
        <div className="flex justify-between items-center mt-4">
          <div>
            <span className="text-gray-500 text-sm">
              Posted by {username} on{" "}
              {new Date(created_at).toLocaleDateString()}
            </span>
          </div>
          <div className="flex items-center">
            <span className="text-green-500 mr-2">{ups_number} 👍</span>
            <span className="text-red-500">{downs_number} 👎</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Post;
