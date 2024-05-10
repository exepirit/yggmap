export interface YggdrasilNode {
  address: string;
  publicKey: string;
  lastSeen: string;
  neighbors?: YggdrasilNode[];
}
