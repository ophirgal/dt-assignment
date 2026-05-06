import { NavLink } from 'react-router-dom'
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar'
import { useStores } from '@/hooks/use-stores'

export function StoreSidebar() {
  const { data: stores = [] } = useStores()

  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Stores</SidebarGroupLabel>
          <SidebarMenu>
            {stores.map((store) => (
              <SidebarMenuItem key={store.id}>
                <SidebarMenuButton asChild>
                  <NavLink to={`/stores/${store.systemName}`}>
                    {store.displayName}
                  </NavLink>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}
