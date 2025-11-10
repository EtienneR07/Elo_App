import { getServerSession } from 'next-auth';
import { authOptions } from '@/app/api/auth/[...nextauth]/route';
import { AuthenticatedButtons, UnauthenticatedButton } from './components/HomeButtons';
import { HamburgerMenu } from './components/HamburgerMenu';
import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Home',
  description: 'Welcome to Elotify - Create leagues, track rankings, and compete with your friends using the ELO rating system.',
};

export default async function Home() {
  const session = await getServerSession(authOptions);

  const jsonLd = {
    '@context': 'https://schema.org',
    '@type': 'WebApplication',
    name: 'Elotify',
    applicationCategory: 'SportsApplication',
    description: 'Track rankings, manage leagues, and compete with friends using ELO rating system.',
    offers: {
      '@type': 'Offer',
      price: '0',
      priceCurrency: 'USD',
    },
    operatingSystem: 'Web',
  };

  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      {session && <HamburgerMenu />}

      <main className="flex min-h-screen items-center justify-center flex-col gap-8 font-sans p-8">
      {session ? (
        <>
          <div className="text-center">
            <h1 className="text-4xl mb-4">
              Welcome, {session.user?.name}!
            </h1>
            <p className="text-gray-600 text-lg">
              You are successfully logged in
            </p>
          </div>

          <div className="p-6 bg-gray-100 rounded-lg max-w-md">
            <p className="my-2"><strong>Name:</strong> {session.user?.name}</p>
            <p className="my-2"><strong>Email:</strong> {session.user?.email}</p>
          </div>

          {/* Client Component for interactivity */}
          <AuthenticatedButtons />
        </>
      ) : (
        <>
          <h1 className="text-4xl">Welcome to Elofy</h1>
          <p className="text-gray-600 text-lg">
            Please sign in to continue
          </p>

          {/* Client Component for interactivity */}
          <UnauthenticatedButton />
        </>
      )}
    </main>
    </>
  );
}
