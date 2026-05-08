import { describe, it, expect, vi, beforeEach } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { MemoryRouter, Route, Routes } from 'react-router-dom'
import StorePanel from './StorePanel'

const mockStores = [
  { id: 'store-1', displayName: 'KFC Times Square', systemName: 'kfc-times-square' },
  { id: 'store-2', displayName: 'KFC Brooklyn', systemName: 'kfc-brooklyn' },
]

const mockForecasts = [
  { storeName: 'KFC Times Square', productName: 'Chicken Sandwich', forecastDate: '2026-01-10', hour: 12, predictedQuantity: 5.2 },
  { storeName: 'KFC Times Square', productName: 'Fries', forecastDate: '2026-01-10', hour: 12, predictedQuantity: 3.7 },
]

vi.mock('@/hooks/use-stores', () => ({
  useStores: () => ({ data: mockStores }),
}))

const mockUseForecast = vi.fn()
vi.mock('@/hooks/use-forecast', () => ({
  useForecast: (...args: unknown[]) => mockUseForecast(...args),
}))

function renderAtRoute(storeSystemName: string) {
  return render(
    <MemoryRouter initialEntries={[`/stores/${storeSystemName}`]}>
      <Routes>
        <Route path="/stores/:storeSystemName" element={<StorePanel />} />
      </Routes>
    </MemoryRouter>
  )
}

beforeEach(() => {
  mockUseForecast.mockReturnValue({ data: mockForecasts })
})

describe('StorePanel', () => {
  it('shows "Store not found" for an unknown store name', () => {
    renderAtRoute('kfc-does-not-exist')
    expect(screen.getByText('Store not found :/')).toBeInTheDocument()
  })

  it('renders the store name and chart for a valid store', () => {
    renderAtRoute('kfc-times-square')
    expect(screen.getByText(/Forecast · KFC Times Square/i)).toBeInTheDocument()
    expect(screen.queryByText('Store not found :/')).not.toBeInTheDocument()
  })

  it('calls useForecast with updated date when date input changes', async () => {
    renderAtRoute('kfc-times-square')
    const input = screen.getByLabelText('Forecast Date')
    await userEvent.clear(input)
    await userEvent.type(input, '2026-02-15')
    expect(mockUseForecast).toHaveBeenCalledWith('store-1', '2026-02-15')
  })
})
