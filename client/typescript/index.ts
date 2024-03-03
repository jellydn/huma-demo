import * as ApiClient from 'itman-api'

const apiClient = new ApiClient.ItmanChannelApiClient({
  environment: 'http://localhost:8888',
})

const health = await apiClient.getHealth()
console.log(health)
