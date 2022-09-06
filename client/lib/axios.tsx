import axios, { AxiosError } from 'axios';
import { getEndPoint, getToken } from './token';

export interface MyAxiosError {
  code: string;
  details: any;
  message: string;
  status: string;
}

export const isAxiosError = (error: any): error is AxiosError => {
  return !!error.isAxiosError;
};

export const asyncApiClient = {
  create: async () => {
    return axios.create({
      baseURL: getEndPoint(),
      responseType: 'json',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${getToken()}`,
      },
    });
  },
};

export async function apiClient() {
  return await asyncApiClient.create();
}
