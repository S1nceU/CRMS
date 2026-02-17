import React, { useCallback, useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import CustomerManagement from './CustomerManagement';
import HistoryManagement from './HistoryManagement';

const Dashboard: React.FC = () => {
  const { username, logout } = useAuth();
  const [activeTab, setActiveTab] = useState<'customers' | 'history'>('customers');
  const [pendingHistoryCustomerId, setPendingHistoryCustomerId] = useState<string | null>(null);

  const handleLogout = async () => {
    await logout();
  };

  const handleCreateHistoryForCustomer = useCallback((customerId: string) => {
    setPendingHistoryCustomerId(customerId);
    setActiveTab('history');
  }, []);

  const handlePrefillHandled = useCallback(() => {
    setPendingHistoryCustomerId(null);
  }, []);

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col gap-4 py-4 sm:flex-row sm:justify-between sm:items-center sm:py-6">
            <div className="flex items-center">
              <h1 className="text-2xl font-bold text-gray-900 sm:text-3xl">CRMS 控制台</h1>
            </div>
            <div className="flex w-full items-center justify-between sm:w-auto sm:justify-start sm:space-x-4">
              <span className="text-sm text-gray-700 sm:text-base">歡迎，{username}</span>
              <button
                onClick={handleLogout}
                className="whitespace-nowrap bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md text-sm font-medium"
              >
                登出
              </button>
            </div>
          </div>
        </div>
      </header>

      {/* Navigation Tabs */}
      <nav className="bg-white border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="overflow-x-auto">
            <div className="flex min-w-max space-x-8">
              <button
                onClick={() => setActiveTab('customers')}
                className={`py-4 px-1 border-b-2 font-medium text-sm ${
                  activeTab === 'customers'
                    ? 'border-indigo-500 text-indigo-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                客戶管理
              </button>
              <button
                onClick={() => setActiveTab('history')}
                className={`py-4 px-1 border-b-2 font-medium text-sm ${
                  activeTab === 'history'
                    ? 'border-indigo-500 text-indigo-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                歷史紀錄管理
              </button>
            </div>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        {activeTab === 'customers' && (
          <CustomerManagement onCreateHistory={handleCreateHistoryForCustomer} />
        )}
        {activeTab === 'history' && (
          <HistoryManagement
            prefillCustomerId={pendingHistoryCustomerId}
            onPrefillHandled={handlePrefillHandled}
          />
        )}
      </main>
    </div>
  );
};

export default Dashboard;
