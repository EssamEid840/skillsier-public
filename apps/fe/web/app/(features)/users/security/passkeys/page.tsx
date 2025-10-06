"use client";
import React from "react";
import { Flags, enableDevDefaults } from "@skillsier/flags";
// ✅ import from ui root
import { PasskeyBadge } from "@skillsier/ui";

export default function PasskeysPage() {
  enableDevDefaults(); // dev-only defaults
  const show = Flags.get<boolean>("users.profile.passkeys", false);

  return (
    <div className="p-6">
      <h2 className="text-xl font-semibold">Passkeys</h2>
      {show ? <PasskeyBadge /> : <p>Passkeys are disabled.</p>}
    </div>
  );
}
