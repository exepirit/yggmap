import { useLocation } from 'preact-iso';

export function Header() {
	const { url } = useLocation();

	return (
		<header>
			<nav className="navbar bg-base-100">
        <ul className="menu">
          <li><a href="/">Home</a></li>
        </ul>
			</nav>
		</header>
	);
}
