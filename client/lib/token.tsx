export const getToken = (): string => {
  const token = localStorage.getItem('token');
  if (!token) return '';
  return token;
};
