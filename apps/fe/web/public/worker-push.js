// placeholder web push logic (wire to your push provider later)
self.addEventListener("push", (event) => {
  const data = event.data?.json?.() ?? {};
  event.waitUntil(
    self.registration.showNotification(data.title || "Skillsier", {
      body: data.body || "",
      data
    })
  );
});
