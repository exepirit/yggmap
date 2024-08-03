import {SearchResults} from "../components/searchResults";
import {SearchBox} from "../components/searchBox";
import {PageSelection} from "../components/pageSelection";
import { useState } from "preact/hooks";

export function SearchPage() {
  const [query, setQuery] = useState("");
  const [page, setPage] = useState(1);

  return (
    <div className="container md:w-1/2 md:mx-auto mx-4">
      <SearchBox onChange={setQuery}/>
      <SearchResults query={query} page={page} />
      <PageSelection selectedPage={page} maxPages={10} onChange={setPage}/>
    </div>
  );
}
