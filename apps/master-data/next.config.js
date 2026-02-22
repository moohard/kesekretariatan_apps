/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  transpilePackages: ["@sikerma/ui", "@sikerma/auth", "@sikerma/shared"],
  images: {
    remotePatterns: [
      {
        protocol: 'http',
        hostname: 'localhost',
      },
      {
        protocol: 'https',
        hostname: 'your-cdn-domain.com',
      },
    ],
  },
  env: {
    NEXT_PUBLIC_APP_NAME: "SIKERMA Master Data",
    NEXT_PUBLIC_KEYCLOAK_URL: process.env.NEXT_PUBLIC_KEYCLOAK_URL || "http://localhost:8081",
    NEXT_PUBLIC_KEYCLOAK_REALM: process.env.NEXT_PUBLIC_KEYCLOAK_REALM || "pengadilan-agama",
    NEXT_PUBLIC_KEYCLOAK_CLIENT_ID: process.env.NEXT_PUBLIC_KEYCLOAK_CLIENT_ID || "master-data-client",
  },
  experimental: {
    ppr: true,
  },
}

export default nextConfig
