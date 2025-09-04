import React, { useState, useEffect } from 'react';
import { apiService, Customer, CustomerRequest, Citizenship } from '../services/api';

const CustomerManagement: React.FC = () => {
  const [customers, setCustomers] = useState<Customer[]>([]);
  const [citizenships, setCitizenships] = useState<Citizenship[]>([]);
  const [loading, setLoading] = useState(true);
  const [citizenshipsLoading, setCitizenshipsLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editingCustomer, setEditingCustomer] = useState<Customer | null>(null);
  const [searchTerm, setSearchTerm] = useState('');
  const [searchType, setSearchType] = useState<'name' | 'nationalId' | 'phone'>('name');
  const [searchError, setSearchError] = useState<string | null>(null);
  const [isSearching, setIsSearching] = useState(false);

  const [formData, setFormData] = useState<CustomerRequest>({
    Name: '',
    Gender: 'Male',
    Birthday: '',
    NationalId: '',
    Address: '',
    PhoneNumber: '',
    CarNumber: '',
    CitizenshipId: 0, // Will be set to first available citizenship
    Note: '',
  });

  const [formErrors, setFormErrors] = useState<{[key: string]: string}>({});

  useEffect(() => {
    loadCustomers();
    loadCitizenships();
  }, []);

  // Debounced search for better UX
  useEffect(() => {
    const term = searchTerm.trim();
    setSearchError(null);
    if (!term) {
      // Reset to full list when cleared
      loadCustomers();
      return;
    }
    // Only auto-search when user typed at least 2 chars
    if (term.length < 2) return;
    const id = setTimeout(() => {
      handleSearch();
    }, 350);
    return () => clearTimeout(id);
  }, [searchTerm, searchType]);

  // When citizenships load and form is open for a new customer, default to Taiwan (TWN)
  useEffect(() => {
    if (showForm && !editingCustomer && (!formData.CitizenshipId || formData.CitizenshipId <= 0) && citizenships.length > 0) {
      const tw = citizenships.find((c) => (c.Alpha3 || '').toUpperCase() === 'TWN');
      setFormData((prev) => ({
        ...prev,
        CitizenshipId: tw ? tw.Id : citizenships[0].Id,
      }));
    }
  }, [citizenships, showForm, editingCustomer]);

  const validateForm = (): boolean => {
    const errors: {[key: string]: string} = {};

    // Name validation (required, string, not empty)
    if (!formData.Name || formData.Name.trim() === '') {
      errors.Name = 'Name is required';
    } else if (formData.Name.length > 100) {
      errors.Name = 'Name must be less than 100 characters';
    }

    // Gender validation (required, must be "Male" or "Female")
    if (!formData.Gender) {
      errors.Gender = 'Gender is required';
    } else if (formData.Gender !== 'Male' && formData.Gender !== 'Female') {
      errors.Gender = 'Gender must be "Male" or "Female"';
    }

    // Birthday validation (required, valid date, not future date)
    if (!formData.Birthday) {
      errors.Birthday = 'Birthday is required';
    } else {
      const birthday = new Date(formData.Birthday);
      const today = new Date();
      if (isNaN(birthday.getTime())) {
        errors.Birthday = 'Invalid date format';
      } else if (birthday > today) {
        errors.Birthday = 'Birthday cannot be in the future';
      } else if (birthday.getFullYear() < 1900) {
        errors.Birthday = 'Birthday year must be after 1900';
      }
    }

    // National ID validation (required, string, not empty, max 100 chars)
    if (!formData.NationalId || formData.NationalId.trim() === '') {
      errors.NationalId = 'National ID is required';
    } else if (formData.NationalId.length > 100) {
      errors.NationalId = 'National ID must be less than 100 characters';
    } else if (!/^[A-Za-z0-9\-_]+$/.test(formData.NationalId)) {
      errors.NationalId = 'National ID can only contain letters, numbers, hyphens, and underscores';
    }

    // Phone Number validation (optional, but if provided, should be valid format)
    if (formData.PhoneNumber && formData.PhoneNumber.trim() !== '') {
      const phoneRegex = /^[\+]?[0-9\-\(\)\s]{7,20}$/;
      if (!phoneRegex.test(formData.PhoneNumber)) {
        errors.PhoneNumber = 'Invalid phone number format';
      }
    }

    // Car Number validation (optional, alphanumeric with hyphens)
    if (formData.CarNumber && formData.CarNumber.trim() !== '') {
      const carRegex = /^[A-Za-z0-9\-]{1,20}$/;
      if (!carRegex.test(formData.CarNumber)) {
        errors.CarNumber = 'Car number can only contain letters, numbers, and hyphens (max 20 chars)';
      }
    }

    // Citizenship validation (required, must be > 0, must exist in list)
    if (!formData.CitizenshipId || formData.CitizenshipId <= 0) {
      errors.CitizenshipId = 'Please select a valid citizenship';
    } else if (!citizenships.find(c => c.Id === formData.CitizenshipId)) {
      errors.CitizenshipId = 'Selected citizenship does not exist';
    }

    // Address validation (optional, max length)
    if (formData.Address && formData.Address.length > 500) {
      errors.Address = 'Address must be less than 500 characters';
    }

    // Note validation (optional, max length)
    if (formData.Note && formData.Note.length > 1000) {
      errors.Note = 'Note must be less than 1000 characters';
    }

    setFormErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const loadCustomers = async () => {
    try {
      console.log('👥 Starting to load customers...');
      const response = await apiService.getCustomers();
      
      // 詳細檢查回應內容
      console.log('📡 Customer API response:', response);
      console.log('📊 Response data:', response.data);
      console.log('📝 Response message:', response.Message);
      console.log('🔍 Data type:', typeof response.data);
      console.log('📏 Is array?', Array.isArray(response.data));
      
      if (response.data && Array.isArray(response.data)) {
        setCustomers(response.data);
        console.log(`✅ Successfully loaded ${response.data.length} customers!`);
        if (response.data.length === 0) {
          console.log('📝 Database is empty - no customers yet');
        }
      } else {
        console.error('❌ Invalid customer data format:', response);
        setCustomers([]);
      }
    } catch (error) {
      console.error('💥 Failed to load customers:', error);
      setCustomers([]);
    } finally {
      setLoading(false);
    }
  };

  const loadCitizenships = async () => {
    try {
      setCitizenshipsLoading(true);
      console.log('🌍 Starting to load citizenships...');
      const response = await apiService.getCitizenships();
      
      // 詳細檢查回應內容
      console.log('📡 Raw API response:', response);
      console.log('📊 Response data:', response.data);
      console.log('📝 Response message:', response.Message);
      console.log('🔍 Data type:', typeof response.data);
      console.log('📏 Is array?', Array.isArray(response.data));
      
      if (response.data && Array.isArray(response.data) && response.data.length > 0) {
        setCitizenships(response.data);
        console.log(`✅ Successfully loaded ${response.data.length} citizenships from API`);
        console.log('📋 First few citizenships:', response.data.slice(0, 3));
      } else if (response.data && !Array.isArray(response.data)) {
        console.error('❌ Response data is not an array:', response.data);
        setFormErrors({ general: 'Invalid citizenship data format received' });
      } else if (response.data && Array.isArray(response.data) && response.data.length === 0) {
        console.warn('⚠️ Empty citizenship array received');
        setFormErrors({ general: 'No citizenship data available' });
      } else {
        console.warn('❌ No citizenship data in response');
        console.log('Full response structure:', JSON.stringify(response, null, 2));
        setFormErrors({ general: 'Failed to load citizenship data' });
      }
    } catch (error) {
      console.error('💥 Failed to load citizenships:', error);
      console.error('Error details:', error.message);
      setFormErrors({ general: 'Failed to load citizenship data. Please try again.' });
    } finally {
      setCitizenshipsLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    // Clear previous errors
    setFormErrors({});
    
    // Validate form
    if (!validateForm()) {
      return;
    }

    try {
      if (editingCustomer) {
        await apiService.updateCustomer({ ...formData, CustomerId: editingCustomer.Id });
      } else {
        await apiService.createCustomer(formData);
      }
      resetForm();
      loadCustomers();
    } catch (error: any) {
      console.error('Failed to save customer:', error);
      
      // Handle specific API errors
      if (error.message && error.message.includes('NationalId')) {
        setFormErrors({ NationalId: 'This National ID already exists' });
      } else {
        setFormErrors({ general: 'Failed to save customer. Please try again.' });
      }
    }
  };

  const handleEdit = (customer: Customer) => {
    setEditingCustomer(customer);
    setFormData({
      CustomerId: customer.Id,
      Name: customer.Name,
      Gender: customer.Gender,
      Birthday: (customer.Birthday || '').split('T')[0],
      NationalId: customer.NationalId,
      Address: customer.Address,
      PhoneNumber: customer.PhoneNumber,
      CarNumber: customer.CarNumber,
      CitizenshipId: customer.CitizenshipId,
      Note: customer.Note,
    });
    setShowForm(true);
  };

  const handleDelete = async (customerId: string) => {
    if (window.confirm('Are you sure you want to delete this customer?')) {
      try {
        await apiService.deleteCustomer(customerId);
        loadCustomers();
      } catch (error) {
        console.error('Failed to delete customer:', error);
      }
    }
  };

  const handleSearch = async () => {
    const term = searchTerm.trim();
    if (!term) {
      loadCustomers();
      return;
    }
    try {
      setIsSearching(true);
      setSearchError(null);
      let response;
      switch (searchType) {
        case 'name':
          response = await apiService.getCustomerByName(term);
          break;
        case 'nationalId':
          response = await apiService.getCustomerByNationalId(term);
          break;
        case 'phone':
          response = await apiService.getCustomerByPhone(term);
          break;
      }

      if (response?.data) {
        const data: any = response.data as any;
        const list = Array.isArray(data) ? data : [data];
        const valid = list.filter((c) => c && typeof c === 'object' && 'Id' in c);
        setCustomers(valid as Customer[]);
      } else {
        setCustomers([]);
      }
    } catch (error) {
      console.error('Search failed:', error);
      setSearchError('Search failed. Please try again.');
      setCustomers([]);
    } finally {
      setIsSearching(false);
    }
  };

  const resetForm = () => {
    setFormData({
      Name: '',
      Gender: 'Male',
      Birthday: '',
      NationalId: '',
      Address: '',
      PhoneNumber: '',
      CarNumber: '',
      CitizenshipId: (citizenships.find(c => (c.Alpha3 || '').toUpperCase() === 'TWN')?.Id) ?? (citizenships.length > 0 ? citizenships[0].Id : 0),
      Note: '',
    });
    setFormErrors({});
    setEditingCustomer(null);
    setShowForm(false);
  };

  const getCitizenshipName = (id: number) => {
    const citizenship = citizenships.find(c => c.Id === id);
    return citizenship ? citizenship.Nation : 'Unknown';
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-center">
          <div className="text-lg text-gray-600">Loading customers...</div>
          {citizenshipsLoading && (
            <div className="text-sm text-gray-500 mt-2">Loading citizenship data...</div>
          )}
        </div>
      </div>
    );
  }

  // Ensure we never render malformed items and guard missing fields
  const safeCustomers: Customer[] = customers
    .filter((c: any) => c && typeof c === 'object' && 'Id' in c)
    .map((c: any) => ({ ...c, Birthday: c.Birthday || '' }));

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold text-gray-900">Customer Management</h2>
        <button
          onClick={() => {
            const tw = citizenships.find((c) => (c.Alpha3 || '').toUpperCase() === 'TWN');
            if (!editingCustomer) {
              setFormData((prev) => ({
                ...prev,
                CitizenshipId: tw ? tw.Id : (citizenships.length > 0 ? citizenships[0].Id : 0),
              }));
            }
            setShowForm(true);
          }}
          className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-md text-sm font-medium"
        >
          Add New Customer
        </button>
      </div>

      {/* Search */}
      <div className="bg-white p-4 rounded-lg shadow">
        <div className="flex gap-4 items-end">
          <div className="flex-1">
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Search Term
            </label>
            <input
              type="text"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === 'Enter') {
                  e.preventDefault();
                  handleSearch();
                }
              }}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
              placeholder="Enter search term..."
            />
            {searchError && (
              <p className="mt-1 text-xs text-red-600">{searchError}</p>
            )}
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Search By
            </label>
            <select
              value={searchType}
              onChange={(e) => setSearchType(e.target.value as 'name' | 'nationalId' | 'phone')}
              className="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
            >
              <option value="name">Name</option>
              <option value="nationalId">National ID</option>
              <option value="phone">Phone</option>
            </select>
          </div>
          <button
            onClick={handleSearch}
            className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md text-sm font-medium"
          >
            {isSearching ? 'Searching...' : 'Search'}
          </button>
          <button
            onClick={() => { setSearchTerm(''); setSearchType('name'); loadCustomers(); }}
            className="bg-gray-600 hover:bg-gray-700 text-white px-4 py-2 rounded-md text-sm font-medium"
          >
            Reset
          </button>
        </div>
      </div>

      {/* Customer Form Modal */}
      {showForm && (
        <div className="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
          <div className="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
            <div className="mt-3">
              <h3 className="text-lg font-medium text-gray-900 mb-4">
                {editingCustomer ? 'Edit Customer' : 'Add New Customer'}
              </h3>
              
              {/* General Error Message */}
              {formErrors.general && (
                <div className="rounded-md bg-red-50 p-4">
                  <div className="text-sm text-red-700">{formErrors.general}</div>
                </div>
              )}
              
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="grid grid-cols-2 gap-4">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Name *
                    </label>
                    <input
                      type="text"
                      required
                      value={formData.Name}
                      onChange={(e) => setFormData({ ...formData, Name: e.target.value })}
                      className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                        formErrors.Name ? 'border-red-500 bg-red-50' : 'border-gray-300'
                      }`}
                      maxLength={100}
                    />
                    {formErrors.Name && (
                      <p className="mt-1 text-xs text-red-600">{formErrors.Name}</p>
                    )}
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Gender *
                    </label>
                    <select
                      value={formData.Gender}
                      onChange={(e) => setFormData({ ...formData, Gender: e.target.value })}
                      className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                        formErrors.Gender ? 'border-red-500 bg-red-50' : 'border-gray-300'
                      }`}
                    >
                      <option value="Male">Male</option>
                      <option value="Female">Female</option>
                    </select>
                    {formErrors.Gender && (
                      <p className="mt-1 text-xs text-red-600">{formErrors.Gender}</p>
                    )}
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Birthday *
                    </label>
                    <input
                      type="date"
                      required
                      value={formData.Birthday}
                      onChange={(e) => setFormData({ ...formData, Birthday: e.target.value })}
                      className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                        formErrors.Birthday ? 'border-red-500 bg-red-50' : 'border-gray-300'
                      }`}
                      max={new Date().toISOString().split('T')[0]}
                      min="1900-01-01"
                    />
                    {formErrors.Birthday && (
                      <p className="mt-1 text-xs text-red-600">{formErrors.Birthday}</p>
                    )}
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      National ID *
                    </label>
                    <input
                      type="text"
                      required
                      value={formData.NationalId}
                      onChange={(e) => setFormData({ ...formData, NationalId: e.target.value })}
                      className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                        formErrors.NationalId ? 'border-red-500 bg-red-50' : 'border-gray-300'
                      }`}
                      maxLength={100}
                      pattern="[A-Za-z0-9\-_]+"
                      placeholder="e.g., A123456789, ID-001"
                    />
                    {formErrors.NationalId && (
                      <p className="mt-1 text-xs text-red-600">{formErrors.NationalId}</p>
                    )}
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Phone Number
                    </label>
                    <input
                      type="tel"
                      value={formData.PhoneNumber}
                      onChange={(e) => setFormData({ ...formData, PhoneNumber: e.target.value })}
                      className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                        formErrors.PhoneNumber ? 'border-red-500 bg-red-50' : 'border-gray-300'
                      }`}
                      placeholder="e.g., +1-555-123-4567"
                    />
                    {formErrors.PhoneNumber && (
                      <p className="mt-1 text-xs text-red-600">{formErrors.PhoneNumber}</p>
                    )}
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                      Car Number
                    </label>
                    <input
                      type="text"
                      value={formData.CarNumber}
                      onChange={(e) => setFormData({ ...formData, CarNumber: e.target.value })}
                      className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                        formErrors.CarNumber ? 'border-red-500 bg-red-50' : 'border-gray-300'
                      }`}
                      maxLength={20}
                      placeholder="e.g., ABC-123"
                    />
                    {formErrors.CarNumber && (
                      <p className="mt-1 text-xs text-red-600">{formErrors.CarNumber}</p>
                    )}
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    Address
                  </label>
                  <input
                    type="text"
                    value={formData.Address}
                    onChange={(e) => setFormData({ ...formData, Address: e.target.value })}
                    className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                      formErrors.Address ? 'border-red-500 bg-red-50' : 'border-gray-300'
                    }`}
                    maxLength={500}
                    placeholder="e.g., 123 Main St, City, State"
                  />
                  {formErrors.Address && (
                    <p className="mt-1 text-xs text-red-600">{formErrors.Address}</p>
                  )}
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    Citizenship *
                  </label>
                  <select
                    value={formData.CitizenshipId}
                    onChange={(e) => setFormData({ ...formData, CitizenshipId: parseInt(e.target.value) })}
                    className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                      formErrors.CitizenshipId ? 'border-red-500 bg-red-50' : 'border-gray-300'
                    }`}
                    disabled={citizenshipsLoading}
                  >
                    {citizenshipsLoading ? (
                      <option value={0}>Loading citizenships...</option>
                    ) : (
                      <>
                        <option value={0}>Please select citizenship...</option>
                        {citizenships.map((citizenship) => (
                          <option key={citizenship.Id} value={citizenship.Id}>
                            {citizenship.Nation} ({citizenship.Alpha3.toUpperCase()})
                          </option>
                        ))}
                      </>
                    )}
                  </select>
                  {formErrors.CitizenshipId && (
                    <p className="mt-1 text-xs text-red-600">{formErrors.CitizenshipId}</p>
                  )}
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    Note
                  </label>
                  <textarea
                    value={formData.Note}
                    onChange={(e) => setFormData({ ...formData, Note: e.target.value })}
                    rows={3}
                    className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 ${
                      formErrors.Note ? 'border-red-500 bg-red-50' : 'border-gray-300'
                    }`}
                    maxLength={1000}
                    placeholder="Additional notes or comments..."
                  />
                  {formErrors.Note && (
                    <p className="mt-1 text-xs text-red-600">{formErrors.Note}</p>
                  )}
                  <p className="mt-1 text-xs text-gray-500">
                    {formData.Note.length}/1000 characters
                  </p>
                </div>
                <div className="flex justify-end space-x-3 pt-4">
                  <button
                    type="button"
                    onClick={resetForm}
                    className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    disabled={citizenshipsLoading}
                    className={`px-4 py-2 text-sm font-medium text-white rounded-md ${
                      citizenshipsLoading 
                        ? 'bg-gray-400 cursor-not-allowed' 
                        : 'bg-indigo-600 hover:bg-indigo-700'
                    }`}
                  >
                    {citizenshipsLoading 
                      ? 'Loading...' 
                      : (editingCustomer ? 'Update' : 'Create')
                    }
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}

      {/* Customer List */}
      <div className="bg-white shadow overflow-hidden sm:rounded-md">
        <ul className="divide-y divide-gray-200">
          {safeCustomers.map((customer) => (
            <li key={customer.Id} className="px-6 py-4">
              <div className="flex items-center justify-between">
                <div className="flex-1 min-w-0">
                  <div className="flex items-center space-x-4">
                    <div className="flex-1">
                      <p className="text-sm font-medium text-gray-900 truncate">
                        {customer.Name}
                      </p>
                      <p className="text-sm text-gray-500">
                        {customer.Gender} • {customer.Birthday.split('T')[0]} • {getCitizenshipName(customer.CitizenshipId)}
                      </p>
                    </div>
                    <div className="flex-1">
                      <p className="text-sm text-gray-900">
                        National ID: {customer.NationalId}
                      </p>
                      <p className="text-sm text-gray-500">
                        Phone: {customer.PhoneNumber}
                      </p>
                    </div>
                    <div className="flex-1">
                      <p className="text-sm text-gray-900">
                        Address: {customer.Address}
                      </p>
                      {customer.CarNumber && (
                        <p className="text-sm text-gray-500">
                          Car: {customer.CarNumber}
                        </p>
                      )}
                    </div>
                  </div>
                  {customer.Note && (
                    <p className="mt-2 text-xs text-gray-500">
                      Note: {customer.Note}
                    </p>
                  )}
                </div>
                <div className="flex space-x-2">
                  <button
                    onClick={() => handleEdit(customer)}
                    className="text-indigo-600 hover:text-indigo-900 text-sm font-medium"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDelete(customer.Id)}
                    className="text-red-600 hover:text-red-900 text-sm font-medium"
                  >
                    Delete
                  </button>
                </div>
              </div>
            </li>
          ))}
        </ul>
        {customers.length === 0 && (
          <div className="text-center py-12">
            <p className="text-gray-500">No customers found</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default CustomerManagement;
