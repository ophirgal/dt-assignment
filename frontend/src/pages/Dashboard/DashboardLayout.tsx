import { Suspense } from 'react'
import { Outlet } from 'react-router-dom'
import { SidebarProvider, SidebarInset } from '@/components/ui/sidebar'
import { Skeleton } from '@/components/ui/skeleton'
import { StoreSidebar } from './StoreSidebar'

export default function DashboardLayout() {
  return (
    <SidebarProvider>
      <Suspense fallback={<SidebarSkeleton />}>
        <StoreSidebar />
      </Suspense>
      <SidebarInset>
        <Suspense fallback={<PanelSkeleton />}>
          <Outlet />
        </Suspense>
      </SidebarInset>
    </SidebarProvider>
  )
}

function SidebarSkeleton() {
  return (
    <div className="flex flex-col gap-2 p-4 w-48">
      {Array.from({ length: 5 }).map((_, i) => (
        <Skeleton key={i} className="h-8 w-full" />
      ))}
    </div>
  )
}

function PanelSkeleton() {
  return (
    <div className="flex items-center justify-center h-full">
      <Skeleton className="h-8 w-8 rounded-full" />
    </div>
  )
}
