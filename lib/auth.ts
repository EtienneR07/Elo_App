import { getServerSession } from "next-auth";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";

/**
 * Get the current user session on the server-side
 * Use this in API routes and Server Components
 */
export async function getCurrentUser() {
  const session = await getServerSession(authOptions);
  return session?.user;
}

/**
 * Get the current user ID on the server-side
 * Throws an error if user is not authenticated
 */
export async function requireAuth() {
  const session = await getServerSession(authOptions);

  console.log("Session in requireAuth:", JSON.stringify(session, null, 2));

  const user = session?.user;

  if (!user || !user.id) {
    console.log("Auth failed - user:", user);
    throw new Error("Unauthorized");
  }

  return user;
}
