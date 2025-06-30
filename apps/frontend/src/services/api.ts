const API_BASE_URL = '/api';

export interface LoginRequest {
  Username: string;
  Password: string;
}

export interface CustomerRequest {
  CustomerId?: string;
  Name: string;
  Gender: string;
  Birthday: string;
  NationalId: string;
  Address: string;
  PhoneNumber: string;
  CarNumber: string;
  CitizenshipId: number;
  Note: string;
}

export interface HistoryRequest {
  HistoryId?: string;
  CustomerId: string;
  Date: string;
  NumberOfPeople: number;
  Price: number;
  Room: string;
  Note: string;
}

export interface Customer {
  Id: string;
  Name: string;
  Gender: string;
  Birthday: string;
  NationalId: string;
  Address: string;
  PhoneNumber: string;
  CarNumber: string;
  CitizenshipId: number;
  Note: string;
}

export interface History {
  Id: string;
  CustomerId: string;
  Date: string;
  NumberOfPeople: number;
  Price: number;
  Room: string;
  Note: string;
}

export interface Citizenship {
  Id: number;
  Nation: string;
  Alpha3: string;
}

export interface ApiResponse<T> {
  Message: string;
  data?: T;
  token?: string;
  // API specific response fields
  citizenships?: Citizenship[];
  citizenship?: Citizenship;
  customers?: Customer[];
  customer?: Customer;
  histories?: History[];
  history?: History;
}

class ApiService {
  private async makeRequest<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    console.log(`🔗 Making API request to: ${API_BASE_URL}${endpoint}`);
    console.log(`📤 Request data:`, data);
    
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: data ? JSON.stringify(data) : undefined,
    });

    console.log(`📥 Response status: ${response.status}`);
    console.log(`📥 Response ok: ${response.ok}`);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.json();
    console.log(`📋 Raw response from ${endpoint}:`, result);
    
    return result;
  }

  // Authentication
  async login(credentials: LoginRequest): Promise<ApiResponse<any>> {
    return this.makeRequest('/userLogin', credentials);
  }

  async logout(): Promise<ApiResponse<any>> {
    return this.makeRequest('/userLogout');
  }

  async authenticate(token: string): Promise<ApiResponse<any>> {
    return this.makeRequest('/userAuthentication', { Token: token });
  }

  // Customers
  async getCustomers(): Promise<ApiResponse<Customer[]>> {
    const response = await this.makeRequest('/customerList');
    // Handle the specific customer response format
    if (response.customers) {
      response.data = response.customers;
    }
    return response;
  }

  async createCustomer(customer: CustomerRequest): Promise<ApiResponse<any>> {
    return this.makeRequest('/customerCre', customer);
  }

  async updateCustomer(customer: CustomerRequest): Promise<ApiResponse<any>> {
    return this.makeRequest('/customerMod', customer);
  }

  async deleteCustomer(customerId: string): Promise<ApiResponse<any>> {
    return this.makeRequest('/customerDel', { CustomerId: customerId });
  }

  async getCustomerById(customerId: string): Promise<ApiResponse<Customer>> {
    const response = await this.makeRequest('/customerID', { CustomerId: customerId });
    // Handle the specific customer response format
    if (response.customer) {
      response.data = response.customer;
    }
    return response;
  }

  async getCustomerByNationalId(nationalId: string): Promise<ApiResponse<Customer>> {
    const response = await this.makeRequest('/customerNationalId', { NationalId: nationalId });
    // Handle the specific customer response format
    if (response.customer) {
      response.data = response.customer;
    }
    return response;
  }

  async getCustomerByName(name: string): Promise<ApiResponse<Customer[]>> {
    const response = await this.makeRequest('/customerName', { Name: name });
    // Handle the specific customer response format
    if (response.customers) {
      response.data = response.customers;
    }
    return response;
  }

  async getCustomerByPhone(phone: string): Promise<ApiResponse<Customer>> {
    const response = await this.makeRequest('/customerPhone', { PhoneNumber: phone });
    // Handle the specific customer response format
    if (response.customer) {
      response.data = response.customer;
    }
    return response;
  }

  // History
  async getHistories(): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest('/historyList');
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  async createHistory(history: HistoryRequest): Promise<ApiResponse<any>> {
    return this.makeRequest('/historyCre', history);
  }

  async updateHistory(history: HistoryRequest): Promise<ApiResponse<any>> {
    return this.makeRequest('/historyMod', history);
  }

  async deleteHistory(historyId: string): Promise<ApiResponse<any>> {
    return this.makeRequest('/historyDel', { HistoryId: historyId });
  }

  async getHistoryById(historyId: string): Promise<ApiResponse<History>> {
    const response = await this.makeRequest('/historyByHistoryId', { HistoryId: historyId });
    // Handle the specific history response format
    if (response.history) {
      response.data = response.history;
    }
    return response;
  }

  async getHistoriesByCustomerId(customerId: string): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest('/historyCustomerId', { CustomerId: customerId });
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  async getHistoriesByDate(date: string): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest('/historyForDate', { Date: date });
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  async getHistoriesByDateRange(startDate: string, endDate: string): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest('/historyForDuring', { startDate, endDate });
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  // Citizenship
  async getCitizenships(): Promise<ApiResponse<Citizenship[]>> {
    const response = await this.makeRequest('/citizenships');
    // Handle the specific citizenship response format
    if (response.citizenships) {
      response.data = response.citizenships;
    }
    return response;
  }

  async getCitizenshipById(id: number): Promise<ApiResponse<Citizenship>> {
    const response = await this.makeRequest('/citizenshipId', { CitizenshipId: id });
    // Handle the specific citizenship response format
    if (response.citizenship) {
      response.data = response.citizenship;
    }
    return response;
  }

  async getCitizenshipByNation(nation: string): Promise<ApiResponse<Citizenship>> {
    const response = await this.makeRequest('/citizenshipNation', { CitizenshipName: nation });
    // Handle the specific citizenship response format
    if (response.citizenship) {
      response.data = response.citizenship;
    }
    return response;
  }
}

export const apiService = new ApiService();