import { zodResolver } from '@hookform/resolvers/zod'
import { createFileRoute, Link } from '@tanstack/react-router'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { useAuth } from '@/hooks/auth/useAuth'

export const Route = createFileRoute('/_auth/sign-in')({
	component: SignIn,
	head: () => ({ meta: [{ title: 'Sign In | ZorpPay' }] }),
})

const signInFormSchema = z.object({
	email: z.email({ message: 'O e-mail é inválido' }),
	password: z.string().min(1, { message: 'A senha é obrigatória' }),
})

export type SignInFormInputs = z.infer<typeof signInFormSchema>

function SignIn() {
	const { signIn } = useAuth()

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<SignInFormInputs>({
		resolver: zodResolver(signInFormSchema),
		defaultValues: {
			email: '',
			password: '',
		},
	})

	const onSubmit = (data: SignInFormInputs) => {
		signIn.mutate(data)
	}

	return (
		<section className="h-screen bg-muted">
			<div className="flex h-full items-center justify-center">
				<div className="flex flex-col items-center gap-6 lg:justify-start">
					<Link to="/">
						<h1 className="font-semibold text-2xl">ZorpPay</h1>
					</Link>

					<form
						onSubmit={handleSubmit(onSubmit)}
						className="flex w-full min-w-sm max-w-sm flex-col items-center gap-y-4 rounded-md border border-muted bg-background px-6 py-8 shadow-md"
					>
						<h1 className="font-semibold text-xl">Login</h1>
						<div className="w-full space-y-1">
							<Input type="email" placeholder="Email" className="text-sm" {...register('email')} />
							{errors.email && <p className="text-red-500 text-sm">{errors.email.message}</p>}
						</div>

						<div className="w-full space-y-1">
							<Input type="password" placeholder="Senha" className="text-sm" {...register('password')} />
							{errors.password && <p className="text-red-500 text-sm">{errors.password.message}</p>}
						</div>
						<Button type="submit" className="w-full">
							Entrar
						</Button>
					</form>
					<div className="flex justify-center gap-1 text-muted-foreground text-sm">
						<p>Precisa de uma conta?</p>
						<Link to="/sign-up" className="font-medium text-primary hover:underline">
							Sign up
						</Link>
					</div>
				</div>
			</div>
		</section>
	)
}
