import { anomalies } from '@/lib/data';
import SeverityBadge from '@/components/shared/SeverityBadge';

export default function AnomaliesList() {
  return (
    <div className="relative bg-gradient-to-b from-[#111827] to-[#0f1420] border border-[#1a2332] rounded-2xl p-6 shadow-lg transition-all hover:border-cyan-400/40 hover:shadow-cyan-500/20">
      
      <div className="absolute inset-0 rounded-2xl bg-gradient-to-tr from-cyan-500/10 via-transparent to-transparent opacity-0 hover:opacity-100 transition-opacity duration-500" />

      
      <h3 className="font-bold text-lg mb-6 tracking-tight text-gray-100 relative z-10">
        Recent Anomalies
      </h3>

      {/* List */}
      <div className="space-y-4 relative z-10">
        {anomalies.map((a) => (
          <div
            key={a.id}
            className="group p-5 bg-[#0a0e1a] border border-[#1a2332] rounded-xl transition-all hover:border-cyan-500/30 hover:bg-[#111827]"
          >
            <div className="flex justify-between items-center mb-2">
              <div className="font-semibold text-gray-200 group-hover:text-cyan-300 transition-colors">
                {a.service}
              </div>
              <SeverityBadge level={a.severity as any} />
            </div>
            <div className="text-sm text-gray-400 font-mono">
              {a.spike} • {a.time}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
