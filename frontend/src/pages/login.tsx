export default function Login() {
	return (
		<div>
			<h1>Login Page</h1>
			<form>
				<input type="email" placeholder="Email" required />
				<input type="password" placeholder="Password" required />
				<button type="submit">Login</button>
			</form>
		</div>
	);
}
