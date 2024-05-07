import { LocationProvider, Router, Route, hydrate, prerender as ssr } from 'preact-iso';

import { Header } from './components/Header.jsx';
import { Home } from './pages/Home';
import { NotFound } from './pages/_404.jsx';
import './style.css';
import {NodeInfoPage} from "./pages/NodeInfo/ui";
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import {NodesPage} from "./pages/Nodes";

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
            <Route path="/nodes" component={NodesPage}/>
            <Route path="/nodes/:publicKey" component={NodeInfoPage} />
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
