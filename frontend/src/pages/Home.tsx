import { AuthenticatedButtons, UnauthenticatedButton } from '../components/AuthButtons';
import { HamburgerMenu } from '../components/HamburgerMenu';
import { useAuth } from '../contexts/AuthContext';

export default function Home() {
  const { user, isAuthenticated } = useAuth();

  return (
    <>
      {isAuthenticated && <HamburgerMenu />}

      <main className="flex min-h-screen items-center justify-center flex-col gap-8 font-sans p-8">
        {isAuthenticated ? (
          <>
            <div className="text-center">
              <h1 className="text-4xl mb-4">
                Welcome, {user?.name}!
              </h1>
              <p className="text-gray-600 text-lg">
                You are successfully logged in
              </p>
            </div>

            <div className="p-6 bg-gray-100 rounded-lg max-w-md">
              <p className="my-2"><strong>Name:</strong> {user?.name}</p>
              <p className="my-2"><strong>Email:</strong> {user?.email}</p>
            </div>

            <AuthenticatedButtons />
          </>
        ) : (
          <>
            <h1 className="text-4xl">Welcome to Elofy</h1>
            <p className="text-gray-600 text-lg">
              Please sign in to continue
            </p>

            <UnauthenticatedButton />
          </>
        )}
      </main>
    </>
  );
}
