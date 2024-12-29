import { useEffect, useState } from "preact/hooks";
import "./style.css";
import ForceGraph3D from 'react-force-graph-3d';
import { Graph } from "./entities/models";
import { getNetworkGraph } from "./entities/api";

export function App() {
  const [ graphData, setGraphData ] = useState<Graph>({
    nodes: [],
    links: []
  });

  useEffect(() => {
    getNetworkGraph()
      .then(graph => setGraphData(graph))
      .catch(error => console.error(error));
  }, []);

  return (
        <main>
          <ForceGraph3D
            graphData={graphData}
            nodeLabel={node => node.id}
          />
        </main>
  );
}
