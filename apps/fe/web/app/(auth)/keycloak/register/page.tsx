import OAuthGoogleButton from "../components/OAuthGoogleButton";
import RegisterForm from "../components/RegisterForm";
import Link from "next/link";

export default function Page() {
  return (
    <main style={{ padding: 24, maxWidth: 480, margin: "0 auto" }}>
      <h1>Create account</h1>
      <RegisterForm />
      <div style={{ margin: "16px 0", textAlign: "center" }}>or</div>
      <OAuthGoogleButton />
      <div style={{ marginTop: 16 }}>
        <span>Already have an account? </span>
        <Link href="/keycloak/login">Sign in</Link>
      </div>
    </main>
  );
}
