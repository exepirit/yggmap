interface ErrorBannerProps {
  text?: string;
}

/**
 * Represents an error banner in the UI.
 * Displays important messages to the user in a alert box format.
 */
export function ErrorBanner(props: ErrorBannerProps) {
  return (
    <div role="alert" className="alert alert-error">
      <span>Error!</span>
      {props.text && <span>{props.text}</span>}
    </div>
  );
}
