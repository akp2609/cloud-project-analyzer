import { Providers } from "@/components/layout/Providers"
import Sidebar from "@/components/layout/Sidebar"
import Header from "@/components/layout/Header"
import "./globals.css"


export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" suppressHydrationWarning>
  <body className="bg-gradient-to-br from-[#0a0e1a] via-[#0d1320] to-[#0a0e1a] text-gray-200 antialiased">
    <Providers>
      <div className="flex h-screen overflow-hidden">
        <Sidebar />
        <div className="flex flex-1 flex-col relative">
          <Header currentTime="" />
          {children}
        </div>
      </div>
    </Providers>
  </body>
</html>

  )
}
