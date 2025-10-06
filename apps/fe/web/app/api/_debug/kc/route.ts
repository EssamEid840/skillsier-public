import { NextResponse } from "next/server";

export async function GET() {
  const keys = [
    "KEYCLOAK_ISSUER_URL",
    "KEYCLOAK_MGMT_CLIENT_ID",
    "KEYCLOAK_MGMT_CLIENT_SECRET",
    "KEYCLOAK_MGMT_REALM",
  ];
  const snapshot = Object.fromEntries(
    keys.map(k => [k, process.env[k] ? "SET" : "MISSING"])
  );
  return NextResponse.json(snapshot);
}
