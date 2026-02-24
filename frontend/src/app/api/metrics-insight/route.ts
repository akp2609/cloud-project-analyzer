import { NextResponse } from 'next/server';
import { fetchMetricsInsight } from '@/lib/services/analysisEngine';

export async function GET(req: Request) {
  const { searchParams } = new URL(req.url);
  const projectId = searchParams.get('project_id');
  const metricType = searchParams.get('metric_type') || "";

  if (!projectId) {
    return NextResponse.json({ error: 'Missing project_id' }, { status: 400 });
  }

  try {
    const data = await fetchMetricsInsight(projectId, metricType);
    return NextResponse.json(data);
  } catch (err) {
    console.error('Error fetching metrics insight:', err);
    return NextResponse.json({ error: 'Failed to fetch metrics insight' }, { status: 500 });
  }
}
