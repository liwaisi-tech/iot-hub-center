import type { ReactNode } from 'react';

interface SummaryCardProps {
  title: string;
  value: string | number;
  icon: ReactNode;
  iconColor?: string;
  valueColor?: string;
}

export const SummaryCard = ({ title, value, icon, iconColor = 'text-green-500', valueColor = 'text-black' }: SummaryCardProps) => {
  return (
    <div className="bg-white rounded-lg p-4 shadow-lg border border-gray-300">
      <div className="flex flex-col space-y-2">
        <div className="text-gray-600 text-sm">{title}</div>
        <div className="flex items-center justify-between">
          <span className={`text-4xl font-semibold ${valueColor}`}>{value}</span>
          <div className={`${iconColor} text-5xl`}>{icon}</div>
        </div>
      </div>
    </div>
  );
}; 