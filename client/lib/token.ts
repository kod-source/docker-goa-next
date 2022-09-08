export const getToken = (): string => {
  const token = localStorage.getItem('token');
  if (!token) return '';
  return token;
};

export const getEndPoint = (): string => {
  const endPoint = process.env.NEXT_PUBLIC_END_POINT;
  if (!endPoint) return '';
  return endPoint;
};
