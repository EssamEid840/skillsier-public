// apps/web/src/app/api/auth/keycloak/oauth/callback/[provider]/route.ts
// OAuth callback endpoint with PKCE verification

import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { exchangeCodeForTokens, verifyToken } from '@/lib/keycloak';

export async function GET(
  request: NextRequest,
  { params }: { params: { provider: string } }
) {
  try {
    const { searchParams } = request.nextUrl;
    const code = searchParams.get('code');
    const state = searchParams.get('state');
    const error = searchParams.get('error');
    const errorDescription = searchParams.get('error_description');

    // Check for OAuth errors
    if (error) {
      console.error('OAuth error:', error, errorDescription);
      return NextResponse.redirect(
        `${request.nextUrl.origin}/login?error=${encodeURIComponent(
          errorDescription || error
        )}`
      );
    }

    // Validate required parameters
    if (!code || !state) {
      return NextResponse.redirect(
        `${request.nextUrl.origin}/login?error=missing_parameters`
      );
    }

    // Get cookies
    const cookieStore = cookies();
    const savedState = cookieStore.get('oauth_state')?.value;
    const codeVerifier = cookieStore.get('code_verifier')?.value;
    const provider = cookieStore.get('oauth_provider')?.value;

    // Verify state to prevent CSRF
    if (!savedState || savedState !== state) {
      console.error('State mismatch - possible CSRF attack');
      return NextResponse.redirect(
        `${request.nextUrl.origin}/login?error=invalid_state`
      );
    }

    // Verify code verifier exists
    if (!codeVerifier) {
      console.error('Code verifier not found');
      return NextResponse.redirect(
        `${request.nextUrl.origin}/login?error=missing_verifier`
      );
    }

    // Exchange authorization code for tokens
    const redirectUri = `${request.nextUrl.origin}/api/auth/keycloak/oauth/callback/${params.provider}`;

    const tokens = await exchangeCodeForTokens({
      code,
      redirectUri,
      codeVerifier,
    });

    // Verify the access token and get user info
    const userInfo = await verifyToken(tokens.access_token);

    // Create response redirecting to dashboard
    const response = NextResponse.redirect(
      `${request.nextUrl.origin}/dashboard`
    );

    // Store tokens in secure cookies
    response.cookies.set('access_token', tokens.access_token, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: tokens.expires_in,
      path: '/',
    });

    response.cookies.set('refresh_token', tokens.refresh_token, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: tokens.refresh_expires_in,
      path: '/',
    });

    // Store user info
    response.cookies.set('user_info', JSON.stringify(userInfo), {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      sameSite: 'lax',
      maxAge: tokens.expires_in,
      path: '/',
    });

    // Clear temporary OAuth cookies
    response.cookies.delete('oauth_state');
    response.cookies.delete('code_verifier');
    response.cookies.delete('oauth_provider');

    return response;
  } catch (error) {
    console.error('OAuth callback error:', error);
    const errorMessage =
      error instanceof Error ? error.message : 'Authentication failed';
    return NextResponse.redirect(
      `${request.nextUrl.origin}/login?error=${encodeURIComponent(
        errorMessage
      )}`
    );
  }
}