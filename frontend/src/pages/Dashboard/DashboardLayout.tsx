import { Outlet } from 'react-router-dom'
import { SidebarProvider, SidebarInset } from '@/components/ui/sidebar'
import { StoreSidebar } from './StoreSidebar'

export default function DashboardLayout() {
  return (
    <SidebarProvider>
      <StoreSidebar />
      <SidebarInset>
        <Outlet />
      </SidebarInset>
    </SidebarProvider>
  )
}
