"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";

export default function VerifyEmailNotice() {
  const [msg, setMsg] = useState<string | null>(null);
  const [err, setErr] = useState<string | null>(null);

  async function resend() {
    setMsg(null); setErr(null);
    const res = await fetch("/api/auth/keycloak/send-verify-email", { method: "POST" });
    if (res.ok) setMsg("Verification email sent.");
    else setErr("Could not send verification email.");
  }

  return (
    <div className="grid gap-3">
      <p>Please check your inbox and click the verification link.</p>
      <Button onClick={resend}>Resend verification email</Button>
      {msg && <div className="text-sm text-emerald-600">{msg}</div>}
      {err && <div className="text-sm text-red-600">{err}</div>}
    </div>
  );
}
