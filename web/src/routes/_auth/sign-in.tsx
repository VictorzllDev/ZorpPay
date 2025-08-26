import { createFileRoute, Link } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export const Route = createFileRoute("/_auth/sign-in")({
	component: SignIn,
});

function SignIn() {
	return (
		<section className="h-screen bg-muted">
			<div className="flex h-full items-center justify-center">
				<div className="flex flex-col items-center gap-6 lg:justify-start">
					<Link to="/">
						<h1 className="font-semibold text-2xl">ZorpPay</h1>
					</Link>

					<div className="flex w-full min-w-sm max-w-sm flex-col items-center gap-y-4 rounded-md border border-muted bg-background px-6 py-8 shadow-md">
						<h1 className="font-semibold text-xl">Login</h1>
						<Input
							type="email"
							placeholder="Email"
							className="text-sm"
							required
						/>
						<Input
							type="password"
							placeholder="Senha"
							className="text-sm"
							required
						/>
						<Button type="submit" className="w-full">
							Entrar
						</Button>
					</div>
					<div className="flex justify-center gap-1 text-muted-foreground text-sm">
						<p>Precisa de uma conta?</p>
						<Link
							to="/sign-up"
							className="font-medium text-primary hover:underline"
						>
							Sign up
						</Link>
					</div>
				</div>
			</div>
		</section>
	);
}
