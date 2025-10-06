// apps/fe/web/app/layout.tsx
import "./globals.css"
import { Toaster } from "@/components/ui/toaster"
import { I18nProvider } from "../src/providers/I18nProvider"
import { ThemeProvider } from "../src/providers/ThemeProvider"
import { AnalyticsProvider } from "../src/providers/AnalyticsProvider"
import { QueryProvider } from "../src/providers/QueryProvider"
import { StateProvider } from "../src/providers/StateProvider"
import { AuthProvider } from "../src/providers/AuthProvider"

// ⬇️ add this
import AppHeaderWeb from "./(features)/users/components/AppHeaderWeb"

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <I18nProvider>
          <ThemeProvider>
            <AnalyticsProvider>
              <QueryProvider>
                <StateProvider>
                  <AuthProvider>
                    {/* ⬇️ NEW global header that matches your (features)/…/Web.tsx convention */}
                    <AppHeaderWeb />
                    {children}
                    <Toaster />
                  </AuthProvider>
                </StateProvider>
              </QueryProvider>
            </AnalyticsProvider>
          </ThemeProvider>
        </I18nProvider>
      </body>
    </html>
  )
}
