// apps/web/src/app/api/health/route.ts
// Health check endpoint for Kubernetes probes

import { NextResponse } from 'next/server';

export async function GET() {
  try {
    // Check if app is healthy
    const health = {
      status: 'healthy',
      timestamp: new Date().toISOString(),
      uptime: process.uptime(),
      environment: process.env.NODE_ENV,
      version: process.env.npm_package_version || '1.0.0',
    };

    return NextResponse.json(health, { status: 200 });
  } catch (error) {
    return NextResponse.json(
      {
        status: 'unhealthy',
        timestamp: new Date().toISOString(),
        error: error instanceof Error ? error.message : 'Unknown error',
      },
      { status: 503 }
    );
  }
}

// Support HEAD requests for basic health checks
export async function HEAD() {
  return new NextResponse(null, { status: 200 });
}