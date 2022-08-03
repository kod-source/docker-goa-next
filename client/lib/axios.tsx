import { AxiosError } from 'axios';

export interface MyAxiosError {
  code: string;
  details: any;
  message: string;
  status: string;
}

export const isAxiosError = (error: any): error is AxiosError => {
  return !!error.isAxiosError;
};
