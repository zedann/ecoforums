// app/settings/page.tsx
"use client"; // This indicates that this component will use client-side hooks

import { useRouter } from "next/navigation"; // Import from 'next/navigation'

const Settings = () => {
  const router = useRouter();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen ">
      <button
        onClick={() => router.back()}
        className="mt-4 bg-gray-400 text-white py-2 px-4 rounded-md"
      >
        Back
      </button>
    </div>
  );
};

export default Settings;
