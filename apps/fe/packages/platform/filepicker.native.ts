/**
 * File picker (native).
 * Stubbed to keep dependencies minimal; wire into expo-document-picker later.
 */
export async function pickFile(): Promise<{ uri: string } | null> {
  throw new Error("filepicker.native.pickFile not implemented — use expo-document-picker and replace this adapter.");
}
