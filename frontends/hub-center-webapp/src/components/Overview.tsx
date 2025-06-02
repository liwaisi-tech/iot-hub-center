import { SummaryCard } from './SummaryCard';
import { BsCheckCircle } from 'react-icons/bs';
import { FaCloudRain } from "react-icons/fa6";
import { AiOutlineWarning } from 'react-icons/ai';
import { TbTemperature } from 'react-icons/tb';

export const Overview = () => {
  return (
    <div className="space-y-6">
      <h2 className="text-2xl font-semibold text-gray-900">Resumen General</h2>
      
      <div className="grid grid-cols-2 md:grid-cols-2 lg:grid-cols-4 gap-4">
        {[
          {
            title: "Zonas Activas",
            value: "4",
            icon: <BsCheckCircle />,
            iconColor: "text-green-500"
          },
          {
            title: "Riego Activo", 
            value: "1",
            icon: <FaCloudRain />,
            iconColor: "text-blue-500",
            valueColor: "text-blue-500"
          },
          {
            title: "Alertas",
            value: "3", 
            icon: <AiOutlineWarning />,
            iconColor: "text-red-500",
            valueColor: "text-red-500"
          },
          {
            title: "Temp. Promedio",
            value: "31°C",
            icon: <TbTemperature />,
            iconColor: "text-orange-500",
            valueColor: "text-orange-500"
          }
        ].map((card, index) => (
          <SummaryCard
            key={index}
            {...card}
          />
        ))}
      </div>
    </div>
  );
};
