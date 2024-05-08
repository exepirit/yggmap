import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import {LocationProvider, Route, Router} from "preact-iso";
import {Header} from "./widgets/header/header";
import {HomePage, NodesPage, NotFoundPage} from "./pages";
import {NodeInfoPage} from "./pages/nodeInfo/ui";
import './style.css';

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
            <Route path="/" component={HomePage} />
            <Route path="/nodes" component={NodesPage}/>
            <Route path="/nodes/:publicKey" component={NodeInfoPage} />
            <Route default component={NotFoundPage} />
          </Router>
        </main>
      </LocationProvider>
    </ApolloProvider>
  );
}