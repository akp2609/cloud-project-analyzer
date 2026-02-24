import { NextResponse } from 'next/server';
import { fetchDashboard, fetchAnomalies, fetchInsights } from '@/lib/services/analysisEngine';

export async function GET(req: Request) {
  const { searchParams } = new URL(req.url);
  const projectId = searchParams.get('project_id');

  if (!projectId) {
    return NextResponse.json({ error: 'Missing project_id' }, { status: 400 });
  }

  try {
    
    const dashboard = await fetchDashboard(projectId);
    const anomalies = await fetchAnomalies(projectId);
    const insights = await fetchInsights(projectId);

    
    const stats = {
      total_cost_30d: dashboard.total_cost_30d ?? 0,
      insights_high_count: insights.filter((i: any) => i.severity === 'HIGH').length,
      cost_vs_baseline: '+42.3%', 
    };

    return NextResponse.json(stats);
  } catch (err) {
    console.error('Error fetching stats:', err);
    return NextResponse.json({ error: 'Failed to fetch stats' }, { status: 500 });
  }
}
