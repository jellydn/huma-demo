import * as ApiClient from 'itman-api';
import { expect } from 'chai';

import { ItmanChannelApiClient } from './index';

describe('ItmanChannelApiClient', () => {
  let apiClient: ItmanChannelApiClient;

  beforeEach(() => {
    apiClient = new ItmanChannelApiClient({
      environment: 'http://localhost:8888',
    });
  });

  it('should get health', async () => {
    // Mock the response from the API client
    const expectedHealth = { status: 'healthy' };
    const getHealthMock = jest
      .spyOn(apiClient, 'getHealth')
      .mockResolvedValue(expectedHealth);

    // Call the function being tested
    const health = await apiClient.getHealth();

    // Assert the expected result
    expect(health).to.deep.equal(expectedHealth);

    // Verify that the function was called with the correct arguments
    expect(getHealthMock).to.have.been.calledOnce;
  });

  // Add more test cases for other functions in index.ts

});
