import ResetPasswordForm from "../../components/ResetPasswordForm";

export default function Page({ params }: { params: { token: string } }) {
  return (
    <main style={{ padding: 24, maxWidth: 480, margin: "0 auto" }}>
      <h1>Reset password</h1>
      <ResetPasswordForm token={params.token} />
    </main>
  );
}
