import { ImageResponse } from "next/og";

export const size = { width: 1200, height: 630 };
export const contentType = "image/png";
export const alt = "Skillsier Profile";

export default function Image() {
  return new ImageResponse(
    (
      <div
        style={{
          fontSize: 48,
          color: "#111",
          width: "100%",
          height: "100%",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          background: "#f5f5f5"
        }}
      >
        Skillsier Profile
      </div>
    ),
    { ...size }
  );
}
