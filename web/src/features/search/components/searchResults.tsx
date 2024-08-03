import { useQuery } from "@apollo/client";
import { searchNodes } from "../api";
import moment from "moment";

interface SearchResultProps {
  query: string;
  page: number;
}

export function SearchResults({ query, page }: SearchResultProps) {
  const pageSize = 10;
  const { data, loading, error } = useQuery(searchNodes, {
    variables: {
      query: query,
      offset: page * pageSize,
      limit: pageSize
    },
  });

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const nodesList = data?.nodesList;

  return (
    <div className="grid gap-4">
      {nodesList && nodesList.items.map((node, index) => (
        <SearchResultCard
          key={index}
          title={node.address}
          publicKey={node.publicKey}
          lastSeen={node.lastSeen}
        />
      ))}
    </div>
  );
}

interface SearchResultCardProps {
  title: string;
  publicKey: string;
  lastSeen: string;
}

function SearchResultCard({ title, publicKey, lastSeen }: SearchResultCardProps) {
  return (
    <div>
      <a href={`/nodes/${publicKey}`} className="link link-primary">
        {title}
      </a>
      <p>Last seen {moment(lastSeen).fromNow().toString()}</p>
    </div>
  );
}