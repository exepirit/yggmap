import {graphql} from "../../gql";

export const getNodeByIdQueryDocument = graphql(`
  query getNodeById($publicKey: String!) {
    getNodeByKey(publicKey: $publicKey) {
      publicKey
      address
      lastSeen
    }
  }
`);