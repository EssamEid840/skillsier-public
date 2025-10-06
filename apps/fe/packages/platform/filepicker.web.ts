export async function pickImage(): Promise<{ uri: string; name?: string; type?: string } | null> {
  return new Promise((resolve) => {
    const input = document.createElement("input");
    input.type = "file";
    input.accept = "image/*";
    input.onchange = () => {
      const file = input.files?.[0];
      if (!file) return resolve(null);
      resolve({ uri: URL.createObjectURL(file), name: file.name, type: file.type });
    };
    input.click();
  });
}
