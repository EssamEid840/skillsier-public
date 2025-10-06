// Light WebAuthn helpers (browser)
export async function createCredential(options: PublicKeyCredentialCreationOptions) {
  const cred = await (navigator as any).credentials.create({ publicKey: options as any });
  return cred as PublicKeyCredential | null;
}

export async function getAssertion(options: PublicKeyCredentialRequestOptions) {
  const cred = await (navigator as any).credentials.get({ publicKey: options as any });
  return cred as PublicKeyCredential | null;
}
