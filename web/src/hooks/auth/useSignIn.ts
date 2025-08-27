import { useMutation } from '@tanstack/react-query'
import { signInWithEmail } from '@/services/auth/signInWithEmail'

export function useSignIn() {
	return useMutation({
		mutationFn: signInWithEmail,
		onError: (error) => {
			console.log('Error ao realizar o login:', error)
		},
		onSuccess: (data) => {
			console.log(data)
			console.log('Login realizado com sucesso')
		},
	})
}
