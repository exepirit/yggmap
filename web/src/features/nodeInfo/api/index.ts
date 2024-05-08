import {graphql} from "../../../gql";

export const getNodeDirectNeighbors = graphql(`
  query getNodeDirectNeighbors($publicKey: String!) {
    getNodeByKey(publicKey: $publicKey) {
      neighbors {
        node {
          publicKey
          address
        }
      }
    }
  }
`);