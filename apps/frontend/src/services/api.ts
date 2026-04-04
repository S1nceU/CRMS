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

// Normalize various backend response shapes into a consistent ApiResponse<T>
function normalizeApiResponse<T>(r: any): ApiResponse<T> {
  if (Array.isArray(r)) {
    return { Message: '', data: r } as ApiResponse<T>;
  }
  if (r && typeof r === 'object') {
    if ('data' in r) return r as ApiResponse<T>;
    // User/auth specific shape
    if ('username' in r) return { ...(r as object), data: (r as any).username } as unknown as ApiResponse<T>;
    // Preserve any additional fields (e.g., token) when Message exists
    if (
      'Message' in r &&
      !('customers' in r) &&
      !('customer' in r) &&
      !('citizenships' in r) &&
      !('citizenship' in r) &&
      !('histories' in r) &&
      !('history' in r)
    ) {
      return r as ApiResponse<T>;
    }
    if ('customers' in r) return { ...(r as object), data: (r as any).customers } as ApiResponse<T>;
    if ('customer' in r) return { ...(r as object), data: (r as any).customer } as ApiResponse<T>;
    if ('citizenships' in r) return { ...(r as object), data: (r as any).citizenships } as ApiResponse<T>;
    if ('citizenship' in r) return { ...(r as object), data: (r as any).citizenship } as ApiResponse<T>;
    if ('histories' in r) return { ...(r as object), data: (r as any).histories } as ApiResponse<T>;
    if ('history' in r) return { ...(r as object), data: (r as any).history } as ApiResponse<T>;
    return { Message: (r as any).Message ?? '', data: r } as ApiResponse<T>;
  }
  return { Message: '', data: r } as ApiResponse<T>;
}

class ApiService {
  private async makeRequest<T>(endpoint: string, data?: any): Promise<ApiResponse<T>> {
    const token = localStorage.getItem('token');

    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
    };
    if (token) {
      headers.Authorization = `Bearer ${token}`;
    }

    console.log(`Making API request to: ${API_BASE_URL}${endpoint}`);
    console.log('Request data:', data);

    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'POST',
      headers,
      credentials: 'include',
      body: data ? JSON.stringify(data) : undefined,
    });

    console.log(`Response status: ${response.status}`);
    console.log(`Response ok: ${response.ok}`);

    if (response.status === 401) {
      localStorage.removeItem('token');
      window.dispatchEvent(new Event('crms:unauthorized'));
      const unauthorizedError: any = new Error('Unauthorized');
      unauthorizedError.status = 401;
      throw unauthorizedError;
    }

    if (!response.ok) {
      const error: any = new Error(`HTTP error! status: ${response.status}`);
      error.status = response.status;
      throw error;
    }

    const result = await response.json();
    console.log(`Raw response from ${endpoint}:`, result);

    return normalizeApiResponse<T>(result);
  }

  // Authentication
  async login(credentials: LoginRequest): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/userLogin', credentials);
  }

  async logout(): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/userLogout');
  }

  async authenticate(token: string): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/userAuthentication', { Token: token });
  }

  // Customers
  async getCustomers(): Promise<ApiResponse<Customer[]>> {
    const response = await this.makeRequest<Customer[]>('/customerList');
    // Handle the specific customer response format
    if (response.customers) {
      response.data = response.customers;
    }
    return response;
  }

  async createCustomer(customer: CustomerRequest): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/customerCre', customer);
  }

  async updateCustomer(customer: CustomerRequest): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/customerMod', customer);
  }

  async deleteCustomer(customerId: string): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/customerDel', { CustomerId: customerId });
  }

  async getCustomerById(customerId: string): Promise<ApiResponse<Customer>> {
    const response = await this.makeRequest<Customer>('/customerID', { CustomerId: customerId });
    // Handle the specific customer response format
    if (response.customer) {
      response.data = response.customer;
    }
    return response;
  }

  async getCustomerByNationalId(nationalId: string): Promise<ApiResponse<Customer>> {
    const response = await this.makeRequest<Customer>('/customerNationalId', { NationalId: nationalId });
    // Handle the specific customer response format
    if (response.customer) {
      response.data = response.customer;
    }
    return response;
  }

  async getCustomerByName(name: string): Promise<ApiResponse<Customer[]>> {
    const response = await this.makeRequest<Customer[]>('/customerName', { Name: name });
    // Handle the specific customer response format
    if (response.customers) {
      response.data = response.customers;
    }
    return response;
  }

  async getCustomerByPhone(phone: string): Promise<ApiResponse<Customer[]>> {
    const response = await this.makeRequest<Customer[]>('/customerPhone', { PhoneNumber: phone });
    // Backend returns a list for phone search
    if ((response as any).customers) {
      response.data = (response as any).customers;
    }
    return response;
  }

  // History
  async getHistories(): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest<History[]>('/historyList');
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  async createHistory(history: HistoryRequest): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/historyCre', history);
  }

  async updateHistory(history: HistoryRequest): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/historyMod', history);
  }

  async deleteHistory(historyId: string): Promise<ApiResponse<any>> {
    return this.makeRequest<any>('/historyDel', { HistoryId: historyId });
  }

  async getHistoryById(historyId: string): Promise<ApiResponse<History>> {
    const response = await this.makeRequest<History>('/historyByHistoryId', { HistoryId: historyId });
    // Handle the specific history response format
    if (response.history) {
      response.data = response.history;
    }
    return response;
  }

  async getHistoriesByCustomerId(customerId: string): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest<History[]>('/historyCustomerId', { CustomerId: customerId });
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  async getHistoriesByDate(date: string): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest<History[]>('/historyForDate', { Date: date });
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  async getHistoriesByDateRange(startDate: string, endDate: string): Promise<ApiResponse<History[]>> {
    const response = await this.makeRequest<History[]>('/historyForDuring', { startDate, endDate });
    // Handle the specific history response format
    if (response.histories) {
      response.data = response.histories;
    }
    return response;
  }

  // Citizenship
  async getCitizenships(): Promise<ApiResponse<Citizenship[]>> {
    const response = await this.makeRequest<Citizenship[]>('/citizenships');
    // Handle the specific citizenship response format
    if (response.citizenships) {
      response.data = response.citizenships;
    }
    return response;
  }

  async getCitizenshipById(id: number): Promise<ApiResponse<Citizenship>> {
    const response = await this.makeRequest<Citizenship>('/citizenshipId', { CitizenshipId: id });
    // Handle the specific citizenship response format
    if (response.citizenship) {
      response.data = response.citizenship;
    }
    return response;
  }

  async getCitizenshipByNation(nation: string): Promise<ApiResponse<Citizenship>> {
    const response = await this.makeRequest<Citizenship>('/citizenshipNation', { CitizenshipName: nation });
    // Handle the specific citizenship response format
    if (response.citizenship) {
      response.data = response.citizenship;
    }
    return response;
  }
}

export const apiService = new ApiService();
