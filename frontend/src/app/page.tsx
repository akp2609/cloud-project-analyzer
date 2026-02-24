'use client';

import { useEffect, useState } from 'react';
import StatsGrid from '@/components/dashboard/StatsGrid';
import MetricsGrid from '@/components/dashboard/MetricsGrid';
import AnomaliesList from '@/components/dashboard/AnomaliesList';
import InsightsList from '@/components/dashboard/InsightsList';

export default function Page() {
  const [currentTime, setCurrentTime] = useState('');
  const projectId = 'job-tracker-app-458110';

  useEffect(() => {
    const tick = () => {
      setCurrentTime(
        new Date().toLocaleString('en-US', {
          month: 'short',
          day: '2-digit',
          year: 'numeric',
          hour: '2-digit',
          minute: '2-digit',
          second: '2-digit',
        })
      );
    };
    tick();
    const id = setInterval(tick, 1000);
    return () => clearInterval(id);
  }, []);

  return (
    <main className="flex-1 flex flex-col relative overflow-hidden bg-gradient-to-br from-gray-950 via-gray-900 to-gray-950 min-h-screen">

      {/* Background glows */}
      <div className="absolute inset-0 pointer-events-none overflow-hidden">
        <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-cyan-500/10 rounded-full blur-[120px]" />
        <div className="absolute bottom-1/3 right-1/4 w-[32rem] h-[32rem] bg-purple-500/10 rounded-full blur-[140px]" />
        <div className="absolute top-2/3 left-1/2 w-64 h-64 bg-blue-500/8 rounded-full blur-[100px]" />
      </div>

      <div className="flex-1 overflow-y-auto relative z-10">
        <div className="w-full px-8 xl:px-12 py-10 space-y-12">

          {/* ── Header ── */}
          <header className="flex items-end justify-between pb-8 border-b border-white/[0.06]">
            <div className="space-y-1.5">
              <p className="text-xs font-semibold tracking-[0.25em] uppercase text-cyan-400/70">
                GCP Monitoring
              </p>
              <h1 className="text-4xl font-bold tracking-tight text-cyan-400 leading-none">
                Project Dashboard
              </h1>
              <p className="text-sm text-gray-500 font-mono pt-0.5">
                {projectId}
              </p>
            </div>

            <div className="flex items-center gap-3">
              {/* Live badge */}
              <div className="flex items-center gap-2 px-3 py-1.5 rounded-full border border-cyan-500/20 bg-cyan-500/5">
                <span className="relative flex h-1.5 w-1.5">
                  <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-cyan-400 opacity-75" />
                  <span className="relative inline-flex rounded-full h-1.5 w-1.5 bg-cyan-400" />
                </span>
                <span className="text-xs text-cyan-400/80 font-semibold tracking-wide">LIVE</span>
              </div>

              {/* Clock */}
              <div className="px-4 py-1.5 rounded-lg bg-white/[0.03] border border-white/[0.08]">
                <span className="text-xs text-gray-400 font-mono tabular-nums tracking-widest">
                  {currentTime}
                </span>
              </div>
            </div>
          </header>

          {/* ── Stats ── */}
          <section>
            <SectionLabel label="Overview" />
            <StatsGrid projectId={projectId} />
          </section>

          
          <section>
            <SectionLabel label="Metrics" />
            <div
              className="grid gap-6"
              style={{ gridTemplateColumns: 'repeat(3, minmax(0, 1fr))' }}
            >
              <MetricsGrid projectId={projectId} />
              <MetricsGrid projectId={projectId} metricType="run.googleapis.com/request_latencies" />
              <MetricsGrid projectId={projectId} metricType="run.googleapis.com/request_count" />
            </div>
          </section>

          
          <section>
            <SectionLabel label="Insights & Anomalies" />
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <AnomaliesList />
              <InsightsList />
            </div>
          </section>

        </div>
      </div>
    </main>
  );
}


function SectionLabel({ label }: { label: string }) {
  return (
    <div className="flex items-center gap-4 mb-5">
      <span className="text-[11px] font-bold tracking-[0.2em] uppercase text-gray-500 whitespace-nowrap">
        {label}
      </span>
      <div className="flex-1 h-px bg-gradient-to-r from-white/[0.08] to-transparent" />
    </div>
  );
}