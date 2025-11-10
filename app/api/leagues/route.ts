import { NextRequest, NextResponse } from "next/server";
import { createLeague, getAllLeagues } from "@/lib/services/leagueService";
import { requireAuth } from "@/lib/auth";

export async function GET() {
  try {
    const leagues = await getAllLeagues();
    return NextResponse.json(leagues);
  } catch (error) {
    console.error("Error fetching leagues:", error);
    return NextResponse.json(
      { error: "Failed to fetch leagues" },
      { status: 500 }
    );
  }
}

export async function POST(request: NextRequest) {
  try {
    const user = await requireAuth();

    const body = await request.json();

    if (!body.name) {
      return NextResponse.json(
        { error: "Name is required" },
        { status: 400 }
      );
    }

    const league = await createLeague({
      name: body.name,
      description: body.description,
      ownerId: user.id,
      playersPerTeam: body.playersPerTeam,
      discipline: body.discipline,
    });

    return NextResponse.json(league, { status: 201 });
  } catch (error) {
    if (error instanceof Error && error.message === "Unauthorized") {
      return NextResponse.json(
        { error: "Unauthorized" },
        { status: 401 }
      );
    }
    console.error("Error creating league:", error);
    return NextResponse.json(
      { error: "Failed to create league" },
      { status: 500 }
    );
  }
}
