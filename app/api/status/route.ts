import { NextResponse } from 'next/server'

export async function GET() {
  return NextResponse.json({
    status: 'running',
    uptime: process.uptime(),
    timestamp: new Date().toISOString()
  })
}
