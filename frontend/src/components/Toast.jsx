import React, { useState, useEffect, useCallback } from 'react';

/**
 * Individual toast notification.
 */
function ToastItem({ id, message, type, onDismiss }) {
  useEffect(() => {
    const timer = setTimeout(() => onDismiss(id), 4000);
    return () => clearTimeout(timer);
  }, [id, onDismiss]);

  return (
    <div
      role="status"
      aria-live="polite"
      className={`toast toast--${type}`}
    >
      <span className="toast__message">{message}</span>
      <button
        className="toast__close"
        aria-label="Dismiss"
        onClick={() => onDismiss(id)}
      >
        ×
      </button>
    </div>
  );
}

let _addToast = null;

/**
 * Imperative helper to show a toast from anywhere.
 * @param {string} message
 * @param {'success'|'error'|'info'} type
 */
export function showToast(message, type = 'info') {
  if (_addToast) _addToast({ message, type });
}

/**
 * Toast container — mount once near the app root.
 */
export function ToastContainer() {
  const [toasts, setToasts] = useState([]);

  const addToast = useCallback((toast) => {
    setToasts((prev) => [...prev, { ...toast, id: Date.now() + Math.random() }]);
  }, []);

  const dismiss = useCallback((id) => {
    setToasts((prev) => prev.filter((t) => t.id !== id));
  }, []);

  useEffect(() => {
    _addToast = addToast;
    return () => { _addToast = null; };
  }, [addToast]);

  return (
    <div className="toast-container" aria-label="Notifications">
      {toasts.map((t) => (
        <ToastItem key={t.id} {...t} onDismiss={dismiss} />
      ))}
    </div>
  );
}

export default ToastContainer;
