import { useSuspenseQuery } from '@tanstack/react-query'

import { API_V1_BASE_URL, ANALYTICS_FORECASTS_ENDPOINT } from '@/api/constants'
import { apiGet } from '@/api/utils'
import type { Forecast } from '@/models/forecast'

async function fetchForecast(storeId: string, date: string): Promise<Forecast[]> {
  return apiGet<Forecast[]>(
    `${API_V1_BASE_URL}${ANALYTICS_FORECASTS_ENDPOINT}?storeId=${storeId}&date=${date}`,
  )
}

export function useForecast(storeId: string, date: string) {
  return useSuspenseQuery({
    queryKey: ['forecasts', storeId, date],
    queryFn: () => fetchForecast(storeId, date),
  })
}
