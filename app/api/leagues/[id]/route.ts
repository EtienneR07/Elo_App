import { NextRequest, NextResponse } from "next/server";
import { getLeague, updateLeague, deleteLeague } from "@/lib/services/leagueService";

export async function GET(
  request: NextRequest,
  { params }: { params: { id: string } }
) {
  try {
    const league = await getLeague(params.id);

    if (!league) {
      return NextResponse.json(
        { error: "League not found" },
        { status: 404 }
      );
    }

    return NextResponse.json(league);
  } catch (error) {
    console.error("Error fetching league:", error);
    return NextResponse.json(
      { error: "Failed to fetch league" },
      { status: 500 }
    );
  }
}

export async function PUT(
  request: NextRequest,
  { params }: { params: { id: string } }
) {
  try {
    const body = await request.json();

    const league = await updateLeague(params.id, {
      name: body.name,
      description: body.description,
      playersPerTeam: body.playersPerTeam,
      discipline: body.discipline,
    });

    return NextResponse.json(league);
  } catch (error) {
    console.error("Error updating league:", error);
    return NextResponse.json(
      { error: "Failed to update league" },
      { status: 500 }
    );
  }
}

export async function DELETE(
  request: NextRequest,
  { params }: { params: { id: string } }
) {
  try {
    await deleteLeague(params.id);
    return NextResponse.json({ message: "League deleted successfully" });
  } catch (error) {
    console.error("Error deleting league:", error);
    return NextResponse.json(
      { error: "Failed to delete league" },
      { status: 500 }
    );
  }
}
