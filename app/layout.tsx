import type { Metadata } from 'next'
import { Providers } from './providers'
import './globals.css'

export const metadata: Metadata = {
  title: {
    default: 'Elo App - Competitive League Management',
    template: '%s | Elo App',
  },
  description: 'Track rankings, manage leagues, and compete with friends using ELO rating system. Perfect for ping pong, foosball, and any competitive sport.',
  keywords: ['elo rating', 'league management', 'sports ranking', 'competition tracker', 'ping pong league', 'foosball league'],
  authors: [{ name: 'Elo App' }],
  creator: 'Elo App',
  openGraph: {
    type: 'website',
    locale: 'en_US',
    url: 'https://elo-app.com',
    title: 'Elo App - Competitive League Management',
    description: 'Track rankings, manage leagues, and compete with friends using ELO rating system.',
    siteName: 'Elo App',
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Elo App - Competitive League Management',
    description: 'Track rankings, manage leagues, and compete with friends using ELO rating system.',
  },
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      'max-video-preview': -1,
      'max-image-preview': 'large',
      'max-snippet': -1,
    },
  },
  verification: {
    // Add your verification tokens here when you set up
    // google: 'your-google-verification-code',
    // yandex: 'your-yandex-verification-code',
  },
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <head>
        <link rel="icon" href="/favicon.ico" sizes="any" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="theme-color" content="#5b6a45" />
      </head>
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
