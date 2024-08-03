import { graphql } from "../../../shared/gql";

export const getNodeDirectNeighbors = graphql(`
  query getNodeDirectNeighbors($publicKey: String!) {
    node(publicKey: $publicKey) {
      neighbors {
        node {
          publicKey
          address
        }
      }
    }
  }
`);
