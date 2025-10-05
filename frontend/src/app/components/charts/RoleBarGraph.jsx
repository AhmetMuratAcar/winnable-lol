"use client";

import { useMemo } from "react";
import { ResponsiveContainer, BarChart, Bar, XAxis, Tooltip } from "recharts";

const ROLE_ICONS = {
  Top: "/images/roles/top.svg",
  Jungle: "/images/roles/jungle.svg",
  Mid: "/images/roles/middle.svg",
  Bot: "/images/roles/bottom.svg",
  Support: "/images/roles/utility.svg",
};

const renderCustomAxisTick = ({ x, y, payload }) => {
  const size = 20;
  const src = ROLE_ICONS[payload.value];
  return (
    <g transform={`translate(${x - size / 2}, ${y + 6})`}>
      <image href={src} xlinkHref={src} width={size} height={size} />
    </g>
  );
};

export default function RoleBarGraph({ roles }) {
  // { top:{wins,losses}, jungle:{wins,losses}, ... }
  const data = useMemo(
    () => [
      { name: "Top", wins: roles?.top?.wins ?? 0, losses: roles?.top?.losses ?? 0 },
      { name: "Jungle", wins: roles?.jungle?.wins ?? 0, losses: roles?.jungle?.losses ?? 0 },
      { name: "Mid", wins: roles?.mid?.wins ?? 0, losses: roles?.mid?.losses ?? 0 },
      { name: "Bot", wins: roles?.bot?.wins ?? 0, losses: roles?.bot?.losses ?? 0 },
      { name: "Support", wins: roles?.utility?.wins ?? 0, losses: roles?.utility?.losses ?? 0 },
    ],
    [roles],
  );

  return (
    <ResponsiveContainer width="100%" height="100%">
      <BarChart data={data} margin={{ top: 8, right: 8, bottom: 0, left: 8 }}>
        <XAxis
          dataKey="name"
          tick={renderCustomAxisTick}
          tickLine={false}
          axisLine={false}
          height={40}
        />
        <Tooltip
          cursor={false}
          offset={20}
          contentStyle={{
            background: "var(--contrast)",
            borderRadius: 8,
            borderColor: "var(--contrast-border)",
          }}
        />

        <Bar
          dataKey="wins"
          stackId="a"
          fill="var(--color-green-500)"
          barSize={20}
          background={{ fill: "var(--contrast-border)" }}
        />
        <Bar dataKey="losses" stackId="a" fill="var(--pastel-red)" barSize={20} />
      </BarChart>
    </ResponsiveContainer>
  );
}
