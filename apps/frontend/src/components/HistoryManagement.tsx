import React, { useState, useEffect } from 'react';
import { apiService, History, HistoryRequest, Customer } from '../services/api';

const HistoryManagement: React.FC = () => {
  const [histories, setHistories] = useState<History[]>([]);
  const [customers, setCustomers] = useState<Customer[]>([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingHistory, setEditingHistory] = useState<History | null>(null);
  const [searchType, setSearchType] = useState<'all' | 'date' | 'dateRange' | 'customer'>('all');
  const [searchDate, setSearchDate] = useState('');
  const [searchStartDate, setSearchStartDate] = useState('');
  const [searchEndDate, setSearchEndDate] = useState('');
  const [selectedCustomerId, setSelectedCustomerId] = useState('');
  const [searchError, setSearchError] = useState<string | null>(null);
  const [isSearching, setIsSearching] = useState(false);

  const [formData, setFormData] = useState<HistoryRequest>({
    CustomerId: '',
    Date: '',
    NumberOfPeople: 1,
    Price: 0,
    Room: '',
    Note: '',
  });

  useEffect(() => {
    loadHistories();
    loadCustomers();
  }, []);

  const loadHistories = async () => {
    try {
      const response = await apiService.getHistories();
      if (response.data) {
        const data: any = response.data as any;
        const list = Array.isArray(data) ? data : [data];
        const valid = list.filter((h: any) => h && typeof h === 'object' && 'Id' in h && 'CustomerId' in h);
        setHistories(valid.map((h: any) => ({ ...h, Date: h.Date || '' })) as History[]);
      } else {
        setHistories([]);
      }
    } catch (error) {
      console.error('Failed to load histories:', error);
    } finally {
      setLoading(false);
    }
  };

  const loadCustomers = async () => {
    try {
      const response = await apiService.getCustomers();
      if (response.data) {
        const data: any = response.data as any;
        const list = Array.isArray(data) ? data : [data];
        const valid = list.filter((c: any) => c && typeof c === 'object' && 'Id' in c);
        setCustomers(valid as Customer[]);
      } else {
        setCustomers([]);
      }
    } catch (error) {
      console.error('Failed to load customers:', error);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (editingHistory) {
        await apiService.updateHistory({ ...formData, HistoryId: editingHistory.Id });
      } else {
        await apiService.createHistory(formData);
      }
      resetForm();
      loadHistories();
    } catch (error) {
      console.error('Failed to save history:', error);
    }
  };

  const handleEdit = (history: History) => {
    setEditingHistory(history);
    setFormData({
      HistoryId: history.Id,
      CustomerId: history.CustomerId,
      Date: (history.Date || '').split('T')[0],
      NumberOfPeople: history.NumberOfPeople,
      Price: history.Price,
      Room: history.Room,
      Note: history.Note,
    });
    setShowForm(true);
  };

  const handleDelete = async (historyId: string) => {
    if (window.confirm('確定要刪除此歷史紀錄嗎？')) {
      try {
        await apiService.deleteHistory(historyId);
        loadHistories();
      } catch (error) {
        console.error('Failed to delete history:', error);
      }
    }
  };

  const handleSearch = async () => {
    try {
      setSearchError(null);
      setIsSearching(true);
      let response;
      switch (searchType) {
        case 'date': {
          if (!searchDate) {
            setSearchError('請選擇日期。');
            return;
          }
          response = await apiService.getHistoriesByDate(searchDate);
          break;
        }
        case 'dateRange': {
          if (!searchStartDate || !searchEndDate) {
            setSearchError('請選擇起訖日期。');
            return;
          }
          if (searchStartDate > searchEndDate) {
            setSearchError('開始日期必須早於結束日期。');
            return;
          }
          response = await apiService.getHistoriesByDateRange(searchStartDate, searchEndDate);
          break;
        }
        case 'customer': {
          if (!selectedCustomerId) {
            setSearchError('請選擇客戶。');
            return;
          }
          response = await apiService.getHistoriesByCustomerId(selectedCustomerId);
          break;
        }
        default:
          await loadHistories();
          return;
      }

      if (response?.data) {
        const data: any = response.data as any;
        const list = Array.isArray(data) ? data : [data];
        const valid = list.filter((h: any) => h && typeof h === 'object' && 'Id' in h && 'CustomerId' in h);
        setHistories(valid.map((h: any) => ({ ...h, Date: h.Date || '' })) as History[]);
      } else {
        setHistories([]);
      }
    } catch (error) {
      console.error('Search failed:', error);
      setSearchError('搜尋失敗，請稍後再試。');
      setHistories([]);
    } finally {
      setIsSearching(false);
    }
  };

  const resetForm = () => {
    setFormData({
      CustomerId: '',
      Date: '',
      NumberOfPeople: 1,
      Price: 0,
      Room: '',
      Note: '',
    });
    setEditingHistory(null);
    setShowForm(false);
  };

  const getCustomerName = (customerId: string) => {
    const customer = customers.find(c => c.Id === customerId);
    return customer ? customer.Name : '未知客戶';
  };

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('zh-TW', {
      style: 'currency',
      currency: 'TWD',
    }).format(amount);
  };

  if (loading) {
    return <div className="flex justify-center items-center h-64">載入中...</div>;
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold text-gray-900">歷史紀錄管理</h2>
        <button
          onClick={() => setShowForm(true)}
          className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-md text-sm font-medium"
        >
          新增紀錄
        </button>
      </div>

      {/* Search */}
      <div className="bg-white p-4 rounded-lg shadow">
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              搜尋方式
            </label>
            <select
              value={searchType}
              onChange={(e) => setSearchType(e.target.value as 'all' | 'date' | 'dateRange' | 'customer')}
              className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            >
              <option value="all">全部紀錄</option>
              <option value="date">依日期</option>
              <option value="dateRange">依日期區間</option>
              <option value="customer">依客戶</option>
            </select>
          </div>

          {searchType === 'date' && (
            <div className="flex gap-4 items-end">
              <div className="flex-1">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  日期
                </label>
                <input
                  type="date"
                  value={searchDate}
                  onChange={(e) => setSearchDate(e.target.value)}
                  onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); handleSearch(); } }}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>
              <button
                onClick={handleSearch}
                className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md text-sm font-medium"
              >
                {isSearching ? '搜尋中...' : '搜尋'}
              </button>
            </div>
          )}

          {searchType === 'dateRange' && (
            <div className="flex gap-4 items-end">
              <div className="flex-1">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  開始日期
                </label>
                <input
                  type="date"
                  value={searchStartDate}
                  onChange={(e) => setSearchStartDate(e.target.value)}
                  onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); handleSearch(); } }}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>
              <div className="flex-1">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  結束日期
                </label>
                <input
                  type="date"
                  value={searchEndDate}
                  onChange={(e) => setSearchEndDate(e.target.value)}
                  onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); handleSearch(); } }}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                />
              </div>
              <button
                onClick={handleSearch}
                className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md text-sm font-medium"
              >
                {isSearching ? '搜尋中...' : '搜尋'}
              </button>
            </div>
          )}

          {searchType === 'customer' && (
            <div className="flex gap-4 items-end">
              <div className="flex-1">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  客戶
                </label>
                <select
                  value={selectedCustomerId}
                  onChange={(e) => setSelectedCustomerId(e.target.value)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                >
                  <option value="">請選擇客戶...</option>
                  {customers.map((customer) => (
                    <option key={customer.Id} value={customer.Id}>
                      {customer.Name}
                    </option>
                  ))}
                </select>
              </div>
              <button
                onClick={handleSearch}
                className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md text-sm font-medium"
              >
                {isSearching ? '搜尋中...' : '搜尋'}
              </button>
            </div>
          )}

          {searchError && (
            <p className="text-sm text-red-600">{searchError}</p>
          )}

          <button
            onClick={() => { setSearchType('all'); setSearchDate(''); setSearchStartDate(''); setSearchEndDate(''); setSelectedCustomerId(''); setSearchError(null); loadHistories(); }}
            className="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-md text-sm font-medium"
          >
            重設
          </button>
        </div>
      </div>

      {/* History Form Modal */}
      {showForm && (
        <div className="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
          <div className="relative top-20 mx-auto p-5 border w-11/12 max-w-lg shadow-lg rounded-md bg-white">
            <div className="mt-3">
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                {editingHistory ? '編輯紀錄' : '新增紀錄'}
              </h3>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    客戶 *
                  </label>
                  <select
                    required
                    value={formData.CustomerId}
                    onChange={(e) => setFormData({ ...formData, CustomerId: e.target.value })}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  >
                    <option value="">請選擇客戶...</option>
                    {customers.map((customer) => (
                      <option key={customer.Id} value={customer.Id}>
                        {customer.Name}
                      </option>
                    ))}
                  </select>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    日期 *
                  </label>
                  <input
                    type="date"
                    required
                    value={formData.Date}
                    onChange={(e) => setFormData({ ...formData, Date: e.target.value })}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    人數 *
                  </label>
                  <input
                    type="number"
                    required
                    min="1"
                    value={formData.NumberOfPeople}
                    onChange={(e) => setFormData({ ...formData, NumberOfPeople: parseInt(e.target.value) })}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    金額 *
                  </label>
                  <input
                    type="number"
                    required
                    min="0"
                    step="0.01"
                    value={formData.Price}
                    onChange={(e) => setFormData({ ...formData, Price: parseFloat(e.target.value) })}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    房號 *
                  </label>
                  <input
                    type="text"
                    required
                    value={formData.Room}
                    onChange={(e) => setFormData({ ...formData, Room: e.target.value })}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    備註
                  </label>
                  <textarea
                    value={formData.Note}
                    onChange={(e) => setFormData({ ...formData, Note: e.target.value })}
                    rows={3}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                  />
                </div>
                <div className="flex justify-end space-x-3 pt-4">
                  <button
                    type="button"
                    onClick={resetForm}
                    className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300"
                  >
                    取消
                  </button>
                  <button
                    type="submit"
                    className="px-4 py-2 text-sm font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700"
                  >
                    {editingHistory ? '更新' : '建立'}
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}

      {/* History List */}
      <div className="bg-white shadow overflow-hidden sm:rounded-md">
        <ul className="divide-y divide-gray-200">
          {histories.map((history) => (
            <li key={history.Id} className="px-6 py-4">
              <div className="flex items-center justify-between">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center space-x-4">
                    <div className="flex-1">
                      <p className="text-sm font-medium text-gray-900">
                        {getCustomerName(history.CustomerId)}
                      </p>
                      <p className="text-sm text-gray-500">
                        {(history.Date ? history.Date.split('T')[0] : '')}
                      </p>
                    </div>
                    <div className="flex-1">
                      <p className="text-sm text-gray-900">
                        房號：{history.Room}
                      </p>
                      <p className="text-sm text-gray-500">
                        人數：{history.NumberOfPeople}
                      </p>
                    </div>
                    <div className="flex-1">
                      <p className="text-sm font-medium text-green-600">
                        {formatCurrency(history.Price)}
                      </p>
                    </div>
                  </div>
                  {history.Note && (
                    <p className="mt-2 text-xs text-gray-500">
                      備註：{history.Note}
                    </p>
                  )}
                </div>
                <div className="flex space-x-2">
                  <button
                    onClick={() => handleEdit(history)}
                    className="text-indigo-600 hover:text-indigo-900 text-sm font-medium"
                  >
                    編輯
                  </button>
                  <button
                    onClick={() => handleDelete(history.Id)}
                    className="text-red-600 hover:text-red-900 text-sm font-medium"
                  >
                    刪除
                  </button>
                </div>
              </div>
            </li>
          ))}
        </ul>
        {histories.length === 0 && (
          <div className="text-center py-12">
            <p className="text-gray-500">查無歷史紀錄</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default HistoryManagement;
