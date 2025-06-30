#!/usr/bin/env node

// 測試 Citizenship API 的簡單腳本
const API_BASE_URL = 'http://localhost:8080/api';

async function testCitizenshipAPI() {
  console.log('🧪 Testing Citizenship API...');
  console.log('================================');

  try {
    // 測試獲取所有國籍
    console.log('📡 Fetching all citizenships...');
    const response = await fetch(`${API_BASE_URL}/citizenships`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
    });

    console.log(`📊 Response status: ${response.status}`);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.json();
    console.log(`✅ API Response received`);
    console.log(`📝 Message: ${result.Message}`);
    
    if (result.data && Array.isArray(result.data)) {
      console.log(`🌍 Found ${result.data.length} citizenships`);
      console.log('\n📋 Sample citizenships:');
      
      // 顯示前 10 個國籍
      result.data.slice(0, 10).forEach((citizenship, index) => {
        console.log(`   ${index + 1}. ${citizenship.Nation} (${citizenship.Alpha3.toUpperCase()}) - ID: ${citizenship.Id}`);
      });

      if (result.data.length > 10) {
        console.log(`   ... and ${result.data.length - 10} more`);
      }

      // 查找一些常見國家
      console.log('\n🔍 Common countries:');
      const commonCountries = ['中國', '台灣', '美國', '日本', '韓國'];
      commonCountries.forEach(country => {
        const found = result.data.find(c => c.Nation.includes(country));
        if (found) {
          console.log(`   ✅ ${country}: ID ${found.Id} (${found.Alpha3.toUpperCase()})`);
        } else {
          console.log(`   ❌ ${country}: Not found`);
        }
      });

    } else {
      console.log('⚠️  No citizenship data in response');
      console.log('Response:', JSON.stringify(result, null, 2));
    }

  } catch (error) {
    console.error('❌ Error testing Citizenship API:', error.message);
    console.log('\n🔧 Troubleshooting:');
    console.log('1. Make sure the API server is running on http://localhost:8080');
    console.log('2. Check if CORS is properly configured');
    console.log('3. Verify the API endpoint /api/citizenships exists');
    console.log('4. Check server logs for any errors');
  }
}

// 測試單個國籍 API
async function testSingleCitizenship() {
  console.log('\n🧪 Testing single citizenship API...');
  try {
    const response = await fetch(`${API_BASE_URL}/citizenshipId`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({ CitizenshipId: 156 }) // 中國
    });

    if (response.ok) {
      const result = await response.json();
      console.log('✅ Single citizenship test successful');
      console.log(`📝 Result: ${JSON.stringify(result.data, null, 2)}`);
    } else {
      console.log(`❌ Single citizenship test failed: ${response.status}`);
    }
  } catch (error) {
    console.log(`❌ Single citizenship test error: ${error.message}`);
  }
}

// 執行測試
async function runAllTests() {
  await testCitizenshipAPI();
  await testSingleCitizenship();
  
  console.log('\n🎯 Summary:');
  console.log('- If you see citizenship data above, the API is working correctly');
  console.log('- The frontend should now load citizenships from the API automatically');
  console.log('- Check the browser console for any additional debug information');
}

runAllTests();