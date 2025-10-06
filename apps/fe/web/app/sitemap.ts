export default function sitemap() {
  return [
    { url: "http://localhost:3000/", lastModified: new Date().toISOString() },
    { url: "http://localhost:3000/users", lastModified: new Date().toISOString() }
  ];
}
