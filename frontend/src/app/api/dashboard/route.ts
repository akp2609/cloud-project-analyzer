import { NextResponse } from "next/server";
import pool from "@/lib/db";

export async function GET(req: Request) {
  try {
    const { searchParams } = new URL(req.url);
    const projectId =
      searchParams.get("project_id") || "test-project";

    const result = await pool.query(
      `
      SELECT COALESCE(SUM(cost), 0) AS total_cost
      FROM cost_daily
      WHERE project_id = $1
      AND usage_date >= CURRENT_DATE - INTERVAL '30 days'
      `,
      [projectId]
    );

    return NextResponse.json({
      ok: true,
      total_cost_30d: result.rows[0].total_cost,
    });
  } catch (err) {
    console.error(err);
    return NextResponse.json(
      { ok: false, error: "Dashboard API failed" },
      { status: 500 }
    );
  }
}
