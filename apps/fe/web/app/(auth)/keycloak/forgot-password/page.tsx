import ForgotPasswordForm from "../components/ForgotPasswordForm";
import Link from "next/link";

export default function Page() {
  return (
    <main style={{ padding: 24, maxWidth: 480, margin: "0 auto" }}>
      <h1>Forgot password</h1>
      <ForgotPasswordForm />
      <div style={{ marginTop: 16 }}>
        <Link href="/keycloak/login">Back to sign in</Link>
      </div>
    </main>
  );
}
