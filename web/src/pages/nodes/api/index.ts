import {graphql} from "../../../gql";

export const getNodes = graphql(`
  query getNodes($limit: Int!, $previous: String) {
    getNodes(limit: $limit, previous: $previous) {
      items {
        publicKey
        address
        lastSeen
      }
    }
  }
`);