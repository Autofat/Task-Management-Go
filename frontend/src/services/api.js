import axios from "axios";

const API_URL = "http://localhost:4000";

const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Auth
export const login = (email, password) =>
  api.post("/auth/login", { email, password });

export const register = (email, password, fullname) =>
  api.post("/users", { email, password, fullname });

export const logout = () => api.post("/auth/logout");

// Projects
export const getProjects = (params) => api.get("/projects", { params });

export const getProjectById = (id) => api.get(`/projects/${id}`);

export const createProject = (title) => api.post("/projects", { title });

export const updateProject = (id, title) =>
  api.put(`/projects/${id}`, { title });

export const deleteProject = (id) => api.delete(`/projects/${id}`);

// Project Members
export const getProjectMembers = (projectId) =>
  api.get(`/projects/${projectId}/members`);

export const inviteMember = (projectId, userId) =>
  api.post(`/projects/${projectId}/members`, { user_id: userId });

export const updateMemberRole = (projectId, userId, role) =>
  api.put(`/projects/${projectId}/members/${userId}`, { role });

export const removeMember = (projectId, userId) =>
  api.delete(`/projects/${projectId}/members/${userId}`);

// Tasks
export const getTasks = (params) => api.get("/tasks", { params });

export const getTaskById = (id, projectId) =>
  api.get(`/tasks/${id}`, { params: { project_id: projectId } });

export const createTask = (taskData) => api.post("/tasks", taskData);

export const updateTask = (id, taskData) => api.put(`/tasks/${id}`, taskData);

export const deleteTask = (id) => api.delete(`/tasks/${id}`);

export default api;
