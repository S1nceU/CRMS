import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { apiService } from '../services/api';

interface AuthContextType {
  isAuthenticated: boolean;
  username: string | null;
  login: (credentials: { Username: string; Password: string }) => Promise<boolean>;
  logout: () => Promise<void>;
  loading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [username, setUsername] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  const clearAuth = () => {
    localStorage.removeItem('token');
    setIsAuthenticated(false);
    setUsername(null);
  };

  useEffect(() => {
    const handleUnauthorized = () => {
      clearAuth();
    };
    window.addEventListener('crms:unauthorized', handleUnauthorized);
    return () => {
      window.removeEventListener('crms:unauthorized', handleUnauthorized);
    };
  }, []);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      setLoading(false);
      return;
    }

    apiService
      .authenticate(token)
      .then((response) => {
        if (response.Message === 'Authentication successfully') {
          setIsAuthenticated(true);
          const name = typeof response.data === 'string' ? response.data : 'user';
          setUsername(name);
          return;
        }
        clearAuth();
      })
      .catch((err) => {
        console.error('Auth check failed:', err);
        clearAuth();
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  const login = async (credentials: { Username: string; Password: string }): Promise<boolean> => {
    try {
      const response = await apiService.login(credentials);
      if (response.token) {
        localStorage.setItem('token', response.token);
        setIsAuthenticated(true);
        const name = typeof response.data === 'string' ? response.data : credentials.Username;
        setUsername(name);
        return true;
      }
      clearAuth();
      return false;
    } catch (error) {
      console.error('Login failed:', error);
      clearAuth();
      return false;
    }
  };

  const logout = async (): Promise<void> => {
    try {
      await apiService.logout();
    } catch (error) {
      console.error('Logout failed:', error);
    } finally {
      clearAuth();
    }
  };

  const value: AuthContextType = {
    isAuthenticated,
    username,
    login,
    logout,
    loading,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
