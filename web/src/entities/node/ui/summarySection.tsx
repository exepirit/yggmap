import moment from "moment/moment";
import {YggdrasilNode} from "../model";

interface SummarySectionProps {
  node: YggdrasilNode;
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
          <td>IPv6 address:</td>
          <td className="break-all"><code>{props.node.address}</code></td>
        </tr>
        <tr>
          <td>Public key:</td>
          <td className="break-all"><code>{props.node.publicKey}</code></td>
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