'use client';

import { useEffect, useState } from 'react';

interface StatsData {
  total_cost_30d: number | string;
  anomalies_count: number;
  insights_high_count: number;
  cost_vs_baseline: string;
}

export default function StatsGrid({ projectId }: { projectId: string }) {
  const [stats, setStats] = useState<StatsData | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        const res = await fetch(`/api/stats?project_id=${projectId}`);
        const data = await res.json();
        setStats(data);
      } catch (err) {
        console.error('Failed to load stats', err);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, [projectId]);

  if (loading) {
    return <div className="text-gray-400">Loading stats...</div>;
  }

  if (!stats) {
    return <div className="text-red-400">No stats available</div>;
  }

  const cards = [
    {
      title: 'Total Cost (30d)',
      value: `${Number(stats.total_cost_30d).toLocaleString()}`,
      subtitle: 'Last 30 days',
      icon: '💰',
    },
    {
      title: 'Active Anomalies',
      value: stats.anomalies_count,
      subtitle: 'Requiring attention',
      icon: '⚠️',
    },
    {
      title: 'High Priority Insights',
      value: stats.insights_high_count,
      subtitle: 'Critical items',
      icon: '💡',
    },
    {
      title: 'Cost vs Baseline',
      value: stats.cost_vs_baseline,
      subtitle: 'Above expected',
      trendColor: 'text-red-400',
      icon: '📈',
    },
  ];

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 mb-12">
      {cards.map((stat, idx) => (
        <div
          key={idx}
          className="relative group bg-gradient-to-b from-[#111827] to-[#0f1420] border border-[#1a2332] rounded-2xl p-6 shadow-lg transition-all duration-300 hover:border-cyan-400/40 hover:shadow-cyan-500/20 hover:-translate-y-1"
        >
          <div className="absolute inset-0 rounded-2xl bg-gradient-to-tr from-cyan-500/10 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-500" />
          <div className="flex justify-between items-start mb-6 relative z-10">
            <div className="text-[11px] uppercase tracking-wider text-gray-400 font-mono font-semibold">
              {stat.title}
            </div>
            <div className="w-10 h-10 rounded-xl bg-cyan-500/10 flex items-center justify-center text-xl shadow-inner">
              {stat.icon}
            </div>
          </div>
          <div className="text-4xl font-extrabold font-mono text-gray-100 mb-3 tracking-tight relative z-10">
            {stat.value}
          </div>
          <div className="flex items-center gap-3 text-sm relative z-10">
            <span className="text-gray-500">{stat.subtitle}</span>
            {stat.trendColor && (
              <span className={`font-mono text-xs font-bold ${stat.trendColor}`}>
                {stat.trendColor === 'text-red-400' ? '+12.3%' : ''}
              </span>
            )}
          </div>
        </div>
      ))}
    </div>
  );
}
