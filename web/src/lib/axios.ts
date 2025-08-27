import axios from 'axios'
import { env } from '@/env'

export const api = axios.create({
	baseURL: env.VITE_API_URL,
	timeout: 10000,
	headers: {
		'Content-Type': 'application/json',
	},
})

api.interceptors.request.use(
	(config) => {
		const token = localStorage.getItem('accessToken')
		if (token) {
			config.headers.Authorization = `Bearer ${token}`
		}
		return config
	},
	(error) => {
		return Promise.reject(error)
	},
)

api.interceptors.response.use(
	(response) => response,
	async (error) => {
		if (error.response?.status === 401) {
			localStorage.removeItem('accessToken')
			window.location.href = '/sign-in'
		}

		return Promise.reject(error)
	},
)
