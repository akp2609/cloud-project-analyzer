export default function SeverityBadge({ level }: { level: 'HIGH' | 'MEDIUM' | 'LOW' }) {
  const map = {
    HIGH: 'bg-red-500/10 text-red-400 border-red-500/20',
    MEDIUM: 'bg-amber-500/10 text-amber-400 border-amber-500/20',
    LOW: 'bg-cyan-500/10 text-cyan-400 border-cyan-500/20',
  };

  return (
    <span className={`px-2 py-1 text-xs font-bold border rounded ${map[level]}`}>
      {level}
    </span>
  );
}
