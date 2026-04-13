import { MetadataRoute } from "next";

export default function manifest(): MetadataRoute.Manifest {
  return {
    name: "寮ランドリー",
    short_name: "ランドリー",
    description: "寮の共同洗濯機の空き状況確認と利用登録",
    start_url: "/",
    display: "standalone",
    background_color: "#f9fbff",
    theme_color: "#f9fbff",
    lang: "ja",
    icons: [
      {
        src: "/icon.svg",
        sizes: "any",
        type: "image/svg+xml"
      }
    ]
  };
}
