'use client';

import { useRouter } from 'next/navigation';
import { signOut } from 'next-auth/react';

export function AuthenticatedButtons() {
  const router = useRouter();

  return (
    <div className="flex gap-4 flex-col items-center">
      <button
        onClick={() => signOut()}
        className="px-8 py-3 bg-error text-white rounded text-base font-medium cursor-pointer hover:bg-error-dark transition-colors"
      >
        Sign Out
      </button>
    </div>
  );
}

export function UnauthenticatedButton() {
  const router = useRouter();

  return (
    <button
      onClick={() => router.push('/login')}
      className="px-8 py-3 bg-primary-500 text-white rounded text-base font-medium cursor-pointer hover:bg-primary-600 transition-colors"
    >
      Sign In
    </button>
  );
}
