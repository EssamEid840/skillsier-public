"use client";

import React, { useRef } from "react";

export function HydrateClient({ children }: { children: React.ReactNode }) {
  // Placeholder for any client hydration you may need
  const mounted = useRef(false);
  if (!mounted.current) mounted.current = true;
  return <>{children}</>;
}
