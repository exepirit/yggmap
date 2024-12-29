export interface Graph {
  nodes: GraphNode[]
  links: GraphLink[]
}

export interface GraphNode {
  id: string
}

export interface GraphLink {
  source: string
  target: string
}
