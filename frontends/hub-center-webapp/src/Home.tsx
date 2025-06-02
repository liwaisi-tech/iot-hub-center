import Header from './components/Header';
import { Overview } from './components/Overview';

function App() {
  return (
    <div className="min-h-screen bg-gray-50 font-body">
      <Header />
      <main className="container mx-auto px-4 py-8">
        <Overview />
      </main>
    </div>
  );
}

export default App;
