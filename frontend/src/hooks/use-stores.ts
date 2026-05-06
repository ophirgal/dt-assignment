
import { ANALYTICS_STORES_ENDPOINT, API_V1_BASE_URL } from '@/api/constants'
import { apiGet } from '@/api/utils'
import type { Store } from '@/models/store'
import { useQuery } from '@tanstack/react-query'

async function fetchStores(): Promise<Store[]> {
    return apiGet<Store[]>(`${API_V1_BASE_URL}${ANALYTICS_STORES_ENDPOINT}`)
}

export function useStores() {
    return useQuery({
        queryKey: ['stores'],
        queryFn: fetchStores,
    })
}