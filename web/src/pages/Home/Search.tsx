export function Search() {
  return (
    <div className="container w-5/6 flex">
      <input type="text" placeholder="Search" className="input input-bordered grow"/>
      <button className="btn btn-primary ml-2">Search</button>
    </div>
  );
}