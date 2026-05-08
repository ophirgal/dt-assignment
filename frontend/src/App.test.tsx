import { describe, it, expect, vi } from 'vitest'
import { render, screen } from '@testing-library/react'
import App from './App'

vi.mock('@/hooks/use-stores', () => ({
  useStores: () => ({ data: [{ id: 'store-1', displayName: 'KFC Times Square', systemName: 'kfc-times-square' }] }),
}))

vi.mock('@/hooks/use-forecast', () => ({
  useForecast: () => ({ data: [] }),
}))

function renderAt(path: string) {
  // App uses BrowserRouter internally, so we override history via window.location
  window.history.pushState({}, '', path)
  return render(<App />)
}

describe('App routing', () => {
  it('redirects unknown routes to the 404 page', () => {
    renderAt('/completely-unknown-path')
    expect(screen.getByText('Page Not Found')).toBeInTheDocument()
  })

  it('redirects /stores to the first store panel', () => {
    renderAt('/stores')
    // After redirect, URL should contain the first store's systemName
    expect(window.location.pathname).toBe('/stores/kfc-times-square')
  })
})
