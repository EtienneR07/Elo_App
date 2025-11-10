'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

export function HamburgerMenu() {
  const [isOpen, setIsOpen] = useState(false);
  const router = useRouter();

  const toggleMenu = () => setIsOpen(!isOpen);

  const handleManageLeagues = () => {
    router.push('/leagues');
    setIsOpen(false);
  };

  return (
    <>
      {/* Hamburger Button */}
      <button
        onClick={toggleMenu}
        className="fixed top-6 left-6 z-50 p-2 rounded-md hover:bg-gray-200 transition-colors"
        aria-label="Menu"
      >
        <div className="w-6 h-5 flex flex-col justify-between">
          <span
            className={`block h-0.5 w-full bg-gray-800 transition-transform duration-300 ${
              isOpen ? 'rotate-45 translate-y-2' : ''
            }`}
          />
          <span
            className={`block h-0.5 w-full bg-gray-800 transition-opacity duration-300 ${
              isOpen ? 'opacity-0' : ''
            }`}
          />
          <span
            className={`block h-0.5 w-full bg-gray-800 transition-transform duration-300 ${
              isOpen ? '-rotate-45 -translate-y-2' : ''
            }`}
          />
        </div>
      </button>

      {/* Menu Overlay */}
      {isOpen && (
        <div
          className="fixed inset-0 bg-black bg-opacity-50 z-40"
          onClick={toggleMenu}
        />
      )}

      {/* Menu Panel */}
      <div
        className={`fixed top-0 left-0 h-full w-64 bg-white shadow-lg z-40 transform transition-transform duration-300 ease-in-out ${
          isOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        <nav className="mt-20 px-4">
          <button
            onClick={handleManageLeagues}
            className="w-full text-left px-4 py-3 rounded-md hover:bg-gray-100 transition-colors text-gray-800 font-medium"
          >
            Manage Leagues
          </button>
        </nav>
      </div>
    </>
  );
}
