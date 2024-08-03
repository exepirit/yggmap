import { graphql } from "../../../shared/gql";

export const getNodeByIdQueryDocument = graphql(`
  query getNodeById($publicKey: String!) {
    node(publicKey: $publicKey) {
      publicKey
      address
      lastSeen
    }
  }
`);
