const Header = () => {
  return (
    <header className="flex items-center justify-between px-6 py-4 bg-white shadow-sm shadow-emerald-200">
      {/* Logo and Title */}
      <div className="flex items-center gap-2">
        <span className="text-2xl">🌱</span>
        <div>
          <h1 className="text-xl font-title font-extrabold">AgroLlano</h1>
          <p className="text-sm text-gray-600 font-subtitle">Sistema de Riego Inteligente</p>
        </div>
      </div>

      {/* Right side items */}
      <div className="flex items-center gap-6 font-body">
        {/* Online Status */}
        <div className="flex items-center gap-2">
          <span className="text-lg">📶</span>
          <span className="text-green-500">En línea</span>
        </div>

        {/* Notifications */}
        <div className="relative">
          <span className="text-xl">🔔</span>
          <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center">
            3
          </span>
        </div>

        {/* Tips */}
        <div className="flex items-center gap-2">
          <span className="text-xl">💡</span>
          <span>Consejos</span>
        </div>

        {/* Settings */}
        <span className="text-xl">⚙️</span>
      </div>
    </header>
  );
};

export default Header; 