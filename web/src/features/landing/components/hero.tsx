export function Hero() {
  // TODO: background image or animation that represents the Yggdrasil network structure,
  // such as a mesmerizing tree-like pattern
  return (
    <div className="hero min-h-80">
      <div className="hero-content text-center">
        <div class="max-w-md">
          <h1 className="text-5xl font-bold">yggmap</h1>
          <p className="py-6">
            Discover nodes, explore connections in the Yggdrasil Network
          </p>
          <a href="/nodes" className="btn btn-primary">
            Explore
          </a>
        </div>
      </div>
    </div>
  );
}
