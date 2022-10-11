import { h } from 'preact';
import { useState } from 'preact/hooks';
import { spanningTreeMock } from './spanningTreeMock';
import { ForceDirectedTree } from './ForceDirectedTree';

export const NetworkSpanningTree = () => {
  const [tree] = useState(spanningTreeMock);

  return (
    <div className='container'>
      <ForceDirectedTree nodes={tree.nodes} links={tree.edges} />
    </div>
  )
};