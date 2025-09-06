const isProd = process.env.NODE_ENV === "production";
// simulating CDN locally
const useCdn = isProd || process.env.NEXT_FORCE_CDN === "1";

/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    if (!useCdn) return [];
    return [
      {
        source: "/images/:path*",
        destination: "https://cdn.winnable.lol/images/:path*",
      },
    ];
  },

  // strong caching for rewritten images in prod
  async headers() {
    if (!useCdn) return [];
    return [
      {
        source: "/images/:path*",
        headers: [
          {
            key: "Cache-Control",
            value: "public, max-age=31536000, immutable",
          },
        ],
      },
    ];
  },
};

export default nextConfig;
