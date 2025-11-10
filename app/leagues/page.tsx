"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";

interface League {
  id: string;
  name: string;
  description?: string;
  playersPerTeam: number;
  discipline?: string;
  ownerId: string;
  createdAt: string;
  updatedAt: string;
}

export default function LeaguesPage() {
  const router = useRouter();
  const { data: session, status } = useSession();
  const [leagues, setLeagues] = useState<League[]>([]);
  const [loading, setLoading] = useState(true);
  const [formData, setFormData] = useState({
    name: "",
    description: "",
    ownerId: session?.user?.id || "",
    playersPerTeam: 1,
    discipline: "",
  });
  const [editingId, setEditingId] = useState<string | null>(null);

  useEffect(() => {
    if (session?.user?.id) {
      setFormData(prev => ({ ...prev, ownerId: session.user.id }));
    }
  }, [session]);

  const fetchLeagues = async () => {
    setLoading(true);
    try {
      const response = await fetch("/api/leagues");
      const data = await response.json();
      setLeagues(data);
    } catch (error) {
      console.error("Error fetching leagues:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchLeagues();
  }, []);

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await fetch("/api/leagues", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        await fetchLeagues();
        setFormData({
          name: "",
          description: "",
          ownerId: session?.user?.id || "",
          playersPerTeam: 1,
          discipline: "",
        });
      }
    } catch (error) {
      console.error("Error creating league:", error);
    }
  };

  const handleUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingId) return;

    try {
      const response = await fetch(`/api/leagues/${editingId}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        await fetchLeagues();
        setEditingId(null);
        setFormData({
          name: "",
          description: "",
          ownerId: session?.user?.id || "",
          playersPerTeam: 1,
          discipline: "",
        });
      }
    } catch (error) {
      console.error("Error updating league:", error);
    }
  };

  // Delete league
  const handleDelete = async (id: string) => {
    if (!confirm("Are you sure you want to delete this league?")) return;

    try {
      const response = await fetch(`/api/leagues/${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        await fetchLeagues();
      }
    } catch (error) {
      console.error("Error deleting league:", error);
    }
  };

  // Start editing
  const startEdit = (league: League) => {
    setEditingId(league.id);
    setFormData({
      name: league.name,
      description: league.description || "",
      ownerId: league.ownerId,
      playersPerTeam: league.playersPerTeam,
      discipline: league.discipline || "",
    });
  };

  // Redirect if not authenticated
  if (status === "unauthenticated") {
    router.push("/");
    return null;
  }

  // Show loading while checking auth
  if (status === "loading") {
    return (
      <div className="container mx-auto p-8">
        <p>Loading...</p>
      </div>
    );
  }

  return (
    <div className="container mx-auto p-8">
      <button
        onClick={() => router.push("/")}
        className="mb-4 px-4 py-2 bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300"
      >
        ← Back to Home
      </button>
      <h1 className="text-3xl font-bold mb-8">League CRUD Demo</h1>

      {/* Form */}
      <div className="bg-white p-6 rounded-lg shadow-md mb-8">
        <h2 className="text-xl font-semibold mb-4">
          {editingId ? "Edit League" : "Create New League"}
        </h2>
        <form onSubmit={editingId ? handleUpdate : handleCreate} className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-1">Name *</label>
            <input
              type="text"
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
              required
              className="w-full px-3 py-2 border rounded-md"
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">Description</label>
            <textarea
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
              className="w-full px-3 py-2 border rounded-md"
              rows={3}
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">Discipline</label>
            <input
              type="text"
              value={formData.discipline}
              onChange={(e) => setFormData({ ...formData, discipline: e.target.value })}
              placeholder="e.g., Ping Pong, Foosball"
              className="w-full px-3 py-2 border rounded-md"
            />
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">Players Per Team</label>
            <input
              type="number"
              value={formData.playersPerTeam}
              onChange={(e) =>
                setFormData({ ...formData, playersPerTeam: parseInt(e.target.value) })
              }
              min={1}
              className="w-full px-3 py-2 border rounded-md"
            />
          </div>

          <div className="flex gap-2">
            <button
              type="submit"
              className="px-4 py-2 bg-primary-500 text-white rounded-md hover:bg-primary-600"
            >
              {editingId ? "Update League" : "Create League"}
            </button>
            {editingId && (
              <button
                type="button"
                onClick={() => {
                  setEditingId(null);
                  setFormData({
                    name: "",
                    description: "",
                    ownerId: session?.user?.id || "",
                    playersPerTeam: 1,
                    discipline: "",
                  });
                }}
                className="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400"
              >
                Cancel
              </button>
            )}
          </div>
        </form>
      </div>

      <div className="bg-white rounded-lg shadow-md">
        <h2 className="text-xl font-semibold p-6 border-b">All Leagues</h2>
        {loading ? (
          <p className="p-6">Loading...</p>
        ) : leagues.length === 0 ? (
          <p className="p-6 text-gray-500">No leagues yet. Create one above!</p>
        ) : (
          <div className="divide-y">
            {leagues.map((league) => (
              <div key={league.id} className="p-6 hover:bg-gray-50">
                <div className="flex justify-between items-start">
                  <div className="flex-1">
                    <h3 className="text-lg font-semibold">{league.name}</h3>
                    {league.description && (
                      <p className="text-gray-600 mt-1">{league.description}</p>
                    )}
                    <div className="flex gap-4 mt-2 text-sm text-gray-500">
                      <span>Discipline: {league.discipline || "N/A"}</span>
                      <span>Players/Team: {league.playersPerTeam}</span>
                    </div>
                    <p className="text-xs text-gray-400 mt-2">ID: {league.id}</p>
                  </div>
                  <div className="flex gap-2">
                    <button
                      onClick={() => startEdit(league)}
                      className="px-3 py-1 bg-primary-500 text-white rounded hover:bg-primary-600"
                    >
                      Edit
                    </button>
                    <button
                      onClick={() => handleDelete(league.id)}
                      className="px-3 py-1 bg-error text-white rounded hover:bg-error-dark"
                    >
                      Delete
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
