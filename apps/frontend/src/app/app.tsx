import React from 'react';
import { AuthProvider, useAuth } from '../contexts/AuthContext';
import LoginPage from '../components/LoginPage';
import Dashboard from '../components/Dashboard';

const AppContent: React.FC = () => {
  const { isAuthenticated, loading } = useAuth();

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-50">
        <div className="text-xl text-gray-600">Loading...</div>
      </div>
    );
  }

  return isAuthenticated ? <Dashboard /> : <LoginPage />;
};

export function App() {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
}

export default App;
