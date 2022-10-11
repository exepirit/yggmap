import { h, Fragment } from 'preact';
import { NetworkSpanningTree } from './features/NetworkSpanningTree/NetworkSpanningTree';
import { NavBar } from './shared/components/NavBar';

export const App = (
  <>
    <NavBar />
    <NetworkSpanningTree />
  </>
)