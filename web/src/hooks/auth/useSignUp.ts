import { useMutation } from '@tanstack/react-query'
import { signUpWithEmail } from '@/services/auth/signUpWithEmail'

export function useSignUp() {
	return useMutation({
		mutationFn: signUpWithEmail,
		onError: (error) => {
			console.log('Error ao realizar o cadastro:', error)
		},
		onSuccess: (data) => {
			console.log(data)
			console.log('Cadastro realizado com sucesso')
		},
	})
}
