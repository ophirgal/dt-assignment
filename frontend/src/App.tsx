import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import DashboardLayout from './pages/Dashboard/DashboardLayout'
import StorePanel from './pages/Dashboard/StorePanel'
import PageNotFound from './pages/PageNotFound/PageNotFound'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<DashboardLayout />}>
          <Route index element={<Navigate to="/stores" replace />} />
          {/* TODO: show spinner and Redirect to the first store loaded */}
          <Route path="stores" element={<Navigate to={`/stores/kfc-times-square`} replace />} />
          <Route path="stores/:storeSystemName" element={<StorePanel />} />
        </Route>
        <Route path="page-not-found" element={<PageNotFound />} />
        <Route path="*" element={<Navigate to="/page-not-found" />} />
      </Routes>
    </BrowserRouter>
  )
}
