import moment from "moment/moment";

interface YggdrasilNodeSummary {
  lastSeen: string;
  publicKey: string;
}

interface SummarySectionProps {
  node: YggdrasilNodeSummary;
}

/**
 * Represents a section of a UI displaying a summary about a specific Yggdrasil node.
 * The last seen time is displayed relative to the current moment in time.
 */
export function SummarySection(props: SummarySectionProps) {
  return (
    <div className="py-7">
      <table className="table">
        <tbody>
          <tr>
            <td>Public key:</td>
            <td><code>{props.node.publicKey}</code></td>
          </tr>
          <tr>
            <td>Last seen:</td>
            <td>{moment(props.node.lastSeen).fromNow()}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}