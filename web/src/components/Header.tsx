import { useLocation } from 'preact-iso';

export function Header() {
	const { url } = useLocation();

	return (
		<header>
      <nav className="navbar bg-base-100">
        <div className="navbar-start">
          <ul className="menu menu-horizontal">
            <li><a href="/">Home</a></li>
            <li><a href="/nodes">Nodes</a></li>
          </ul>
        </div>
      </nav>
    </header>
);
}
