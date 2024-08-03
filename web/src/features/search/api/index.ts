import { graphql } from "../../../shared/gql";

export const searchNodes = graphql(`
  query searchNodes($query: String!, $offset: Int!, $limit: Int!) {
    nodesList(query: { keyOrAddress: $query }, offset: $offset, limit: $limit) {
      items {
        address
        publicKey
        lastSeen
      }
    }
  }
`);
