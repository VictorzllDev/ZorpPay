import { api } from '@/lib/axios'
import type { IAuthResponse } from '@/types/auth'

export interface ISignInWithEmail {
	email: string
	password: string
}

export async function signInWithEmail({ email, password }: ISignInWithEmail): Promise<IAuthResponse> {
	const { data } = await api.post<IAuthResponse>('/api/v1/auth/signin', {
		email,
		password,
	})

	return data
}
