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
import { STORES } from '@/data/stores'

export function StoreSidebar() {
  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Stores</SidebarGroupLabel>
          <SidebarMenu>
            {STORES.map((store) => (
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
