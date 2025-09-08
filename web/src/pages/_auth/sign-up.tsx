import { zodResolver } from '@hookform/resolvers/zod'
import { createFileRoute, Link } from '@tanstack/react-router'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { useAuth } from '@/hooks/auth/useAuth'

export const Route = createFileRoute('/_auth/sign-up')({
	component: SignUp,
})

const signUpFormSchema = z.object({
	name: z.string().min(3, { message: 'Mínimo 3 caracteres' }).max(50, { message: 'Máximo 50 caracteres' }),
	email: z.email({ message: 'O e-mail é inválido' }),
	password: z.string().min(8, { message: 'Mínimo 6 caracteres' }).max(30, { message: 'Máximo 30 caracteres' }),
})

export type SignUpFormInputs = z.infer<typeof signUpFormSchema>

function SignUp() {
	const { signUp } = useAuth()

	const {
		register,
		handleSubmit,
		formState: { errors },
	} = useForm<SignUpFormInputs>({
		resolver: zodResolver(signUpFormSchema),
		defaultValues: {
			name: '',
			email: '',
			password: '',
		},
	})

	const onSubmit = (data: SignUpFormInputs) => {
		signUp.mutate(data)
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
						<h1 className="font-semibold text-xl">Registre-se</h1>
						<div className="w-full space-y-1">
							<Input type="text" placeholder="Nome" className="text-sm" {...register('name')} />
							{errors.name && <p className="text-red-500 text-sm">{errors.name.message}</p>}
						</div>
						<div className="w-full space-y-1">
							<Input type="email" placeholder="Email" className="text-sm" {...register('email')} />
							{errors.email && <p className="text-red-500 text-sm">{errors.email.message}</p>}
						</div>
						<div className="w-full space-y-1">
							<Input type="password" placeholder="Password" className="text-sm" {...register('password')} />
							{errors.password && <p className="text-red-500 text-sm">{errors.password.message}</p>}
						</div>

						<Button type="submit" className="w-full">
							Registrar
						</Button>
					</form>
					<div className="flex justify-center gap-1 text-muted-foreground text-sm">
						<p>Já é um usuário?</p>
						<Link to="/sign-in" className="font-medium text-primary hover:underline">
							Login
						</Link>
					</div>
				</div>
			</div>
		</section>
	)
}
