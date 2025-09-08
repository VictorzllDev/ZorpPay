import { resolve } from 'node:path'
import tailwindcss from '@tailwindcss/vite'
import viteReact from '@vitejs/plugin-react'
import { tanstackRouter } from '@tanstack/router-plugin/vite'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		tanstackRouter({
			target: 'react',
			autoCodeSplitting: true,
			generatedRouteTree: './src/route-tree.gen.ts',
			routesDirectory: './src/pages',
			routeToken: '_layout',
		}),
		,
		viteReact(),
		tailwindcss(),
	],
	resolve: {
		alias: {
			'@': resolve(__dirname, './src'),
		},
	},
})
