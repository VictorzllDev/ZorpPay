import type { QueryClient } from '@tanstack/react-query'
import { createRootRouteWithContext, HeadContent, Outlet } from '@tanstack/react-router'
import { AuthProvider } from '@/contexts/AuthContext'

interface MyRouterContext {
	queryClient: QueryClient
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
	component: RootComponent,
})

function RootComponent() {
	return (
		<AuthProvider>
			<HeadContent />
			<Outlet />
		</AuthProvider>
	)
}
