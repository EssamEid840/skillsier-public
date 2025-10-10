// apps/web/src/app/api/auth/keycloak/oauth/start/[provider]/route.ts
// OAuth initiation endpoint with PKCE support

import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import crypto from 'crypto';
import { buildAuthorizationUrl, generateCodeVerifier } from '@/lib/keycloak';

export async function GET(
  request: NextRequest,
  { params }: { params: { provider: string } }
) {
  try {
    const { provider } = params;

    // Validate provider
    const validProviders = ['google', 'github', 'local'];
    if (!validProviders.includes(provider)) {
      return NextResponse.json(
        { error: 'Invalid provider' },
        { status: 400 }
      );
    }

    // Generate state and code verifier for PKCE
    const state = crypto.randomBytes(32).toString('base64url');
    const codeVerifier = generateCodeVerifier();

    // Build redirect URI
    const redirectUri = `${request.nextUrl.origin}/api/auth/keycloak/oauth/callback/${provider}`;

    // Build authorization URL
    const authUrl = buildAuthorizationUrl({
      redirectUri,
      state,
      codeVerifier,
      idpHint: provider !== 'local' ? provider : undefined,
    });

    // Create response
    const response = NextResponse.redirect(authUrl);

    // Store state and code verifier in encrypted cookies
    const cookieStore = cookies();
    
    // Set secure, httpOnly cookies
    response.cookies.set('oauth_state', state, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: 600, // 10 minutes
      path: '/',
    });

    response.cookies.set('code_verifier', codeVerifier, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: 600, // 10 minutes
      path: '/',
    });

    response.cookies.set('oauth_provider', provider, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: 600,
      path: '/',
    });

    return response;
  } catch (error) {
    console.error('OAuth start error:', error);
    return NextResponse.json(
      { error: 'Authentication initiation failed' },
      { status: 500 }
    );
  }
}

// Handle POST for programmatic OAuth initiation
export async function POST(
  request: NextRequest,
  { params }: { params: { provider: string } }
) {
  return GET(request, { params });
}