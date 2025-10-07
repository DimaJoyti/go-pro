"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function Lesson1Page() {
  const router = useRouter();

  useEffect(() => {
    // Redirect to the dynamic lesson page
    router.replace("/learn/lesson-1");
  }, [router]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-background via-background to-accent/5">
      <div className="container-responsive padding-responsive-y">
        <div className="flex items-center justify-center min-h-[60vh]">
          <div className="text-center">
            <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
            <p className="text-responsive text-muted-foreground">Redirecting to lesson...</p>
          </div>
        </div>
      </div>
    </div>
  );
}
