import {useEffect} from "preact/hooks";
import * as d3 from "d3";
import {createRef} from "react";

interface GraphNode extends d3.SimulationNodeDatum {
  id: string
  group: string
}

interface ForceDirectedGraphProps {
  nodes: GraphNode[]
  links: d3.SimulationLinkDatum<GraphNode>[]

  width: number
  height: number

}

export function ForceDirectedGraph(props: ForceDirectedGraphProps) {
  const svgRef = createRef();

  useEffect(() => {
    // Specify the color scale.
    const color = d3.scaleOrdinal(d3.schemeCategory10);

    // Create a simulation with several forces.
    const simulation = d3.forceSimulation(props.nodes)
      .force("link", d3.forceLink(props.links).
      id((d: GraphNode) => d.id))
      .force("charge", d3.forceManyBody())
      .force("center", d3.forceCenter(props.width / 2, props.height / 2))
      .on("tick", ticked);

    // Select the SVG container.
    const svg = d3.select(svgRef.current)
      .attr("viewBox", [0, 0, props.width, props.height]);

    // Add a line for each link, and a circle for each node.
    const link = svg.append("g")
      .attr("stroke", "#999")
      .attr("stroke-opacity", 0.6)
      .selectAll()
      .data(props.links)
      .join("line")
      .attr("stroke-width", "1px");

    const node = svg.append("g")
      .attr("stroke", "#fff")
      .attr("stroke-width", 1.5)
      .selectAll()
      .data(props.nodes)
      .join("circle")
      .attr("r", 5)
      .attr("fill", node => color(node.group));

    node.append("title")
      .text((d: GraphNode) => d.id);

    // Set the position attributes of links and nodes each time the simulation ticks.
    function ticked() {
      link
        .attr("x1", d => (d.source as GraphNode).x)
        .attr("y1", d => (d.source as GraphNode).y)
        .attr("x2", d => (d.target as GraphNode).x)
        .attr("y2", d => (d.target as GraphNode).y);

      node
        .attr("cx", d => d.x)
        .attr("cy", d => d.y);
    }

    return () => {
      simulation.stop();
    }
  }, []);

  return <svg ref={svgRef}/>
}