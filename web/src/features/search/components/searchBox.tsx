interface SearchBoxProps {
  onChange?: (query: string) => void;
}

export function SearchBox({onChange}: SearchBoxProps) {
  return (
    <input
      type="text"
      placeholder="Enter address or public key"
      className="input input-bordered w-full mb-8"
      onChange={(ev) => onChange?.(ev.currentTarget.value)}
    />
  );
}
