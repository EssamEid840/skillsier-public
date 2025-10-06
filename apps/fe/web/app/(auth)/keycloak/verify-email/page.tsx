import VerifyEmailNotice from "../components/VerifyEmailNotice";

export default function Page() {
  return (
    <main style={{ padding: 24, maxWidth: 480, margin: "0 auto" }}>
      <h1>Verify your email</h1>
      <VerifyEmailNotice />
    </main>
  );
}
