import { useQuery } from "@apollo/client";
import { getNodes } from "../api";
import moment from "moment/moment";

export function NodesPage() {
  const { data } = useQuery(getNodes, {
    variables: {
      limit: 100,
    },
  });

  return (
    <div className="container md:w-1/2 md:mx-auto mx-4 grid gap-4">
      {data &&
        data.getNodes.items.map((node) => (
          <div>
            <a href={`/nodes/${node.publicKey}`} className="link link-primary">
              {node.address}
            </a>
            <p>Last seen {moment(node.lastSeen).fromNow().toString()}</p>
          </div>
        ))}
    </div>
  );
}
