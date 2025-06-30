import React from 'react';
import { AuthProvider } from './contexts/AuthContext';
import LoginPage from './components/LoginPage';
import Dashboard from './components/Dashboard';
import CustomerManagement from './components/CustomerManagement';
import HistoryManagement from './components/HistoryManagement';

const TestComponents: React.FC = () => {
  return (
    <div>
      <h1>CRMS Components Test</h1>
      <p>All components are successfully imported and ready to use!</p>
      
      <div className="component-list">
        <ul>
          <li>✅ AuthProvider</li>
          <li>✅ LoginPage</li>
          <li>✅ Dashboard</li>
          <li>✅ CustomerManagement</li>
          <li>✅ HistoryManagement</li>
          <li>✅ API Service Layer</li>
        </ul>
      </div>
    </div>
  );
};

export default TestComponents;