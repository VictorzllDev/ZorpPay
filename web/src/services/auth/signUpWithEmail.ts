import { api } from '@/lib/axios'
import type { IAuthResponse } from '@/types/auth'

export interface ISignUpWithEmail {
	name: string
	email: string
	password: string
}

export async function signUpWithEmail({ name, email, password }: ISignUpWithEmail): Promise<IAuthResponse> {
	const { data } = await api.post<IAuthResponse>('/api/v1/auth/signup', {
		name,
		email,
		password,
	})

	return data
}
