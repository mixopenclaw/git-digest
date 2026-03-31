import React from 'react';

/**
 * Global error boundary that catches render errors in child components.
 */
export class ErrorBoundary extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error) {
    return { hasError: true, error };
  }

  componentDidCatch(error, info) {
    console.error('[ErrorBoundary]', error, info.componentStack);
  }

  handleReset = () => {
    this.setState({ hasError: false, error: null });
  };

  render() {
    if (this.state.hasError) {
      return (
        <div role="alert" className="error-boundary">
          <h2>Something went wrong</h2>
          <p>An unexpected error occurred. Please reload the page.</p>
          <pre className="error-boundary__detail">
            {this.state.error && this.state.error.toString()}
          </pre>
          <button onClick={this.handleReset}>Try Again</button>
        </div>
      );
    }
    return this.props.children;
  }
}

export default ErrorBoundary;
