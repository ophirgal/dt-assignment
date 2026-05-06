import { useQuery } from '@tanstack/react-query'

import { API_V1_BASE_URL, ANALYTICS_FORECASTS_ENDPOINT } from '@/api/constants'
import { apiGet } from '@/api/utils'
import type { Forecast } from '@/models/forecast'

async function fetchForecast(storeSystemName: string, date: string): Promise<Forecast[]> {
  return apiGet<Forecast[]>(
    `${API_V1_BASE_URL}${ANALYTICS_FORECASTS_ENDPOINT}?store=${storeSystemName}&date=${date}`,
  )
}

export function useForecast(storeSystemName: string, date: string) {
  return useQuery({
    queryKey: ['forecasts', storeSystemName, date],
    queryFn: () => fetchForecast(storeSystemName, date),
    enabled: Boolean(storeSystemName && date),
  })
}
