import { Suspense } from 'react'
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import DashboardLayout from './pages/Dashboard/DashboardLayout'
import StorePanel from './pages/Dashboard/StorePanel'
import PageNotFound from './pages/PageNotFound/PageNotFound'
import { useStores } from './hooks/use-stores'

function DefaultStoreRedirect() {
  const { data: stores } = useStores()
  return <Navigate to={`/stores/${stores[0].systemName}`} replace />
}

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<DashboardLayout />}>
          <Route index element={<Navigate to="/stores" replace />} />
          <Route path="stores" element={<Suspense><DefaultStoreRedirect /></Suspense>} />
          <Route path="stores/:storeSystemName" element={<StorePanel />} />
        </Route>
        <Route path="page-not-found" element={<PageNotFound />} />
        <Route path="*" element={<Navigate to="/page-not-found" />} />
      </Routes>
    </BrowserRouter>
  )
}
