'use client';

import { useEffect, useState } from 'react';

interface MetricInsight {
  project_id: string;
  metric_type: string;
  avg?: number;
  max?: number;
  min?: number;
  count?: number;
}

export default function MetricsGrid({ projectId, metricType }: { projectId: string, metricType?: string }) {
  const [insights, setInsights] = useState<MetricInsight[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        const url = `/api/metrics-insight?project_id=${projectId}${metricType ? `&metric_type=${metricType}` : ''}`;
        const res = await fetch(url);
        const data = await res.json();
        setInsights(Array.isArray(data) ? data : [data]);
      } catch (err) {
        console.error('Failed to load metrics insights', err);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, [projectId, metricType]);

  if (loading) {
    return <div className="text-gray-400">Loading metrics...</div>;
  }

  return (
  <div className="w-full">
    {insights.map((m, idx) => (
      <div
        key={idx}
        className="relative w-full h-full min-h-[200px] flex flex-col justify-between bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 border border-gray-700 rounded-2xl p-6 shadow-lg transition-all duration-300 hover:border-cyan-400/40 hover:shadow-cyan-500/20 hover:-translate-y-1"
      >
        <div className="absolute inset-0 rounded-2xl bg-gradient-to-tr from-purple-500/10 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500" />

        <div className="flex justify-between items-start mb-6 relative z-10">
          <div className="text-xs uppercase tracking-wider text-gray-400 font-mono font-semibold">
            {metricType ? m.metric_type : 'Total Metrics'}
          </div>
          <div className="w-10 h-10 rounded-xl bg-purple-500/10 flex items-center justify-center text-xl shadow-inner">
            📊
          </div>
        </div>

        {metricType === 'run.googleapis.com/request_count' ? (
          <div className="text-3xl font-extrabold font-mono text-cyan-400 mb-3 tracking-tight relative z-10">
            Count: {m.count ?? 'N/A'}
          </div>
        ) : (
          <>
            <div className="text-2xl font-extrabold font-mono text-gray-100 mb-3 tracking-tight relative z-10">
              Avg: {m.avg !== undefined ? m.avg.toFixed(2) : 'N/A'}
            </div>
            <div className="flex flex-col gap-1 text-sm relative z-10 text-gray-400 font-mono">
              <span>Max: {m.max !== undefined ? m.max.toFixed(2) : 'N/A'}</span>
              <span>Min: {m.min !== undefined ? m.min.toFixed(2) : 'N/A'}</span>
              <span>Count: {m.count ?? 'N/A'}</span>
            </div>
          </>
        )}
      </div>
    ))}
  </div>
);
}
