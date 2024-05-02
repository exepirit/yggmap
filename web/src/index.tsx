import { LocationProvider, Router, Route, hydrate, prerender as ssr } from 'preact-iso';

import { Header } from './components/Header.jsx';
import { Home } from './pages/Home';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import {NodeInfo} from "./pages/NodeInfo";
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";

export function App() {
  const client = new ApolloClient({
    uri: '/graphql',
    cache: new InMemoryCache(),
  });

	return (
    <ApolloProvider client={client}>
      <LocationProvider>
        <Header />
        <main>
          <Router>
            <Route path="/" component={Home} />
            <Route path="/nodes/:publicKey" component={NodeInfo} />
            <Route default component={NotFound} />
          </Router>
        </main>
      </LocationProvider>
    </ApolloProvider>
	);
}

if (typeof window !== 'undefined') {
	hydrate(<App />, document.getElementById('app'));
}

export async function prerender(data) {
	return await ssr(<App {...data} />);
}
