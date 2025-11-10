import { prisma } from "@/lib/prisma";

export interface CreateLeagueInput {
  name: string;
  description?: string;
  ownerId: string;
  playersPerTeam?: number;
  discipline?: string;
}

export interface UpdateLeagueInput {
  name?: string;
  description?: string;
  playersPerTeam?: number;
  discipline?: string;
}

export async function createLeague(data: CreateLeagueInput) {
  return prisma.league.create({
      data: {
          name: data.name,
          description: data.description,
          ownerId: data.ownerId,
          playersPerTeam: data.playersPerTeam ?? 1,
          discipline: data.discipline,
      },
  });
}

export async function getLeague(id: string) {
  return prisma.league.findUnique({
      where: {id},
      include: {
          owner: true,
          seasons: true,
          players: true,
      },
  });
}

export async function getAllLeagues() {
  return prisma.league.findMany({
      include: {
          owner: true,
          _count: {
              select: {
                  seasons: true,
                  players: true,
              },
          },
      },
  });
}

export async function updateLeague(id: string, data: UpdateLeagueInput) {
  return prisma.league.update({
      where: {id},
      data,
  });
}

export async function deleteLeague(id: string) {
  return prisma.league.delete({
      where: {id},
  });
}
