import { graphql } from "../../../gql";

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
