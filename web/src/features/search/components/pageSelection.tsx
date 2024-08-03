interface PageSelectionProps {
  selectedPage: number
  maxPages: number
  onChange: (index: number) => void
}

export function PageSelection({ selectedPage, maxPages, onChange }: PageSelectionProps) {
  const handleClick = (page: number) => {
    if (page !== selectedPage) {
      onChange(page);
    }
  };

  return (
    <div className="join">
      {[...Array(maxPages)].map((_, index) => (
        <button
          key={index + 1}
          className={`join-item btn ${selectedPage === index + 1 ? 'btn-active' : ''}`}
          onClick={() => handleClick(index + 1)}
        >
          {index + 1}
        </button>
      ))}
    </div>
  );
}
