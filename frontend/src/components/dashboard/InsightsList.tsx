import SeverityBadge from '@/components/shared/SeverityBadge';

const insights = [
  {
    id: 1,
    title: 'Unusual Storage Growth Pattern',
    description:
      'Cloud Storage costs increased significantly due to legacy backup retention.',
    severity: 'HIGH',
    time: '2h ago',
  },
  {
    id: 2,
    title: 'Compute Training Cost Spike',
    description:
      'ML training jobs ran with idle instances during off-hours.',
    severity: 'MEDIUM',
    time: '5h ago',
  },
  {
    id: 3,
    title: 'Cloud Run Traffic Surge',
    description:
      'New enterprise client increased API traffic substantially.',
    severity: 'HIGH',
    time: '1d ago',
  },
];

export default function InsightsList() {
  return (
    <div className="relative bg-gradient-to-b from-[#111827] to-[#0f1420] border border-[#1a2332] rounded-2xl p-6 shadow-lg transition-all hover:border-cyan-400/40 hover:shadow-cyan-500/20">
      {/* Glow accent */}
      <div className="absolute inset-0 rounded-2xl bg-gradient-to-tr from-purple-500/10 via-transparent to-transparent opacity-0 hover:opacity-100 transition-opacity duration-500" />

      {/* Title */}
      <h3 className="font-bold text-lg mb-6 tracking-tight text-gray-100 relative z-10">
        Latest Insights
      </h3>

      {/* List */}
      <div className="space-y-4 relative z-10">
        {insights.map((insight) => (
          <div
            key={insight.id}
            className="group p-5 bg-[#0a0e1a] border border-[#1a2332] rounded-xl transition-all hover:border-cyan-500/30 hover:bg-[#111827]"
          >
            <div className="flex justify-between items-start mb-3">
              <div className="font-semibold text-sm text-gray-200 group-hover:text-cyan-300 transition-colors pr-3">
                {insight.title}
              </div>
              <SeverityBadge level={insight.severity as any} />
            </div>

            <p className="text-sm text-gray-400 mb-3 line-clamp-2">
              {insight.description}
            </p>

            <div className="text-xs text-gray-500 font-mono">
              {insight.time}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
