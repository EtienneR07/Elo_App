'use client'

import { useSession, signOut } from 'next-auth/react'
import { useRouter } from 'next/navigation'

export default function Home() {
  const { data: session, status } = useSession()
  const router = useRouter()

  if (status === 'loading') {
    return (
      <main style={{
        display: 'flex',
        minHeight: '100vh',
        alignItems: 'center',
        justifyContent: 'center',
        fontFamily: 'system-ui, sans-serif'
      }}>
        <p>Loading...</p>
      </main>
    )
  }

  return (
    <main style={{
      display: 'flex',
      minHeight: '100vh',
      alignItems: 'center',
      justifyContent: 'center',
      flexDirection: 'column',
      gap: '2rem',
      fontFamily: 'system-ui, sans-serif',
      padding: '2rem'
    }}>
      {session ? (
        <>
          <div style={{ textAlign: 'center' }}>
            <h1 style={{ fontSize: '2.5rem', marginBottom: '1rem' }}>
              Welcome, {session.user?.name}!
            </h1>
            <p style={{ color: '#666', fontSize: '1.1rem' }}>
              You are successfully logged in
            </p>
          </div>

          <div style={{
            padding: '1.5rem',
            backgroundColor: '#f5f5f5',
            borderRadius: '8px',
            maxWidth: '400px'
          }}>
            <p style={{ margin: '0.5rem 0' }}><strong>Name:</strong> {session.user?.name}</p>
            <p style={{ margin: '0.5rem 0' }}><strong>Email:</strong> {session.user?.email}</p>
          </div>

          <button
            onClick={() => signOut()}
            style={{
              padding: '0.75rem 2rem',
              backgroundColor: '#dc3545',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              fontSize: '1rem',
              fontWeight: '500',
              cursor: 'pointer'
            }}
          >
            Sign Out
          </button>
        </>
      ) : (
        <>
          <h1 style={{ fontSize: '2.5rem' }}>Welcome to Elo App</h1>
          <p style={{ color: '#666', fontSize: '1.1rem' }}>
            Please sign in to continue
          </p>
          <button
            onClick={() => router.push('/login')}
            style={{
              padding: '0.75rem 2rem',
              backgroundColor: '#1976d2',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              fontSize: '1rem',
              fontWeight: '500',
              cursor: 'pointer'
            }}
          >
            Sign In
          </button>
        </>
      )}
    </main>
  )
}
