import {graphql} from "../../../gql";

export const getNodeByIdQueryDocument = graphql(`
  query getNodeById($publicKey: String!) {
    getNodeByKey(publicKey: $publicKey) {
      publicKey
      address
      lastSeen
    }
  }
`);

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
`)