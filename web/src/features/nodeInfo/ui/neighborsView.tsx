import {useQuery} from "@apollo/client";
import {getNodeDirectNeighbors} from "../api";
import {ForceDirectedGraph} from "../../../widgets/forceDirectedGraph";

interface NeighborsWidgetProps {
  node: {
    publicKey: string
    address: string
  }
}

export function NeighborsView(props: NeighborsWidgetProps) {
  const { data } = useQuery(getNodeDirectNeighbors, {variables: {
    publicKey: props.node.publicKey
  }});

  function renderGraph() {
    const nodes = [
      {id: props.node.publicKey, group: "self"},
      ...data.getNodeByKey.neighbors.map(link => ({
        id: link.node.publicKey,
        group: "others"
      }))
    ]
    const links = data.getNodeByKey.neighbors.map(link => ({
      source: props.node.publicKey,
      target: link.node.publicKey
    }));

    return <ForceDirectedGraph nodes={nodes} links={links} width={250} height={250}/>
  }

  return (
    <>
      <h3 className="pb-2 text-lg">Neighbors</h3>
      <div className="rounded-lg border border-slate-700 xl:w-2/5 lg:w-2/3">
        {data && renderGraph()}
      </div>
    </>
  );
}