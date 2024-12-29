import { Graph } from '../models';

export async function getNetworkGraph(): Promise<Graph> {
  const response = await fetch('/api/networkGraph');
  return await response.json();
}
