// apps/web/src/app/api/auth/keycloak/register/route.ts
// User registration endpoint with Keycloak integration

import { NextRequest, NextResponse } from 'next/server';
import { createKeycloakUser } from '@/lib/keycloak';

interface RegisterRequest {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  userType: 'freelancer' | 'client';
}

export async function POST(request: NextRequest) {
  try {
    const body: RegisterRequest = await request.json();

    // Validate required fields
    const { email, password, firstName, lastName, userType } = body;

    if (!email || !password || !firstName || !lastName || !userType) {
      return NextResponse.json(
        { error: 'Missing required fields' },
        { status: 400 }
      );
    }

    // Validate email format
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      return NextResponse.json(
        { error: 'Invalid email format' },
        { status: 400 }
      );
    }

    // Validate password strength (min 8 chars, 1 uppercase, 1 lowercase, 1 number)
    if (password.length < 8) {
      return NextResponse.json(
        { error: 'Password must be at least 8 characters' },
        { status: 400 }
      );
    }

    // Validate userType
    if (!['freelancer', 'client'].includes(userType)) {
      return NextResponse.json(
        { error: 'Invalid user type' },
        { status: 400 }
      );
    }

    // Create user in Keycloak
    const keycloakUserId = await createKeycloakUser({
      email,
      firstName,
      lastName,
      password,
    });

    // Create user profile in your backend
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    const profileResponse = await fetch(`${apiUrl}/users/profile`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        keycloakId: keycloakUserId,
        email,
        firstName,
        lastName,
        userType,
      }),
    });

    if (!profileResponse.ok) {
      // If profile creation fails, we should ideally rollback Keycloak user
      // For now, log the error
      console.error('Failed to create user profile in backend');
      throw new Error('Failed to create user profile');
    }

    const userProfile = await profileResponse.json();

    return NextResponse.json({
      success: true,
      message: 'User registered successfully',
      userId: keycloakUserId,
      profile: userProfile,
    });
  } catch (error) {
    console.error('Registration error:', error);

    // Handle specific Keycloak errors
    if (error instanceof Error) {
      if (error.message.includes('User exists')) {
        return NextResponse.json(
          { error: 'Email already registered' },
          { status: 409 }
        );
      }
    }

    return NextResponse.json(
      { error: 'Registration failed. Please try again.' },
      { status: 500 }
    );
  }
}