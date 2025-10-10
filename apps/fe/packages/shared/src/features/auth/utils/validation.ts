// packages/shared/src/features/auth/utils/validation.ts
// Authentication validation utilities

export interface ValidationResult {
  isValid: boolean;
  errors: string[];
}

/**
 * Validate email format
 */
export function validateEmail(email: string): ValidationResult {
  const errors: string[] = [];
  
  if (!email) {
    errors.push('Email is required');
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    errors.push('Invalid email format');
  }
  
  return {
    isValid: errors.length === 0,
    errors,
  };
}

/**
 * Validate password strength
 */
export function validatePassword(password: string): ValidationResult {
  const errors: string[] = [];
  
  if (!password) {
    errors.push('Password is required');
  } else {
    if (password.length < 8) {
      errors.push('Password must be at least 8 characters');
    }
    if (!/[A-Z]/.test(password)) {
      errors.push('Password must contain at least one uppercase letter');
    }
    if (!/[a-z]/.test(password)) {
      errors.push('Password must contain at least one lowercase letter');
    }
    if (!/[0-9]/.test(password)) {
      errors.push('Password must contain at least one number');
    }
  }
  
  return {
    isValid: errors.length === 0,
    errors,
  };
}

/**
 * Get password strength level
 */
export function getPasswordStrength(password: string): {
  level: 'weak' | 'medium' | 'strong';
  score: number;
} {
  let score = 0;
  
  if (password.length >= 8) score++;
  if (password.length >= 12) score++;
  if (/[A-Z]/.test(password)) score++;
  if (/[a-z]/.test(password)) score++;
  if (/[0-9]/.test(password)) score++;
  if (/[^A-Za-z0-9]/.test(password)) score++;
  
  if (score <= 2) return { level: 'weak', score };
  if (score <= 4) return { level: 'medium', score };
  return { level: 'strong', score };
}

/**
 * Validate name
 */
export function validateName(name: string, fieldName: string = 'Name'): ValidationResult {
  const errors: string[] = [];
  
  if (!name) {
    errors.push(`${fieldName} is required`);
  } else if (name.length < 2) {
    errors.push(`${fieldName} must be at least 2 characters`);
  } else if (name.length > 50) {
    errors.push(`${fieldName} must be less than 50 characters`);
  }
  
  return {
    isValid: errors.length === 0,
    errors,
  };
}

/**
 * Validate registration data
 */
export function validateRegistration(data: {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  confirmPassword?: string;
}): ValidationResult {
  const errors: string[] = [];
  
  // Validate email
  const emailResult = validateEmail(data.email);
  errors.push(...emailResult.errors);
  
  // Validate password
  const passwordResult = validatePassword(data.password);
  errors.push(...passwordResult.errors);
  
  // Validate password confirmation
  if (data.confirmPassword && data.password !== data.confirmPassword) {
    errors.push('Passwords do not match');
  }
  
  // Validate names
  const firstNameResult = validateName(data.firstName, 'First name');
  errors.push(...firstNameResult.errors);
  
  const lastNameResult = validateName(data.lastName, 'Last name');
  errors.push(...lastNameResult.errors);
  
  return {
    isValid: errors.length === 0,
    errors,
  };
}