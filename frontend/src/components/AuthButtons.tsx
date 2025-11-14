import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

export function AuthenticatedButtons() {
  const { logout } = useAuth();

  return (
    <div className="flex gap-4 flex-col items-center">
      <button
        onClick={logout}
        className="px-8 py-3 bg-red-500 text-white rounded text-base font-medium cursor-pointer hover:bg-red-600 transition-colors"
      >
        Sign Out
      </button>
    </div>
  );
}

export function UnauthenticatedButton() {
  const navigate = useNavigate();

  return (
    <button
      onClick={() => navigate('/login')}
      className="px-8 py-3 bg-blue-500 text-white rounded text-base font-medium cursor-pointer hover:bg-blue-600 transition-colors"
    >
      Sign In
    </button>
  );
}
