import { useQuery } from "@apollo/client";
import { getNodeByIdQueryDocument } from "../api";
import { ErrorBanner } from "../../../widgets/errorBanner";
import { SummarySection, NeighborsView } from "../../../entities/node/ui";

interface NodeInfoPageProps {
  publicKey: string;
}

export function NodeInfoPage(props: NodeInfoPageProps) {
  const { data, loading, error } = useQuery(getNodeByIdQueryDocument, {
    variables: { publicKey: props.publicKey },
  });

  return (
    <div className="container mx-auto px-4">
      <a className="link" href="/nodes">
        Back to Search
      </a>
      <div className="pt-4">
        {loading && (
          <span className="mx-auto loading loading-dots loading-lg"></span>
        )}
        {error && (
          <ErrorBanner text={error.networkError?.message || error.message} />
        )}
        {!loading && !error && (
          <>
            <h3 className="text-lg">{data.getNodeByKey.address}</h3>
            <span className="badge badge-md badge-primary badge-outline">
              Node
            </span>
            <SummarySection node={data.getNodeByKey} />
            <NeighborsView node={data.getNodeByKey} />
          </>
        )}
      </div>
    </div>
  );
}
