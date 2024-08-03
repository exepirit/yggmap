import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import { LocationProvider, Route, Router } from "preact-iso";
import "./style.css";
import {NotFoundPage} from "./features/status/pages";
import {SearchPage} from "./features/search/pages";
import {Header} from "./shared/components";
import {HomePage} from "./features/landing/pages";
import {NodeInfoPage} from "./features/node-info/pages";

export function App() {
  const client = new ApolloClient({
    uri: "/graphql",
    cache: new InMemoryCache(),
  });

  return (
    <ApolloProvider client={client}>
      <LocationProvider>
        <Header />
        <main>
          <Router>
            <Route path="/" component={HomePage} />
            <Route path="/nodes" component={SearchPage} />
            <Route path="/nodes/:publicKey" component={NodeInfoPage} />
            <Route default component={NotFoundPage} />
          </Router>
        </main>
      </LocationProvider>
    </ApolloProvider>
  );
}
