"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function ResetPasswordForm({ token }: { token: string }) {
  const [password, setPassword] = useState("");
  const [confirm, setConfirm] = useState("");
  const [msg, setMsg] = useState<string | null>(null);
  const [err, setErr] = useState<string | null>(null);

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setMsg(null); setErr(null);
    if (password !== confirm) { setErr("Passwords do not match"); return; }
    const res = await fetch("/api/auth/keycloak/reset-password", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ token, password }),
    });
    if (res.ok) setMsg("Password updated. You can sign in now.");
    else setErr("Reset failed");
  }

  return (
    <form onSubmit={onSubmit} className="grid gap-3">
      <div className="grid gap-1">
        <Label htmlFor="pw">New password</Label>
        <Input id="pw" type="password" value={password} onChange={e=>setPassword(e.target.value)} />
      </div>
      <div className="grid gap-1">
        <Label htmlFor="confirm">Confirm password</Label>
        <Input id="confirm" type="password" value={confirm} onChange={e=>setConfirm(e.target.value)} />
      </div>
      {msg && <div className="text-sm text-emerald-600">{msg}</div>}
      {err && <div className="text-sm text-red-600">{err}</div>}
      <Button type="submit">Update password</Button>
    </form>
  );
}
