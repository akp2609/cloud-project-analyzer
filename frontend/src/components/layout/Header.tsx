export default function Header({ currentTime }: { currentTime: string }) {
  return (
    <header className="relative h-16 bg-gradient-to-r from-[#0f1420] via-[#111827] to-[#0f1420] border-b border-[#1a2332] flex items-center justify-between px-8 shadow-md">
      {/* Glow accent */}
      <div className="absolute inset-0 bg-gradient-to-r from-cyan-500/5 via-transparent to-purple-500/5 opacity-0 hover:opacity-100 transition-opacity duration-500" />

      {/* Project info */}
      <div className="relative z-10">
        <div className="text-xs uppercase tracking-wider text-gray-500 font-mono">
          Active Project
        </div>
        <div className="font-semibold text-gray-100 tracking-tight">
          Core Analytics Platform
        </div>
      </div>

      {/* Current time */}
      <div className="relative z-10 text-right font-mono text-sm text-gray-400 bg-[#0a0e1a] px-3 py-1 rounded-lg border border-[#1a2332] shadow-inner">
        {currentTime}
      </div>
    </header>
  );
}
