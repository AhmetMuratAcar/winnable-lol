"use client";

import { useMemo } from "react";
import { PieChart, Pie, Label, Tooltip, Cell } from "recharts";

export default function WinLossDonut({ wins, losses }) {
  const total = (wins || 0) + (losses || 0);

  const chartData = useMemo(
    () => [
      { name: "Wins", value: wins || 0, class: "fill-green-500" },
      { name: "Losses", value: losses || 0, class: "fill-(--pastel-red)" },
    ],
    [wins, losses],
  );

  const winRate = useMemo(() => {
    if (!total) return null;
    return Math.round(((wins || 0) / total) * 100);
  }, [wins, total]);

  return (
    <div className="flex flex-col items-center">
      <PieChart width={100} height={100}>
        <Pie
          data={chartData}
          dataKey="value"
          nameKey="name"
          innerRadius={30}
          outerRadius={40}
          paddingAngle={2}
          stroke="none"
        >
          {chartData.map((entry, index) => (
            <Cell key={index} className={entry.class} />
          ))}

          <Label
            value={winRate == null ? "—" : `${winRate}%`}
            position="center"
            className={`font-bold ${winRate != null && winRate < 50 ? "fill-red-500" : "fill-green-500"}`}
          />
        </Pie>
        {/* <Tooltip
          formatter={(value, name) => [`${value}`, name]}
          contentStyle={{
            background: "var(--contrast)",
            borderRadius: "4px",
            borderColor: "var(--contrast-border)",
          }}
          itemStyle={{ color: "#fff" }}
          position={{ x: 100, y: 20 }}
        /> */}
      </PieChart>
    </div>
  );
}
