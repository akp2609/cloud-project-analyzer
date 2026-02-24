'use client';

import { useState } from 'react';

const NAV = [
  { id: 'overview', label: 'Overview', icon: '📊' },
  { id: 'anomalies', label: 'Cost Anomalies', icon: '⚠️' },
  { id: 'insights', label: 'Insights', icon: '💡' },
  { id: 'projects', label: 'Projects', icon: '📁' },
];

export default function Sidebar() {
  const [active, setActive] = useState('overview');

  return (
    <aside className="relative w-[260px] bg-gradient-to-b from-[#0f1420] to-[#0a0e1a] border-r border-[#1a2332] flex flex-col shadow-lg">
      {/* Glow accent */}
      <div className="absolute inset-0 bg-gradient-to-tr from-cyan-500/10 via-transparent to-purple-500/10 opacity-0 hover:opacity-100 transition-opacity duration-500" />

      {/* Logo / Title */}
      <div className="h-16 flex items-center px-6 border-b border-[#1a2332] relative z-10">
        <h1 className="text-xl font-extrabold text-cyan-400 tracking-tight">
          CloudWatch
        </h1>
      </div>

      {/* Navigation */}
      <nav className="flex-1 px-4 py-6 space-y-2 relative z-10">
        {NAV.map((item) => (
          <div
            key={item.id}
            onClick={() => setActive(item.id)}
            className={`flex items-center gap-3 px-4 py-2.5 rounded-lg cursor-pointer transition-all duration-300
              ${
                active === item.id
                  ? 'bg-cyan-500/10 text-cyan-400 shadow-inner'
                  : 'text-gray-400 hover:bg-[#141b2b] hover:text-gray-200'
              }`}
          >
            <span className="text-lg">{item.icon}</span>
            <span className="font-medium tracking-wide">{item.label}</span>
          </div>
        ))}
      </nav>

      {/* Footer */}
      <div className="px-6 py-4 border-t border-[#1a2332] text-xs text-gray-500 font-mono relative z-10">
        v1.0.0 • Dashboard
      </div>
    </aside>
  );
}
