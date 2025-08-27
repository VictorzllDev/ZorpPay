import type { UseMutationResult } from '@tanstack/react-query'
import { createContext, useMemo, useState } from 'react'
import { useSignIn } from '@/hooks/auth/useSignIn'
import { useSignUp } from '@/hooks/auth/useSignUp'
import type { ISignInWithEmail } from '@/services/auth/signInWithEmail'
import type { ISignUpWithEmail } from '@/services/auth/signUpWithEmail'
import type { IAuthResponse, IUser } from '@/types/auth'

interface AuthContextType {
	user: IUser | null
	signIn: UseMutationResult<IAuthResponse, Error, ISignInWithEmail, unknown>
	signUp: UseMutationResult<IAuthResponse, Error, ISignUpWithEmail, unknown>
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
	const [user, _setUser] = useState<IUser | null>(null)

	const signIn = useSignIn()
	const signUp = useSignUp()

	const value = useMemo(
		() => ({
			user,
			signIn,
			signUp,
		}),
		[user, signIn, signUp],
	)

	return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}
