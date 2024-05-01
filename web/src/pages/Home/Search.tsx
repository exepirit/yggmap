export function Search() {
  return (
    <div className="container lg:w-2/3 flex">
      <input type="text" placeholder="Search nodes, public keys, addresses..." className="input input-bordered grow"/>
      <button className="btn btn-primary ml-2">Search</button>
    </div>
  );
}