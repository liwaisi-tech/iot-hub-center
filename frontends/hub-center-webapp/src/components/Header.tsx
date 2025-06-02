import { WiDayFog } from "react-icons/wi";
import { GrWifi } from "react-icons/gr";
import { PiBellRingingFill } from "react-icons/pi";
import { HiOutlineSparkles } from "react-icons/hi";
import { VscSettingsGear } from "react-icons/vsc";

const Header = () => {
  return (
    <header className="flex items-center justify-between px-2 sm:px-6 py-1 bg-white shadow-sm shadow-emerald-200">
      {/* Logo and Title */}
      <div className="flex items-center gap-2">
        <WiDayFog className="text-6xl sm:text-8xl text-blue-700" />
        <div>
          <h1 className="text-xl sm:text-4xl font-title font-bold">Liwaisi IoT</h1>
          <p className="hidden sm:block text-sm text-gray-600 font-subtitle">Sistema de cultivos y riegos inteligentes</p>
        </div>
      </div>

      {/* Right side items */}
      <div className="flex items-center gap-3 sm:gap-6 font-body">
        {/* Online Status */}
        <div className="flex items-center gap-2">
          <GrWifi className="text-2xl text-green-600" />
          <span className="hidden sm:inline text-green-600">En línea</span>
        </div>

        {/* Notifications */}
        <div className="relative flex items-center gap-2 border border-gray-200 rounded-xl p-2">
          <PiBellRingingFill className="text-2xl text-gray-800" />
          <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center">
            3
          </span>
        </div>

        {/* Tips */}
        <div className="flex items-center gap-2 border border-gray-200 rounded-xl p-2">
          <HiOutlineSparkles className="text-2xl text-gray-800" />
          <span className="hidden sm:inline">Consejos</span>
        </div>

        {/* Settings */}
        <VscSettingsGear className="text-2xl text-gray-800" />
      </div>
    </header>
  );
};

export default Header; 