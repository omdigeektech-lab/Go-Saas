export const API_URL = 'http://localhost:8080';

export const getHeaders = () => {
  const token = localStorage.getItem('token');
  return {
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  };
};

export const fetchTasks = async () => {
  const res = await fetch(`${API_URL}/tasks`, { headers: getHeaders() });
  return res.json();
};

export const createTask = async (data: any) => {
  const res = await fetch(`${API_URL}/tasks`, {
    method: 'POST',
    headers: getHeaders(),
    body: JSON.stringify(data),
  });
  return res.json();
};
