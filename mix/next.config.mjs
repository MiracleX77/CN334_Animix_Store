/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: false,
  images: {
    domains: ['images.unsplash.com','via.placeholder.com'],
  },
  typescript: {
    ignoreBuildErrors: true,
  },
};

export default nextConfig;
