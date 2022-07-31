import axios from "axios";

const instance = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  timeout: 5000,
});

instance.interceptors.request.use((config) => {
  if (config?.headers?.Authorization) {
    return config;
  }

  const token = localStorage.getItem("token");

  return {
    ...config,
    headers: {
      ...config.headers,
      Authorization: `Bearer ${token}`,
    },
  };
});

instance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    const message = error?.response?.data?.error
      ? { message: error?.response?.data?.error }
      : error;

    return Promise.reject(message);
  }
);

export default instance;
