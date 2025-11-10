import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Manage Leagues',
  description: 'Create and manage your competitive leagues. Track player rankings, record game results, and view detailed statistics.',
};

export default function LeaguesLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return children;
}
