/** @type {import('next').NextConfig} */
const nextConfig = {};

// next.config.js
module.exports = {
  webpack: (config, { isServer }) => {
    if (!isServer) {
      config.node = {
        process: "mock", // This mocks the process object for the client side
      };
    }
    return config;
  },
  env: {
    MY_ENV_VARIABLE: process.env.MY_ENV_VARIABLE,
  },
};

export default nextConfig;
