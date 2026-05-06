import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import DashboardLayout from './pages/Dashboard/DashboardLayout'
import { STORES } from './data/stores' // TODO: get data from backend
import StorePanel from './pages/Dashboard/StorePanel'
import PageNotFound from './pages/PageNotFound/PageNotFound'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<DashboardLayout />}>
          <Route index element={<Navigate to={`/stores/${STORES[0].systemName}`} replace />} />
          <Route path="stores" element={<Navigate to={`/stores/${STORES[0].systemName}`} replace />} />
          <Route path="stores/:storeSystemName" element={<StorePanel />} />
        </Route>
        <Route path="page-not-found" element={<PageNotFound />} />
        <Route path="*" element={<Navigate to="/page-not-found" />} />
      </Routes>
    </BrowserRouter>
  )
}
