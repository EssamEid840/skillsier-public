import { ApiError } from './base-client';

export const handleApiError = (error: unknown): never => {
  if (error instanceof ApiError) {
    throw error;
  }

  if (error instanceof Error) {
    throw new ApiError(0, 'UNKNOWN_ERROR', error.message);
  }

  throw new ApiError(0, 'UNKNOWN_ERROR', 'An unknown error occurred');
};

export const isApiError = (error: unknown): error is ApiError => {
  return error instanceof ApiError;
};

export const getErrorMessage = (error: unknown): string => {
  if (isApiError(error)) {
    return error.message;
  }

  if (error instanceof Error) {
    return error.message;
  }

  return 'An unknown error occurred';
};