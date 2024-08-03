import {useState} from "preact/hooks";
import type {DocumentNode, OperationVariables, TypedDocumentNode} from "@apollo/client/core";
import type {NoInfer, QueryHookOptions} from "@apollo/client/react/types/types";
import {useQuery} from "@apollo/client";

type PaginatedVariables = OperationVariables | {
  limit: number;
  offset: number;
};

export function usePagination<TData, TVariables extends PaginatedVariables = PaginatedVariables>
  (
    query: DocumentNode | TypedDocumentNode<TData, TVariables>,
    options?: QueryHookOptions<NoInfer<TData>, NoInfer<TVariables>>
  ) {
  const pageSize = 10;
  const [page, setPage] = useState(0);

  const variables = {
    ...options?.variables,
    limit: pageSize,
    offset: pageSize * page,
  };

  const { data } = useQuery(query, { ...options, variables });

  return {
    data: data,
    setPage: setPage,
  }
}